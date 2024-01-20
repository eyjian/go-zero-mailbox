package logic

import (
	"context"

	"mooon-mailbox/internal/svc"
	"mooon-mailbox/pb/mooon-mailbox"

	"github.com/zeromicro/go-zero/core/logx"
)

type DeleteLetterLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewDeleteLetterLogic(ctx context.Context, svcCtx *svc.ServiceContext) *DeleteLetterLogic {
	return &DeleteLetterLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

func (l *DeleteLetterLogic) DeleteLetter(in *mooon_mailbox.DeleteLetterReq) (*mooon_mailbox.DeleteLetterResp, error) {
	// todo: add your logic here and delete this line

	return &mooon_mailbox.DeleteLetterResp{}, nil
}
