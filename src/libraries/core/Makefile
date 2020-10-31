.PHONY: proto
proto: contracts
	@echo "setting up service schema definition via protobuf"
	protoc -I. \
			-I$(GOPATH)/src \
			-I=$(GOPATH)/src/github.com/infobloxopen/protoc-gen-gorm \
			-I=$(GOPATH)/src/github.com/infobloxopen/atlas-app-toolkit \
			-I=$(GOPATH)/src/github.com/lyft/protoc-gen-validate/validate/validate.proto \
			-I=$(GOPATH)/src/github.com/infobloxopen/protoc-gen-gorm/options \
			-I=$(GOPATH)/src/github.com/protobuf/src/google/protobuf/timestamp.proto \
			--proto_path=${GOPATH}/src/github.com/gogo/protobuf/protobuf \
            --govalidators_out=./schema/models/ \
			--go_out=plugins=grpc:./schema/models/  --gorm_out="engine=postgres:./schema/models/" ./schema/proto/models/*.proto
.PHONY: contracts
contracts:
		@echo "setting up service level contracts with the other microservices via proto definitions"
		protoc -I. \
    			-I=$(GOPATH)/src \
    			--go_out=plugins=grpc:./schema/models/ ./schema/proto/contracts/*.proto
