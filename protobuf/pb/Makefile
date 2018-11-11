
api.pb.go: api.proto Makefile
	protoc --go_out=plugins=grpc:. api.proto

clean:
	rm -f api.pb.go
