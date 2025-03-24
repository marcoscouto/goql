package goql

import (
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
		return strings.ReplaceAll(content, "%s", "@p"), args
	case Oracle:
		return strings.ReplaceAll(content, "%s", ":param"), args
	default:
		result := content
		for i := 1; i <= strings.Count(content, "%s"); i++ {
			result = strings.Replace(result, "%s", "$"+strconv.Itoa(i), 1)
		}
		return result, args
	}
}
