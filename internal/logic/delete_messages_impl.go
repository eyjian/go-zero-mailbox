// Package logic
// Written by yijian on 2024/01/20
package logic

import (
	"fmt"
	"github.com/zeromicro/go-zero/core/logc"
	"mooon-mailbox/mooonmailbox"
	"mooon-mailbox/pb/mooon-mailbox"
)

func deleteMessages(l *DeleteMessagesLogic, in *mooon_mailbox.DeleteMessagesReq) (*mooon_mailbox.DeleteMessagesResp, error) {
	sql := getDeleteSql(in)
	dbResult, err := l.svcCtx.CachedConn.ExecNoCacheCtx(l.ctx, sql)
	if err != nil {
		logc.Errorf(l.ctx, "Exec %s error: %s\n", sql, err.Error())
		return nil, err
	} else {
		var rowsAffected int64 = 0
		rowsAffected, _ = dbResult.RowsAffected()
		logc.Infof(l.ctx, "Exec %s success: RowsAffected=%d\n", sql, rowsAffected)
		return &mooonmailbox.DeleteMessagesResp{
			Recipient: in.Recipient,
		}, nil
	}
}

func getDeleteSql(in *mooonmailbox.DeleteMessagesReq) string {
	inStr := getInStr(in.LettersIdList)
	return fmt.Sprintf(
		"DELETE FROM t_mooon_mailbox WHERE f_recipient='%s' AND f_id IN (%s)",
		in.Recipient,
		inStr)

}
