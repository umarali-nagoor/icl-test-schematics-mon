package collector

import (
	"time"

	"go.uber.org/zap"
)

func (c *Collector) run() {
	ticker := time.NewTicker(30 * time.Second)
	for range ticker.C {
		if err := c.scrape(); err != nil {
			c.logger.Warn("scrape failed", zap.Error(err))
		}
	}
}

func (c *Collector) collectAll(targets []*Target) {
	for {
		for _, t := range targets {
			m, err := c.scrape(t)
			if err != nil {
				c.log.Errorf("scrape failed for target %+v: %v", t, err)
				continue
			}
			c.log.Infof("collected metric: %+v", m)
		}
		time.Sleep(10 * time.Second)
	}
}
