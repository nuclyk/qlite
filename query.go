package main

import (
	"strings"
)

const (
	SELECT = "SELECT"
	INSERT = "INSERT"
)

type Query struct {
	queryType  string
	columns    []string
	table      string
	values     []string
	conditions []string
}

func NewQuery() *Query {
	return &Query{
		columns: make([]string, 0),
		values:  make([]string, 0),
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
	if len(q.conditions) != 0 {
		q.conditions = append(q.conditions, "AND")
		q.conditions = append(q.conditions, cond)
	} else {
		q.conditions = append(q.conditions, "WHERE")
		q.conditions = append(q.conditions, cond)
	}
	q.values = append(q.values, value)
	return q
}

func (q *Query) Or(cond, value string) *Query {
	if len(q.conditions) != 0 {
		q.conditions = append(q.conditions, "OR")
		q.conditions = append(q.conditions, cond)
	} else {
		q.conditions = append(q.conditions, "WHERE")
		q.conditions = append(q.conditions, cond)
	}
	q.values = append(q.values, value)
	return q
}

func (q *Query) Build() string {
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

	for _, cond := range q.conditions {
		sql = append(sql, cond)
	}
	return strings.Join(sql, " ")
}
