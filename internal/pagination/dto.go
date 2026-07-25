package pagination

type Metadata struct {
	Page       int `json:"page"`
	Limit      int `json:"limit"`
	TotalItems int `json:"total_items"`
	TotalPages int `json:"total_pages"`
}

type Response[T any] struct {
	Items      []T      `json:"items"`
	Pagination Metadata `json:"pagination"`
}
