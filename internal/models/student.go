package models

import "context"

type Student struct {
	ID   int64  `json:"id"`
	Name string `json:"name"`
}

type CreateStudentInput struct {
	Name     string `json:"name" validate:"required"`
	Password string `json:"password" validate:"required"`
}

func (i *CreateStudentInput) Validate() error {
	return validator.Struct(i)
}

type StudentUsecase interface {
	GetGradeBySubjectID(ctx context.Context, subjectID int64) (*Grade, error)
	GetAllGrade(ctx context.Context) ([]Grade, error)
}

type StudentRepository interface {
	FindByStudentAndSubjectID(ctx context.Context, studentID, subjectID int64) (*Grade, error)
	GetAllGrade(ctx context.Context, id int64) ([]Grade, error)
	FindByID(ctxt context.Context, id int64) (*Student, error)
}
