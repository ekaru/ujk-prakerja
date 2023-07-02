package helpers

import "net/http"

type BaseResponse struct {
	Status  string      `json:"status"`
	Message string      `json:"message"`
	Data    interface{} `json:"data"`
}

func ResponseDefault(code int, message string, data interface{}) BaseResponse {
	status := "Success"
	switch code {
	case http.StatusOK:
		status = "Success"
	case http.StatusCreated:
		status = "Created"
	case http.StatusBadRequest:
		status = "Bad Request"
	case http.StatusConflict:
		status = "Conflict"
	case http.StatusUnauthorized:
		status = "Unauthorized"
	case http.StatusForbidden:
		status = "Forbidden"
	case http.StatusNotFound:
		status = "Not Found"
	case http.StatusInternalServerError:
		status = "Internal Server Error"
	default:
		status = "Error"
	}

	return BaseResponse{
		Status:  status,
		Message: message,
		Data:    data,
	}
}
