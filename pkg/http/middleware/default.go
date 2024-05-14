package middleware

import (
	"net/http"

	"github.com/rs/zerolog"
	"github.com/rs/zerolog/log"
)

// EmptyLoggingEnricher is a no-op log enricher designed to be used with the [Logging] middleware.
var EmptyLoggingEnricher = func(c *zerolog.Context) {}

// Logging is a middleware that adds a logger enriched with details of the HTTP request to the request's context.
func Logging(h http.Handler, enricher func(c *zerolog.Context)) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		l := log.With().
			Str("url", r.URL.String()).
			Str("method", r.Method)

		if enricher != nil {
			enricher(&l)
		}

		r = r.WithContext(l.Logger().WithContext(r.Context()))

		h.ServeHTTP(w, r)
	})
}

// Recovery recovers any panics, logs an error and executes the provided action.
func Recovery(h http.Handler, onPanic http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		l := zerolog.Ctx(r.Context())

		defer func() {
			if err := recover(); err != nil {
				l.Error().Any("panic", err).Msg("recovered from panic")
				onPanic.ServeHTTP(w, r)
			}
		}()

		h.ServeHTTP(w, r)
	})
}
