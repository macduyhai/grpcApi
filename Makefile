gen-api :
	protoc apiproto/apipb.proto --go_out=plugins=grpc:. 