package http

type (
	// Router .
	Router interface {
		BasePath() string
		SetBasePath(bp string)
	}

	// BaseRouter .
	BaseRouter struct {
		basePath string
	}
)

// BasePath .
func (br *BaseRouter) BasePath() string {
	return br.basePath
}

// SetBasePath .
func (br *BaseRouter) SetBasePath(bp string) {
	br.basePath = bp
}
