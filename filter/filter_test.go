package filter_test

import (
	"testing"

	"github.com/alecthomas/participle/v2"
	"github.com/kiyo5hi/go-lib/filter"
	"github.com/samber/lo"
	"github.com/stretchr/testify/assert"
)

func TestValue_ParseString(t *testing.T) {
	parser, err := participle.Build[filter.Value]()
	assert.NoError(t, err)

	v, err := parser.ParseString("", `""`)
	assert.NoError(t, err)
	assert.Equal(t, `""`, v.Primitive())

	v, err = parser.ParseString("", `"hello"`)
	assert.NoError(t, err)
	assert.Equal(t, `"hello"`, v.Primitive())

	_, err = parser.ParseString("", `"`)
	assert.ErrorContains(t, err, "literal not terminate")
}

func TestValue_ParseInt(t *testing.T) {
	parser, err := participle.Build[filter.Value]()
	assert.NoError(t, err)

	v, err := parser.ParseString("", `0`)
	assert.NoError(t, err)
	assert.Equal(t, int64(0), v.Primitive())

	v, err = parser.ParseString("", `911`)
	assert.NoError(t, err)
	assert.Equal(t, int64(911), v.Primitive())
}

func TestValue_ParseFloat(t *testing.T) {
	parser, err := participle.Build[filter.Value]()
	assert.NoError(t, err)

	v, err := parser.ParseString("", `0.0`)
	assert.NoError(t, err)
	assert.Equal(t, float64(0.0), v.Primitive())

	v, err = parser.ParseString("", `119.911`)
	assert.NoError(t, err)
	assert.Equal(t, float64(119.911), v.Primitive())
}

func TestValue_Boolean(t *testing.T) {
	parser, err := participle.Build[filter.Value]()
	assert.NoError(t, err)

	v, err := parser.ParseString("", `false`)
	assert.NoError(t, err)
	assert.Equal(t, false, v.Primitive())

	v, err = parser.ParseString("", `true`)
	assert.NoError(t, err)
	assert.Equal(t, true, v.Primitive())

	_, err = parser.ParseString("", `tru`)
	assert.ErrorContains(t, err, "unexpected token")
}

func TestValue_Null(t *testing.T) {
	parser, err := participle.Build[filter.Value]()
	assert.NoError(t, err)

	v, err := parser.ParseString("", `null`)
	assert.NoError(t, err)
	assert.Equal(t, nil, v.Primitive())

	_, err = parser.ParseString("", `nul`)
	assert.ErrorContains(t, err, "unexpected token")
}

func TestFilter_Parse(t *testing.T) {
	parser, err := participle.Build[filter.Filter]()
	assert.NoError(t, err)

	v, err := parser.ParseString("", `a = "1"`)
	assert.NoError(t, err)
	assert.Equal(t, "a", v.Identifier)
	assert.Equal(t, filter.ComparisonOperatorEqual, v.Operator)
	assert.Equal(t, `"1"`, v.Value.Primitive())

	v, err = parser.ParseString("", `a >= 911.119`)
	assert.NoError(t, err)
	assert.Equal(t, "a", v.Identifier)
	assert.Equal(t, filter.ComparisonOperatorGreaterOrEqual, v.Operator)
	assert.Equal(t, float64(911.119), v.Value.Primitive())
}

func TestFilterExpression_Parse(t *testing.T) {
	parser, err := participle.Build[filter.FilterExpression]()
	assert.NoError(t, err)

	v, err := parser.ParseString("", `
	and(
		or(
			not(
				c = 100
			),
			b = "2"
		),
    a >= 1
	)`)
	assert.NoError(t, err)

	expected := &filter.FilterExpression{
		Logic: filter.LogicalOperatorAnd,
		Filters: []*filter.FilterOrFilterExpression{
			{
				FilterExpression: &filter.FilterExpression{
					Logic: filter.LogicalOperatorOr,
					Filters: []*filter.FilterOrFilterExpression{
						{
							FilterExpression: &filter.FilterExpression{
								Logic: filter.LogicalOperatorNot,
								Filters: []*filter.FilterOrFilterExpression{
									{
										Filter: &filter.Filter{
											Identifier: "c",
											Operator:   filter.ComparisonOperatorEqual,
											Value: &filter.Value{
												Int: lo.ToPtr(int64(100)),
											},
										},
									},
								},
							},
						},
						{
							Filter: &filter.Filter{
								Identifier: "b",
								Operator:   filter.ComparisonOperatorEqual,
								Value: &filter.Value{
									String: lo.ToPtr(`"2"`),
								},
							},
						},
					},
				},
			},
			{
				Filter: &filter.Filter{
					Identifier: "a",
					Operator:   filter.ComparisonOperatorGreaterOrEqual,
					Value: &filter.Value{
						Int: lo.ToPtr(int64(1)),
					},
				},
			},
		},
	}
	assert.Equal(t, expected, v)
}
