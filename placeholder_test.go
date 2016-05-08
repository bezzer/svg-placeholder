package main

import (
  "fmt"
  "log"
  "testing"
  "net/http"
  "net/http/httptest"
)

func TestSvg(t *testing.T) {
  req, err := http.NewRequest("GET", "http://example.com/svg/100/100", nil)
	if err != nil {
		log.Fatal(err)
	}
  w := httptest.NewRecorder()
  svg(w, req)

  fmt.Printf("%d - %s", w.Code, w.Body.String())
}
