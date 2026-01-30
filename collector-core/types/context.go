package types

import (
	"sync"

	cron "github.com/robfig/cron/v3"
)

type Context struct {
	CronHandle  *cron.Cron
	Params      *Conf
	Config      *Config
	Result      map[string]string
	ResultMutex sync.RWMutex
}

func (ctc *Context) GetParams() *Conf {
	return ctc.Params
}
