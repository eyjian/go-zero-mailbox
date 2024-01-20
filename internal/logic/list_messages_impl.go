// Package logic
// Written by yijian on 2024/01/20
package logic

import (
	"fmt"
	"github.com/zeromicro/go-zero/core/logc"
	"mooon-mailbox/model"
	"mooon-mailbox/pb/mooon_mailbox"
	"strconv"
)

func listMessages(l *ListMessagesLogic, in *mooon_mailbox.ListMessagesReq) (*mooon_mailbox.ListMessagesResp, error) {
	var letters []model.TMooonMailbox
	var nextPageStart string

	// StartLetterId
	startLetterId := in.StartLetterId
	if len(startLetterId) == 0 || startLetterId == "0" {
		startLetterId = "1000000000000"
	}
	_, err := strconv.ParseInt(startLetterId, 10, 64)
	if err != nil {
		logc.Errorf(l.ctx, "Invalid parameter[start_letter_id]: (%s)%s", startLetterId, err.Error())
		return nil, err
	}

	// PageSize
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
			"%s"+
			"f_id<%s "+
			"ORDER BY f_id DESC "+
			"LIMIT %d",
		in.Recipient,
		getState(in.ListAction),
		startLetterId,
		in.PageSize)
	err = l.svcCtx.CachedConn.QueryRowsNoCacheCtx(l.ctx, &letters, sql)
	if err != nil {
		logc.Errorf(l.ctx, "Exec %s error: %s", sql, err.Error())
		return nil, err
	} else {
		var messages mooon_mailbox.ListMessagesResp
		logc.Infof(l.ctx, "Exec %s success\n", sql)

		// 将 letters 转换为 Letter 类型的切片
		var letterList []*mooon_mailbox.Letter
		for _, letter := range letters {
			l := &mooon_mailbox.Letter{
				LetterId:    strconv.FormatInt(letter.FId, 10),
				DeliverTime: letter.FDeliverTime.Format("2006-01-02 15:04:05"),
				ArrivalTime: letter.FArrivalTime.Format("2006-01-02 15:04:05"),
				LetterBody:  letter.FLetterBody,
				LetterState: mooon_mailbox.Letter_LetterState(letter.FState),
			}
			letterList = append(letterList, l)
			nextPageStart = strconv.FormatInt(letter.FId, 10)
		}

		// 将 letters 赋值给 messages
		messages = mooon_mailbox.ListMessagesResp{
			Recipient:     in.Recipient,
			Letters:       letterList,
			NextPageStart: nextPageStart,
		}
		return &messages, nil
	}
}

func getState(listAction mooon_mailbox.ListMessagesReq_ListAction) string {
	if listAction == mooon_mailbox.ListMessagesReq_LA_ALL {
		return ""
	} else if listAction == mooon_mailbox.ListMessagesReq_LA_ONLY_UNREAD {
		return fmt.Sprintf("f_state=%d AND ", int(mooon_mailbox.Letter_LETTER_UNREAD))
	} else {
		return fmt.Sprintf("f_state=%d AND ", int(mooon_mailbox.Letter_LETTER_READ))
	}
}
