package router

import (
	"context"
	"github.com/jackc/pgx/v4"
	"github.com/spf13/afero"

	"github.com/OmegaVoid/omega-inventory/internal/generated"
	"github.com/OmegaVoid/omega-inventory/pkg/graph"
	"github.com/gin-gonic/gin"
	"github.com/jackc/pgx/v4/log/zerologadapter"
	"github.com/jackc/pgx/v4/pgxpool"
	"github.com/rs/zerolog/log"
	"github.com/spf13/viper"

	"github.com/99designs/gqlgen/graphql/handler"
	"github.com/99designs/gqlgen/graphql/playground"
)

func GraphqlHandler(f *afero.Afero) gin.HandlerFunc {
	// NewExecutableSchema and Config are in the generated.go file
	// Resolver is in the resolver.go file

	logger := zerologadapter.NewLogger(log.Logger)
	dbpool, err := pgxpool.Connect(context.Background(), viper.GetString("database"))
	if err != nil {
		log.Fatal().Err(err).Msg("unable to connect to database")
	}
	connConf := dbpool.Config().ConnConfig
	connConf.Logger = logger
	connConf.LogLevel = pgx.LogLevelTrace

	h := handler.NewDefaultServer(
		generated.NewExecutableSchema(
			generated.Config{
				Resolvers: &graph.Resolver{
					DbPool: dbpool,
					Fs:     f,
				},
			},
		),
	)

	return func(c *gin.Context) {
		h.ServeHTTP(c.Writer, c.Request)
	}
}

func PlaygroundHandler() gin.HandlerFunc {
	h := playground.Handler("GraphQL", "/query")

	return func(c *gin.Context) {
		h.ServeHTTP(c.Writer, c.Request)
	}
}
