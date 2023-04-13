package helpers

type ResponseDefault struct {
	Status  string      `json:"status"`
	Message string      `json:"message"`
	Data    interface{} `json:"data"`
}

// APIResponse is for generating template responses
func APIResponse(message string, status string, data interface{}) ResponseDefault {

	return ResponseDefault{
		Status:  status,
		Message: message,
		Data:    data,
	}
}
