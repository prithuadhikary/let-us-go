package errors

type ServiceError interface {
	Error() string
	GetErrorCode() string
	GetMessage() string
	GetStatus() int
}

type serviceError struct {
	ErrorCode string
	Message   string
	Status    int
}

func (serviceErr *serviceError) GetMessage() string {
	return serviceErr.Message
}

func (serviceErr *serviceError) GetErrorCode() string {
	return serviceErr.ErrorCode
}

func (serviceErr *serviceError) GetStatus() int {
	return serviceErr.Status
}

func (serviceErr *serviceError) Error() string {
	return serviceErr.Message
}

// NewServiceError creates a new ServiceError implementation
// with the passed error code and message and status.
func NewServiceError(code string, message string, status int) ServiceError {
	return &serviceError{code, message, status}
}
