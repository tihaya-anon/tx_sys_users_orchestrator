package common

type ValidationError struct {
	Field string `json:"field,omitempty"`
	Msg   string `json:"msg,omitempty"`
}
