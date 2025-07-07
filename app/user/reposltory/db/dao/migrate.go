package dao

import "micro-memorandum/app/user/reposltory/db/model"

func migration() {
	_db.Set(`gorm:table_options`, "charset=utf8mb4").
		AutoMigrate(&model.User{})
}
