
up:
	docker-compose up -d --build

down:
	docker-compose down --remove-orphans

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
	go run cmd/server/main.go

build:
	go build -o server cmd/server/main.go

cover:
	go test -short -tags=unit -race -coverprofile=coverage.out ./...
	go tool cover -func=coverage.out
	rm coverage.out

cover-html:
	go test -short -tags=unit -race -coverprofile=coverage.out ./...
	go tool cover -html=coverage.out
	rm coverage.out

migration-up:
	migrate -database ${DATABASE_URI} -path migrations up

migration-down:
	migrate -database ${DATABASE_URI} -path migrations down

mock-generate:
	mockery --dir=internal/domain/port --all --with-expecter

protoc:
	protoc --go_out=. --go_opt=paths=source_relative --go-grpc_out=. --go-grpc_opt=paths=source_relative ./pkg/proto/app.proto

ssl-generate:
	openssl req -newkey rsa:4096 -nodes -keyout local.key -x509 -sha256 -days 365 -subj "/C=US/ST=WA/L=SEATTLE/O=MyCompany/OU=MyDivision/CN=*.domain.test" -addext "subjectAltName = DNS:*.domain.test, DNS:localhost, DNS:127.0.0.1, DNS:mail.domain.test" -out local.crt