package errors

const (
	ErrorFindingPlanetByName = 1
	ErrorFindingPlanetById   = 2
	ErrorCreatingPlanet      = 3
	ErrorDeletingPlanet      = 4
)

//AppError -
type AppError struct { //nolint
	HTTPCode int    `json:"-"`
	Code     string `json:"errCode"`
	Message  string `json:"errMessage,omitempty"`
}

func (e AppError) Error() string {
	return e.Message
}

// TypedError is used to send to interface errors with code
type TypedError struct {
	Code    int         `json:"code"`
	Message string      `json:"message"`
	Stack   interface{} `json:"stack"`
}
