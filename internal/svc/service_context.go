package svc

import (
	"mooon-mailbox/internal/config"
	"mooon-mailbox/model"
)

type ServiceContext struct {
	Config       config.Config
	MailboxModel model.TMooonMailboxModel
}

func NewServiceContext(c config.Config) *ServiceContext {
	return &ServiceContext{
		Config: c,
	}
}
