package paginationHelper

import (
	"math"
)

func NewPagination(totalRow int64, page int, limit int) (*PaginationV1, error) {
	totalPage := math.Ceil(float64(totalRow) / float64(limit))
	nextPage := page + 1
	if nextPage > int(totalPage) {
		nextPage = 0
	}

	return &PaginationV1{
		Page:      page,
		TotalRow:  totalRow,
		TotalPage: int64(totalPage),
		NextPage:  nextPage,
	}, nil
}
