// Copyright 2020 Luca Stasio. All rights reserved.
// Use of this source code is governed by a MIT
// license that can be found in the LICENSE file.

// Package core implements core components of gogitsu lib.
//
// controller.go defines the Controller interface and the
// base controller implementation struct from which derive concrete
// controllers into apps.
package core

// Controller .
type Controller interface {
	BasePath() string
	Route(r Router)
}

// BaseController .
type BaseController struct {
	Path string
}

// NewBaseController returns a new BaseController instance
// useful in composition.
func NewBaseController(path string) BaseController {
	return BaseController{Path: path}
}

// BasePath .
func (bc BaseController) BasePath() string {
	return bc.Path
}

// Route .
func (bc BaseController) Route(r Router) {
}
