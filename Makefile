# Written by yijian on 2024/01/20

all: mooon_mailbox_service

mooon_mailbox_service: mooonmailbox.go
	go build -o $@ $<

.PHONY: clean rpc sql tidy

clean:
	rm -f mooon_mailbox_service

rpc:
	goctl rpc protoc mooon-mailbox.proto --go_out=./pb --go-grpc_out=./pb --zrpc_out=. --style go_zero

sql:
	goctl model mysql ddl --src mooon-mailbox.sql --dir ./model -c --style go_zero

tidy:
	go mod tidy
