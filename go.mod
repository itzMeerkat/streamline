module gitee.com/fat_marmota/streamline

go 1.15

require (
	gitee.com/fat_marmota/infra/log v0.0.0-20210224080616-e18c97687651
	go.uber.org/zap v1.16.0
)

replace gitee.com/fat_marmota/infra/log v0.0.0 => ../fat_marmota/infra/log
