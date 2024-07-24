package filter

import "github.com/samber/lo"

// Primitive returns the value in primitive Go types
func (v *Value) Primitive() any {
	if v.String != nil {
		return string(*v.String)
	}
	if v.Int != nil {
		return int64(*v.Int)
	}
	if v.Float != nil {
		return float64(*v.Float)
	}
	if v.Boolean != nil {
		return bool(*v.Boolean)
	}
	return nil
}

// AndI returns a FilterOrFilterExpression which is an intermediate part of a filter expression with and logic
func AndI(filters ...*FilterOrFilterExpression) *FilterOrFilterExpression {
	return &FilterOrFilterExpression{
		FilterExpression: AndR(filters...),
	}
}

// OrI returns a FilterOrFilterExpression which is an intermediate part of a filter expression with or logic
func OrI(filters ...*FilterOrFilterExpression) *FilterOrFilterExpression {
	return &FilterOrFilterExpression{
		FilterExpression: OrR(filters...),
	}
}

// NotI returns a FilterOrFilterExpression which is an intermediate part of a filter expression with not logic
func NotI(filter *FilterOrFilterExpression) *FilterOrFilterExpression {
	return &FilterOrFilterExpression{
		FilterExpression: NotR(filter),
	}
}

// AndR returns a FilterExpression which can be the root of a filter expression with and logic
func AndR(filters ...*FilterOrFilterExpression) *FilterExpression {
	return &FilterExpression{
		Logic:   LogicalOperatorAnd,
		Filters: filters,
	}
}

// OrR returns a FilterExpression which can be the root of a filter expression with or logic
func OrR(filters ...*FilterOrFilterExpression) *FilterExpression {
	return &FilterExpression{
		Logic:   LogicalOperatorOr,
		Filters: filters,
	}
}

// NotR returns a FilterExpression which can be the root of a filter expression with not logic
func NotR(filter *FilterOrFilterExpression) *FilterExpression {
	return &FilterExpression{
		Logic:   LogicalOperatorNot,
		Filters: []*FilterOrFilterExpression{filter},
	}
}

// NewFilter is a shorthand for FilterExpression initialization
func NewFilterExpression(logic LogicalOperator, filters ...*FilterOrFilterExpression) *FilterOrFilterExpression {
	return &FilterOrFilterExpression{
		FilterExpression: &FilterExpression{
			Logic:   logic,
			Filters: filters,
		},
	}
}

// NewFilter is a shorthand for Filter initialization
func NewFilter(id string, op ComparisonOperator, val *Value) *FilterOrFilterExpression {
	return &FilterOrFilterExpression{
		Filter: &Filter{
			Identifier: id,
			Operator:   op,
			Value:      val,
		},
	}
}

// String creates a string value
func String(s string) *Value {
	return &Value{
		String: &s,
	}
}

// Int creates an int64 value
func Int(i int64) *Value {
	return &Value{
		Int: &i,
	}
}

// Float creates a float64 value
func Float(f float64) *Value {
	return &Value{
		Float: &f,
	}
}

// Boolean creates a bool value
func Boolean(b bool) *Value {
	return &Value{
		Boolean: lo.ToPtr(boolean(b)),
	}
}

// Null creates a nil value
func Null() *Value {
	return &Value{
		Null: lo.ToPtr(null(struct{}{})),
	}
}
