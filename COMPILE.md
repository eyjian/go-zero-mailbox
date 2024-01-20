第一次执行"make"编译时，如果遇到下列错误，请执行下"make tidy"后再执行"make"编译：

```text
mooonmailbox.go:17:2: missing go.sum entry for module providing package github.com/zeromicro/go-zero/core/conf; to add:
        go mod download github.com/zeromicro/go-zero
mooonmailbox.go:18:2: missing go.sum entry for module providing package github.com/zeromicro/go-zero/core/service; to add:
        go mod download github.com/zeromicro/go-zero
mooonmailbox.go:7:2: missing go.sum entry for module providing package github.com/zeromicro/go-zero/core/stores/sqlc; to add:
        go mod download github.com/zeromicro/go-zero
mooonmailbox.go:8:2: missing go.sum entry for module providing package github.com/zeromicro/go-zero/core/stores/sqlx; to add:
        go mod download github.com/zeromicro/go-zero
mooonmailbox.go:19:2: missing go.sum entry for module providing package github.com/zeromicro/go-zero/zrpc; to add:
        go mod download github.com/zeromicro/go-zero
mooonmailbox.go:20:2: missing go.sum entry for module providing package google.golang.org/grpc; to add:
        go mod download google.golang.org/grpc
mooonmailbox.go:21:2: missing go.sum entry for module providing package google.golang.org/grpc/reflection; to add:
        go mod download google.golang.org/grpc
internal/config/config.go:4:2: missing go.sum entry for module providing package github.com/zeromicro/go-zero/core/stores/cache (imported by mooon-mailbox/internal/config); to add:
        go get mooon-mailbox/internal/config
internal/logic/delete_messages_impl.go:7:2: missing go.sum entry for module providing package github.com/pkg/errors (imported by mooon-mailbox/internal/logic); to add:
        go get mooon-mailbox/internal/logic
internal/logic/delete_messages_impl.go:8:2: missing go.sum entry for module providing package github.com/zeromicro/go-zero/core/logc (imported by mooon-mailbox/internal/logic); to add:
        go get mooon-mailbox/internal/logic
internal/logic/delete_messages_logic.go:9:2: missing go.sum entry for module providing package github.com/zeromicro/go-zero/core/logx (imported by mooon-mailbox/internal/logic); to add:
        go get mooon-mailbox/internal/logic
model/t_mooon_mailbox_model_gen.go:12:2: missing go.sum entry for module providing package github.com/zeromicro/go-zero/core/stores/builder (imported by mooon-mailbox/model); to add:
        go get mooon-mailbox/model
model/t_mooon_mailbox_model_gen.go:16:2: missing go.sum entry for module providing package github.com/zeromicro/go-zero/core/stringx (imported by mooon-mailbox/model); to add:
        go get mooon-mailbox/model
pb/mooon-mailbox/mooon-mailbox_grpc.pb.go:14:2: missing go.sum entry for module providing package google.golang.org/grpc/codes (imported by mooon-mailbox/pb/mooon-mailbox); to add:
        go get mooon-mailbox/pb/mooon-mailbox
pb/mooon-mailbox/mooon-mailbox_grpc.pb.go:15:2: missing go.sum entry for module providing package google.golang.org/grpc/status (imported by mooon-mailbox/pb/mooon-mailbox); to add:
        go get mooon-mailbox/pb/mooon-mailbox
pb/mooon-mailbox/mooon-mailbox.pb.go:12:2: missing go.sum entry for module providing package google.golang.org/protobuf/reflect/protoreflect (imported by mooon-mailbox/pb/mooon-mailbox); to add:
        go get mooon-mailbox/pb/mooon-mailbox
pb/mooon-mailbox/mooon-mailbox.pb.go:13:2: missing go.sum entry for module providing package google.golang.org/protobuf/runtime/protoimpl (imported by mooon-mailbox/pb/mooon-mailbox); to add:
        go get mooon-mailbox/pb/mooon-mailbox
make: *** [mooonmailbox_service] Error 1
```
