package utils

import (
	"math"
	"user_service/internal/models"
	"user_service/proto/common"
)

func ToPagination(page, pageSize int32) models.Pagination {
	return models.Pagination{
		Limit:  pageSize,
		Offset: (page - 1) * pageSize,
	}
}

func ToPageInfo(page, pageSize, totalItems int32) *common.PageInfo {
	totalPages := int32(math.Ceil(float64(totalItems) / float64(pageSize)))

	return &common.PageInfo{
		TotalItems: totalItems,
		Page:       page,
		TotalPages: totalPages,
		PageSize:   pageSize,
		HasPrev:    page > 1,
		HasNext:    page < totalPages,
	}
}
