package cron

import (
	"github.com/robfig/cron"
	"webesapp/dao/mysql"
)

func Cron() {
	c := cron.New()
	c.AddFunc("* * * * *", func() { mysql.SqlV6Qurey() })
	c.Start()
	select {}
}
