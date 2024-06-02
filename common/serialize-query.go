package common

import (
	"strings"

	"gorm.io/gorm"
)

func getQueryKeyValue(key, value string) (string, string) {
	switch {
	case strings.Contains(key, "__exact"):
		key = strings.Replace(key, "__exact", "", 1)

	case strings.Contains(key, "__contains"):
		key = strings.Replace(key, "__contains", " LIKE ? ", 1)
		value = "%%" + value + "%%"

	case strings.Contains(key, "__startswith"):
		key = strings.Replace(key, "__startswith", " LIKE ? ", 1)
		value = value + "%%"

	case strings.Contains(key, "__endswith"):
		key = strings.Replace(key, "__endswith", " LIKE ?", 1)
		value = "%%" + value

	case strings.Contains(key, "__gt"):
		key = strings.Replace(key, "__gt", " > ? ", 1)

	case strings.Contains(key, "__gte"):
		key = strings.Replace(key, "__gte", " >= ? ", 1)

	case strings.Contains(key, "__lt"):
		key = strings.Replace(key, "__lt", " < ? ", 1)

	case strings.Contains(key, "__lte"):
		key = strings.Replace(key, "__lte", " <= ? ", 1)
	}

	return key, value
}

func Filter(db *gorm.DB, filter string) *gorm.DB {
	fields := strings.Split(filter, ",")
	for _, field := range fields {
		values := strings.Split(field, ":")
		if len(values) >= 2 {
			key, value := getQueryKeyValue(values[0], values[1])
			db = db.Where(key, value)
		}
	}

	return db
}
