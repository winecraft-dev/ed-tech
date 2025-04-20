package sections

import (
	"net/http"
	"strconv"

	"github.com/labstack/echo/v4"
	"github.com/winecraft-dev/ed-tech/sections/models"
	"github.com/winecraft-dev/ed-tech/service"
)

type Service struct {
	Repository
}

func NewService(db Repository) *Service {
	return &Service{
		Repository: db,
	}
}

func (s *Service) CreateSection(c echo.Context) error {
	ctx := c.Request().Context()

	input := new(models.CreateSection)
	if err := c.Bind(input); err != nil {
		return service.ErrMalformedInput.WithInternal(err)
	}

	sectionID, err := s.Repository.CreateSection(ctx, *input)
	if err != nil {
		return service.ErrDatabaseError.WithInternal(err)
	}

	return c.JSON(http.StatusCreated, models.CreatedSection{
		ID: *sectionID,
	})
}

func (s *Service) GetSections(c echo.Context) error {
	ctx := c.Request().Context()

	sections, err := s.Repository.GetSections(ctx)
	if err != nil {
		return service.ErrDatabaseError.WithInternal(err)
	}

	return c.JSON(http.StatusOK, sections)
}

func (s *Service) GetSection(c echo.Context) error {
	ctx := c.Request().Context()

	sectionID, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		return service.ErrMalformedInput.WithInternal(err)
	}

	section, err := s.Repository.GetPopulatedSection(ctx, sectionID)
	if err != nil {
		return service.ErrDatabaseError.WithInternal(err)
	}
	if section == nil {
		return c.JSON(http.StatusNotFound, section)
	}
	return c.JSON(http.StatusOK, section)
}
