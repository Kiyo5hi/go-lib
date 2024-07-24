package filter

import "github.com/alecthomas/participle/v2"

func NewParser() *participle.Parser[FilterExpression] {
	parser, _ := participle.Build[FilterExpression]()
	return parser
}
