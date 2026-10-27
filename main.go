package main

import (
	"flag"
	"fmt"
	"log/slog"
	"net/http"
	"os"
	"snot/config"
	"snot/web"
)

func killProgram(err error) {
	_, _ = fmt.Fprintln(os.Stderr, err.Error())
	os.Exit(1)
}

func main() {
	flag.Parse() // add flags later

	conf := config.New()
	err := conf.IsValid()
	if err != nil {
		killProgram(err)
	}
	slog.Debug("Parsed config:", conf)

	renderer, err := web.NewRenderer()
	if err != nil {
		killProgram(err)
	}

	mux := http.NewServeMux()

	mux.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		renderer.Home(conf, w, r)
	})
	mux.Handle("/static/", http.StripPrefix("/static/", http.FileServer(http.Dir("static"))))

	slog.Info("Listening on " + "http://" + conf.ListenAddr)
	killProgram(http.ListenAndServe(conf.ListenAddr, mux))
}
