package filter

import (
	"fmt"

	"gorm.io/gorm/clause"
)

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

var logicalOperatorToGorm = map[LogicalOperator]func(...clause.Expression) clause.Expression{
	LogicalOperatorAnd: clause.And,
	LogicalOperatorOr:  clause.Or,
	LogicalOperatorNot: clause.Not,
}

func (fe *FilterExpression) ToGorm() (clause.Expression, error) {
	exprBuilder, ok := logicalOperatorToGorm[fe.Logic]
	if !ok {
		return nil, fmt.Errorf("unsupported logical operator: %s", string(fe.Logic))
	}

	exprs := []clause.Expression{}
	for _, ffe := range fe.Filters {
		if ffe.FilterExpression != nil {
			expr, err := ffe.FilterExpression.ToGorm()
			if err != nil {
				return nil, err
			}
			exprs = append(exprs, expr)
		} else {
			f := ffe.Filter
			gormOp, ok := comparisonOperatorToGorm[f.Operator]
			if !ok {
				return nil, fmt.Errorf("unsupported comparison operator: %s", string(f.Operator))
			}
			expr := gormOp(f.Identifier, f.Value)
			exprs = append(exprs, expr)
		}
	}

	return exprBuilder(exprs...), nil
}
