package svc

import (
	"github.com/zeromicro/go-zero/core/stores/sqlc"
	"mooon-mailbox/internal/config"
	"mooon-mailbox/model"
)

type ServiceContext struct {
	Config       config.Config
	MailboxModel model.TMooonMailboxModel
	CachedConn   sqlc.CachedConn
}

func NewServiceContext(c config.Config) *ServiceContext {
	return &ServiceContext{
		Config: c,
	}
}
