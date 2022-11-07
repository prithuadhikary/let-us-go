package model

type ApiError struct {
	ErrorCode string `json:"errorCode"`
	Message   string `json:"message"`
}
