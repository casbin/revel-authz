package authz

import (
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/revel/revel"
)

var testFilters = []revel.Filter{
	AuthzFilter,
	func(c *revel.Controller, fc []revel.Filter) {
		c.RenderHTML("OK.")
	},
}

func testRequest(t *testing.T, user string, path string, method string, code int) {
	r, _ := http.NewRequest(method, path, nil)
	r.SetBasicAuth(user, "123")
	w := httptest.NewRecorder()
	c := revel.NewController(revel.NewRequest(r), revel.NewResponse(w))

	testFilters[0](c, testFilters)

	if c.Response.Status != code {
		t.Errorf("%s, %s, %s: %d, supposed to be %d", user, path, method, c.Response.Status, code)
	}
}

func TestBasic(t *testing.T) {
	testRequest(t, "alice", "/dataset1/resource1", "GET", 200)
	testRequest(t, "alice", "/dataset1/resource1", "POST", 200)
	testRequest(t, "alice", "/dataset1/resource2", "GET", 200)
	testRequest(t, "alice", "/dataset1/resource2", "POST", 403)
}

func TestPathWildcard(t *testing.T) {
	testRequest(t, "bob", "/dataset2/resource1", "GET", 200)
	testRequest(t, "bob", "/dataset2/resource1", "POST", 200)
	testRequest(t, "bob", "/dataset2/resource1", "DELETE", 200)
	testRequest(t, "bob", "/dataset2/resource2", "GET", 200)
	testRequest(t, "bob", "/dataset2/resource2", "POST", 403)
	testRequest(t, "bob", "/dataset2/resource2", "DELETE", 403)

	testRequest(t, "bob", "/dataset2/folder1/item1", "GET", 403)
	testRequest(t, "bob", "/dataset2/folder1/item1", "POST", 200)
	testRequest(t, "bob", "/dataset2/folder1/item1", "DELETE", 403)
	testRequest(t, "bob", "/dataset2/folder1/item2", "GET", 403)
	testRequest(t, "bob", "/dataset2/folder1/item2", "POST", 200)
	testRequest(t, "bob", "/dataset2/folder1/item2", "DELETE", 403)
}

func TestRBAC(t *testing.T) {
	// cathy can access all /dataset1/* resources via all methods because it has the dataset1_admin role.
	testRequest(t, "cathy", "/dataset1/item", "GET", 200)
	testRequest(t, "cathy", "/dataset1/item", "POST", 200)
	testRequest(t, "cathy", "/dataset1/item", "DELETE", 200)
	testRequest(t, "cathy", "/dataset2/item", "GET", 403)
	testRequest(t, "cathy", "/dataset2/item", "POST", 403)
	testRequest(t, "cathy", "/dataset2/item", "DELETE", 403)
}
