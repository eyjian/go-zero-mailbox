# Written by yijian on 2024/01/20

.PHONY: tidy

rpc:
	goctl rpc protoc mooon-mailbox.proto --go_out=./pb --go-grpc_out=./pb --zrpc_out=. --style go_zero

tidy:
	go mod tidy
