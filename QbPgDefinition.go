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
func (qb *PGQueryBuilder) Union(qb1, qb2 QueryBuilder) QueryBuilder {
	qb.Tokens = append(qb.Tokens,
		qb1.String(),
		"UNION",
		qb2.String())
	return qb
}

// Intersect for postgres INTERSECT clause
func (qb *PGQueryBuilder) Intersect(qb1, qb2 QueryBuilder) QueryBuilder {
	qb.Tokens = append(qb.Tokens,
		qb1.String(),
		"INTERSECT",
		qb2.String())
	return qb
}

// Except for postgres EXCEPT clause
func (qb *PGQueryBuilder) Except(qb1, qb2 QueryBuilder) QueryBuilder {
	qb.Tokens = append(qb.Tokens,
		qb1.String(),
		"EXCEPT",
		qb2.String())
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
