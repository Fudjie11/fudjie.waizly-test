package err

import (
	"fmt"
	"net/http"

	"github.com/pkg/errors"
	"google.golang.org/grpc/codes"
)

type ErrType string

const (
	ErrBadRequest   ErrType = "BadRequest"
	ErrUnauthorized ErrType = "Unauthorized"
	ErrInternal     ErrType = "Internal"
	ErrForbidden    ErrType = "Forbidden"
	ErrNotFound     ErrType = "NotFound"
)

const (
	defaultErrType     ErrType    = ErrInternal
	grpcDefaultErrCode codes.Code = codes.Internal
	httpDefaultErrCode int        = http.StatusInternalServerError
)

var grpcErrCode = map[ErrType]codes.Code{
	ErrBadRequest:   codes.InvalidArgument,
	ErrUnauthorized: codes.Unauthenticated,
	ErrInternal:     codes.Internal,
	ErrForbidden:    codes.PermissionDenied,
	ErrNotFound:     codes.NotFound,
}

var httpErrCode = map[ErrType]int{
	ErrBadRequest:   http.StatusBadRequest,
	ErrUnauthorized: http.StatusUnauthorized,
	ErrInternal:     http.StatusInternalServerError,
	ErrForbidden:    http.StatusForbidden,
	ErrNotFound:     http.StatusNotFound,
}

type (
	ErrorOpts struct {
		Message string
		Type    ErrType
		Cause   error
		Fields  map[string]interface{}
	}

	CustErr struct {
		GRPCCode codes.Code
		HTTPCode int
		ErrorOpts
	}
)

func (err CustErr) Error() string {
	if err.Cause == nil {
		return ""
	}

	return err.Cause.Error()
}

func (err CustErr) CustError() string {
	bcoz := ""
	fields := ""
	if err.Cause != nil {
		bcoz = fmt.Sprint(" because (", err.Cause.Error(), ")")
		if len(err.Fields) > 0 {
			fields = fmt.Sprintf(" with Fields {%+v}", err.Fields)
		}
	}

	return fmt.Sprint(err.Message, bcoz, fields)
}

// NewError - function for initializing error
func NewError(opts ErrorOpts) CustErr {
	return CustErr{
		GRPCCode:  getGRPCStatusCode(opts.Type),
		HTTPCode:  getHttpStatusCode(opts.Type),
		ErrorOpts: opts,
	}
}

// GetError - get error with CustErr type
func GetError(e error) CustErr {
	errc := errors.Cause(e)
	switch errc.(type) {
	case CustErr:
		return e.(CustErr)
	case *CustErr:
		return e.(CustErr)
	default:
		return CustErr{
			ErrorOpts: ErrorOpts{
				Cause: e,
				Type:  ErrInternal,
			},
			GRPCCode: getGRPCStatusCode(defaultErrType),
			HTTPCode: getHttpStatusCode(defaultErrType),
		}
	}
}

func getGRPCStatusCode(errType ErrType) codes.Code {
	if errType != "" {
		return grpcErrCode[errType]
	} else {
		return grpcDefaultErrCode
	}
}

func getHttpStatusCode(errType ErrType) int {
	if errType != "" {
		return httpErrCode[errType]
	} else {
		return httpDefaultErrCode
	}
}
