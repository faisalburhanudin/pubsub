gen-proto:
	protoc -I ping/ ping/ping.proto --go_out=plugins=grpc:ping