default: build

build:
	go build -o ldapctl main.go

build-linux:
	CGO_ENABLED=0 GOARCH=amd64 GOOS=linux go build -o ldapctl main.go

build-linux-arm:
	CGO_ENABLED=0 GOARCH=arm64 GOOS=linux go build -o ldapctl main.go