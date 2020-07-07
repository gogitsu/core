package core

/*
// register the handlers or handler_funcs by name in a map:
handlerMap := make(map[string]*http.Handler)
// OR
handlerFuncMap := make(map[string]func(http.ResponseWriter, *http.Request))

handlerMap["myHandler"] = myHandler

// now you can iterate over you config values and assign them to a router
for path, handler := range routes {
    myRouter.Handler(path, handlerMap[handler])
}
*/
type Controller interface {
	BasePath() string
	Route(r *Router)
}

// BaseController .
type BaseController struct {
	basePath string
}

// BasePath .
func (bc BaseController) BasePath() string {
	return bc.basePath
}

// Route .
func (bc BaseController) Route(r *Router) {

}
