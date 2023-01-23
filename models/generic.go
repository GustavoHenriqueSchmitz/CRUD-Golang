package models

type Response struct {
	Data    interface{}
	Message string
	Error   bool
}
