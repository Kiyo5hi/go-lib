package filter

import (
	"fmt"

	"go.mongodb.org/mongo-driver/bson"
)

var comparisonOperatorToMongo = map[ComparisonOperator]string{
	ComparisonOperatorEqual:          "$eq",
	ComparisonOperatorGreater:        "$gt",
	ComparisonOperatorGreaterOrEqual: "$gte",
	ComparisonOperatorLess:           "$lt",
	ComparisonOperatorLessOrEqual:    "$lte",
}

var logicalOperatorToMongo = map[LogicalOperator]string{
	LogicalOperatorAnd: "$and",
	LogicalOperatorOr:  "$or",
	LogicalOperatorNot: "$not",
}

func (fe *FilterExpression) ToMongoQuery() (bson.M, error) {
	logicalOperator, ok := logicalOperatorToMongo[fe.Logic]
	if !ok {
		return nil, fmt.Errorf("unsupported logical operator: %s", string(fe.Logic))
	}

	subqueries := []bson.M{}
	for _, ffe := range fe.Filters {
		if ffe.FilterExpression != nil {
			subquery, err := ffe.FilterExpression.ToMongoQuery()
			if err != nil {
				return nil, err
			}
			subqueries = append(subqueries, subquery)
		} else {
			f := ffe.Filter
			mongoOp, ok := comparisonOperatorToMongo[f.Operator]
			if !ok {
				return nil, fmt.Errorf("unsupported comparison operator: %s", string(f.Operator))
			}
			subquery := bson.M{
				f.Identifier: bson.M{
					mongoOp: f.Value.Primitive(),
				},
			}
			subqueries = append(subqueries, subquery)
		}
	}

	return bson.M{
		logicalOperator: subqueries,
	}, nil
}
