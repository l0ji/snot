package main

import (
	"flag"
	"fmt"
	"html/template"
	"log/slog"
	"net/http"
	"os"
	"snot/config"
)

type IndexData struct {
	Test string
}

func killProgram(err error) {
	_, _ = fmt.Fprintln(os.Stderr, err.Error())
	os.Exit(1)
}

func main() {
	//configFile := flag.String("config", "conf.ini", "Configuration file")
	flag.Parse()

	conf := config.New()
	err := conf.IsValid()
	if err != nil {
		killProgram(err)
	}
	slog.Debug("Parsed config:", conf)

	mux := http.NewServeMux()
	tmpl, err := template.ParseFiles("templates/base.tmpl")
	if err != nil {
		killProgram(err)
	}
	mux.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		data := IndexData{"yes!"}
		err := tmpl.Execute(w, data)
		if err != nil {
			return
		}
	})
	slog.Info("Listening on " + "http://" + conf.ListenAddr)
	killProgram(http.ListenAndServe(conf.ListenAddr, mux))
}
