package helper

func FormatResponse(message string, data any) map[string]any {
	var response = map[string]any{}
	response["message"] = message
	if data != nil {
		response["data"] = data
	}
	return response
}
