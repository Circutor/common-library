// Copyright (c) 2021 Circutor S.A. All rights reserved.

package errors

import "github.com/pkg/errors"

// errFound return error.
type errFound struct {
	error
}

// WrapErrFound wrap a error.
func WrapErrFound(err error, format string, args ...interface{}) error {
	return &errFound{errors.Wrapf(err, format, args...)}
}

// NewErrFound create a new error.
func NewErrFound(format string, args ...interface{}) error {
	return &errFound{errors.Errorf(format, args...)}
}

// ErrMessage struct content message error.
type ErrMessage struct {
	Message string `json:"message"`
}

// NewErrMessage create a error message.
func NewErrMessage(message string) ErrMessage {
	return ErrMessage{Message: message}
}
