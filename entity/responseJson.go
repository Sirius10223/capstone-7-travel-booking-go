package entity

type ResponseJson struct {
	Message string `json:"message"`
	Data    any    `json:"data"`
}