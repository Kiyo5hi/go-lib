package filter

type FilterExpresssionBuilder[T any] interface {
	LogicalOperatorMapper[T]
	ComparisonOperatorMapper[T]
}

type LogicalOperatorMapper[T any] interface {
	LogicalOperator(LogicalOperator) (func(...T) T, error)
}

type ComparisonOperatorMapper[T any] interface {
	ComparisonOperator(ComparisonOperator) (func(string, *Value) T, error)
}

func Build[T any](feb FilterExpresssionBuilder[T], fe *FilterExpression) (t T, err error) {
	logical, err := feb.LogicalOperator(fe.Logic)
	if err != nil {
		return t, err
	}

	subqueries := []T{}
	for _, ffe := range fe.Filters {
		if ffe.FilterExpression != nil {
			query, err := Build(feb, ffe.FilterExpression)
			if err != nil {
				return t, err
			}
			subqueries = append(subqueries, query)
		} else {
			f := ffe.Filter
			comparison, err := feb.ComparisonOperator(f.Operator)
			if err != nil {
				return t, err
			}
			query := comparison(f.Identifier, f.Value)
			subqueries = append(subqueries, query)
		}
	}

	return logical(subqueries...), nil
}
