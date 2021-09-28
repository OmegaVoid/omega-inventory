package graph

import (
	"github.com/jackc/pgx/v4/pgxpool"
	"github.com/spf13/afero"
)

// This file will not be regenerated automatically.
//
// It serves as dependency injection for your app, add any dependencies you require here.

type Resolver struct {
	Fs     *afero.Afero
	DbPool *pgxpool.Pool
}
