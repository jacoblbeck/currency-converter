.PHONY: start
start:
	docker compose up -d --build

.PHONY: stop
stop:
	docker compose down

.PHONY: restart
restart: stop start
