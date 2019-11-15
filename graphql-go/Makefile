OUTPUT = ./cmd/graphql/main
PACKAGED_TEMPLATE = packaged.yml

.PHONY: test
test:
	go test ./...

.PHONY: clean
clean:
	rm -f $(OUTPUT) $(PACKAGED_TEMPLATE)
	rm -f ./internal/app/graphql/schema/bindata.go

generate:
	go generate ./internal/app/graphql/schema

graphql: generate
	go build -o $(OUTPUT) ./cmd/graphql/main.go

.PHONY: lambda
lambda:
	GOOS=linux GOARCH=amd64 $(MAKE) graphql

.PHONY: build
build: clean lambda

.PHONY: local
local: build
	sam local start-api
