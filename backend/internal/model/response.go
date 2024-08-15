package model

import "github.com/google/uuid"

type ResponseStatus struct {
	Success      bool   `json:"success"`
	ResponseTime int64  `json:"response_time"`
	Latency      int64  `json:"latency"`
	ErrorCode    string `json:"error_code"`
	ErrorMessage string `json:"error_message"`
}

type WriteResponse struct {
	DataId  uuid.UUID `json:"data_id"`
	Success bool      `json:"success"`
}
