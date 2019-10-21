package app

import (
	"github.com/growerlab/backend/app/model/db"
	"github.com/growerlab/backend/app/service/graphql/think"
	"github.com/growerlab/backend/app/utils/conf"
)

// 需要初始化的全局数据放在这里
//	eg. onStart(job.Work)
//
func init() {
	onStart(conf.LoadConfig)
	onStart(db.InitDatabase)
	onStart(think.InitGraphQL)
}

func onStart(fn func() error) {
	if err := fn(); err != nil {
		panic(err)
	}
}
