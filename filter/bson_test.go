package filter_test

import (
	"testing"

	"github.com/kiyo5hi/go-lib/filter"
	"github.com/stretchr/testify/assert"
	"go.mongodb.org/mongo-driver/bson"
)

func TestFilterExpression_ToBson(t *testing.T) {
	fe := filter.AndR(
		filter.OrI(
			filter.NotI(
				filter.NewFilter("c", filter.ComparisonOperatorEqual, filter.Int(100)),
			),
			filter.NewFilter("b", filter.ComparisonOperatorEqual, filter.String("2")),
			filter.NewFilter("d", filter.ComparisonOperatorLike, filter.String("HeLlO")),
		),
		filter.NewFilter("a", filter.ComparisonOperatorGreaterOrEqual, filter.Int(1)),
	)

	query, err := fe.ToBson()
	assert.NoError(t, err)
	assert.Equal(t, bson.M{
		"$and": []bson.M{{
			"$or": []bson.M{{
				"$not": []bson.M{{
					"c": bson.M{
						"$eq": int64(100),
					},
				},
				},
			}, {
				"b": bson.M{
					"$eq": "2",
				},
			}, {
				"d": bson.M{
					"$regex":   ".*hello.*",
					"$options": "i",
				},
			}},
		}, {
			"a": bson.M{
				"$gte": int64(1),
			},
		}},
	}, query)
}
