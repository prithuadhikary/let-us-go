package errors

import "net/http"

func InternalServerError() ServiceError {
	return NewServiceError(
		"server.error",
		"Internal Server Error.",
		http.StatusInternalServerError,
	)
}
