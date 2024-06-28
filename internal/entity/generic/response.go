package generic

type Response[T any] struct {
	Data    T      `json:"data"`
	Message string `json:"message"`
	Success bool   `json:"success"`
}
