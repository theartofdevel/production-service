package postgresql

import (
	"errors"
	"fmt"
	"github.com/jackc/pgconn"
	"strings"
)

func ParsePgError(err error) error {
	var pgErr *pgconn.PgError
	if errors.Is(err, pgErr) {
		pgErr = err.(*pgconn.PgError)
		return fmt.Errorf("database error. message:%s, detail:%s, where:%s, sqlstate:%s", pgErr.Message, pgErr.Detail, pgErr.Where, pgErr.SQLState())
	}
	return err
}

func PrettySQL(q string) string {
	return strings.ReplaceAll(strings.ReplaceAll(q, "\t", ""), "\n", " ")
}
