package querybuilder

// QueryBuilder interface for the querybuilder
type QueryBuilder interface {
	String() string
	Select(fields ...string) QueryBuilder
	OrderBy(fields ...string) QueryBuilder
	// SelectDistinct(fields ...string) QueryBuilder
	Where(fields ...string) QueryBuilder
	Limit(fields ...string) QueryBuilder
	Offset(fields ...string) QueryBuilder
	Fetch(N int64) QueryBuilder
	In(fields ...string) QueryBuilder
	Between(value, low, high string) QueryBuilder
	Like(value string, fields ...string) QueryBuilder
	IsNull(value string) QueryBuilder
	As(alias string) QueryBuilder
	// InnerJoin(fields ...string) QueryBuilder
	// LeftJoin(fields ...string) QueryBuilder
	// SelfJoin(fields ...string) QueryBuilder
	// FullOuterJoin(fields ...string) QueryBuilder
	// CrossJoin(fields ...string) QueryBuilder
	// NaturalJoin(fields ...string) QueryBuilder
	GroupBy(fields ...string) QueryBuilder
	Having(fields ...string) QueryBuilder
	Union(qb1, qb2 QueryBuilder) QueryBuilder
	Intersect(qb1, qb2 QueryBuilder) QueryBuilder
	Except(qb1, qb2 QueryBuilder) QueryBuilder
	// GroupingSets(fields ...string) QueryBuilder
	// Cube(fields ...string) QueryBuilder
	// RollUp(fields ...string) QueryBuilder
	// Any(fields ...string) QueryBuilder
	// All(fields ...string) QueryBuilder
	Exists(query string) QueryBuilder
	And(cond string) QueryBuilder
	Or(cond string) QueryBuilder
}

// NewQb returns an instance of querybuilder
func NewQb() (qb QueryBuilder) {
	qb = new(PGQueryBuilder)
	return
}
