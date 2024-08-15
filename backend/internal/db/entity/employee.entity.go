package entity

import (
	"time"

	"github.com/google/uuid"
)

type Employee struct {
	EmployeeId uuid.UUID `json:"employee_id" db:"employee_id"`
	Name       string    `json:"name" db:"name"`
	JobTitle   string    `json:"job_title" db:"job_title"`
	Salary     float32   `json:"salary" db:"salary"`
	Department string    `json:"department" db:"department"`
	JoinedDate time.Time `json:"joined_date" db:"joined_date"`
}

func NewEmployee() *Employee {
	instance := &Employee{}

	return instance
}
