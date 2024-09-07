package filter

import (
	"fmt"

	"gorm.io/gorm/clause"
)

type GormBuilder struct{}

var logicalOperatorToGorm = map[LogicalOperator]func(...clause.Expression) clause.Expression{
	LogicalOperatorAnd: clause.And,
	LogicalOperatorOr:  clause.Or,
	LogicalOperatorNot: clause.Not,
}

var comparisonOperatorToGorm = map[ComparisonOperator]func(string, *Value) clause.Expression{
	ComparisonOperatorEqual: func(field string, value *Value) clause.Expression {
		return &clause.Eq{Column: field, Value: value.Primitive()}
	},
	ComparisonOperatorGreater: func(field string, value *Value) clause.Expression {
		return &clause.Gt{Column: field, Value: value.Primitive()}
	},
	ComparisonOperatorGreaterOrEqual: func(field string, value *Value) clause.Expression {
		return &clause.Gte{Column: field, Value: value.Primitive()}
	},
	ComparisonOperatorLess: func(field string, value *Value) clause.Expression {
		return &clause.Lt{Column: field, Value: value.Primitive()}
	},
	ComparisonOperatorLessOrEqual: func(field string, value *Value) clause.Expression {
		return &clause.Lte{Column: field, Value: value.Primitive()}
	},
}

func (*GormBuilder) LogicalOperator(op LogicalOperator) (func(...clause.Expression) clause.Expression, error) {
	query, ok := logicalOperatorToGorm[op]
	if !ok {
		return nil, fmt.Errorf("GORM does not support logical operator=%s", string(op))
	}
	return func(e ...clause.Expression) clause.Expression {
		return query(e...)
	}, nil
}

func (*GormBuilder) ComparisonOperator(op ComparisonOperator) (func(string, *Value) clause.Expression, error) {
	query, ok := comparisonOperatorToGorm[op]
	if !ok {
		return nil, fmt.Errorf("GORM does not support comparison operator=%s", string(op))
	}
	return func(s string, v *Value) clause.Expression {
		return query(s, v)
	}, nil
}

func (fe *FilterExpression) ToGorm() (clause.Expression, error) {
	return Build(&GormBuilder{}, fe)
}
