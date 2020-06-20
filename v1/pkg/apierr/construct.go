package apierr

func newApiError(stat, msg, reqId string) ApiError {
	return ApiError{stat, msg, reqId}
}

func NewBadRequestError(msg, reqId string) ApiError {
	return newApiError(Status400, msg, reqId)
}

func NewUnauthorizedError(msg, reqId string) ApiError {
	return newApiError(Status401, msg, reqId)
}

func NewForbiddenError(msg, reqId string) ApiError {
	return newApiError(Status403, msg, reqId)
}

func NewNotFoundError(msg, reqId string) ApiError {
	return newApiError(Status404, msg, reqId)
}

func NewBadMethodError(msg, reqId string) ApiError {
	return newApiError(Status405, msg, reqId)
}

func NewRateLimitError(msg, reqId string) ApiError {
	return newApiError(Status429, msg, reqId)
}

func NewServerError(msg, reqId string) ApiError {
	return newApiError(Status500, msg, reqId)
}
