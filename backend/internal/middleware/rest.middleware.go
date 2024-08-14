package middleware

import (
	"encoding/json"
	"net/http"
	"strconv"
	"strings"
	"time"

	// pb "bluebird.tech/kirim/protobuf/tms/v1"
	iHelper "fudjie.waizly/backend-test/internal/helper"
	iModel "fudjie.waizly/backend-test/internal/model"
	pkgErr "fudjie.waizly/backend-test/library/err"
	"github.com/gofiber/fiber/v2"
	"github.com/google/uuid"
)

// HttpMiddleware is an interface that defines middleware for handling request and response
type HttpMiddleware interface {
	ResponseDataMiddleware(fc *fiber.Ctx) error // use this middleware before call the handler function
	RequestDataMiddleware(fc *fiber.Ctx) error
}

// Module is a struct that implements HttpMiddleware
type Module struct{}

// NewHttpMiddleware returns a new instance of HttpMiddleware
func NewHttpMiddleware() HttpMiddleware {
	return &Module{}
}

// setDefaultContext sets the request start time in the context
func (m *Module) setDefaultContext(fc *fiber.Ctx) {
	startTime := time.Now()
	//  set request start time to context
	fc.Locals("req_start_time", startTime)
}

// setAppContext sets the request app data in the context
func (m *Module) setAppContext(fc *fiber.Ctx) {
	defaultUserId := "65a06585-1625-448d-8c2c-5a705a04797f" // id user from tms_user

	parsedUserId, err := uuid.Parse(fc.Get("x-user-id", defaultUserId))
	if err != nil {
		parsedUserId = uuid.Nil
	}

	parsedAppId := fc.Get("x-app-id")

	parsedCustomerId, err := uuid.Parse(fc.Get("x-customer-id"))
	if err != nil {
		parsedCustomerId = uuid.Nil
	}

	parsedTenantId, err := uuid.Parse(fc.Get("x-tenant-id"))
	if err != nil {
		parsedTenantId = uuid.Nil
	}

	appData := iModel.ContextAppData{
		UserId:     parsedUserId,
		TenantId:   parsedTenantId,
		CustomerId: parsedCustomerId,
		AppId:      parsedAppId,
		LanguageId: "id",
	}

	fc.Locals("REQUEST_DATA", appData)
}

// measuresLatency measures the latency of the request by retrieving the start time from the context
func (m *Module) measuresLatency(fc *fiber.Ctx) int64 {
	var (
		latency int64 = 0
	)
	// validate request time from context if exist calculate the latency
	if startTime, ok := fc.Locals("req_start_time").(time.Time); ok {
		// measures the time elapsed from startTime to the present.
		latency = time.Since(startTime).Microseconds()
	}
	return latency
}

// handleErrorResponse processes the error and returns a JSON response with error information
func (m *Module) handleErrorResponse(fc *fiber.Ctx, latency int64, err error) error {
	baseRes := iHelper.JSONResponse{}
	cusErr := pkgErr.GetError(err)
	message := cusErr.Error()
	if cusErr.Message != "" {
		message = cusErr.Message
	}
	resStatus := &iModel.ResponseStatus{
		Latency:      latency,
		ErrorMessage: message,
		ErrorCode:    strconv.FormatInt(int64(cusErr.HTTPCode), 10),
	}
	baseRes.ResponseStatus = resStatus
	fc.Response().SetStatusCode(cusErr.HTTPCode)
	fc.JSON(baseRes)
	return nil
}

// handleResponseJson processes the JSON response by adding latency information and success status
func (m *Module) handleResponseJson(fc *fiber.Ctx, latency int64) error {
	baseRes := &iHelper.JSONResponse{}
	if err := json.Unmarshal(fc.Response().Body(), baseRes); err != nil {
		return err
	}
	resStatus := &iModel.ResponseStatus{
		Latency: latency,
		Success: true,
	}
	baseRes.ResponseStatus = resStatus
	fc.Response().SetStatusCode(http.StatusOK)
	fc.JSON(baseRes)
	return nil
}

// RequestDataMiddleware sets the request start time in the context
func (m *Module) RequestDataMiddleware(fc *fiber.Ctx) error {
	m.setDefaultContext(fc)
	m.setAppContext(fc)
	return fc.Next()
}

// ResponseDataMiddleware processes the response by adding latency information and handling errors if any
func (m *Module) ResponseDataMiddleware(fc *fiber.Ctx) error {
	err := fc.Next()

	// calculate latency for get latency value
	latency := m.measuresLatency(fc)
	// if content type json, only need to set data latency and is success true.
	if contentType := string(fc.Response().Header.ContentType()); strings.Contains(contentType, "application/json") {
		err = m.handleResponseJson(fc, latency)
	}

	// if error is not nil, generates a json response status according to the data from custom error
	if err != nil {
		err = m.handleErrorResponse(fc, latency, err)
	}

	return err
}
