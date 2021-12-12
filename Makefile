run:
	echo "RUN"
	go run cmd/scanner/main.go

build:
	mkdir -p "bin"
	go build -o bin/scanner cmd/scanner/main.go

lint:
	golangci-lint run