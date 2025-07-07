package dao

import "micro-memorandum/app/task/reposltory/db/model"

func migration() {
	_db.Set(`gorm:table_options`, "charset=utf8mb4").
		AutoMigrate(&model.Task{})
}
