package proxy

import (
	"io"
	"net/http"
)

func handleHTTP(w http.ResponseWriter, r *http.Request, c *http.Client) (err error) {
	resp, err := c.Do(r)
	if err != nil {
		return err
	}

	defer resp.Body.Close()
	copyHeader(w.Header(), resp.Header)
	w.WriteHeader(resp.StatusCode)
	io.Copy(w, resp.Body)
	return
}

func copyHeader(dst, src http.Header) {
	for k, v1 := range src {
		for _, v2 := range v1 {
			dst.Add(k, v2)
		}
	}
}
