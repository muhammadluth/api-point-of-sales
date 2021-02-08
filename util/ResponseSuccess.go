package util

import (
	"api-point-of-sales/model"
	"math"
)

func ResponseSuccessWithPagination(totalData, limit, page float64,
	data interface{}) model.ResponseSuccessWithPagination {
	var totalPage float64

	if limit != 0 {
		totalPage = math.Ceil(totalData / limit)
	} else {
		totalPage = limit
	}

	return model.ResponseSuccessWithPagination{
		TotalData: int(totalData),
		TotalPage: int(totalPage),
		Limit:     int(limit),
		Page:      int(page),
		Data:      data,
	}
}
