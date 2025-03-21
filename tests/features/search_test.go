package main

import (
	"context"
	"os"
	"testing"
	"time"

	"github.com/chromedp/chromedp"
	"github.com/cucumber/godog"
)

var (
	browserCtx context.Context
	cancel     context.CancelFunc
	baseURL    = "http://localhost:8080" // fallback if env var is not set
	nameOnPage string
)

func theServerIsRunning() error {
	if envURL := os.Getenv("E2E_BASE_URL"); envURL != "" {
		baseURL = envURL
	}

	// Set up Chromedp with custom headless config (Docker-friendly)
	opts := append(chromedp.DefaultExecAllocatorOptions[:],
		chromedp.ExecPath("/usr/bin/chromium"),
		chromedp.Flag("headless", true),
		chromedp.Flag("disable-gpu", true),
		chromedp.Flag("no-sandbox", true),
	)

	allocCtx, _ := chromedp.NewExecAllocator(context.Background(), opts...)
	browserCtx, cancel = chromedp.NewContext(allocCtx)

	// Navigate to base URL
	return chromedp.Run(browserCtx,
		chromedp.Navigate(baseURL),
		chromedp.WaitVisible(`input[name="q"]`, chromedp.ByQuery),
	)
}

func iSearchFor(name string) error {
	return chromedp.Run(browserCtx,
		chromedp.WaitVisible(`input[name="q"]`),
		chromedp.SendKeys(`input[name="q"]`, name),
		chromedp.Sleep(500*time.Millisecond),
	)
}

func iClickOnTheFirstResult() error {
	return chromedp.Run(browserCtx,
		chromedp.Click(`#results a`, chromedp.NodeVisible),
		chromedp.WaitVisible(`h1.title`),
	)
}

func iShouldSeeThePersonsNameAs(expected string) error {
	return chromedp.Run(browserCtx,
		chromedp.Text(`h1.title`, &nameOnPage),
	)
}

func InitializeScenario(ctx *godog.ScenarioContext) {
	ctx.Step(`^the server is running$`, theServerIsRunning)
	ctx.Step(`^I search for "([^"]*)"$`, iSearchFor)
	ctx.Step(`^I click on the first result$`, iClickOnTheFirstResult)
	ctx.Step(`^I should see the person's name as "([^"]*)"$`, iShouldSeeThePersonsNameAs)

	ctx.AfterScenario(func(sc *godog.Scenario, err error) {
		if cancel != nil {
			cancel()
		}
	})
}

func TestFeatures(t *testing.T) {
	opts := godog.Options{
		Format: "pretty",
		Paths:  []string{"./search.feature"},
	}

	status := godog.TestSuite{
		Name:                "search",
		ScenarioInitializer: InitializeScenario,
		Options:             &opts,
	}.Run()

	if status != 0 {
		t.Fail()
	}
}
