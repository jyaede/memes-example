package main

import (
	"encoding/json"
	"net/http"
)

func jsonWithCode(w http.ResponseWriter, body interface{}, code int) {
	b, err := json.Marshal(body)
	w.Header().Set("content-type", "application/json")
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		return
	}
	w.WriteHeader(code)
	w.Write(b)
}

//ErrorMessage ...
type ErrorMessage struct {
	Status int    `json:"status"`
	Error  string `json:"error"`
}

func jsonError(w http.ResponseWriter, err error, code int) {
	em := ErrorMessage{
		Status: code,
		Error:  err.Error(),
	}
	b, _ := json.Marshal(em)
	w.Header().Set("content-type", "application/json")
	w.WriteHeader(code)
	w.Write(b)
}
