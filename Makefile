.PHONY: cli bot testc

cli: 
	@./scripts/run-cli.sh

test:
	@./scripts/run-test.sh

bot: 
	@./scripts/run-bot.sh