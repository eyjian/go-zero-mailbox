// Package logic
// Written by yijian on 2024/01/20
package logic

import (
	"github.com/zeromicro/go-zero/core/logc"
	"mooon-mailbox/model"
	mooonmailbox "mooon-mailbox/pb/mooon-mailbox"
	"strconv"
	"time"
)

func deliverMessage(l *DeliverMessageLogic, in *mooonmailbox.DeliverMessageReq) (*mooonmailbox.DeliverMessageResp, error) {
	now := time.Now()
	letter := model.TMooonMailbox{
		FRecipient:   in.Recipient,
		FDeliverTime: now,
		FArrivalTime: now,
		FState:       int64(mooonmailbox.Letter_LETTER_UNREAD),
		FLetterBody:  in.LetterBody,
	}
	dbResult, err := l.svcCtx.MailboxModel.Insert(l.ctx, &letter)
	if err != nil {
		logc.Errorf(l.ctx, "Insert %s to table error: %s\n", in.String(), err.Error())
		return nil, err
	} else {
		var rowsAffected int64 = -1
		var lastInsertId int64 = -1
		rowsAffected, _ = dbResult.RowsAffected()
		lastInsertId, _ = dbResult.LastInsertId()
		logc.Infof(l.ctx, "Insert %s success: RowsAffected=%d, LastInsertId=%d\n", in.String(), rowsAffected, lastInsertId)
		return &mooonmailbox.DeliverMessageResp{
			Recipient: in.Recipient,
			LetterId:  strconv.FormatInt(lastInsertId, 10),
		}, nil
	}
}
