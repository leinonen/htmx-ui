test:
	docker compose up --build --abort-on-container-exit --exit-code-from tests

run:
	docker compose up --build api html
