package infra_errors

type CustomError struct {
	StatusCode int
	Messages   []*CustomErrorMessage
}

type CustomErrorMessage struct {
	Message string
	Field   string
}

func NewCustomError(err error, statusCode int, field string) *CustomError {
	var customErrorMessages []*CustomErrorMessage

	customErrorMessages = append(customErrorMessages, NewCustomErrorMessage(err, field))

	return &CustomError{
		StatusCode: statusCode,
		Messages:   customErrorMessages,
	}
}

func NewCustomErrorMessage(err error, field string) *CustomErrorMessage {
	return &CustomErrorMessage{
		Message: err.Error(),
		Field:   field,
	}
}

func NewCustomErrorArray(customErrorMessages []*CustomErrorMessage, statusCode int) *CustomError {
	return &CustomError{
		StatusCode: statusCode,
		Messages:   customErrorMessages,
	}
}
