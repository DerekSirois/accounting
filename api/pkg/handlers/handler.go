package handlers

import (
	"encoding/json"
	"net/http"
)

type JsonResponse struct {
	Error   bool   `json:"error"`
	Message string `json:"message"`
	Data    any    `json:"data,omitempty"`
}

func writeJson(w http.ResponseWriter, status int, data any) {
	out, _ := json.Marshal(data)

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(status)

	_, _ = w.Write(out)
}

func readJson(r *http.Request, data any) error {
	err := json.NewDecoder(r.Body).Decode(data)
	return err
}

func errorJson(w http.ResponseWriter, err error, status int) {
	res := JsonResponse{
		Error:   true,
		Message: err.Error(),
	}
	writeJson(w, status, res)
}

func Index(w http.ResponseWriter, r *http.Request) {
	res := JsonResponse{
		Error:   false,
		Message: "Welcome",
	}
	writeJson(w, http.StatusOK, res)
}
