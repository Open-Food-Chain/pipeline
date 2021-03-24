package cron_trigger

import (
	"github.com/robfig/cron/v3"
	"github.com/The-New-Fork/pipeline/pkg/domain"
)

type Trigger struct {
	config         *Config
	stub           domain.Stub
	cron           *cron.Cron
	RequestChannel chan *domain.Request
}

func NewTrigger() *Trigger {
	return &Trigger{}
}
