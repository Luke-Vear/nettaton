package cloudplatform

import "errors"

var (
	ErrUserNotFoundInDatabase    = errors.New("user not found in database")
	ErrUserAlreadyExists         = errors.New("user already exists in database")
	ErrUserNotSpecified          = errors.New("user not specified")
	ErrRequiredFieldNotInRequest = errors.New("required request field empty")
	ErrInvalidQuestionKind       = errors.New("invalid question kind")
	ErrClaimNotFoundInJWT        = errors.New("claim not found in jwt")
	ErrInvalidJWT                = errors.New("invalid jwt")
)
