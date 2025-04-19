package database

import (
	"context"
	"fmt"

	"github.com/winecraft-dev/ed-tech/sections/models"
)

func (c *Connection) GetSection(ctx context.Context, id int) (*models.PopulatedSection, error) {
	query := `
			SELECT
				id, school_code, section, homeroom, homeroom_teacher,

				number, s.email as student_email, preferred_name,
				given_name, last_name, parent_email, parent_number
	 		FROM sections c
			LEFT JOIN students s
			ON s.section_id=c.id
			WHERE c.id=$1
		`

	var results []models.JoinedSectionStudent
	err := c.db.SelectContext(ctx, &results, query, id)
	if err != nil {
		return nil, fmt.Errorf("%w: %w", ErrRetrieving, err)
	}

	if len(results) <= 0 {
		return nil, nil
	}

	var out models.PopulatedSection
	out.Section = results[0].Section
	for _, r := range results {
		student := r.ExtractStudent()
		if student != nil {
			out.Students = append(out.Students, *student)
		}
	}

	return &out, nil
}
