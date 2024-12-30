package main

var (
	ErrInternalServer     = New(500, "internal server error")
	ErrUnimplemented      = New(501, "unimplemented")
	ErrForbidden          = New(403, "forbidden")
	ErrUnauthorized       = New(401, "unauthorized")
	ErrNotFound           = New(404, "not found")
	ErrInvalidCredentials = New(401, "invalid credentials")
)
