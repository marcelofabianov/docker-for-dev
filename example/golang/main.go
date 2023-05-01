package main

import (
	"encoding/json"
	"net/http"
)

type Response struct {
	Data   []interface{} `json:"data"`
	Status Status        `json:"status"`
}

type Status struct {
	Code    int    `json:"code"`
	Message string `json:"message"`
	Success bool   `json:"success"`
}

var response = Response{
	Data: []interface{}{},
	Status: Status{
		Code:    200,
		Message: "OK",
		Success: true,
	},
}

func main() {
	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		json.NewEncoder(w).Encode(response)
	})

	http.ListenAndServe(":8080", nil)
}
