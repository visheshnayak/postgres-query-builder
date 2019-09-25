package querybuilder

// QueryBuilder interface for the querybuilder
type QueryBuilder interface {
	String() string
	Select(fields ...string) QueryBuilder
	OrderBy(fields ...string) QueryBuilder
	Distinct(fields ...string) QueryBuilder
	Where(fields ...string) QueryBuilder
	Limit(fields ...string) QueryBuilder
	Offset(fields ...string) QueryBuilder
	Fetch(N int64) QueryBuilder
	In(fields ...string) QueryBuilder
	Between(value, low, high string) QueryBuilder
	Like(value string, fields ...string) QueryBuilder
	IsNull(value string) QueryBuilder
	As(alias string) QueryBuilder
	From(value string) QueryBuilder
	InnerJoin(value string) QueryBuilder
	LeftJoin(table string) QueryBuilder
	FullOuterJoin(table string) QueryBuilder
	CrossJoin(table string) QueryBuilder
	NaturalJoin(joinType, table string) QueryBuilder
	GroupBy(fields ...string) QueryBuilder
	Having(fields ...string) QueryBuilder
	Union() QueryBuilder
	Intersect() QueryBuilder
	Except() QueryBuilder
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
