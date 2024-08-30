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

	hj, ok := w.(http.Hijacker)
	if !ok {
		return errors.New("not a hijacker")
	}

	s, _, err := hj.Hijack()
	if err != nil {
		return
	}
	defer s.Close()

	if err = r.Write(d); err != nil {
		return
	}

	c := make(chan error, 2)
	f := func(dst io.Writer, src io.Reader) {
		_, err := io.Copy(dst, src)
		c <- err
	}
	go f(s, d)
	go f(d, s)
	return <-c
}
