# Swagger UI

[![GoDoc](https://godoc.org/github.com/lazada/swgui?status.svg)](https://godoc.org/github.com/lazada/swgui)

Package `swgui` (Swagger UI) provide a HTTP handler to serve Swagger UI.
All assets are embedded in GO source code, so just build and run.

## How to use

```go
package main

import (
    "http"

    "github.com/lazada/swgui"
)

func main() {
    http.Handle("/", swgui.NewHandler("Page title", "path/to/swagger.json", "/"))
    http.ListenAndServe(":8080", nil)
}
```

## Run as standalone server

Install `swgui-server`

    go get github.com/lazada/swgui/...

Start server

    swgui-server -port 8080
