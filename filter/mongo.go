package filter

import (
	"fmt"

	"go.mongodb.org/mongo-driver/bson"
)

var comparisonOperatorToMongo = map[ComparisonOperator]func(string, *Value) bson.M{
	ComparisonOperatorEqual: func(field string, value *Value) bson.M {
		return bson.M{
			field: bson.M{
				"$eq": value.Primitive(),
			},
		}
	},
	ComparisonOperatorGreater: func(field string, value *Value) bson.M {
		return bson.M{
			field: bson.M{
				"$gt": value.Primitive(),
			},
		}
	},
	ComparisonOperatorGreaterOrEqual: func(field string, value *Value) bson.M {
		return bson.M{
			field: bson.M{
				"$gte": value.Primitive(),
			},
		}
	},
	ComparisonOperatorLess: func(field string, value *Value) bson.M {
		return bson.M{
			field: bson.M{
				"$lt": value.Primitive(),
			},
		}
	},
	ComparisonOperatorLessOrEqual: func(field string, value *Value) bson.M {
		return bson.M{
			field: bson.M{
				"$lte": value.Primitive(),
			},
		}
	},
}

var logicalOperatorToMongo = map[LogicalOperator]func(...bson.M) bson.M{
	LogicalOperatorAnd: func(m ...bson.M) bson.M {
		return bson.M{
			"$and": m,
		}
	},
	LogicalOperatorOr: func(m ...bson.M) bson.M {
		return bson.M{
			"$or": m,
		}
	},
	LogicalOperatorNot: func(m ...bson.M) bson.M {
		return bson.M{
			"$not": m,
		}
	},
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
			subquery := mongoOp(f.Identifier, f.Value)
			subqueries = append(subqueries, subquery)
		}
	}

	return logicalOperator(subqueries...), nil
}
