package identity

import (
	"github.com/google/uuid"
)

type Generator struct {
}

func (g *Generator) GenerateUUIDv4String() string {
	return uuid.NewString()
}
