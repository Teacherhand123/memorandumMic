package dao

import (
	"context"
	"micro-memorandum/app/user/reposltory/db/model"

	"gorm.io/gorm"
)

// 定义对数据库user表的curd操作
type UserDao struct {
	*gorm.DB
}

// 每个请求、每个操作可以有独立的超时时间、取消信号、trace id 等信息。
// 适合并发场景，比如 Web 服务每个 HTTP 请求都用自己的 ctx，互不影响。
// 可以精细控制每个操作的生命周期，便于资源回收和错误追踪。
// 使用同一个 ctx

// 所有操作共享同一个上下文，适合生命周期完全一致的批量操作。
// 如果某个操作被取消，所有用这个 ctx 的操作都会被取消。
// 灵活性较差，不适合高并发和独立请求场景。
func NewUserDao(ctx context.Context) *UserDao {
	if ctx == nil {
		ctx = context.Background()
	}
	return &UserDao{NewDBClinet(ctx)}
}

func (dao *UserDao) FindUserByUserName(userName string) (r *model.User, err error) {
	err = dao.Model(&model.User{}).
		Where("user_name = ?", userName).First(&r).Error
	return r, err
}

func (dao *UserDao) CreateUser(in *model.User) (err error) {
	return dao.Model(&model.User{}).Create(in).Error
}
