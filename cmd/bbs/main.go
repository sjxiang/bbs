package main

import (
	"errors"
	"flag"
	"fmt"
	"log"
	"net/http"
	"os"

	"github.com/sjxiang/bbs/web"
)


func main() {
	if err := run(); err != nil {
		_, _ = fmt.Fprintln(os.Stderr, err)
		os.Exit(1)
	}
}

func run() error {

	var addr string

	fs := flag.NewFlagSet("bbs", flag.ExitOnError)
	fs.StringVar(&addr, "addr", ":4000", "HTTP service address")
	
	if err := fs.Parse(os.Args[1:]); err != nil {
		return fmt.Errorf("parse flags: %w", err)
	}

	logger := log.New(os.Stderr, "", log.Ldate|log.Ltime|log.Llongfile)

	handler := &web.Handler{
		Logger: *logger,
	}

	srv := &http.Server{
		Handler: handler,
		Addr:    addr,
	}
	defer srv.Close()

	logger.Printf("listening on %s\n", addr)

	err := srv.ListenAndServe()
	if err != nil && !errors.Is(err, http.ErrServerClosed) {
		return fmt.Errorf("listen and serve: %w", err)
	}

	return nil
}