package qlite

// Add WHERE clause to the query.
// If WHERE is already in the query then it will ad AND.
func (q *Query) Where(condition, value string) *Query {
	if len(q.where) != 0 {
		q.where = append(q.where, "AND", condition)
	} else {
		q.where = append(q.where, "WHERE", condition)
	}
	q.values = append(q.values, value)
	return q
}

// Adds OR to the WHERE clause
func (q *Query) OrWhere(cond, value string) *Query {
	if len(q.where) != 0 {
		q.where = append(q.where, "OR")
		q.where = append(q.where, cond)
		q.values = append(q.values, value)
	}
	return q
}
