package scope

import (
	helper "main/helper/struct"
	"math"

	"gorm.io/gorm"
)

func Paginate(query *gorm.DB, pagination *helper.Pagination) func(db *gorm.DB) *gorm.DB {
	var totalRows int64
	query.Count(&totalRows)
	pagination.TotalRows = totalRows
	totalPages := int(math.Ceil(float64(totalRows) / float64(pagination.Limit)))
	pagination.TotalPages = totalPages
	return func(db *gorm.DB) *gorm.DB {
		return db.Offset(pagination.GetOffset()).Limit(pagination.GetLimit()).Order(pagination.GetSort())
	}
}
