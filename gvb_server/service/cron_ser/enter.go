package cron_ser

import (
	"time"

	"github.com/robfig/cron/v3"
	"github.com/sirupsen/logrus"
)

func CronInit() {
	timezone, err := time.LoadLocation("Asia/Shanghai")
	if err != nil {
		logrus.Error(err.Error())
		return
	}
	Cron := cron.New(cron.WithSeconds(), cron.WithLocation(timezone))
	Cron.AddFunc("*/10 * * * * *", SyncArticleDate)
	Cron.AddFunc("*/10 * * * * *", SyncCommentData)
	Cron.Start()
}
