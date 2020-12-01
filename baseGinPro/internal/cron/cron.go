// 定时任务
package cron

import (
	"baseGinPro/internal/cronlib"
)

// 初始化定时任务
func InitCron() {

	for srv, spec := range SpecList {
		// create job
		job, err := cronlib.NewJobModel(
			spec,
			Tasks[srv],
		)
		if err != nil {
			panic(err.Error())
		}

		// register srvName -> job
		err = Cron.Register(srv, job)
		if err != nil {
			panic(err.Error())
		}
	}

	Cron.Start()
}
