package classroom

import (
	"context"

	"github.com/winecraft-dev/ed-tech/classroom/models"
)

type Repository interface {
	GetSeats(ctx context.Context, c, s int) ([]models.Seat, error)
	AssignSeat(context.Context, int, models.AssignSeat) (*models.Seat, error)
	UnassignSeat(ctx context.Context, c, s int) error
}
