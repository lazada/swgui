// +build swguicdn

package v3

import "net/http"

var staticServer http.Handler

const (
	assetsBase  = "https://cdnjs.cloudflare.com/ajax/libs/swagger-ui/3.35.1/"
	faviconBase = "https://petstore.swagger.io/"
)
