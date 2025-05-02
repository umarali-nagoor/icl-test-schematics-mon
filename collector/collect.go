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
