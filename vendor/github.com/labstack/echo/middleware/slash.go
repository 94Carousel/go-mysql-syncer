package middleware

import (
	"github.com/labstack/echo"
)

type (
	// TrailingSlashConfig defines the config for TrailingSlash middleware.
	TrailingSlashConfig struct {
		// Status code to be used when redirecting the request.
		// Optional, but when provided the request is redirected using this code.
		RedirectCode int `json:"redirect_code"`
	}
)

// AddTrailingSlash returns a root level (before router) middleware which adds a
// trailing slash to the request `URL#Path`.
//
// Usage `Echo#Pre(AddTrailingSlash())`
func AddTrailingSlash() echo.MiddlewareFunc {
	return AddTrailingSlashWithConfig(TrailingSlashConfig{})
}

// AddTrailingSlashWithConfig returns a AddTrailingSlash middleware from config.
// See `AddTrailingSlash()`.
func AddTrailingSlashWithConfig(config TrailingSlashConfig) echo.MiddlewareFunc {
	return func(next echo.HandlerFunc) echo.HandlerFunc {
		return func(c echo.Context) error {
			req := c.Request()
			url := req.URL()
			path := url.Path()
			qs := url.QueryString()
			if path != "/" && path[len(path)-1] != '/' {
				path += "/"
				uri := path
				if qs != "" {
					uri += "?" + qs
				}

				// Redirect
				if config.RedirectCode != 0 {
					return c.Redirect(config.RedirectCode, uri)
				}

				// Forward
				req.SetURI(uri)
				url.SetPath(path)
			}
			return next(c)
		}
	}
}

// RemoveTrailingSlash returns a root level (before router) middleware which removes
// a trailing slash from the request URI.
//
// Usage `Echo#Pre(RemoveTrailingSlash())`
func RemoveTrailingSlash() echo.MiddlewareFunc {
	return RemoveTrailingSlashWithConfig(TrailingSlashConfig{})
}

// RemoveTrailingSlashWithConfig returns a RemoveTrailingSlash middleware from config.
// See `RemoveTrailingSlash()`.
func RemoveTrailingSlashWithConfig(config TrailingSlashConfig) echo.MiddlewareFunc {
	return func(next echo.HandlerFunc) echo.HandlerFunc {
		return func(c echo.Context) error {
			req := c.Request()
			url := req.URL()
			path := url.Path()
			qs := url.QueryString()
			l := len(path) - 1
			if l >= 0 && path != "/" && path[l] == '/' {
				path = path[:l]
				uri := path
				if qs != "" {
					uri += "?" + qs
				}

				// Redirect
				if config.RedirectCode != 0 {
					return c.Redirect(config.RedirectCode, uri)
				}

				// Forward
				req.SetURI(uri)
				url.SetPath(path)
			}
			return next(c)
		}
	}
}
