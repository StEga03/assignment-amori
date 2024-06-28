-include .envrc

# Run from docker entirely
.PHONY: run
run:
	docker-compose --profile dev up --build

# Stop local services from docker compose
.PHONY: stop
stop:
	@docker-compose down