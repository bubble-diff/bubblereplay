update-idl:
	@rm -rf ./idl
	@echo "step1> fetching idl repo..."
	@git clone --depth=1 https://github.com/bubble-diff/IDL.git idl
	@rm -rf ./idl/.git
	@rm ./idl/.gitignore
	@echo "step2> compile idl..."
	@protoc --go_out=. idl/*.proto
	@go mod tidy

build:
	@mkdir -p output/
	@cp conf/* output/
	@go build -o output/bubblereplay

run: build
	@cd output && ./bubblereplay