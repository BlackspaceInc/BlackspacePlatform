build:
	go build -o bin/main main.go

create-service:
	./bin/main generateBlackspaceServiceTemplate -a "REST" -s "test"

run:
	go run main.go

compile:
	echo "Compiling for every OS and Platform"
	GOOS=linux GOARCH=arm go build -o bin/main-linux-arm main.go
	GOOS=linux GOARCH=arm64 go build -o bin/main-linux-arm64 main.go
	GOOS=freebsd GOARCH=386 go build -o bin/main-freebsd-386 main.go

generate: build create-service

all: build compile create-service
