package classroom

import (
	"net/http"
	"strconv"

	"github.com/labstack/echo/v4"
	"github.com/winecraft-dev/ed-tech/classroom/models"
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

func (s *Service) AssignSeat(c echo.Context) error {
	ctx := c.Request().Context()

	cid, err := strconv.Atoi(c.Param("classroom"))
	if err != nil {
		return service.ErrMalformedInput.WithInternal(err)
	}

	input := new(models.AssignSeat)
	if err := c.Bind(input); err != nil {
		return service.ErrMalformedInput.WithInternal(err)
	}

	seat, err := s.Repository.AssignSeat(ctx, cid, *input)
	if err != nil {
		return service.ErrDatabaseError.WithInternal(err)
	}

	return c.JSON(http.StatusCreated, seat)
}

func (s *Service) UnassignSeat(c echo.Context) error {
	ctx := c.Request().Context()

	cid, err := strconv.Atoi(c.Param("classroom"))
	if err != nil {
		return service.ErrMalformedInput.WithInternal(err)
	}
	snum, err := strconv.Atoi(c.Param("student"))
	if err != nil {
		return service.ErrMalformedInput.WithInternal(err)
	}

	if err := s.Repository.UnassignSeat(ctx, cid, snum); err != nil {
		return service.ErrDatabaseError.WithInternal(err)
	}

	return c.JSON(http.StatusOK, nil)
}

func (s *Service) GetClassroom(c echo.Context) error {
	_ = c.Request().Context()

	_ = c.Param("section")

	return nil
}
