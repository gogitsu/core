package core

// Controller .
type Controller interface {
	BasePath() string
	Route(r Router)
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
func (bc BaseController) Route(r Router) {
}
