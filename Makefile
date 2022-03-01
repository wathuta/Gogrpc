gen:
		protoc -I=./proto --go_out=./ ./proto/currency.proto
newgen:
	protoc -I=./proto --go_out=pb --go_opt=paths=source_relative --go-grpc_out=pb --go-grpc_opt=paths=source_relative proto/currency.proto
funcGen:
	protoc --go_out=plugins=grpc:../ *.proto