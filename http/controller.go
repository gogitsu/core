package http

import "github.com/gogitsu/core/http/router"

// Controller define a controller interface.
type Controller interface {
	Route(r router.Router) Controller
}
