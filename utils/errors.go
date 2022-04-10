package utils

import "fmt"

type ErrorResponse struct {
	StatusCode int    `json:"-"`
	Request    string `json:"request"`
	ErrorMsg   string `json:"error"`
}

func (e *ErrorResponse) Error() string {
	return fmt.Sprintf("response error statusCode: %d, message: %s, request: %s", e.StatusCode, e.ErrorMsg, e.Request)
}

func (e *ErrorResponse) GetErrorMsg() string {
	return e.ErrorMsg
}

func (e *ErrorResponse) GetRequest() string {
	return e.Request
}

func (e *ErrorResponse) GetStatusCode() int {
	return e.StatusCode
}
