package proxy

import (
	"net/http"
	"strings"
)

func HTTP(w http.ResponseWriter, r *http.Request, c *http.Client) error {
	if c == nil {
		c = http.DefaultClient
	}
	return handleHTTP(w, r, c)
}

func WS(w http.ResponseWriter, r *http.Request) error {
	if !strings.HasSuffix(r.Host, ":http") {
		r.Host = r.Host + ":http"
	}
	return handleWS(w, r)
}
