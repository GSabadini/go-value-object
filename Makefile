#!make

test:
	go test -cover ./...

code-review: fmt vet test

fmt:
	go fmt ./...

vet:
	go vet ./...

ci:
	docker run --rm -it -v $(PWD):/app -w /app golangci/golangci-lint:v1.24.0 \
	golangci-lint run
    --exclude-use-default=false \
		--enable=golint \
		--enable=gocyclo \
		--enable=goconst \
		--enable=unconvert \
		./...