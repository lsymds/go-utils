package http_api

import (
	"encoding/json"
	"io"
	"net/http"
	"net/http/httptest"
	"testing"
)

type response struct {
	Name string
}

func readBodyResponse(r io.Reader) *response {
	var rs response

	json.NewDecoder(r).Decode(&rs)

	return &rs
}

func TestResponseMethodsReturnCorrectResponses(t *testing.T) {
	er := response{Name: "Bob Smith"}

	tt := []struct {
		name   string
		method func(http.ResponseWriter, interface{})
		status int
	}{
		{name: "ok", method: ok, status: 200},
		{name: "notAuthorized", method: notAuthorized, status: 401},
		{name: "forbidden", method: forbidden, status: 403},
		{name: "notFound", method: notFound, status: 404},
		{name: "internalServerError", method: internalServerError, status: 500},
	}

	for _, test := range tt {
		t.Run(test.name, func(t *testing.T) {
			w := httptest.NewRecorder()

			test.method(w, er)

			if w.Code != test.status {
				t.Errorf("expected status %d, got: %d", test.status, w.Code)
			} else if r := readBodyResponse(w.Body); r.Name != "Bob Smith" {
				t.Errorf("expected name 'Bob Smith', got: %s", r.Name)
			}
		})
	}
}
