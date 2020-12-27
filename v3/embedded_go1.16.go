// +build go1.16,!swguicdn

package v3

import (
	"github.com/swaggest/swgui/v3/static"
	"github.com/vearutop/statigz"
)

var staticServer = statigz.FileServer(static.FS)

const (
	assetsBase  = "{{ .BasePath }}"
	faviconBase = "{{ .BasePath }}"
)
