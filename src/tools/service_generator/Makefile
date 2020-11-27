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

gqlproto:
	protoc -I. \
	-I$(GOPATH)/src \
	-I=$(GOPATH)/src/github.com/infobloxopen/protoc-gen-gorm \
	-I=$(GOPATH)/src/github.com/infobloxopen/atlas-app-toolkit \
	-I=$(GOPATH)/src/github.com/lyft/protoc-gen-validate/validate/validate.proto \
	-I=$(GOPATH)/src/github.com/infobloxopen/protoc-gen-gorm/options \
	-I=$(GOPATH)/src/github.com/protobuf/src/google/protobuf/timestamp.proto \
	--gogoopsee_out=plugins=grpc+graphql,Mopsee/protobuf/opsee.proto=github.com/opsee/protobuf/opseeproto,Mgoogle/protobuf/descriptor.proto=github.com/gogo/protobuf/protoc-gen-gogo/descriptor:./$(target) --proto_path=$(GOPATH)/src:. $(source)

protogen: gqlproto
	@echo "setting up service schema definition via protobuf"
	protoc -I. \
			-I$(GOPATH)/src \
			-I=$(GOPATH)/src/github.com/infobloxopen/protoc-gen-gorm \
			-I=$(GOPATH)/src/github.com/infobloxopen/atlas-app-toolkit \
			-I=$(GOPATH)/src/github.com/lyft/protoc-gen-validate/validate/validate.proto \
			-I=$(GOPATH)/src/github.com/infobloxopen/protoc-gen-gorm/options \
			-I=$(GOPATH)/src/github.com/protobuf/src/google/protobuf/timestamp.proto \
			--proto_path=${GOPATH}/src/github.com/gogo/protobuf/protobuf \
			--govalidators_out=./$(target) \
			--go_out=plugins=grpc:./$(target) --gorm_out="engine=postgres:./$(target)" $(source)

updateSchema: protogen gqlproto
	go run github.com/99designs/gqlgen

all: build compile create-service
