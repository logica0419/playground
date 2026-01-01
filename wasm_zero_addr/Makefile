# ----------------- Go -----------------

main-go.wasm: main.go
	tinygo build -target wasip1 -opt=0 -o main-go.wasm main.go

main-go.wat: main-go.wasm
	wasm2wat main-go.wasm -o main-go.wat

.PHONY: run-go
run-go: main-go.wasm main-go.wat
	wasmtime run --invoke "test_zero" main-go.wasm

# ----------------- Rust -----------------

target/wasm32-wasip1/debug/main-rs.wasm: main.rs Cargo.toml
	cargo build --target wasm32-wasip1

main-rs.wat: target/wasm32-wasip1/debug/main-rs.wasm
	wasm2wat target/wasm32-wasip1/debug/main-rs.wasm -o main-rs.wat

.PHONY: run-rs
run-rs: target/wasm32-wasip1/debug/main-rs.wasm main-rs.wat
	wasmtime run --invoke "test_zero" target/wasm32-wasip1/debug/main-rs.wasm

# ----------------- Other -----------------

.PHONY: all
all: run-go run-rs

.PHONY: test
test:
	@echo "No tests defined"

.PHONY: clean
clean:
	@rm -rf target/wasm32-wasip1/debug
	@rm -f *.wasm
	@rm -f *.wat
