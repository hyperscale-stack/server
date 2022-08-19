// Copyright 2022 Axel Etcheverry. All rights reserved.
// Use of this source code is governed by a MIT
// license that can be found in the LICENSE file.

package server

// Controller interface
//
//go:generate go run -mod=mod github.com/vektra/mockery/v2 --inpackage --case underscore --name=Controller
type Controller interface {
	Mount(r *Router)
}
