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
	go run cmd/client/main.go

cover:
	go test -short -tags=unit -race -coverprofile=coverage.out ./...
	go tool cover -func=coverage.out
	rm coverage.out

cover-html:
	go test -short -tags=unit -race -coverprofile=coverage.out ./...
	go tool cover -html=coverage.out
	rm coverage.out
