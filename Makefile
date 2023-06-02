.PHONY: clean
clean:
	@rm -r build

.PHONY: build
build: 
	go build -o build/program/app cmd/cli/main.go 

.PHONY: run
run:
	make build
	./build/program/app

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

.PHONY: deploy
deploy:
	docker build --tag busr-bot .
	docker run busr-bot &