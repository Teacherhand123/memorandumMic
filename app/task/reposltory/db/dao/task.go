package dao

import (
	"context"
	"micro-memorandum/app/task/reposltory/db/model"
	"micro-memorandum/idl/pb"

	"gorm.io/gorm"
)

// 定义对数据库user表的curd操作
type TaskDao struct {
	*gorm.DB
}

// 每个请求、每个操作可以有独立的超时时间、取消信号、trace id 等信息。
// 适合并发场景，比如 Web 服务每个 HTTP 请求都用自己的 ctx，互不影响。
// 可以精细控制每个操作的生命周期，便于资源回收和错误追踪。
// 使用同一个 ctx

// 所有操作共享同一个上下文，适合生命周期完全一致的批量操作。
// 如果某个操作被取消，所有用这个 ctx 的操作都会被取消。
// 灵活性较差，不适合高并发和独立请求场景。
func NewTaskDao(ctx context.Context) *TaskDao {
	if ctx == nil {
		ctx = context.Background()
	}
	return &TaskDao{NewDBClinet(ctx)}
}

func (dao *TaskDao) CreateTask(data *model.Task) error {
	return dao.Model(&model.Task{}).Create(data).Error
}

func (dao *TaskDao) LimitTaskByUserId(userId uint64, start, limit int) (r []*model.Task, count int64, err error) {
	err = dao.Model(&model.Task{}).Offset(start).Limit(limit).
		Where("uid = ?", userId).Find(&r).Error
	if err != nil {
		return
	}

	err = dao.Model(&model.Task{}).Where("uid = ?", userId).Count(&count).Error
	return
}

func (dao *TaskDao) GetTaskByTaskIdAndUserId(id, userId uint64) (r *model.Task, err error) {
	err = dao.Model(&model.Task{}).Where("id = ? AND uid = ?", id, userId).First(&r).Error
	return
}

func (dao *TaskDao) UpdateTask(req *pb.TaskRequest) (r *model.Task, err error) {
	err = dao.Model(&model.Task{}).Where("id = ? AND uid = ?", req.Id, req.Uid).First(&r).Error
	if err != nil {
		return
	}

	r.Title = req.Title
	r.Status = int(req.Status)
	r.Content = req.Content

	return r, dao.Save(&r).Error
}

func (dao *TaskDao) DeleteTask(id, userId uint64) (err error) {
	err = dao.Model(&model.Task{}).Where("id = ? AND uid = ?", id, userId).Delete(&model.Task{}).Error
	return
}
