package sections

import (
	"context"

	"github.com/winecraft-dev/ed-tech/sections/models"
)

type Repository interface {
	GetSections(context.Context) ([]models.Section, error)
	GetPopulatedSection(context.Context, int) (*models.PopulatedSection, error)
	CreateSection(context.Context, models.CreateSection) (*int, error)
}
