gen:
	wire gen ./...

tidy:
	go mod tidy -v

run:
	go run main.go

build:
	go build
