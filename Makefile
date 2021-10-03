lint:
	go fmt ./...


build:
	go build -v .


test:
	go test -v ./...


run:
	go run main.go