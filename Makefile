build:
	@mkdir -p output/
	@cp conf/* output/
	@go build -o output/bubblereplay

run: build
	@cd output && ./bubblereplay