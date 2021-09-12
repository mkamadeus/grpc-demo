proto:
	protoc \
		--go_out=. \
		--go_opt=paths=source_relative \
		--go-grpc_out=. \
		--go-grpc_opt=paths=source_relative \
		protobuf/weather.proto

	python -m grpc_tools.protoc \
		-I ./protobuf \
		--python_out=. \
		--python-grpc_out=. \
		protobuf/weather.proto
.PHONY: proto