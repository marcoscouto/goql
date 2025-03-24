package goql

import (
	"reflect"
	"testing"
)

func TestBuildSQLStatement(t *testing.T) {
	tests := []struct {
		name           string
		databaseDriver string
		content        string
		args           []any
		expectedQuery  string
		expectedArgs   []any
	}{
		{
			name:           "MySQL",
			databaseDriver: MySQL,
			content:        "SELECT * FROM users WHERE id = %s",
			args:           []any{1},
			expectedQuery:  "SELECT * FROM users WHERE id = ?",
			expectedArgs:   []any{1},
		},
		{
			name:           "SQLite",
			databaseDriver: SQLite,
			content:        "SELECT * FROM users WHERE id = %s",
			args:           []any{1},
			expectedQuery:  "SELECT * FROM users WHERE id = ?",
			expectedArgs:   []any{1},
		},
		{
			name:           "SQLServer",
			databaseDriver: SQLServer,
			content:        "SELECT * FROM users WHERE id = %s",
			args:           []any{1},
			expectedQuery:  "SELECT * FROM users WHERE id = @p",
			expectedArgs:   []any{1},
		},
		{
			name:           "Oracle",
			databaseDriver: Oracle,
			content:        "SELECT * FROM users WHERE id = %s",
			args:           []any{1},
			expectedQuery:  "SELECT * FROM users WHERE id = :param",
			expectedArgs:   []any{1},
		},
		{
			name:           "PostgreSQL",
			databaseDriver: PostgreSQL,
			content:        "SELECT * FROM users WHERE id = %s",
			args:           []any{1},
			expectedQuery:  "SELECT * FROM users WHERE id = $1",
			expectedArgs:   []any{1},
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			g := New(tt.databaseDriver)
			query, queryArgs := g.BuildSQLStatement(tt.content, tt.args...)
			if query != tt.expectedQuery {
				t.Errorf("expected query %s, got %s", tt.expectedQuery, query)
			}
			if !reflect.DeepEqual(queryArgs, tt.expectedArgs) {
				t.Errorf("expected args %v, got %v", tt.expectedArgs, queryArgs)
			}
		})
	}
}
