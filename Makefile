
.PHONY: wasm-binarytree

test:
	go test -v .

run:
	go run .

wasm-binarytree:
	tinygo build -wasm-abi=generic -target=wasi -o binarytree.wasm binarytree.go
