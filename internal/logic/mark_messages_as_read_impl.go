// Package logic
// Written by yijian on 2024/01/20
package logic

import (
	"fmt"
	"github.com/zeromicro/go-zero/core/logc"
	mooonmailbox "mooon-mailbox/pb/mooon-mailbox"
)

func markMessagesAsRead(l *MarkMessagesAsReadLogic, in *mooonmailbox.MarkMessagesAsReadReq) (*mooonmailbox.MarkMessagesAsReadResp, error) {
	sql := getSql(in)
	dbResult, err := l.svcCtx.CachedConn.ExecNoCacheCtx(l.ctx, sql)
	if err != nil {
		logc.Errorf(l.ctx, "Exec %s error: %s\n", sql, err.Error())
		return nil, err
	} else {
		var rowsAffected int64 = 0
		rowsAffected, _ = dbResult.RowsAffected()
		logc.Infof(l.ctx, "Exec %s success: RowsAffected=%d\n", sql, rowsAffected)
		return &mooonmailbox.MarkMessagesAsReadResp{
			Recipient: in.Recipient,
		}, nil
	}
}

func getSql(in *mooonmailbox.MarkMessagesAsReadReq) string {
	inStr := getInStr(in.LettersIdList)
	return fmt.Sprintf(
		"UPDATE t_mooon_mailbox SET f_state=%d WHERE f_recipient='%s' AND f_id IN (%s)",
		mooonmailbox.ListMessagesReq_LA_ONLY_READ,
		in.Recipient,
		inStr)

}
