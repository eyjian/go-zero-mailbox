package logic

import (
	"context"

	"mooon-mailbox/internal/svc"
	"mooon-mailbox/pb/mooon-mailbox"

	"github.com/zeromicro/go-zero/core/logx"
)

type MarkMessageAsReadLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewMarkMessageAsReadLogic(ctx context.Context, svcCtx *svc.ServiceContext) *MarkMessageAsReadLogic {
	return &MarkMessageAsReadLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

func (l *MarkMessageAsReadLogic) MarkMessageAsRead(in *mooon_mailbox.MarkAsReadReq) (*mooon_mailbox.MarkAsReadResp, error) {
	// todo: add your logic here and delete this line

	return &mooon_mailbox.MarkAsReadResp{}, nil
}
