package qlite

import (
	"strconv"
	"strings"
)

const (
	SELECT           = "SELECT"
	INSERT           = "INSERT"
	ASC    Direction = "ASC"
	DESC   Direction = "DESC"
)

type Query struct {
	queryType string
	distinct  bool
	columns   []string
	table     string
	values    []any
	where     []string
	groupBy   []string
	having    []string
	orderBy   []string
	limit     *int
}

type Direction string

func NewQuery() *Query {
	return &Query{
		columns: make([]string, 0),
		where:   make([]string, 0),
		groupBy: make([]string, 0),
	}
}

// Builds and returns sql string
func (q *Query) String() string {
	var sql []string
	switch q.queryType {
	case SELECT:
		sql = append(sql, SELECT)
		if q.distinct {
			sql = append(sql, "DISTINCT")
		}
		sql = append(sql, strings.Join(q.columns, ", "))
		if q.table != "" {
			sql = append(sql, "FROM")
			sql = append(sql, q.table)
		}
	default:
		sql = append(sql, "")
	}

	for _, cond := range q.where {
		sql = append(sql, cond)
	}

	if len(q.groupBy) != 0 {
		sql = append(sql, "GROUP BY", strings.Join(q.groupBy, ", "))
	}

	for _, cond := range q.having {
		sql = append(sql, cond)
	}

	if len(q.orderBy) != 0 {
		sql = append(sql, "ORDER BY", q.orderBy[0], q.orderBy[1])
	}

	if q.limit != nil {
		sql = append(sql, "LIMIT", strconv.Itoa(*q.limit))

	}

	return strings.Join(sql, " ")
}

func (q *Query) GetValues() []any {
	return q.values
}
