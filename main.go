package main

import (
	"embed"
	"fmt"
	"html/template"
	"math"
	"net/http"
	"os"
	"time"
)

type days struct {
	Days string
}

func main() {
	//go:embed assets/*
	var assets embed.FS

	//go:embed index.html
	var index string

	port := os.Getenv("PORT")
	if port == "" {
		port = "80"
	}

	mux := http.NewServeMux()

	mux.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		tmpl, err := template.New("index").Parse(index)
		if err != nil {
			http.Error(w, "internal server error", http.StatusInternalServerError)
			return
		}
		tmpl.Execute(w, days{Days: getDaysPassed()})
	})

	mux.Handle("/assets/", http.FileServer(http.FS(assets)))

	err := http.ListenAndServe(fmt.Sprintf(":%s", port), mux)
	if err != nil {
		panic(err)
	}
}

func getDaysPassed() string {
	eventTime := os.Getenv("EVENT")
	if eventTime == "" {
		eventTime = "2020-10-14T15:04:05+02:00"
	}

	parsedTime, err := time.Parse(time.RFC3339, eventTime)
	if err != nil {
		parsedTime, _ = time.Parse(time.RFC3339, "2020-10-14T15:04:05+02:00")
	}

	return fmt.Sprintf("%.0f", math.Floor(time.Now().Sub(parsedTime).Hours()/24))
}
