package main

import "net/http"

type jsonResponse struct {
	Error   bool   `json:"error"`
	Message string `json:"message"`
	// поле может быть пустым
	Data any `json:"data,omitempty"`
}

func (app *Config) readJSON(w http.ResponseWriter, r *http.Request, data any) error {
	// ограничение на максимальный размер json в 1мбайт
	maxBytes := 1048576
}
