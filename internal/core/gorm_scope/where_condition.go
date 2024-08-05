package gormscope

import (
	"gorm.io/gorm"
)

type ConditionType string

const (
	AndCondition ConditionType = "AND"
	OrCondition  ConditionType = "OR"
)

type Condition struct {
	Type  ConditionType
	Query string
	Args  interface{}
}

func WhereCondition(conditions ...Condition) func(db *gorm.DB) *gorm.DB {
	return func(db *gorm.DB) *gorm.DB {
		query := db

		for _, condition := range conditions {
			switch condition.Type {
			case AndCondition:
				query = query.Where(condition.Query, condition.Args)
			case OrCondition:
				query = query.Or(condition.Query, condition.Args)
			}
		}

		return query
	}
}
