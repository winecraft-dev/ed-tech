package students

import (
	"net/http"

	"github.com/labstack/echo/v4"
	"github.com/winecraft-dev/ed-tech/service"
	"github.com/winecraft-dev/ed-tech/students/models"
)

type Service struct {
	Repository
}

func NewService(db Repository) *Service {
	return &Service{
		Repository: db,
	}
}

func (s *Service) CreateStudent(c echo.Context) error {
	ctx := c.Request().Context()

	input := new(models.Student)
	if err := c.Bind(input); err != nil {
		return service.ErrMalformedInput.WithInternal(err)
	}

	studentNum, err := s.Repository.CreateStudent(ctx, *input)
	if err != nil {
		return service.ErrDatabaseError.WithInternal(err)
	}

	return c.JSON(http.StatusCreated, models.CreatedStudent{
		Number: *studentNum,
	})
}

func (s *Service) GetStudent(c echo.Context) error {
	return nil
}
