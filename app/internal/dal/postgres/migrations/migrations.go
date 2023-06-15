package migrations

import "embed"

//go:embed *.sql
var Content embed.FS
