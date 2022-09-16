package models

import "context"

type Teacher struct {
	ID int64 `json:"id"`
	Name string `json:"name"`
}

type CreateTeacherInput struct {
	Name string `json:"name"`
}

type TeacherUsecase interface {
	FindByID(ctx context.Context, id int64) (*Teacher, error)
	GradeByStudentID(ctx context.Context, input *CreateGradeInput) error
}

type TeacherRepository interface {
	FindByID(ctx context.Context, id int64) (*Teacher, error)
	GradeByStudentID(ctx context.Context, grade *Grade) error
}