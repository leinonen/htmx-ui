test:
	docker compose up --build --abort-on-container-exit --exit-code-from tests

run:
	docker compose up --build api html

build:
	docker compose build

dev:
	docker compose -f compose.dev.yaml up --build

dev-down:
	docker compose -f compose.dev.yaml down