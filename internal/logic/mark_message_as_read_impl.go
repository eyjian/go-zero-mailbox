// Package logic
// Written by yijian on 2024/01/20
package logic

import (
	"fmt"
	"github.com/zeromicro/go-zero/core/logc"
	mooonmailbox "mooon-mailbox/pb/mooon-mailbox"
	"strconv"
)

func markMessageAsRead(l *MarkMessageAsReadLogic, in *mooonmailbox.MarkAsReadReq) (*mooonmailbox.MarkAsReadResp, error) {
	// todo: add your logic here and delete this line
	sql := getSql(in)
	dbResult, err := l.svcCtx.CachedConn.ExecNoCacheCtx(l.ctx, sql)
	if err != nil {
		logc.Errorf(l.ctx, "Insert %s to table error: %s\n", in.String(), err.Error())
		return nil, err
	} else {
		var rowsAffected int64 = 0
		rowsAffected, _ = dbResult.RowsAffected()
		logc.Infof(l.ctx, "Exec %s success: RowsAffected=%d\n", sql, rowsAffected)
		return &mooonmailbox.MarkAsReadResp{
			Recipient: in.Recipient,
		}, nil
	}
}

func getSql(in *mooonmailbox.MarkAsReadReq) string {
	inStr := ""
	for index, letterId := range in.LettersIdList {
		if index == 0 {
			inStr = strconv.FormatInt(letterId, 10)
		} else {
			inStr = inStr + "," + strconv.FormatInt(letterId, 10)
		}
	}
	return fmt.Sprintf(
		"UPDATE t_mooon_mailbox SET f_state=%d WHERE f_recipient='%s' AND f_id IN (%s)",
		mooonmailbox.ListMessagesReq_LA_ONLY_READ,
		in.Recipient,
		inStr)

}
