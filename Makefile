
build:
	go mod tidy
	go build .

vet:
	go vet ./...

test:
	go test -v ./...