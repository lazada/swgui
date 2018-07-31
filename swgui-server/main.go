package main

import (
	"flag"
	"fmt"
	"net/http"

	"github.com/swaggest/swgui/v2"
	"github.com/swaggest/swgui/v3"
)

var (
	host    string
	port    uint
	title   string
	version = "v3"
)

func init() {
	flag.StringVar(&host, "host", "", "Host")
	flag.UintVar(&port, "port", 8080, "Port")
	flag.StringVar(&title, "title", "API Document", "Page title")
	flag.StringVar(&version, "version", "v3", "Swagger UI version (v2/v3)")

	flag.Parse()
}

func main() {
	var h http.Handler
	if version != "v3" {
		h = v2.NewHandler(title, "/swagger.json", "/")
	} else {
		h = v3.NewHandler(title, "/swagger.json", "/")
	}
	http.Handle("/", h)
	fmt.Printf("Listening at %s:%d\n", host, port)
	http.ListenAndServe(fmt.Sprintf("%s:%d", host, port), nil)
}
