package err

import (
	"net/http"
	"testing"

	"github.com/stretchr/testify/assert"
	"google.golang.org/grpc/codes"
)

func TestError(t *testing.T) {
	tests := []struct {
		name             string
		opts             ErrorOpts
		expectedGRPCCode codes.Code
		expectedHTTPCode int
		expectedMessage  string
	}{
		{
			name: "Test Bad Request Error",
			opts: ErrorOpts{
				Message: "bad request",
				Type:    ErrBadRequest,
			},
			expectedGRPCCode: codes.InvalidArgument,
			expectedHTTPCode: http.StatusBadRequest,
			expectedMessage:  "bad request",
		},
		{
			name: "Test Unauthorized Error",
			opts: ErrorOpts{
				Message: "unauthorized",
				Type:    ErrUnauthorized,
			},
			expectedGRPCCode: codes.Unauthenticated,
			expectedHTTPCode: http.StatusUnauthorized,
			expectedMessage:  "unauthorized",
		},
		{
			name: "Test Internal Error",
			opts: ErrorOpts{
				Message: "internal error",
				Type:    ErrInternal,
			},
			expectedGRPCCode: codes.Internal,
			expectedHTTPCode: http.StatusInternalServerError,
			expectedMessage:  "internal error",
		},
		{
			name: "Test Forbidden Error",
			opts: ErrorOpts{
				Message: "forbidden error",
				Type:    ErrForbidden,
			},
			expectedGRPCCode: codes.PermissionDenied,
			expectedHTTPCode: http.StatusForbidden,
			expectedMessage:  "forbidden error",
		},
		{
			name: "Test Error with cause and fields",
			opts: ErrorOpts{
				Message: "error with cause and fields",
				Type:    ErrBadRequest,
				Cause:   NewError(ErrorOpts{Message: "cause"}),
				Fields:  map[string]interface{}{"field": "value"},
			},
			expectedGRPCCode: codes.InvalidArgument,
			expectedHTTPCode: http.StatusBadRequest,
			expectedMessage:  "error with cause and fields because {cause} with Fields {map[field:value]}",
		},
		{
			name: "Test Default Error",
			opts: ErrorOpts{
				Message: "default error",
			},
			expectedGRPCCode: codes.Internal,
			expectedHTTPCode: http.StatusInternalServerError,
			expectedMessage:  "default error",
		},
	}

	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			err := NewError(test.opts)
			err = GetError(err)
			assert.Equal(t, test.expectedGRPCCode, err.GRPCCode)
			assert.Equal(t, test.expectedHTTPCode, err.HTTPCode)
			assert.Equal(t, test.expectedMessage, err.Message)
		})
	}
}
