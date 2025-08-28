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

func (q *Query) Select(columns ...string) *Query {
	q.queryType = SELECT
	if len(columns) == 0 {
		q.columns = []string{"*"}
	} else {
		q.columns = columns
	}
	return q
}

func (q *Query) Distinct() *Query {
	q.distinct = true
	return q
}

func (q *Query) From(table string) *Query {
	q.table = table
	return q
}

// Adds WHERE clause
func (q *Query) Where(condition, value string) *Query {
	if len(q.where) != 0 {
		q.where = append(q.where, "AND", condition)
	} else {
		q.where = append(q.where, "WHERE", condition)
	}
	q.values = append(q.values, value)
	return q
}

func (q *Query) OrWhere(cond, value string) *Query {
	if len(q.where) != 0 {
		q.where = append(q.where, "OR")
		q.where = append(q.where, cond)
		q.values = append(q.values, value)
	}
	return q
}

func (q *Query) GroupBy(columns ...string) *Query {
	q.groupBy = append(q.groupBy, columns...)
	return q
}

func (q *Query) Having(condition, value string) *Query {
	if len(q.having) != 0 {
		q.having = append(q.having, "AND", condition)
	} else {
		q.having = append(q.having, "HAVING", condition)
	}
	q.values = append(q.values, value)
	return q
}

func (q *Query) OrHaving(condition, value string) *Query {
	if len(q.having) != 0 {
		q.having = append(q.having, "OR", condition)
		q.values = append(q.values, value)
	}
	return q
}

func (q *Query) OrderBy(column string, dir Direction) *Query {
	if len(q.orderBy) == 0 {
		q.orderBy = append(q.orderBy, column, string(dir))
	}
	return q
}

func (q *Query) Limit(limit int) *Query {
	q.limit = &limit
	return q
}

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
