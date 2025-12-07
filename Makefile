gen:
	@protoc \
	--proto_path=protobuf "protobuf/coffee_shop.proto" \
	--go_out=./coffeeshop_proto --go_opt=paths=source_relative \
	--go-grpc_out=./coffeeshop_proto --go-grpc_opt=paths=source_relative