# Swagger UI

[![GoDoc](https://godoc.org/github.com/swaggest/swgui?status.svg)](https://godoc.org/github.com/swaggest/swgui)

Package `swgui` (Swagger UI) provides HTTP handler to serve Swagger UI.
All assets are embedded in Go source code, so just build and run.

Static assets for `v3` are built from Swagger UI [v3.20.7](https://github.com/swagger-api/swagger-ui/releases/tag/v3.20.7).

## How to use

```go
package main

import (
    "net/http"

    "github.com/swaggest/swgui/v3"
)

func main() {
    http.Handle("/", v3.NewHandler("My API", "/swagger.json", "/"))
    http.ListenAndServe(":8080", nil)
}
```

## Run as standalone server

Install `swgui-server`

    go get github.com/swaggest/swgui/...

Start server

    swgui-server -port 8080
