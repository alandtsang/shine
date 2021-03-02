package tpl

func RouterTemplate() []byte {
	return []byte(`// Package router is responsible for initializing the router.
/*
{{ .Copyright }}
{{ if .License.Header }}{{ .License.Header }}{{ end }}
*/
package router

import (
	"net/http"
	"time"

	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"

    "{{ .PkgName }}/internal/router/routes"
    "{{ .PkgName }}/internal/validator"
)

// Router is responsible for managing api routing.
type Router struct {
	e *echo.Echo
}

// New return a new router instance.
func New() *Router {
	e := echo.New()
	e.Validator = validator.New()

	// Middleware
	e.Use(middleware.Logger())
	e.Use(middleware.Recover())

	return &Router{
		e: e,
	}
}

// ListenAndServe starts http web server.
func (r *Router) ListenAndServe(addr string) error {
    routes.SetRoutes(r.e)
	server := &http.Server{
		Addr:         addr,
		ReadTimeout:  30 * time.Second,
		WriteTimeout: 30 * time.Second,
	}
	return r.e.StartServer(server)
}
`)
}

func RoutesTemplate() []byte {
	return []byte(`// Package routes handles http routing.
/*
{{ .Copyright }}
{{ if .License.Header }}{{ .License.Header }}{{ end }}
*/
package routes

import (
    "net/http"

	"github.com/labstack/echo/v4"
)

// SetRoutes sets all routes.
func SetRoutes(e *echo.Echo) {
    e.GET("", hello)
}

func hello(c echo.Context) error {
    return c.JSON(http.StatusOK, "hello")
}
`)
}
