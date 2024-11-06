# Build an Exchanger service
.PHONY: build
build:
	@ go build -o ./bin/server ./main.go

# Generate gRPC code from specification proto-exchanger/exchanger/v1
.PHONY: gen-proto
gen-proto:
	@ rm -rf gen/exchange
	@ protoc --proto_path=proto-exchanger --go_out=gen \
	 --go_opt=paths=source_relative \
	 --go-grpc_out=gen --go-grpc_opt=paths=source_relative \
	  proto-exchanger/exchanger/v1/exchanger.proto
