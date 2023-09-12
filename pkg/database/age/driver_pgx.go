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

func (driver PgxDriver) CreateGraph(ctx context.Context, graph string) error {
	sqlQuery := `SELECT * FROM ag_catalog.create_graph($1);`
	_, err := driver.Connection.ExecEx(ctx, sqlQuery, &pgx.QueryExOptions{}, graph)
	return err
}

func (driver PgxDriver) Exec(ctx context.Context, graph, query string, args ...any) (Result, error) {
	sqlQuery := buildCypherQuery(graph, query, 1)
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

func (driver PgxDriver) Query(graph, query string, retCount int, args ...any) AgtypeScanner {
	return &PgxAgtypeScanner{
		conn:        driver.Connection,
		graph:       graph,
		cypherQuery: query,
		returnCount: retCount,
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
	returnCount        int
	args               []any
	err                error
	cursor             *pgx.Rows
}

func (scanner *PgxAgtypeScanner) prepare(ctx context.Context) bool {
	sqlQuery := buildCypherQuery(scanner.graph, scanner.cypherQuery, scanner.returnCount)
	scanner.cursor, scanner.err = scanner.conn.QueryEx(ctx, sqlQuery, &pgx.QueryExOptions{}, scanner.args...)
	return scanner.err == nil && scanner.cursor.Err() == nil
}

func (scanner *PgxAgtypeScanner) Scan(ctx context.Context, dest ...*Agtype) {
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

func (scanner *PgxAgtypeScanner) Next(ctx context.Context) bool {
	if scanner.cursor == nil && !scanner.prepare(ctx) {
		return false
	}

	return scanner.err == nil && scanner.cursor.Next()
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
