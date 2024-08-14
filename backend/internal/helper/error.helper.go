package helper

import (
	"database/sql"

	pkgErr "fudjie.waizly/backend-test/library/err"
)

/*
If there's only one message (index 0), it is the general error message
*/
func getMessage(s []string) string {
	msg := ""
	for i, v := range s {
		if i == 0 {
			msg = v
			break
		}
	}

	return msg
}

func NewBadRequestErr(err error, s ...string) pkgErr.CustErr {
	msg := getMessage(s)
	return pkgErr.NewError(pkgErr.ErrorOpts{
		Cause:   err,
		Message: msg,
		Type:    pkgErr.ErrBadRequest,
	})
}

func NewUnauthorizedErr(err error, s ...string) pkgErr.CustErr {
	msg := getMessage(s)
	return pkgErr.NewError(pkgErr.ErrorOpts{
		Cause:   err,
		Message: msg,
		Type:    pkgErr.ErrUnauthorized,
	})
}

func NewInternalServerErr(err error, s ...string) pkgErr.CustErr {
	msg := getMessage(s)
	return pkgErr.NewError(pkgErr.ErrorOpts{
		Cause:   err,
		Message: msg,
		Type:    pkgErr.ErrInternal,
	})
}

func NewForbiddenErr(err error, s ...string) pkgErr.CustErr {
	msg := getMessage(s)
	return pkgErr.NewError(pkgErr.ErrorOpts{
		Cause:   err,
		Message: msg,
		Type:    pkgErr.ErrForbidden,
	})
}

func NewNotFoundErr(err error, s ...string) pkgErr.CustErr {
	msg := getMessage(s)
	return pkgErr.NewError(pkgErr.ErrorOpts{
		Cause:   err,
		Message: msg,
		Type:    pkgErr.ErrNotFound,
	})
}

func NewSqlErr(err error, s ...string) pkgErr.CustErr {
	msg := getMessage(s)
	errOpts := pkgErr.ErrorOpts{
		Cause:   err,
		Message: msg,
		Type:    pkgErr.ErrInternal,
	}

	if err == sql.ErrNoRows {
		errOpts.Type = pkgErr.ErrNotFound
	}
	return pkgErr.NewError(errOpts)
}
