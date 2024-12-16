package utils

type BasicResponse[T any] struct {
	Success bool   `json:"success"`
	Message string `json:"message"` // *string에서 string으로 변경
	Data    T      `json:"data"`
}

type ArrayBasicResponse[T any] struct {
	Success bool   `json:"success"`
	Message string `json:"message"` // *string에서 string으로 변경
	Data    []T    `json:"data"`
}
