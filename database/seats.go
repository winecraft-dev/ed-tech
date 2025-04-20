package database

import (
	"context"

	"github.com/winecraft-dev/ed-tech/classroom/models"
)

func (c *Connection) AssignSeat(
	ctx context.Context,
	class int,
	input models.AssignSeat,
) (*models.Seat, error) {
	return nil, nil
}

func (c *Connection) UnassignSeat(
	ctx context.Context,
	class int,
	student int,
) error {
	return nil
}

func (c *Connection) GetSeats(
	ctx context.Context,
	class int,
	student int,
) ([]models.Seat, error) {
	return []models.Seat{}, nil
}
