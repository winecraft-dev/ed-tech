package database

import (
	"context"
	"fmt"

	"github.com/winecraft-dev/ed-tech/students/models"
)

func (c *Connection) CreateStudent(ctx context.Context, i models.Student) (*int, error) {
	query := `
		INSERT INTO students (
			number, email, section_id, preferred_name, given_name,
			last_name, parent_email, parent_number
		) VALUES (
			:number, :email, :section_id, :preferred_name, :given_name,
			:last_name, :parent_email, :parent_number
		) RETURNING number`

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
