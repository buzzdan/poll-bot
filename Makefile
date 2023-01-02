test-colors:
	go get -u github.com/kyoh86/richgo && richgo test ./...

tidy:
	go mod tidy

test:
	go test ./...