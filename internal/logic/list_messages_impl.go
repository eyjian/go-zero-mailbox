// Package logic
// Written by yijian on 2024/01/20
package logic

import (
	"fmt"
	"github.com/zeromicro/go-zero/core/logc"
	"mooon-mailbox/model"
	mooonmailbox "mooon-mailbox/pb/mooon-mailbox"
)

func listMessages(l *ListMessagesLogic, in *mooonmailbox.ListMessagesReq) (*mooonmailbox.ListMessagesResp, error) {
	var letters []model.TMooonMailbox
	var nextPageStart int64 = 0

	pageStart := in.PageStart
	if pageStart == 0 {
		pageStart = 1000000000000
	}
	pageSize := in.PageSize
	if pageSize > 10 {
		pageSize = 10
	} else if pageSize < 1 {
		pageSize = 10
	}

	sql := fmt.Sprintf(
		"SELECT "+
			"f_id,"+
			"f_recipient,"+
			"f_deliver_time,"+
			"f_arrival_time,"+
			"f_state,"+
			"f_letter_body "+
			"FROM t_mooon_mailbox "+
			"WHERE "+
			"f_recipient='%s' AND "+
			"f_id<%d "+
			"ORDER BY f_id DESC "+
			"LIMIT %d",
		in.Recipient,
		pageStart,
		in.PageSize)
	err := l.svcCtx.CachedConn.QueryRowsNoCacheCtx(l.ctx, &letters, sql)
	if err != nil {
		logc.Errorf(l.ctx, "Exec %s error: %s", sql, err.Error())
		return nil, err
	} else {
		var messages mooonmailbox.ListMessagesResp
		logc.Infof(l.ctx, "Exec %s success\n", sql)

		// 将 letters 转换为 Letter 类型的切片
		var letterList []*mooonmailbox.Letter
		for _, letter := range letters {
			l := &mooonmailbox.Letter{
				LetterId:    letter.FId,
				DeliverTime: letter.FDeliverTime.Format("2006-01-02 15:04:05"),
				ArrivalTime: letter.FArrivalTime.Format("2006-01-02 15:04:05"),
				LetterBody:  letter.FLetterBody,
			}
			letterList = append(letterList, l)
			if l.LetterId > nextPageStart {
				nextPageStart = l.LetterId
			}
		}

		// 将 letters 赋值给 messages
		messages = mooonmailbox.ListMessagesResp{
			Recipient:     letters[0].FRecipient, // 假设收件人都是相同的
			Letters:       letterList,
			NextPageStart: nextPageStart,
		}
		return &messages, nil
	}
}
