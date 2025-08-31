package models

type ErrorHandler string

const (
	InternalServerError ErrorHandler = "internal server error"
	UnauthorizedError   ErrorHandler = "user not authenticated"
)
