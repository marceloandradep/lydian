package plgparser

import "errors"

var (
	HeaderNotFound = errors.New("header not found")
	InvalidSyntax  = errors.New("invalid syntax")
)
