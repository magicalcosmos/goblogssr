package util

import (
	"fmt"
)

const (
	ErrFail = 1
)

func NewErrorWithCode(code int32, text string) error {
	return &errorWithCode{c: code, s: text}
}

func GetErrorCode(err error) int32 {
	if err == nil {
		return 0
	}

	if e, ok := err.(*errorWithCode); ok {
		return e.Code()
	} else {
		return ErrFail
	}
}

type errorWithCode struct {
	c int32
	s string
}

func (e *errorWithCode) Error() string {
	return fmt.Sprintf("error: code = %d, desc = %s", e.c, e.s)
}

func (e *errorWithCode) Code() int32 {
	return e.c
}
