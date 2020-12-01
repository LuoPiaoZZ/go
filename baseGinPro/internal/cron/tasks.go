package cron

import (
	"baseGinPro/internal/cron/tasks"
	"baseGinPro/internal/cronlib"
)

var (

	Cron = cronlib.New()

	SpecList = map[string]string{
		"EveryDayPrintDate":   "10 34 16 1/1 * ?", //每天凌晨0点0分10秒
	}

	Tasks = map[string]func(){
		"EveryDayPrintDate":   func() { tasks.PrintDate() }, //打印每天日期
	}
)
