package model

import (
	"fmt"
	"strings"
)

type PaginationAndSearch struct {
	Limit  int32  `bson:"limit" json:"limit"`
	Offset int32  `bson:"offset" json:"offset"`
	Search string `bson:"search" json:"search"`
}

type Pagination struct {
	TotalRows int32 `bson:"total_rows" json:"total_rows"`
	Limit     int32 `bson:"limit" json:"limit"`
	Offset    int32 `bson:"offset" json:"offset"`
}

type ListPaging[T any] struct {
	Data       []T        `db:"data" json:"data"`
	Pagination Pagination `db:"pagination" json:"pagination"`
}

func (pg PaginationAndSearch) BuildPaginationAndSearchQuery(includeSearch bool) string {
	var (
		sb              strings.Builder
		paginationQuery string
		whereQuery      string
	)

	if includeSearch {
		whereQuery = ` WHERE LOWER(name) LIKE ('%' || LOWER(?) || '%') `
	}

	if pg.Limit > 0 {
		paginationQuery = fmt.Sprintf(" LIMIT %d OFFSET %d", pg.Limit, pg.Offset)
	}

	sb.WriteString(whereQuery)
	sb.WriteString(paginationQuery)

	return sb.String()
}

func NewListPaging[T any](data []T, totalRows int32, limit int32, offset int32) ListPaging[T] {
	var defLimit int32 = 10
	var defOffset int32 = 0

	result := ListPaging[T]{
		Data: data,
		Pagination: Pagination{
			TotalRows: int32(totalRows),
			Limit:     defLimit,
			Offset:    defOffset,
		},
	}

	result.Pagination.Limit = int32(limit)
	result.Pagination.Offset = int32(offset)

	return result
}
