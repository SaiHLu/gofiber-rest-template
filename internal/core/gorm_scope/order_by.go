package gormscope

import (
	"strings"

	"gorm.io/gorm"
)

func OrderBy(filter string) func(db *gorm.DB) *gorm.DB {
	return func(db *gorm.DB) *gorm.DB {
		fields := strings.Split(filter, ",")
		for _, order := range fields {
			values := strings.Split(order, ":")
			if len(values) >= 2 {
				if values[1] == "desc" || values[1] == "asc" {
					db = db.Order(values[0] + " " + values[1])
				}
			}
		}

		return db
	}
}
