package models

type Student struct {
	Number int     `json:"number" db:"number"`
	Email  *string `json:"email" db:"email"`

	SectionID *int `json:"section_id" db:"section_id"`

	PreferredName *string `json:"preferred_name" db:"preferred_name"`
	GivenName     string  `json:"given_name" db:"given_name"`
	LastName      string  `json:"last_name" db:"last_name"`

	ParentEmail  *string `json:"parent_email" db:"parent_email"`
	ParentNumber *string `json:"parent_number" db:"parent_number"`
}

type CreatedStudent struct {
	Number int
}
