package storage

import (
	"production_service/pkg/client/postgresql"
	"production_service/pkg/logging"
)

type storage struct {
	client postgresql.Client
	logger *logging.Logger
}
