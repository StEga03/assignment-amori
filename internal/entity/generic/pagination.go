package generic

type PaginationResponse struct {
	Limit  int   `json:"limit"`
	Offset int   `json:"offset"`
	Total  int64 `json:"totalData"`
	Page   int   `json:"totalPage"`
}

type SearchRequest struct {
	ID     uint64 `json:"id"`
	Limit  int    `json:"limit"`
	Offset int    `json:"offset"`
}
