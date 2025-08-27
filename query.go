package main

import (
	"strings"
)

const (
	SELECT = "SELECT"
	INSERT = "INSERT"
)

type Query struct {
	queryType string
	columns   []string
	table     string
	values    []any
	where     []string
	groupBy   []string
	having    []string
}

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

func (q *Query) From(table string) *Query {
	q.table = table
	return q
}

func (q *Query) Where(cond, value string) *Query {
	if len(q.where) != 0 {
		q.where = append(q.where, "AND")
		q.where = append(q.where, cond)
	} else {
		q.where = append(q.where, "WHERE")
		q.where = append(q.where, cond)
	}
	q.values = append(q.values, value)
	return q
}

func (q *Query) Or(cond, value string) *Query {
	if len(q.where) != 0 {
		q.where = append(q.where, "OR")
		q.where = append(q.where, cond)
	} else {
		q.where = append(q.where, "WHERE")
		q.where = append(q.where, cond)
	}
	q.values = append(q.values, value)
	return q
}

func (q *Query) GroupBy(columns ...string) *Query {
	q.groupBy = append(q.groupBy, columns...)
	return q
}

func (q *Query) Having(condition, value string) *Query {
	if len(q.having) != 0 {
		q.having = append(q.having, "AND")
		q.having = append(q.having, condition)
	} else {
		q.having = append(q.having, "HAVING")
		q.having = append(q.having, condition)
	}
	q.values = append(q.values, value)
	return q
}

func (q *Query) OrHaving(cond, value string) *Query {
	if len(q.having) != 0 {
		q.having = append(q.having, "OR")
		q.having = append(q.having, cond)
	} else {
		q.having = append(q.having, "HAVING")
		q.having = append(q.having, cond)
	}
	q.values = append(q.values, value)
	return q
}

func (q *Query) String() string {
	var sql []string
	switch q.queryType {
	case SELECT:
		sql = append(sql, SELECT)
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

	return strings.Join(sql, " ")
}
