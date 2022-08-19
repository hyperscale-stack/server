# Hyperscale Server [![Last release](https://img.shields.io/github/release/hyperscale-stack/server.svg)](https://github.com/hyperscale-stack/server/releases/latest) [![Documentation](https://godoc.org/github.com/hyperscale-stack/server?status.svg)](https://godoc.org/github.com/hyperscale-stack/server)

[![Go Report Card](https://goreportcard.com/badge/github.com/hyperscale-stack/server)](https://goreportcard.com/report/github.com/hyperscale-stack/server)

| Branch | Status                                                                                                                                                                     | Coverage                                                                                                                                               |
| ------ | -------------------------------------------------------------------------------------------------------------------------------------------------------------------------- | ------------------------------------------------------------------------------------------------------------------------------------------------------ |
| master | [![Build Status](https://github.com/hyperscale-stack/server/workflows/Go/badge.svg?branch=master)](https://github.com/hyperscale-stack/server/actions?query=workflow%3AGo) | [![Coveralls](https://img.shields.io/coveralls/hyperscale-stack/server/master.svg)](https://coveralls.io/github/hyperscale-stack/server?branch=master) |

The Hyperscale server library provides a simple server over Fiber.

## Example

```go
package main

import (
    "github.com/hyperscale-stack/server"
)

func main() {
    app := server.New()

    app.AddController(&myController{})

    app.Listen(":3000")
}

```

## License

Hyperscale Server is licensed under [the MIT license](LICENSE.md).
