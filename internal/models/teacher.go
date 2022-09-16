package models

import "context"

type Teacher struct {
	ID   int64  `json:"id"`
	Name string `json:"name"`
}

type CreateTeacherInput struct {
	Name     string `json:"name" validate:"required"`
	Password string `json:"password" validate:"required"`
}

func (i *CreateTeacherInput) Validate() error {
	return validator.Struct(i)
}

type TeacherUsecase interface {
	GradeByStudentID(ctx context.Context, input *CreateGradeInput) (*Grade, error)
}

type TeacherRepository interface {
	FindByID(ctx context.Context, id int64) (*Teacher, error)
}
