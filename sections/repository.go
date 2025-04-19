package sections

import (
	"context"

	"github.com/winecraft-dev/ed-tech/sections/models"
)

type Repository interface {
	GetSections(context.Context) ([]models.Section, error)
	GetSection(context.Context, int) (*models.PopulatedSection, error)
	CreateSection(context.Context, models.CreateSection) (*int, error)
}
