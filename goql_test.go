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
			content:        "SELECT * FROM users WHERE id = %s AND name = %s",
			args:           []any{1, "John"},
			expectedQuery:  "SELECT * FROM users WHERE id = ? AND name = ?",
			expectedArgs:   []any{1, "John"},
		},
		{
			name:           "SQLite",
			databaseDriver: SQLite,
			content:        "SELECT * FROM users WHERE id = %s AND name = %s",
			args:           []any{1, "John"},
			expectedQuery:  "SELECT * FROM users WHERE id = ? AND name = ?",
			expectedArgs:   []any{1, "John"},
		},
		{
			name:           "SQLServer",
			databaseDriver: SQLServer,
			content:        "SELECT * FROM users WHERE id = %s AND name = %s",
			args:           []any{1, "John"},
			expectedQuery:  "SELECT * FROM users WHERE id = @p1 AND name = @p2",
			expectedArgs:   []any{1, "John"},
		},
		{
			name:           "Oracle",
			databaseDriver: Oracle,
			content:        "SELECT * FROM users WHERE id = %s AND name = %s",
			args:           []any{1, "John"},
			expectedQuery:  "SELECT * FROM users WHERE id = :1 AND name = :2",
			expectedArgs:   []any{1, "John"},
		},
		{
			name:           "PostgreSQL",
			databaseDriver: PostgreSQL,
			content:        "SELECT * FROM users WHERE id = %s AND name = %s",
			args:           []any{1, "John"},
			expectedQuery:  "SELECT * FROM users WHERE id = $1 AND name = $2",
			expectedArgs:   []any{1, "John"},
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
