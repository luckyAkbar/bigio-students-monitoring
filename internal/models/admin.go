package models

import "context"

type Admin struct {
	ID   int64  `json:"id"`
	Name string `json:"name"`
}

type AdminUsecase interface {
	FindByID(ctx context.Context, id int64) (*Admin, error)
	CreateTeacher(ctx context.Context, input *CreateTeacherInput) (*Teacher, error)
	CreateStudent(ctx context.Context, input *CreateStudentInput) (*Student, error)
	CreateSubject(ctx context.Context, input *CreateSubjectInput) (*Subject, error)
}

type AdminRepository interface {
	FindByID(ctx context.Context, id int64) (*Admin, error)
	Create(ctx context.Context, admin *Admin) error
	CreateTeacher(ctx context.Context, teacher *Teacher, user *User) error
	CreateStudent(ctx context.Context, student *Student, user *User) error
	CreateSubject(ctx context.Context, subject *Subject) error
}
