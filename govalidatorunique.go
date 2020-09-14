package govalidatorunique

import (
	"errors"
	"fmt"
	"strings"

	"github.com/jinzhu/gorm"
)

// UniqueRule to check data is unique or not in db
type UniqueRule struct {
	db       *gorm.DB
	ruleName string
}

// Rule is func to register as custom rule
func (r *UniqueRule) Rule(field string, rule string, message string, value interface{}) error {
	var queryRow *gorm.DB
	var total int

	query := `SELECT COUNT(*) as total FROM %s WHERE %s = $1`
	params := strings.Split(strings.TrimPrefix(rule, fmt.Sprintf("%s:", r.ruleName)), ",")
	fmt.Println(params)

	if len(params) == 2 {
		query = fmt.Sprintf(query, params[0], params[1])
		queryRow = r.db.Raw(query, value)
	} else if len(params) == 4 {
		query += ` AND %s != ?`
		query = fmt.Sprintf(query, params[0], params[1], params[2])
		queryRow = r.db.Raw(query, value, params[3])
	} else {
		return fmt.Errorf("Arguments not enough")
	}

	queryRow.Count(&total)
	fmt.Println(total)

	if total > 0 {
		if message != "" {
			return errors.New(message)
		}

		return fmt.Errorf("The %s has already been taken", field)
	}

	return nil
}

// NewUniqueRule to create instance
func NewUniqueRule(db *gorm.DB, ruleName string) *UniqueRule {
	return &UniqueRule{db, ruleName}
}
