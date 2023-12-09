package helper

func FormatResponse(message string, data any) map[string]any {
	var response = map[string]any{}
	response["message"] = message
	if data != nil {
		response["data"] = data
	}
	return response
}

func FormatResponseValidation(message string, msgErr any) map[string]any {
	var response = map[string]any{}
	response["message"] = message
	if msgErr != nil {
		response["error"] = msgErr
	}
	return response
}

type ApiResponse[T any] struct {
	Status  int    `json:"status,omitempty"`
	Message string `json:"message,omitempty"`
	Data    T      `json:"data,omitempty"`
}
