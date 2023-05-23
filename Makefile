.PHONY: dev

dev: 
	@./scripts/dev.sh

.PHONY: bot

bot: 
	@./scripts/bot.sh

.PHONY: test

test:
	go clean -testcache 
	go mod tidy
	go test -cover -v ./...