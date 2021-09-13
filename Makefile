proto:
	protoc --proto_path=./schemas \
	--go_out=server \
	--go_opt=Mweather.proto=./schemas \
	--go-grpc_out=server/schemas \
	--go-grpc_opt=Mweather.proto=./schemas \
	schemas/weather.proto

	# python -m grpc_tools.protoc \
	# 	-I ./protobuf \
	# 	--python_out=. \
	# 	--python-grpc_out=. \
	# 	protobuf/weather.proto
.PHONY: proto