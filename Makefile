##Proto generate code
protoc:
	@echo "Build protobuf file"
	@protoc --go_out=. --go_opt=paths=source_relative \
		--go-grpc_out=. --go-grpc_opt=paths=source_relative \
		rpc/rka/rka.proto