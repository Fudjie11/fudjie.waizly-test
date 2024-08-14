package model

type Pagination struct {
	TotalRows int32 `bson:"total_rows" json:"total_rows"`
	Limit     int32 `bson:"limit" json:"limit"`
	Offset    int32 `bson:"offset" json:"offset"`
	Page      int32 `bson:"page" json:"page"`
}

type ListPagingWithoutTotalRows[T any] struct {
	Data   []T `bson:"data" json:"data"`
	Limit  int `bson:"limit" json:"limit"`
	Offset int `bson:"offset" json:"offset"`
	Page   int `bson:"page" json:"page"`
}

type ListPaging[T any] struct {
	Data       []T        `db:"data" json:"data"`
	Pagination Pagination `db:"pagination" json:"pagination"`
}

func NewListPaging[T any](data []T, totalRows int, limit, offset, page *int) ListPaging[T] {

	var defLimit int32 = 10
	var defPage int32 = 1
	var defOffset int32 = 0

	result := ListPaging[T]{
		Data: data,
		Pagination: Pagination{
			TotalRows: int32(totalRows),
			Limit:     defLimit,
			Offset:    defOffset,
			Page:      defPage,
		},
	}

	if limit != nil {
		result.Pagination.Limit = int32(*limit)
	}
	if offset != nil {
		result.Pagination.Offset = int32(*offset)
	}
	if page != nil {
		result.Pagination.Page = int32(*page)
	}

	return result
}
