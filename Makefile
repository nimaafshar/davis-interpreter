.PHONY: build
build:
	go build -o bin/decoder cmd/decoder/main.go
	go build -o bin/interpreter cmd/interpreter/main.go

.PHONY: tests
tests:
	cat tests/interpreter/input.txt | ./bin/interpreter | diff tests/interpreter/output.txt -
	cat tests/decoder/input.txt | ./bin/decoder | diff tests/decoder/output.txt -