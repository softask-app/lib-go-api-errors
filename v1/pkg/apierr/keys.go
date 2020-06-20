package apierr

const (
	JsKeyMessage   = "message"
	JsKeyRequestId = "requestId"
	JsKeyStatus    = "status"
)

const (
	Status400 = "bad-request"
	Status401 = "unauthorized"
	Status403 = "forbidden"
	Status404 = "not-found"
	Status405 = "invalid-http-method"
	Status429 = "rate-limited"
	Status500 = "server-error"
)