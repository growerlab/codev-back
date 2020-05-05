package app

import (
	"github.com/growerlab/backend/app/common/mq"
	"github.com/growerlab/backend/app/common/notify"
	"github.com/growerlab/backend/app/common/permission"
	"github.com/growerlab/backend/app/model/db"
	"github.com/growerlab/backend/app/utils/conf"
)

// 需要初始化的全局数据放在这里
//	eg. onStart(job.Work)
//
func init() {
	onStart(conf.LoadConfig)
	onStart(db.InitMemDB)
	onStart(db.InitDatabase)
	onStart(notify.InitNotify)
	// onStart(events.InitEvents)
	onStart(permission.InitPermission)
	onStart(mq.InitMQ)
}

func onStart(fn func() error) {
	if err := fn(); err != nil {
		panic(err)
	}
}
