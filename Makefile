# Written by yijian on 2024/01/20

all: mooon_mailbox_service

mooon_mailbox_service: mooon_mailbox.go internal/logic/deliver_message_impl.go internal/logic/list_messages_impl.go
	go build -o $@ $<

.PHONY: clean rpc sql tidy

clean:
	rm -f mooon_mailbox_service

rpc:
	goctl rpc protoc mooon_mailbox.proto --go_out=./pb --go-grpc_out=./pb --zrpc_out=. --style go_zero

sql:
	goctl model mysql ddl --src mooon_mailbox.sql --dir ./model -c --style go_zero

tidy:
	go mod tidy
