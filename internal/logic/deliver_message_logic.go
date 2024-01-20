package logic

import (
	"context"

	"mooon-mailbox/internal/svc"
	"mooon-mailbox/pb/mooon_mailbox"

	"github.com/zeromicro/go-zero/core/logx"
)

type DeliverMessageLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewDeliverMessageLogic(ctx context.Context, svcCtx *svc.ServiceContext) *DeliverMessageLogic {
	return &DeliverMessageLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

func (l *DeliverMessageLogic) DeliverMessage(in *mooon_mailbox.DeliverMessageReq) (*mooon_mailbox.DeliverMessageResp, error) {
	// todo: add your logic here and delete this line
	return deliverMessage(l, in)
	//return &mooon_mailbox.DeliverMessageResp{}, nil
}
