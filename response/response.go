package response

import "net/http"

// Response is template for http response
type Response struct {
	Code    int         `json:"code"`
	Status  string      `json:"status"`
	Message string      `json:"message"`
	Data    interface{} `json:"data"`
}

// Ok return response with status 200 OK
func Ok(data interface{}, msg string) *Response {
	return &Response{
		Code:    http.StatusOK,
		Status:  "OK",
		Message: msg,
		Data:    data,
	}
}

// BadRequest return response with status 400 Bad Request
func BadRequest(msg string) *Response {
	return &Response{
		Code:    http.StatusBadRequest,
		Status:  "Bad Request",
		Message: msg,
	}
}

// Unauthorized return response with status 401 Unauthorized
func Unauthorized(msg string) *Response {
	return &Response{
		Code:    http.StatusUnauthorized,
		Status:  "Unauthorized",
		Message: msg,
	}
}

// PaymentRequired return response with status 402 Payment Required
func PaymentRequired(msg string) *Response {
	return &Response{
		Code:    http.StatusPaymentRequired,
		Status:  "Payment Required",
		Message: msg,
	}
}

// Forbidden return response with status 403 Forbidden
func Forbidden(msg string) *Response {
	return &Response{
		Code:    http.StatusForbidden,
		Status:  "Forbidden",
		Message: msg,
	}
}

// NotFound return response with status 404 Not Found
func NotFound(msg string) *Response {
	return &Response{
		Code:    http.StatusNotFound,
		Status:  "Not Found",
		Message: msg,
	}
}

// InternalServerErr return response with status 500 Internal Server Error
func InternalServerErr(msg string) *Response {
	return &Response{
		Code:    http.StatusInternalServerError,
		Status:  "Internal Server Error",
		Message: msg,
	}
}

// ServiceUnavailable return response with status 503 Service Unavailable
func ServiceUnavailable(msg string) *Response {
	return &Response{
		Code:    http.StatusServiceUnavailable,
		Status:  "Service Unavailable",
		Message: msg,
	}
}
