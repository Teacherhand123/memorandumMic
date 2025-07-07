package script

import (
	"context"
	"micro-memorandum/app/task/reposltory/mq/task"
)

func TaskCreateSync(ctx context.Context) {
	tSync := new(task.SyncTask)
	err := tSync.RunTaskService(ctx)
	if err != nil {
		return
	}
}
