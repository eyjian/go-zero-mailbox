package logic

import (
	"context"

	"mooon-mailbox/internal/svc"
	"mooon-mailbox/pb/mooon-mailbox"

	"github.com/zeromicro/go-zero/core/logx"
)

type DeleteMessagesLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewDeleteMessagesLogic(ctx context.Context, svcCtx *svc.ServiceContext) *DeleteMessagesLogic {
	return &DeleteMessagesLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

func (l *DeleteMessagesLogic) DeleteMessages(in *mooon_mailbox.DeleteMessagesReq) (*mooon_mailbox.DeleteMessagesResp, error) {
	// todo: add your logic here and delete this line

	return &mooon_mailbox.DeleteMessagesResp{}, nil
}
