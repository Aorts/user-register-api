dev: 
	nodemon --exec go run --tags dynamic $(shell pwd)/cmd/main.go --signal SIGTERM

proto:
	protoc pkg/**/pb/*.proto --go-grpc_out=.