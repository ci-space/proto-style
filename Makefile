lint:
	golangci-lint run --fix

build:
	go build main.go

.PHONY: e2e
e2e:
	go build -o e2e/protostyle main.go
	cd e2e && go test

.PHONY: test
test:
	go test ./...
