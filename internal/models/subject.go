package models

import "context"

type Subject struct {
	ID int64 `json:"id"`
	Name string `json:"name"`
	TeacherID int64 `json:"teacher_id"`
}

type CreateSubjectInput struct {
	Name string `json:"name"`
	TeacherID int64 `json:"teacher_id"`
}

type SubjectRepository interface {
	FindByID(ctx context.Context, id int64) (*Subject, error)
}