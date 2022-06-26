package apperr

import "fmt"

type ApplicationError struct {
	Code    string
	Message string
	Err     error
}

func (e *ApplicationError) Error() string {
	var errStr string
	if e.Err != nil {
		errStr = fmt.Sprintf("%s", e.Err.Error())
	}
	return fmt.Sprintf("code=%s, err=%s, msg=%s", e.Code, errStr, e.Message)
}

func (e *ApplicationError) ResponseError() string {
	return fmt.Sprintf("%s (code: %s)", e.Message, e.Code)
}
