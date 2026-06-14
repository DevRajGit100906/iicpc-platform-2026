.DEFAULT_GOAL := help
up:
	docker compose up -d
down:
	docker compose down
test:
	go test ./...
race:
	go test -race ./...
build:
	go build ./...
vet:
	go vet ./...
fmt:
	gofmt -w .
proto:
	cd proto && buf generate
engine:
	go run ./sut/reference-engine
orchestrator:
	go run ./services/orchestrator
leaderboard:
	go run ./services/leaderboard
ingester:
	go run ./services/ingester
frontend:
	cd frontend && python -m http.server 3000
help:
	@grep -E '^[a-zA-Z_-]+:.*?## .*$$' $(MAKEFILE_LIST) | awk 'BEGIN{FS=":.*?## "}{printf "  \033[36m%-14s\033[0m %s\n", $$1, $$2}'
.PHONY: up down test race build vet fmt proto engine orchestrator leaderboard ingester frontend help