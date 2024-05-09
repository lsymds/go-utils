package middleware_test

import (
	"io"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/lsymds/go-utils/pkg/http/middleware"
	"github.com/rs/zerolog"
)

func TestLoggingMiddlewareWorks(t *testing.T) {
	var logger *zerolog.Logger

	middleware.
		Logging(
			http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
				logger = zerolog.Ctx(r.Context())
			}),
			nil,
		).
		ServeHTTP(httptest.NewRecorder(), httptest.NewRequest("GET", "https://google.com", nil))

	if logger == nil {
		t.Errorf("expected logger to be present")
	}
}

func TestRecoveryMiddlewareWorks(t *testing.T) {
	recorder := httptest.NewRecorder()

	middleware.
		Recovery(
			http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
				panic("woops")
			}),
			http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
				w.Write([]byte("recovered"))
			}),
		).
		ServeHTTP(recorder, httptest.NewRequest("GET", "https://google.com", nil))

	b, err := io.ReadAll(recorder.Body)
	if err != nil {
		t.Errorf("reading body: %v", err)
	} else if string(b) != "recovered" {
		t.Errorf("expected body to be 'recovered', got: %s", string(b))
	}
}
