package response

import (
	"devbook_backend/src/models"
	"encoding/json"
	"net/http"
)

func sendResponse(w http.ResponseWriter, message string, statusCode int, data interface{}) {
	response := models.Response{
		Status:  statusCode,
		Message: message,
		Data:    nil,
	}

	if data != nil {
		response.Data = data
	}

	// if err := json.NewEncoder(response.Data).Encode(response); err != nil {
	// 	http.Error(w, "Error encoding response: "+err.Error(), http.StatusInternalServerError)
	// }

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(statusCode)

	if err := json.NewEncoder(w).Encode(response); err != nil {
		http.Error(w, "Error encoding response: "+err.Error(), http.StatusInternalServerError)
	}
}

func Success(w http.ResponseWriter, statusCode int, data interface{}) {
	sendResponse(w, "Request was successful", statusCode, data)
}

func Error(w http.ResponseWriter, message string, statusCode int, data interface{}) {
	sendResponse(w, message, statusCode, data)
}
