package models

import "context"

type Subject struct {
	ID        int64  `json:"id"`
	Name      string `json:"name"`
	TeacherID int64  `json:"teacher_id"`
}

type CreateSubjectInput struct {
	Name      string `json:"name" validate:"required"`
	TeacherID int64  `json:"teacher_id" validate:"required"`
}

func (i *CreateSubjectInput) Validate() error {
	return validator.Struct(i)
}

type SubjectRepository interface {
	FindByID(ctx context.Context, id int64) (*Subject, error)
	Create(ctx context.Context, subject *Subject) error
}
