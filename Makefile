gen:
	#protoc --proto_path=proto proto/*.proto --go_out=plugins=grpc:pb
	protoc ./proto/*.proto --go_out=./pb
clean:
	rm ./pb/pb/*.go
