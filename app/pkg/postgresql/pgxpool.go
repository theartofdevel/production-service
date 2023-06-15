package psql

import (
	"context"
	"log"
	"strings"
	"time"

	"github.com/jackc/pgx/v5"
	"github.com/jackc/pgx/v5/pgconn"

	"github.com/jackc/pgx/v5/pgxpool"
)

type Client interface {
	Close()
	Acquire(ctx context.Context) (*pgxpool.Conn, error)
	AcquireFunc(ctx context.Context, f func(*pgxpool.Conn) error) error
	AcquireAllIdle(ctx context.Context) []*pgxpool.Conn
	Stat() *pgxpool.Stat
	Exec(ctx context.Context, sql string, arguments ...any) (pgconn.CommandTag, error)
	Query(ctx context.Context, sql string, args ...any) (pgx.Rows, error)
	QueryRow(ctx context.Context, sql string, args ...any) pgx.Row
	Begin(ctx context.Context) (pgx.Tx, error)
	BeginTx(ctx context.Context, txOptions pgx.TxOptions) (pgx.Tx, error)
}

type PgConfig struct {
	username string
	password string
	host     string
	port     string
	database string
}

// NewClient creates new postgres client.
func NewClient(
	ctx context.Context,
	maxAttempts int,
	maxDelay time.Duration,
	dsn string,
	binary bool,
) (pool *pgxpool.Pool, err error) {
	err = DoWithAttempts(func() error {
		ctxWto, cancel := context.WithTimeout(ctx, 5*time.Second) //nolint:gomnd
		defer cancel()

		pgxCfg, parseConfigErr := pgxpool.ParseConfig(dsn)
		if parseConfigErr != nil {
			log.Fatalf("Unable to parse config: %v\n", parseConfigErr)
		}

		if binary {
			pgxCfg.ConnConfig.DefaultQueryExecMode = pgx.QueryExecModeCacheDescribe
		}

		pool, parseConfigErr = pgxpool.NewWithConfig(ctxWto, pgxCfg)
		if parseConfigErr != nil {
			log.Println("Failed to connect to postgres... Going to do the next attempt")

			return parseConfigErr
		}

		return nil
	}, maxAttempts, maxDelay)

	if err != nil {
		log.Fatal("All attempts are exceeded. Unable to connect to postgres")
	}

	return pool, nil
}

func DoWithAttempts(fn func() error, maxAttempts int, delay time.Duration) error {
	var err error

	for maxAttempts > 0 {
		if err = fn(); err != nil {
			time.Sleep(delay)
			maxAttempts--

			continue
		}

		return nil
	}

	return err
}

func (c *PgConfig) ConnStringFromCfg() string {
	url := strings.Builder{}
	url.WriteString("postgresql://")
	url.WriteString(c.username)
	url.WriteString(":")
	url.WriteString(c.password)
	url.WriteString("@")
	url.WriteString(c.host)
	url.WriteString(":")
	url.WriteString(c.port)
	url.WriteString("/")
	url.WriteString(c.database)

	return url.String()
}
