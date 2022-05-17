package postgresql

import (
	"errors"
	"fmt"

	"github.com/jackc/pgconn"
)

// ParsePgError TODO Задача №3. Вынести этот метод куда-то из utils если вы видите в этом смысл
func ParsePgError(err error) error {
	var pgErr *pgconn.PgError
	if errors.Is(err, pgErr) {
		pgErr = err.(*pgconn.PgError)
		return fmt.Errorf("database error. message:%s, detail:%s, where:%s, sqlstate:%s",
			pgErr.Message, pgErr.Detail, pgErr.Where, pgErr.SQLState())
	}
	return err
}
