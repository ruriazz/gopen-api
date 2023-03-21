package paginationHelper

type PaginationV1 struct {
	Page      int   `json:"page"`
	NextPage  int   `json:"nextPage"`
	TotalRow  int64 `json:"totalRow"`
	TotalPage int64 `json:"totalPage"`
}
