package filter_test

import (
	"sync"
	"testing"

	"github.com/kiyo5hi/go-lib/filter"
	"github.com/stretchr/testify/assert"
	"gorm.io/gorm"
	"gorm.io/gorm/clause"
	"gorm.io/gorm/schema"
	"gorm.io/gorm/utils/tests"
)

var db, _ = gorm.Open(tests.DummyDialector{}, nil)

// Ref: https://github.com/go-gorm/gorm/blob/master/clause/clause_test.go
func buildStmt(clauses []clause.Interface) *gorm.Statement {
	var (
		buildNames    []string
		buildNamesMap = map[string]bool{}
		user, _       = schema.Parse(&tests.User{}, &sync.Map{}, db.NamingStrategy)
		stmt          = gorm.Statement{DB: db, Table: user.Table, Schema: user, Clauses: map[string]clause.Clause{}}
	)

	for _, c := range clauses {
		if _, ok := buildNamesMap[c.Name()]; !ok {
			buildNames = append(buildNames, c.Name())
			buildNamesMap[c.Name()] = true
		}

		stmt.AddClause(c)
	}

	stmt.Build(buildNames...)

	return &stmt
}

func TestFilterExpression_ToGorm(t *testing.T) {
	fe := filter.AndR(
		filter.OrI(
			filter.NotI(
				filter.NewFilter("c", filter.ComparisonOperatorEqual, filter.Int(100)),
			),
			filter.NewFilter("b", filter.ComparisonOperatorEqual, filter.String("2")),
		),
		filter.NewFilter("a", filter.ComparisonOperatorGreaterOrEqual, filter.Int(1)),
	)

	query, err := fe.ToGorm()
	assert.NoError(t, err)

	clauses := []clause.Interface{clause.Select{}, clause.From{}, clause.Where{
		Exprs: []clause.Expression{query},
	}}
	stmt := buildStmt(clauses)

	assert.Equal(t, "SELECT * FROM `users` WHERE (`c` <> ? OR `b` = ?) AND `a` >= ?", stmt.SQL.String())
	assert.Equal(t, []any{int64(100), "2", int64(1)}, stmt.Vars)
}
