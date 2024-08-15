package entity

import (
	"github.com/google/uuid"
)

type Sales struct {
	SalesId    uuid.UUID `json:"sales_id" db:"sales_id"`
	EmployeeId uuid.UUID `json:"employee_id" db:"employee_id"`
	Sales      uuid.UUID `json:"sales" db:"sales"`
}

func NewSales() *Sales {
	instance := &Sales{}

	return instance
}
