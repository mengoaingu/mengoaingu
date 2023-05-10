package repository

import (
	"backend/internal/quizzes/domain"
	"context"

	"gorm.io/gorm"
)

var _ QuizReportRepository = (*quizReportRepository)(nil)

type quizReportRepository struct {
	db *gorm.DB
}

func NewQuizReportRepository(db QuizzesDB) QuizReportRepository {
	return &quizReportRepository{
		db: db,
	}
}

// FindByID implements QuizReportRepository
func (*quizReportRepository) FindByID(ctx context.Context, id int64) (*domain.QuizReport, error) {
	panic("unimplemented")
}

// Save implements QuizReportRepository
func (*quizReportRepository) Save(ctx context.Context, quizReport *domain.QuizReport) error {
	panic("unimplemented")
}

// Update implements QuizReportRepository
func (*quizReportRepository) Update(ctx context.Context, id int64, quizReport *domain.QuizReport) error {
	panic("unimplemented")
}

// // Delete implements TaskRepository
// func (r *taskRepository) Delete(ctx context.Context, task_id int64, user_id string) error {
// 	err := r.db.WithContext(ctx).Where("id = ?", task_id).Where("user_id = ?", user_id).Delete(&domain.Task{}).Error
// 	return err
// }

// // Find implements TaskRepository
// func (r *taskRepository) Find(ctx context.Context, name string, date int64, user_id string) ([]*domain.Task, int64, error) {
// 	var tasks []*domain.Task
// 	var count int64
// 	query := r.db.WithContext(ctx).Model(&tasks).Where("user_id = ?", user_id)
// 	if name != "" {
// 		query.Where("name LIKE ?", "%"+name+"%")
// 	}
// 	if date != 0 {
// 		query.Where("created_at >= ?", time.UnixMilli(date))
// 		query.Where("created_at < ?", time.UnixMilli(date).Add(time.Duration(
// 			time.Hour*24,
// 		)))
// 	}
// 	err := query.Count(&count).Find(&tasks).Error
// 	return tasks, count, err
// }

// // FindByID implements TaskRepository
// func (r *taskRepository) FindByID(ctx context.Context, task_id int64, user_id string) (*domain.Task, error) {
// 	var task domain.Task
// 	err := r.db.WithContext(ctx).Where("id = ?", task_id).Where("user_id = ?", user_id).Find(&task).Error
// 	return &task, err
// }

// // Save implements TaskRepository
// func (r *taskRepository) Save(ctx context.Context, task *domain.Task) error {
// 	err := r.db.WithContext(ctx).Create(&task).Error
// 	return err
// }

// // Update implements TaskRepository
// func (r *taskRepository) Update(ctx context.Context, task_id int64, user_id string, task *domain.Task) (*domain.Task, error) {
// 	err := r.db.WithContext(ctx).Where("id = ?", task_id).Where("user_id = ?", user_id).Updates(&task).Error
// 	return task, err
// }
