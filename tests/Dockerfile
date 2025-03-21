FROM golang:1.24.1

# Install dependencies
RUN apt-get update && apt-get install -y \
    chromium \
    xvfb \
    libx11-xcb1 \
    libxcb-dri3-0 \
    libgbm1 \
    libgtk-3-0 \
    libnss3 \
    fonts-liberation \
    curl \
    && rm -rf /var/lib/apt/lists/*

# Set up workdir
WORKDIR /app
COPY . .

# Get dependencies
RUN go mod tidy

# Use xvfb-run to fake display for chromedp
CMD xvfb-run go test ./features -v
