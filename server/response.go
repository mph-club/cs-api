package server

import "net/http"

type successResponse struct {
	Data    map[string]interface{} `json:"data"`
	Message string                 `json:"message"`
}

type errorResponse struct {
	Error   map[string]interface{} `json:"error"`
	Message string                 `json:"message"`
}

func makeOKResponse(responseMap map[string]interface{}) (int, interface{}) {
	return http.StatusOK, &successResponse{
		Data:    responseMap,
		Message: "success",
	}
}

func makeErrorResponse(statusCode int, responseMap map[string]interface{}) (int, interface{}) {
	return statusCode, &errorResponse{
		Error:   responseMap,
		Message: "fail",
	}
}
