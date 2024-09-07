package filter

type FilterExpression struct {
	Logic   LogicalOperator             `parser:"@( 'and' | 'or' | 'not' )"`
	Filters []*FilterOrFilterExpression `parser:"'(' @@ (',' @@ )* ')'"`
}

type FilterOrFilterExpression struct {
	Filter           *Filter           `parser:"  @@"`
	FilterExpression *FilterExpression `parser:"| @@"`
}

type LogicalOperator string

const (
	LogicalOperatorAnd LogicalOperator = `and`
	LogicalOperatorOr  LogicalOperator = `or`
	LogicalOperatorNot LogicalOperator = `not`
)

type Filter struct {
	Identifier string             `parser:"@Ident"`
	Operator   ComparisonOperator `parser:"@( '=' | '>''=' | '<''=' | '>' | '<' | '~' )"`
	Value      *Value             `parser:"@@"`
}

type ComparisonOperator string

const (
	ComparisonOperatorEqual          ComparisonOperator = `=`
	ComparisonOperatorGreater        ComparisonOperator = `>`
	ComparisonOperatorGreaterOrEqual ComparisonOperator = `>=`
	ComparisonOperatorLess           ComparisonOperator = `<`
	ComparisonOperatorLessOrEqual    ComparisonOperator = `<=`
	ComparisonOperatorLike           ComparisonOperator = "~"
)

type Value struct {
	String  *string  `parser:"  @String"`
	Int     *int64   `parser:"| @Int"`
	Float   *float64 `parser:"| @Float"`
	Boolean *boolean `parser:"| @( 'true' | 'false' )"`
	Null    *null    `parser:"| @( 'null' )"`
}

type boolean bool

func (b *boolean) Capture(values []string) error {
	*b = values[0] == "true"
	return nil
}

type null struct{}

func (n *null) Capture(values []string) error {
	*n = struct{}{}
	return nil
}
