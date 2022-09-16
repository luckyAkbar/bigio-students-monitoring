package models

import (
	"context"
	"errors"
	"strings"
)

type Mark = string

var (
	MarkA Mark = "A"
	MarkB Mark = "B"
	MarkC Mark = "C"
	MarkD Mark = "D"
	MarkE Mark = "E"
)

func ConvertStringToMark(s string) (Mark, error) {
	mark := strings.ToUpper(s)
	switch mark {
	default:
		return MarkE, errors.New("string value is not convertible to mark")
	case "A":
		return MarkA, nil
	case "B":
		return MarkB, nil
	case "C":
		return MarkC, nil
	case "D":
		return MarkD, nil
	case "E":
		return MarkD, nil
	}
}

type Grade struct {
	ID        int64 `json:"id"`
	StudentID int64 `json:"student_id"`
	TeacherID int64 `json:"teacher_id"`
	SubjectID int64 `json:"subject_id"`
	Mark      Mark  `json:"mark"`
	Value     int   `json:"value"`
}

type CreateGradeInput struct {
	StudentID int64 `json:"student_id" validate:"required"`
	SubjectID int64 `json:"subject_id" validate:"required"`
	Mark      Mark  `json:"mark" validate:"required"`
	Value     int   `json:"value" validate:"required"`
}

func (i *CreateGradeInput) Validate() error {
	return validator.Struct(i)
}

type GradeRepository interface {
	FindByID(ctx context.Context, id int64) (*Grade, error)
	Create(ctx context.Context, grade *Grade) error
}
