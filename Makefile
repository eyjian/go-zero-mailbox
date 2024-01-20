# Written by yijian on 2024/01/20

.PHONY: rpc sql tidy

rpc:
	goctl rpc protoc mooon-mailbox.proto --go_out=./pb --go-grpc_out=./pb --zrpc_out=. --style go_zero

sql:
	goctl model mysql ddl --src mooon-mailbox.sql --dir .

tidy:
	go mod tidy
