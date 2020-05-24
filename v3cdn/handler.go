package v3cdn

import (
	"encoding/json"
	"html/template"
	"net/http"

	"github.com/swaggest/swgui"
)

// Handler handle swagger UI request
type Handler struct {
	swgui.Config

	ConfigJson template.JS

	tpl *template.Template
}

// NewHandler returns a HTTP handler for swagger UI
func NewHandler(title, swaggerJSONPath string, basePath string) *Handler {
	h := &Handler{}
	h.Title = title
	h.SwaggerJSON = swaggerJSONPath
	h.BasePath = basePath

	j, _ := json.Marshal(h.Config)
	h.ConfigJson = template.JS(j)
	h.tpl, _ = template.New("index").Parse(indexTpl)
	return h
}

// NewHandlerWithConfig returns a HTTP handler for swagger UI
func NewHandlerWithConfig(config swgui.Config) *Handler {
	h := &Handler{
		Config: config,
	}
	j, _ := json.Marshal(h.Config)
	h.ConfigJson = template.JS(j)
	h.tpl, _ = template.New("index").Parse(indexTpl)
	return h
}

// ServeHTTP implement http.Handler interface, to handle swagger UI request
func (h *Handler) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "text/html")
	if err := h.tpl.Execute(w, h); err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
	}
}
