package agent

import (
	"context"
	"github.com/tovenja/cron/v3"
	"hippo-data-acquisition/commons/logger"
)

var (
	agentCron *cron.Cron
	entryID   cron.EntryID
)

type Cron struct {
	spec string
}

func (c *Cron) Start(handle func()) error {
	agentCron = cron.New(cron.WithSeconds())
	var err error = nil
	entryID, err = agentCron.AddFunc(c.spec, handle)
	if err != nil {
		logger.LogError("agent", "启动定时器失败！")
		return err
	}

	agentCron.Start()
	return nil
}

func (c *Cron) Stop() context.Context {
	if agentCron != nil {
		return agentCron.Stop()
	}

	return nil
}

func (c *Cron) GetAgentCron() *cron.Cron {
	return agentCron
}
