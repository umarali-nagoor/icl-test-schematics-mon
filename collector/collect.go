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

func (c *Collector) logAlert(alert *Alert) {
	b, err := json.Marshal(alert)
	if err != nil {
		c.log.Errorf("failed to marshal alert: %v", err)
		return
	}
	c.log.Infof("alert fired: %s", string(b))
}
