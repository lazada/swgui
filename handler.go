package swgui

import (
	"html/template"
	"net/http"
)

var staticServer = http.FileServer(assetFS())

// Handler handle swagger UI request
type Handler struct {
	Title       string // title of index file
	SwaggerJSON string // path to swagger.json document specification
	BasePath    string // base path to docs

	JsonEditor     bool // Enable visual json editor support (experimental, can fail with complex schemas)

	tpl          *template.Template
	staticServer http.Handler
}

// NewHandler returns a HTTP handler for swagger UI
func NewHandler(title, swaggerJSONPath string, basePath string) *Handler {
	hdl := &Handler{
		Title:       title,
		SwaggerJSON: swaggerJSONPath,
		BasePath:    basePath,
	}

	hdl.tpl, _ = template.New("index").Parse(indexTpl)
	hdl.staticServer = http.StripPrefix(basePath, staticServer)
	return hdl
}

// NewHandlerWithConfig returns a HTTP handler for swagger UI
func NewHandlerWithConfig(handler Handler) *Handler {
	hdl := &handler
	hdl.tpl, _ = template.New("index").Parse(indexTpl)
	hdl.staticServer = http.StripPrefix(handler.BasePath, staticServer)
	return hdl
}

// ServeHTTP implement http.Handler interface, to handle swagger UI request
func (h *Handler) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	if r.URL.Path != h.BasePath {
		h.staticServer.ServeHTTP(w, r)
		return
	}

	if err := h.tpl.Execute(w, h); err != nil {
		http.NotFound(w, r)
	}
}
