package errors

//AppError -
type AppError struct { //nolint
	HTTPCode int    `json:"-"`
	Code     string `json:"errCode"`
	Message  string `json:"errMessage,omitempty"`
}

func (e AppError) Error() string {
	return e.Message
}
