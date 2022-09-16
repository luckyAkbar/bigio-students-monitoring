package models

import "context"

type Mark = string
var (
	MarkA Mark = "A"
	MarkB Mark = "B"
	MarkC Mark = "C"
	MarkD Mark = "D"
	MarkE Mark = "E"
)

type Grade struct {
	ID int64 `json:"id"`
	StudentID int64 `json:"student_id"`
	TeacherID int64 `json:"teacher_id"`
	SubjectID int64 `json:"subject_id"`
	Mark Mark `json:"mark"`
	Value int `json:"value"`
}

type CreateGradeInput struct {
	StudentID int64 `json:"student_id"`
	TeacherID int64 `json:"teacher_id"`
	SubjectID int64 `json:"subject_id"`
	Mark Mark `json:"mark"`
	Value int `json:"value"`
}

type GradeRepository interface {
	FindByID(ctx context.Context, id int64) (*Grade, error)
}