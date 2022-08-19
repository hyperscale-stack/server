// Copyright 2022 Axel Etcheverry. All rights reserved.
// Use of this source code is governed by a MIT
// license that can be found in the LICENSE file.

package server

import (
	"github.com/gofiber/fiber/v2"
)

// Router struct
type Router struct {
	*fiber.App
}

// New constructor
func New(config ...fiber.Config) *Router {
	return &Router{
		App: fiber.New(config...),
	}
}

// AddController to Router
func (r *Router) AddController(controller Controller) {
	controller.Mount(r)
}
