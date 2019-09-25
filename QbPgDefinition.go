package querybuilder

import (
	"strconv"
	"strings"
)

// PGQueryBuilder struct that implements QueryBuilder
type PGQueryBuilder struct {
	Tokens []string
}

// CommaSpace the attacher
const CommaSpace = ", "

// String returns the querybuilder overall result
func (qb *PGQueryBuilder) String() string {
	return strings.Join(qb.Tokens, " ")
}

// Select for postgres select statement
func (qb *PGQueryBuilder) Select(fields ...string) QueryBuilder {
	qb.Tokens = append(qb.Tokens,
		"SELECT",
		strings.Join(fields, CommaSpace))
	return qb
}

// Distinct for postgres SELECT DISTINCT clause
func (qb *PGQueryBuilder) Distinct(fields ...string) QueryBuilder {
	qb.Tokens = append(qb.Tokens,
		"DISTINCT",
		strings.Join(fields, CommaSpace))
	return qb
}

// OrderBy for postgres order by statement
func (qb *PGQueryBuilder) OrderBy(fields ...string) QueryBuilder {
	qb.Tokens = append(qb.Tokens,
		"ORDER BY",
		strings.Join(fields, CommaSpace))
	return qb
}

// Where for postgres where clause
func (qb *PGQueryBuilder) Where(fields ...string) QueryBuilder {
	qb.Tokens = append(qb.Tokens,
		"WHERE",
		strings.Join(fields, CommaSpace))
	return qb
}

// Limit for postgres limit clause
func (qb *PGQueryBuilder) Limit(fields ...string) QueryBuilder {
	qb.Tokens = append(qb.Tokens,
		"LIMIT",
		strings.Join(fields, CommaSpace))
	return qb
}

// Offset for postgres offset clause
func (qb *PGQueryBuilder) Offset(fields ...string) QueryBuilder {
	qb.Tokens = append(qb.Tokens,
		"OFFSET",
		strings.Join(fields, CommaSpace))
	return qb
}

// Fetch for postgres FETCH clause
func (qb *PGQueryBuilder) Fetch(N int64) QueryBuilder {
	qb.Tokens = append(qb.Tokens,
		"FETCH",
		"FIRST",
		strconv.FormatInt(N, 10),
		"ROW",
		"ONLY")
	return qb
}

// In for postgres IN clause
func (qb *PGQueryBuilder) In(fields ...string) QueryBuilder {
	qb.Tokens = append(qb.Tokens,
		"IN",
		strings.Join(fields, CommaSpace))
	return qb
}

// Between for postgres BETWEEN clause
func (qb *PGQueryBuilder) Between(value, low, high string) QueryBuilder {
	qb.Tokens = append(qb.Tokens, value, low, "BETWEEN", high)
	return qb
}

// Like for postgres LIKE clause
func (qb *PGQueryBuilder) Like(value string, fields ...string) QueryBuilder {
	qb.Tokens = append(qb.Tokens,
		value,
		"LIKE",
		strings.Join(fields, CommaSpace))
	return qb
}

// IsNull for postgres ISNULL clause
func (qb *PGQueryBuilder) IsNull(value string) QueryBuilder {
	qb.Tokens = append(qb.Tokens, value, "IS", "NULL")
	return qb
}

// As for postgres AS clause
func (qb *PGQueryBuilder) As(alias string) QueryBuilder {
	qb.Tokens = append(qb.Tokens, "AS", alias)
	return qb
}

// From for postgres FROM clause
func (qb *PGQueryBuilder) From(table string) QueryBuilder {
	qb.Tokens = append(qb.Tokens, "FROM", table)
	return qb
}

// InnerJoin for postgres INNER JOIN clause
func (qb *PGQueryBuilder) InnerJoin(table string) QueryBuilder {
	qb.Tokens = append(qb.Tokens, "INNER JOIN", table)
	return qb
}

// LeftJoin for postgres LEFT JOIN clause
func (qb *PGQueryBuilder) LeftJoin(table string) QueryBuilder {
	qb.Tokens = append(qb.Tokens, "LEFT JOIN", table)
	return qb
}

// FullOuterJoin for postgres FULL OUTER JOIN clause
func (qb *PGQueryBuilder) FullOuterJoin(table string) QueryBuilder {
	qb.Tokens = append(qb.Tokens, "FULL OUTER JOIN", table)
	return qb
}

// CrossJoin for postgres CROSS JOIN clause
func (qb *PGQueryBuilder) CrossJoin(table string) QueryBuilder {
	qb.Tokens = append(qb.Tokens, "CROSS JOIN", table)
	return qb
}

// NaturalJoin for postgres NATURAL [INNER|OUTER|LEFT] JOIN clause
func (qb *PGQueryBuilder) NaturalJoin(joinType, table string) QueryBuilder {
	qb.Tokens = append(qb.Tokens,
		"NATURAL", "[", joinType, "]", "JOIN", table)
	return qb
}

// GroupBy for postgres GROUP BY clause
func (qb *PGQueryBuilder) GroupBy(fields ...string) QueryBuilder {
	qb.Tokens = append(qb.Tokens,
		"GROUP BY",
		strings.Join(fields, CommaSpace))
	return qb
}

// Having for postgres HAVING clause
func (qb *PGQueryBuilder) Having(fields ...string) QueryBuilder {
	qb.Tokens = append(qb.Tokens,
		"HAVING",
		strings.Join(fields, CommaSpace))
	return qb
}

// Union for postgres UNION clause
func (qb *PGQueryBuilder) Union() QueryBuilder {
	qb.Tokens = append(qb.Tokens, "UNION")
	return qb
}

// Intersect for postgres INTERSECT clause
func (qb *PGQueryBuilder) Intersect() QueryBuilder {
	qb.Tokens = append(qb.Tokens, "INTERSECT")
	return qb
}

// Except for postgres EXCEPT clause
func (qb *PGQueryBuilder) Except() QueryBuilder {
	qb.Tokens = append(qb.Tokens, "EXCEPT")
	return qb
}

// Exists for postgres EXISTS clause
func (qb *PGQueryBuilder) Exists(query string) QueryBuilder {
	qb.Tokens = append(qb.Tokens,
		"EXISTS", "(", query, ")")
	return qb
}

// And for postgres AND clause
func (qb *PGQueryBuilder) And(cond string) QueryBuilder {
	qb.Tokens = append(qb.Tokens, "AND", cond)
	return qb
}

// Or for postgres OR clause
func (qb *PGQueryBuilder) Or(cond string) QueryBuilder {
	qb.Tokens = append(qb.Tokens, "OR", cond)
	return qb
}
