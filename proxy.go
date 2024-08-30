package proxy

import (
	"net/http"
)

func HTTP(w http.ResponseWriter, r *http.Request, c *http.Client) error {
	if c == nil {
		c = http.DefaultClient
	}
	return handleHTTP(w, r, c)
}

// WS note: if request URL contain domain name (not ip address) than
// add suffix :http/:https (for func net.Dial)
func WS(w http.ResponseWriter, r *http.Request) error {
	return handleWS(w, r)
}
