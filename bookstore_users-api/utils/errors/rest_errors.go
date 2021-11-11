package errors

type RestErr struct {
	Message string `json:"message"`
	Status int32 `json:"code"`
	Error string `json:"error"`
}

func NewBadRequestError(message string) *RestErr {
	return &RestErr{
		Message: message,
		Status: 400,
		Error: "invalid_request",
	}
}

func NewNotFoundError(message string) *RestErr {
	return &RestErr{
		Message: message,
		Status: 404,
		Error: "not_found",
	}
}