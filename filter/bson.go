package filter

import (
	"fmt"

	"go.mongodb.org/mongo-driver/bson"
)

type BsonBuilder struct{}

var comparisonOperatorToBson = map[ComparisonOperator]func(string, *Value) bson.M{
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

var logicalOperatorToBson = map[LogicalOperator]func(...bson.M) bson.M{
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

func (*BsonBuilder) LogicalOperator(op LogicalOperator) (func(...bson.M) bson.M, error) {
	query, ok := logicalOperatorToBson[op]
	if !ok {
		return nil, fmt.Errorf("GORM does not support logical operator=%s", string(op))
	}
	return func(e ...bson.M) bson.M {
		return query(e...)
	}, nil
}

func (*BsonBuilder) ComparisonOperator(op ComparisonOperator) (func(string, *Value) bson.M, error) {
	query, ok := comparisonOperatorToBson[op]
	if !ok {
		return nil, fmt.Errorf("GORM does not support comparison operator=%s", string(op))
	}
	return func(s string, v *Value) bson.M {
		return query(s, v)
	}, nil
}

func (fe *FilterExpression) ToBson() (bson.M, error) {
	return Build(&BsonBuilder{}, fe)
}
