package common

type Pagination struct {
	Page    int `form:"page" default:"1" json:"page"`
	PerPage int `form:"perpage" default:"10" json:"perpage"`
}

type PaginationResp struct {
	Pagination
	Data  interface{}
	Total int64
}
