package middleware

import (
	"bytes"
	"log"
	"net/http"
	"net/http/httptest"
	"reflect"
	"testing"
)

func expect(t *testing.T, a interface{}, b interface{}) {
	if a != b {
		t.Errorf("Expected %v (type %v) - Got %b (type %v)", b, reflect.TypeOf(b), a, reflect.TypeOf(a))
	}
}

func TestRecovery(t *testing.T) {
	buff := bytes.NewBufferString("")
	recorder := httptest.NewRecorder()

	rec := NewRecovery()
	rec.Logger = log.New(buff, "[ak1m1tsu] ", 0)

	h := New()
	h.Use(rec)
	h.UseHandler(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		panic("here is a panic!")
	}))
	h.ServeHTTP(recorder, (*http.Request)(nil))
	expect(t, recorder.Header().Get("Content-Type"), "text/plain; charset=utf-8")
	expect(t, recorder.Code, http.StatusInternalServerError)
}
