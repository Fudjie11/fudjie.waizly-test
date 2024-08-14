package enum

import "golang.org/x/exp/slices"

const (
	Pending  string = "pending"
	Active   string = "active"
	Inactive string = "inactive"
)

func IsValidUserType(s string) bool {
	stringArr := []string{Pending, Active, Inactive}
	return slices.Contains(stringArr, s)
}

const (
	InternalEmployee string = "BBG"
	ExternalEmployee string = "logistic"
)
