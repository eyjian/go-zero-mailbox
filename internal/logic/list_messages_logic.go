package logic

import (
	"context"

	"mooon-mailbox/internal/svc"
	"mooon-mailbox/pb/mooon_mailbox"

	"github.com/zeromicro/go-zero/core/logx"
)

type ListMessagesLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewListMessagesLogic(ctx context.Context, svcCtx *svc.ServiceContext) *ListMessagesLogic {
	return &ListMessagesLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

func (l *ListMessagesLogic) ListMessages(in *mooon_mailbox.ListMessagesReq) (*mooon_mailbox.ListMessagesResp, error) {
	// todo: add your logic here and delete this line
	return listMessages(l, in)
	//return &mooon_mailbox.ListMessagesResp{}, nil
}
