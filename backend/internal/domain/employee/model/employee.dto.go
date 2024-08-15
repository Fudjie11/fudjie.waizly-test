package model

import (
	"time"

	"github.com/google/uuid"
)

type EmployeeDto struct {
	EmployeeId uuid.UUID `json:"employee_id" db:"employee_id"`
	Name       string    `json:"name" db:"name"`
	JobTitle   string    `json:"job_title" db:"job_title"`
	Salary     float32   `json:"salary" db:"salary"`
	Department string    `json:"department" db:"department"`
	JoinedDate time.Time `json:"joined_date" db:"joined_date"`
}

type AddEmployeeDto struct {
	Name       string  `json:"name" db:"name"`
	JobTitle   string  `json:"job_title" db:"job_title"`
	Salary     float32 `json:"salary" db:"salary"`
	Department string  `json:"department" db:"department"`
}

type UpdateEmployeeDto struct {
	EmployeeId uuid.UUID `json:"employee_id" db:"employee_id"`
	Name       string    `json:"name" db:"name"`
	JobTitle   string    `json:"job_title" db:"job_title"`
	Salary     float32   `json:"salary" db:"salary"`
	Department string    `json:"department" db:"department"`
}
