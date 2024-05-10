package errors_validation

type CustomErrorMessage struct {
	Message string
	Field   string
}

func NewCustomErrorMessage(err error, field string) *CustomErrorMessage {
	return &CustomErrorMessage{
		Message: err.Error(),
		Field:   field,
	}
}
