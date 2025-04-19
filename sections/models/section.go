package models

import "github.com/winecraft-dev/ed-tech/students/models"

type Section struct {
	ID int `json:"id" db:"id"`

	SchoolCode string `json:"school_code" db:"school_code"`
	Section    int    `json:"section" db:"section"`
	Homeroom   int    `json:"homeroom" db:"homeroom"`

	HomeroomTeacher *string `json:"homeroom_teacher" db:"homeroom_teacher"`
}

type PopulatedSection struct {
	Section
	Students []models.Student `json:"students"`
}

type CreateSection struct {
	SchoolCode string `json:"school_code" db:"school_code"`
	Section    int    `json:"section" db:"section"`
	Homeroom   int    `json:"homeroom" db:"homeroom"`

	HomeroomTeacher *string `json:"homeroom_teacher" db:"homeroom_teacher"`
}

type CreatedSection struct {
	ID int `json:"id" db:"id"`
}

type JoinedSectionStudent struct {
	Section

	Number        *int    `db:"number"`
	Email         *string `db:"student_email"`
	PreferredName *string `db:"preferred_name"`
	GivenName     *string `db:"given_name"`
	LastName      *string `db:"last_name"`
	ParentEmail   *string `db:"parent_email"`
	ParentNumber  *string `db:"parent_number"`
}

func (j JoinedSectionStudent) ExtractStudent() *models.Student {
	if j.Number == nil {
		return nil
	}
	return &models.Student{
		Number:        *j.Number,
		Email:         j.Email,
		SectionID:     &j.Section.ID,
		PreferredName: j.PreferredName,
		GivenName:     *j.GivenName,
		LastName:      *j.LastName,
		ParentEmail:   j.ParentEmail,
		ParentNumber:  j.ParentNumber,
	}
}
