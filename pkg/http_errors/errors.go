package httpErrors

import "errors"

var (
	ErrBadRequest            = errors.New("Bad request")
	ErrWrongCredential       = errors.New("Wrond Credentials")
	ErrNotFount              = errors.New("Not found")
	ErrUnauthorized          = errors.New("Unauthorized")
	ErrForbidden             = errors.New("Forbidden")
	ErrPermissionDenied      = errors.New("Permission Denied")
	ErrNotRequiredFields     = errors.New("Not requiresd fields")
	ErrBadQueryParams        = errors.New("Bad query params")
	ErrInternalServerError   = errors.New("Internal server error")
	ErrRequestTimeoutError   = errors.New("Request timeout error")
	ErrInvalidJWTToken       = errors.New("Invalid JWT Token")
	ErrInvalidJWTClaims      = errors.New("Invalid JWT Claims")
	ErrNotAllowedImageHeader = errors.New("Not allowed image header")
	ErrNoCookie              = errors.New("No cookie")
	ErrInvalidUUID           = errors.New("Invalid UUID")
)
