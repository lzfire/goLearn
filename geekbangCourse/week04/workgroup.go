package main

import (
	"context"
	"fmt"
	"net/http"
)

func server(addr string, handler http.Handler, stop <-chan struct{}) error {
	s := http.Server{
		Addr:    addr,
		Handler: handler,
	}
	go func() {
		<-stop
		s.Shutdown(context.Background())
	}()
	return s.ListenAndServe()

}
func main() {
	done := make(chan error, 2)
	stop := make(chan struct{})
	mux := http.NewServeMux()
	mux.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		fmt.Fprintln(w, "Hello lzfire")
	})
	go func() {
		done <- server("0.0.0.0:8080", mux, stop)
	}()
	go func() {
		done <- server("127.0.0.1:8081", http.DefaultServeMux, stop)
	}()

	var stopped bool
	for i := 0; i < cap(done); i++ {
		if err := <-done; err != nil {
			fmt.Printf("err:%v\n", err)
		}
		if !stopped {
			stopped = true
			close(stop)
		}
	}
}
