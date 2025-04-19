package database

import (
	"context"
	"fmt"

	"github.com/winecraft-dev/ed-tech/sections/models"
)

func (c *Connection) GetSections(ctx context.Context) ([]models.Section, error) {
	query := `SELECT * FROM sections`

	var sections []models.Section
	err := c.db.SelectContext(ctx, &sections, query)
	if err != nil {
		return nil, fmt.Errorf("%w: %w", ErrRetrieving, err)
	}

	return sections, nil
}

func (c *Connection) CreateSection(ctx context.Context, i models.CreateSection) (*int, error) {
	query := `
		INSERT INTO sections (
			school_code, section, homeroom, homeroom_teacher
		) VALUES (
			:school_code, :section, :homeroom, :homeroom_teacher
		) RETURNING id`

	rows, err := c.db.NamedQueryContext(ctx, query, i)
	if err != nil {
		return nil, fmt.Errorf("%w: %w", ErrCreate, err)
	}
	defer rows.Close()

	var id int
	if rows.Next() {
		err = rows.Scan(&id)
	}
	return &id, err
}
