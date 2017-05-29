package authz

import (
	"github.com/casbin/casbin"
	"github.com/revel/revel"
)

// AuthzFilter enables the authorization based on Casbin.
//
// Usage:
//  1) Add `authz.AuthzFilter` to the app's filters (it must come after the authentication).
//  2) Init the Casbin enforcer.
func AuthzFilter(c *revel.Controller, fc []revel.Filter) {
	e := casbin.NewEnforcer("authz_model.conf", "authz_policy.csv")

	if !CheckPermission(e, c.Request) {
		c.Result = c.Forbidden("Access denied by the Authz plugin.")
		return
	} else {
		fc[0](c, fc[1:])
	}
}

// GetUserName gets the user name from the request.
// Currently, only HTTP basic authentication is supported
func GetUserName(r *revel.Request) string {
	username, _, _ := r.BasicAuth()
	return username
}

// CheckPermission checks the user/method/path combination from the request.
// Returns true (permission granted) or false (permission forbidden)
func CheckPermission(e *casbin.Enforcer, r *revel.Request) bool {
	user := GetUserName(r)
	method := r.Method
	path := r.URL.Path
	return e.Enforce(user, path, method)
}
