package gql_server

import (
	"io"
	"time"

	"github.com/OmegaVoid/omega-inventory/pkg/middlewares"
	"github.com/OmegaVoid/omega-inventory/pkg/router"
	"github.com/gin-contrib/gzip"
	"github.com/gin-contrib/logger"
	"github.com/gin-gonic/gin"
	"github.com/rs/zerolog"
)

func GinServer(log zerolog.Logger) *gin.Engine {
	r := gin.New()

	r.Use(
		logger.SetLogger(
			logger.WithUTC(true),
			logger.WithLogger(
				func(c *gin.Context, out io.Writer, latency time.Duration) zerolog.Logger {
					return log.With().
						Timestamp().
						Int("status", c.Writer.Status()).
						Str("method", c.Request.Method).
						Str("path", c.Request.URL.Path).
						Str("ip", c.ClientIP()).
						Dur("latency", latency).
						Str("user_agent", c.Request.UserAgent()).
						Logger()
				},
			),
		),
	)
	r.Use(middlewares.GinContextToContextMiddleware())
	r.Use(gin.Recovery())
	r.Use(gzip.Gzip(gzip.DefaultCompression))

	r.POST("/query", router.GraphqlHandler())
	r.GET("/", router.PlaygroundHandler())
	return r
}
