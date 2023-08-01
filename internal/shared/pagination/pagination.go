package pagination

import (
	"net/http"
	"strconv"
	"strings"
)

func ConvertToInt(s string) (int, error) {
	return strconv.Atoi(s)
}

type Pagination struct {
	Page   int    `validate:"required"`
	Limit  int    `validate:"required"`
	Offset int    `db:"offset"`
	Field  string `db:"field"`
	Sort   string `db:"sort"`
}

func NewPaginationQuery(page, limit int, field, sort string) Pagination {
	pg := Pagination{Page: page, Limit: limit, Offset: (page - 1) * limit, Field: field, Sort: sort}
	return pg
}

func GetSortDirection(s string) string {
	switch strings.ToLower(s) {
	case "asc":
		return "ASC"
	case "desc":
		return "DESC"
	default:
		return "ASC"
	}
}

func CheckFieldQuery(s string, def string) string {
	switch strings.ToLower(s) {
	case "":
		return def
	default:
		return s
	}
}

func ParseQueryParams(r *http.Request, key string) string {
	return r.URL.Query().Get(key)
}
