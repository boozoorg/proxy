package proxy

import (
	"errors"
	"io"
	"net"
	"net/http"
)

func handleWS(w http.ResponseWriter, r *http.Request) (err error) {
	d, err := net.Dial("tcp", r.Host)
	if err != nil {
		return
	}
	defer d.Close()

	h, ok := w.(http.Hijacker)
	if !ok {
		return errors.New("not an hijacker")
	}

	hj, _, err := h.Hijack()
	if err != nil {
		return
	}
	defer hj.Close()

	if err = r.Write(d); err != nil {
		return
	}

	c := make(chan error, 2)
	f := func(w io.Writer, r io.Reader) {
		_, err := io.Copy(w, r)
		c <- err
	}
	go f(d, hj)
	go f(hj, d)
	err = <-c
	return
}
