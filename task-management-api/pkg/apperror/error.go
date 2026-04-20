package apperror

type AppError struct {
	HTTPCode int    `json:"-"`
	Code     string `json:"code"`
	Message  string `json:"message"`
}

func (e *AppError) Error() string { return e.Message }

func New(httpCode int, code, message string) *AppError {
	return &AppError{httpCode, code, message}
}

var (
	ErrNotFound       = &AppError{404, "NOT_FOUND", "resource not found"}
	ErrUnauthorized   = &AppError{401, "UNAUTHORIZED", "unauthorized"}
	ErrForbidden      = &AppError{403, "FORBIDDEN", "you don't have permission"}
	ErrBadRequest     = &AppError{400, "BAD_REQUEST", "bad request"}
	ErrInternalServer = &AppError{500, "INTERNAL_ERROR", "internal server error"}
	ErrEmailExists    = &AppError{409, "EMAIL_EXISTS", "email already exists"}
	ErrInvalidCreds   = &AppError{401, "INVALID_CREDENTIALS", "invalid email or password"}
)