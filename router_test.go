// Copyright 2022 Axel Etcheverry. All rights reserved.
// Use of this source code is governed by a MIT
// license that can be found in the LICENSE file.

package server

import (
	"testing"

	"github.com/stretchr/testify/mock"
)

func TestRouterAddController(t *testing.T) {
	router := New()

	controllerMock := &MockController{}

	controllerMock.On("Mount", mock.Anything)

	router.AddController(controllerMock)

	controllerMock.AssertExpectations(t)
}
