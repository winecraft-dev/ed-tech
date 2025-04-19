package students

import (
	"context"

	"github.com/winecraft-dev/ed-tech/students/models"
)

type Repository interface {
	CreateStudent(context.Context, models.Student) (*int, error)
}
