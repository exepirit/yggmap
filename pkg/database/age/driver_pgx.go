package age

import (
	"context"
	"github.com/jackc/pgx"
	"strconv"
	"strings"
)

type PgxDriver struct {
	Connection *pgx.Conn
}

func (driver PgxDriver) Exec(ctx context.Context, graph, query string, args ...any) (Result, error) {
	sqlQuery := buildCypherQuery(graph, query, 0)
	return driver.Connection.ExecEx(ctx, sqlQuery, &pgx.QueryExOptions{}, args...)
}

func (driver PgxDriver) QueryOne(ctx context.Context, graph, query string, args ...any) (Agtype, error) {
	sqlQuery := buildCypherQuery(graph, query, 1)
	row := driver.Connection.QueryRowEx(ctx, sqlQuery, &pgx.QueryExOptions{}, args...)

	var result string
	if err := row.Scan(&result); err != nil {
		return Vertex{}, err
	}

	return newUnmarshaller().unmarshal(result)
}

func (driver PgxDriver) Query(graph, query string, args ...any) AgtypeScanner {
	return &PgxAgtypeScanner{
		conn:        driver.Connection,
		graph:       graph,
		cypherQuery: query,
		args:        args,
	}
}

func buildCypherQuery(graph, cypherQuery string, retCount int) string {
	builder := strings.Builder{}
	builder.WriteString("select * from ag_catalog.cypher('")
	builder.WriteString(graph)
	builder.WriteString("', $$")
	builder.WriteString(cypherQuery)
	builder.WriteString("$$)")

	if retCount > 0 {
		builder.WriteString(" as (")
		for i := 0; i < retCount; i++ {
			builder.WriteRune('v')
			builder.WriteString(strconv.Itoa(i))
			builder.WriteString(" agtype")
			if i < retCount-1 {
				builder.WriteString(", ")
			}
		}
		builder.WriteString(")")
	}

	return builder.String()
}

type PgxAgtypeScanner struct {
	conn               *pgx.Conn
	graph, cypherQuery string
	args               []any
	err                error
	cursor             *pgx.Rows
}

func (scanner *PgxAgtypeScanner) prepare(ctx context.Context, retCount int) bool {
	sqlQuery := buildCypherQuery(scanner.graph, scanner.cypherQuery, retCount)
	scanner.cursor, scanner.err = scanner.conn.QueryEx(ctx, sqlQuery, &pgx.QueryExOptions{}, scanner.args...)
	return scanner.err != nil
}

func (scanner *PgxAgtypeScanner) Scan(ctx context.Context, dest ...*Agtype) {
	if scanner.cursor == nil && !scanner.prepare(ctx, len(dest)) {
		return
	}

	retValuesStr := make([]string, len(dest))
	scanDest := make([]any, len(retValuesStr))
	for i := 0; i < len(retValuesStr); i++ {
		scanDest[i] = &retValuesStr[i]
	}

	scanner.err = scanner.cursor.Scan(scanDest...)
	if scanner.err != nil || scanner.cursor.Err() != nil {
		return
	}

	for i, resultPtr := range dest {
		*resultPtr, scanner.err = newUnmarshaller().unmarshal(retValuesStr[i])
		if scanner.err != nil {
			return
		}
	}
}

func (scanner *PgxAgtypeScanner) Next() bool {
	if scanner.cursor == nil || scanner.err != nil {
		return scanner.err != nil
	}
	return scanner.cursor.Next()
}

func (scanner *PgxAgtypeScanner) Err() error {
	if scanner.cursor == nil || scanner.err != nil {
		return scanner.err
	}
	return scanner.cursor.Err()
}

func (scanner *PgxAgtypeScanner) Close() {
	scanner.cursor.Close()
}
