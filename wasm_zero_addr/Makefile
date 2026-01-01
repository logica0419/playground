main.wasm: main.go
	tinygo build -target wasip1 -opt=0 main.go

PHONY: run
run: main.wasm
	wasmtime run --invoke "test_zero" main.wasm

PHONY: clean
clean:
	rm -f main.wasm
