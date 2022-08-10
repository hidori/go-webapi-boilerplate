package errors

import (
	"errors"
	"fmt"
)

var (
	ErrBadRequest   = NewBadRequestError(fmt.Errorf("bad request"))
	ErrUnauthorized = NewUnauthorizedError(errors.New("unauthorized"))
	ErrForbidden    = NewForbiddenError(errors.New("forbidden"))
	ErrNotFound     = NewNotFoundError(errors.New("not found"))
)

type BadRequestError struct {
	error
}

func NewBadRequestError(err error) *BadRequestError {
	return &BadRequestError{
		error: err,
	}
}

func (e *BadRequestError) Error() string {
	if e.error != nil {
		return fmt.Sprintf("%s: %s", ErrBadRequest.Error(), e.Error())
	}

	return ErrBadRequest.Error()
}

func (e *BadRequestError) Unwrap() error {
	return e.error
}

type UnauthorizedError struct {
	error
}

func NewUnauthorizedError(err error) *UnauthorizedError {
	return &UnauthorizedError{}
}

func (e *UnauthorizedError) Error() string {
	if e.error != nil {
		return fmt.Sprintf("%s: %s", ErrUnauthorized.Error(), e.Error())
	}

	return ErrUnauthorized.Error()
}

func (e *UnauthorizedError) Unwrap() error {
	return e.error
}

type ForbiddenError struct {
	error
}

func NewForbiddenError(err error) *ForbiddenError {
	return &ForbiddenError{
		error: err,
	}
}

func (e *ForbiddenError) Error() string {
	if e.error != nil {
		return fmt.Sprintf("%s: %s", ErrForbidden.Error(), e.Error())
	}

	return ErrForbidden.Error()
}

func (e *ForbiddenError) Unwrap() error {
	return e.error
}

type NotFoundError struct {
	error
}

func NewNotFoundError(err error) *NotFoundError {
	return &NotFoundError{
		error: err,
	}
}

func (e *NotFoundError) Error() string {
	if e.error != nil {
		return fmt.Sprintf("%s: %s", ErrNotFound.Error(), e.Error())
	}

	return ErrNotFound.Error()
}

func (e *NotFoundError) Unwrap() error {
	return e.error
}
