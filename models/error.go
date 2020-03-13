package models

// ErrorStatus is used to relay error information to the client
type ErrorStatus struct {
	Status int    `json:"status"`
	Detail string `json:"detail"`
}
