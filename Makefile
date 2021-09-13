proto-weather:
	protoc --proto_path=./schemas \
		--go_out=server \
		--go_opt=Mweather.proto=./schemas \
		--go-grpc_out=server/schemas \
		--go-grpc_opt=Mweather.proto=./schemas \
		schemas/weather.proto

proto-student:
	protoc --proto_path=./schemas \
		--go_out=server \
		--go_opt=Mstudent.proto=./schemas \
		--go-grpc_out=server/schemas \
		--go-grpc_opt=Mstudent.proto=./schemas \
		schemas/student.proto

	# python -m grpc_tools.protoc \
	# 	-I ./schemas \
	# 	--python_out=client/schemas \
	# 	--grpc_python_out=client/schemas \
	# 	schemas/student.proto

.PHONY: proto-weather
.PHONY: proto-student