.PHONY: init up down provider consumer ingestor

init:
	@cp .env.dist .env

up:
	@docker compose up -d --no-build --remove-orphans ${s}

down:
	@docker compose down --remove-orphans ${s}

provider:
	@go run ./provider/

consumer:
	@go run main.go start consumer --topic=${t}

ingestor:
	@go run main.go start ingestor
