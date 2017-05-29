Revel-authz [![GoDoc](https://godoc.org/github.com/casbin/revel-authz?status.svg)](https://godoc.org/github.com/casbin/revel-authz)
======

Revel-authz is an authorization middleware for [Revel](https://github.com/revel/revel), it's based on [https://github.com/casbin/casbin](https://github.com/casbin/casbin).

## Installation

    go get github.com/casbin/revel-authz

## Simple Example

```Go
package main

import (
	"net/http"
	"net/http/httptest"

    "github.com/casbin/casbin"
	"github.com/casbin/revel-authz"
	"github.com/revel/revel"
)

var testFilters = []revel.Filter{
	authz.AuthzFilter,
	func(c *revel.Controller, fc []revel.Filter) {
		c.RenderHTML("OK.")
	},
}

func main() {
	r, _ := http.NewRequest("GET", "/dataset1/resource1", nil)
    	r.SetBasicAuth("alice", "123")
    	w := httptest.NewRecorder()
    	c := revel.NewController(revel.NewRequest(r), revel.NewResponse(w))
    
    	testFilters[0](c, testFilters)
}
```

## Documentation

The authorization determines a request based on ``{subject, object, action}``, which means what ``subject`` can perform what ``action`` on what ``object``. In this plugin, the meanings are:

1. ``subject``: the logged-on user name
2. ``object``: the URL path for the web resource like "dataset1/item1"
3. ``action``: HTTP method like GET, POST, PUT, DELETE, or the high-level actions you defined like "read-file", "write-blog"


For how to write authorization policy and other details, please refer to [the Casbin's documentation](https://github.com/casbin/casbin).

## Getting Help

- [Casbin](https://github.com/casbin/casbin)

## License

This project is under MIT License. See the [LICENSE](LICENSE) file for the full license text.
