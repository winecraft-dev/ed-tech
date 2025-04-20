package models

type Seat struct {
	Classroom     int `db:"classroom_id"`
	StudentNumber int `db:"student_number"`

	Table int `db:"table"`
}

type AssignSeat struct {
	StudentNumber int  `json:"student_number"`
	Table         *int `json:"table"`
}
