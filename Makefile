build-wasm: ## Build the WebAssembly binary
	GOOS=js GOARCH=wasm go build -v -o main.wasm
dist:
	mkdir -p dist && cp index.html dist/ && cp wasm_exec.js dist/ && cp main.wasm dist/ && cp -R sounds/ dist/sounds/