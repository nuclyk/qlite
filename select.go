package qlite

// Adds SELECT to the query.
// If there are columns then it will add '*' to the query.
func (q *Query) Select(columns ...string) *Query {
	q.queryType = SELECT
	if len(columns) == 0 {
		q.columns = []string{"*"}
	} else {
		q.columns = columns
	}
	return q
}

// Add DISTINCT after SELECT
func (q *Query) Distinct() *Query {
	q.distinct = true
	return q
}

func (q *Query) From(table string) *Query {
	q.table = table
	return q
}

// Adds GROUP BY (columns)
func (q *Query) GroupBy(columns ...string) *Query {
	q.groupBy = append(q.groupBy, columns...)
	return q
}

// Adds HAVING to the query
// If HAVING is already in the query and it adds AND to the expression.
func (q *Query) Having(condition, value string) *Query {
	if len(q.having) != 0 {
		q.having = append(q.having, "AND", condition)
	} else {
		q.having = append(q.having, "HAVING", condition)
	}
	q.values = append(q.values, value)
	return q
}

// Adds OR to HAVING
func (q *Query) OrHaving(condition, value string) *Query {
	if len(q.having) != 0 {
		q.having = append(q.having, "OR", condition)
		q.values = append(q.values, value)
	}
	return q
}

// Adds ORDER BY to the query
func (q *Query) OrderBy(column string, dir Direction) *Query {
	if len(q.orderBy) == 0 {
		q.orderBy = append(q.orderBy, column, string(dir))
	}
	return q
}

// Adds LIMIT to the query
func (q *Query) Limit(limit int) *Query {
	q.limit = &limit
	return q
}
