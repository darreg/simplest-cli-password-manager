NOW := $(shell date +'%Y/%m/%d')

lint:
	golangci-lint run

fix:
	golangci-lint run --fix

unit-test:
	go test -v -tags=unit ./...

integration-test:
	go test -tags=integration -v ./...

race:
	go test -v -race -count=1 ./...

run:
	go run -ldflags "-X main.buildVersion=v1.0.0 -X main.buildDate=$(NOW)" cmd/client/main.go

build:
	go build -o client -ldflags "-X main.buildVersion=v1.0.0 -X main.buildDate=$(NOW)" cmd/client/main.go

cover:
	go test -short -tags=unit -race -coverprofile=coverage.out ./...
	go tool cover -func=coverage.out
	rm coverage.out

cover-html:
	go test -short -tags=unit -race -coverprofile=coverage.out ./...
	go tool cover -html=coverage.out
	rm coverage.out

mock-generate:
	mockery --dir=internal/domain/port --all --with-expecter