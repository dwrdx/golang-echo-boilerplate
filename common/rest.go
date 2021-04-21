package common

import (
	"regexp"
	"strings"

	"github.com/go-playground/validator/v10"
)

var BodyValidator *validator.Validate

// Filter is a common structure holds the filter information for RESTapi
type Filter struct {
	Name  string
	Value string
}

// PaginationSQL generates SQL string to do paganition
// offset and limit will be used in SQL query string if they are valid
func PaginationSQL(offset string, limit string) string {
	sqlString := ""

	if limit != "" {
		if CheckNumber(limit) {
			// limit is a number
			sqlString = sqlString + " LIMIT " + limit
		}
	}

	if offset != "" {
		if CheckNumber(offset) {
			// offset is a number
			sqlString = sqlString + " OFFSET " + offset
		}
	}

	return sqlString
}

// ExtractFilters from a given filer string, if there is no valid filter found
// it returns nil
func ExtractFilters(filterStr string) []Filter {
	var filters []Filter
	r := regexp.MustCompile(`([a-zA-Z0-9_-]+:([a-zA-Z0-9_-]|\p{Han})+),?`)

	matches := r.FindAllStringSubmatch(filterStr, -1)
	if matches != nil {
		for _, v := range matches {
			tmp := v[1]
			s := strings.Split(tmp, ":")
			f := Filter{Name: s[0], Value: s[1]}
			filters = append(filters, f)
		}
		if len(filters) != 0 {
			return filters
		} else {
			return nil
		}
	} else {
		return nil
	}
}

func FilterSQL() {

}
