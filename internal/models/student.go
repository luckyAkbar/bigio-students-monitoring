package models

import "context"

type Student struct {
	ID int64 `json:"id"`
	Name string `json:"name"`
}

type CreateStudentInput struct {
	Name string `json:"name"`
}

type StudentUsecase interface {
	GetGradeBySubjectID(ctxt context.Context, subjectID int64) (int, error)
	FindByID(ctxt context.Context, id int64) (*Student, error)
}

type StudentRepository interface {
	GetGradeBySubjectID(ctxt context.Context, subjectID int64) (int, error)
	FindByID(ctxt context.Context, id int64) (*Student, error)
}