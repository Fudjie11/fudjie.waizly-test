package helper

import (
	"context"
	"net/http"
	"strconv"
	"time"

	iModel "fudjie.waizly/backend-test/internal/model"
	"github.com/gofiber/fiber/v2"
)

type JSONResponse struct {
	Data           interface{} `json:"data,omitempty"`
	Pagination     interface{} `json:"pagination,omitempty"`
	ResponseStatus interface{} `json:"response_status,omitempty"`
}

type ResponseStatus struct {
	Success      bool   `json:"success,omitempty"`
	Latency      int64  `json:"latency,omitempty"`
	StatusCode   string `json:"status_code,omitempty"`
	ErrorMessage string `json:"error_message,omitempty"`
}

func getAppData(cRequestData any) map[string]interface{} {
	appData := map[string]interface{}{}

	if cRequestData == nil {
		return appData
	}

	mRequestData := cRequestData.(iModel.ContextAppData)
	appData["LanguageId"] = mRequestData.LanguageId
	appData["AppId"] = mRequestData.AppId
	appData["UserId"] = mRequestData.UserId
	appData["TenantId"] = mRequestData.TenantId
	appData["CustomerId"] = mRequestData.CustomerId
	appData["TimeZoneId"] = mRequestData.TimeZoneId
	appData["TimeZoneOffset"] = mRequestData.TimeZoneOffset

	return appData
}

func NewContextFromHttp(c *fiber.Ctx) context.Context {
	appData := getAppData(c.Locals("REQUEST_DATA"))
	newCtx := context.WithValue(c.Context(), "APP_DATA", appData)
	return newCtx
}

func NewOKRequestResp(c *fiber.Ctx, message string, data interface{}, startTime time.Time) error {
	res := NewJSONResponse().WithData(data)
	res.ResponseStatus = NewResponseStatus().WithSuccess(true).WithLatency(startTime).WithStatusCode(http.StatusOK)
	return c.JSON(res)
}

func NewBadRequestResp(c *fiber.Ctx, message string, startTime time.Time) error {
	res := NewJSONResponse()
	res.ResponseStatus = NewResponseStatus().WithSuccess(false).WithLatency(startTime).WithStatusCode(http.StatusBadRequest)
	c.Status(http.StatusBadRequest)

	return c.JSON(res)
}

func New404RequestResp(c *fiber.Ctx, message string, startTime time.Time) error {
	res := NewJSONResponse()
	res.ResponseStatus = NewResponseStatus().WithSuccess(false).WithLatency(startTime).WithStatusCode(http.StatusNotFound)
	c.Status(http.StatusNotFound)
	return c.JSON(res)
}

func NewJSONResponse() JSONResponse {
	return JSONResponse{}
}

func (r JSONResponse) WithData(data interface{}) JSONResponse {
	r.Data = data
	return r
}

func NewResponseStatus() ResponseStatus {
	return ResponseStatus{}
}

func (r ResponseStatus) WithSuccess(success bool) ResponseStatus {
	r.Success = success
	return r
}

func (r ResponseStatus) WithLatency(start time.Time) ResponseStatus {
	end := time.Now()
	r.Latency = end.Sub(start).Microseconds()
	return r
}

func (r ResponseStatus) WithStatusCode(errorCode int) ResponseStatus {
	r.StatusCode = strconv.FormatInt(int64(errorCode), 10)
	return r
}

func (r ResponseStatus) WithErrorMessage(errorMessage string) ResponseStatus {
	r.ErrorMessage = errorMessage
	return r
}

func NewOkResponse(c *fiber.Ctx, data interface{}) error {
	return c.JSON(fiber.Map{"data": data})
}

func NewOkPagingResponse(c *fiber.Ctx, dataWithPaging interface{}) error {
	return c.JSON(dataWithPaging)
}
