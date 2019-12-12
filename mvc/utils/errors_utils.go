package utils

// ApplicationError : struct for the errors
type ApplicationError struct {
	Message    string `json:"message"`
	StatusCode int    `json:"status_code"`
	Code       string `json:"code"`
}

// ApplicationErrorCreator : method for the ApplicationError
func (e *ApplicationError) ApplicationErrorCreator(message string, statuscode int, code string) {
	e.Message = message
	e.StatusCode = statuscode
	e.Code = code

}
