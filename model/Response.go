package model

type Response struct {
	Status string `json:"status"`
	Error string `json:"error,omitempty"`
}
