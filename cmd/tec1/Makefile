.PHONY: all
all:
	go build

.PHONY: wasm
wasm:
	GOOS=js GOARCH=wasm go build -o tec1.wasm .

.PHONY: clean
clean:
	go clean
	-rm history.txt
	-rm *.wasm
