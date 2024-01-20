package logic

import (
	"context"

	"mooon-mailbox/internal/svc"
	"mooon-mailbox/pb/mooon_mailbox"

	"github.com/zeromicro/go-zero/core/logx"
)

type MarkMessagesAsReadLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewMarkMessagesAsReadLogic(ctx context.Context, svcCtx *svc.ServiceContext) *MarkMessagesAsReadLogic {
	return &MarkMessagesAsReadLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

func (l *MarkMessagesAsReadLogic) MarkMessagesAsRead(in *mooon_mailbox.MarkMessagesAsReadReq) (*mooon_mailbox.MarkMessagesAsReadResp, error) {
	// todo: add your logic here and delete this line
	return markMessagesAsRead(l, in)
	//return &mooon_mailbox.MarkMessagesAsReadResp{}, nil
}
