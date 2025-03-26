package goql

import (
	"fmt"
	"strconv"
	"strings"
)

const (
	MySQL      = "mysql"
	PostgreSQL = "postgres"
	SQLite     = "sqlite"
	SQLServer  = "sqlserver"
	Oracle     = "oracle"
)

type GoQL interface {
	BuildSQLStatement(content string, args ...any) (query string, queryArgs []any)
}

type goql struct {
	databaseDriver string
}

func New(databaseDriver string) GoQL {
	return &goql{
		databaseDriver: databaseDriver,
	}
}

func (g *goql) BuildSQLStatement(content string, args ...any) (query string, queryArgs []any) {
	switch g.databaseDriver {
	case MySQL, SQLite:
		return strings.ReplaceAll(content, "%s", "?"), args
	case SQLServer:
		return buildNumericPlaceholders(content, "@p"), args
	case Oracle:
		return buildNumericPlaceholders(content, ":"), args
	default:
		return buildNumericPlaceholders(content, "$"), args
	}
}

func buildNumericPlaceholders(content, prefix string) string {
	result := content
	for i := 1; i <= strings.Count(content, "%s"); i++ {
		result = strings.Replace(result, "%s", fmt.Sprint(prefix, strconv.Itoa(i)), 1)
	}
	return result
}
