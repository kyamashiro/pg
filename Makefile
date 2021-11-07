deps:
	@echo "Downloading dependencies..."
	@GO111MODULE=on go mod download

test: deps
	@echo "Running tests..."
	go test -v ./...

build:
	go build -o gp main.go
