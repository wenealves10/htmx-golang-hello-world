package main

import (
	"net/http"

	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
	"github.com/wenealves10/htmx-golang-hello-world/template"
	"golang.org/x/time/rate"
)

func main() {
	e := echo.New()
	e.Pre(middleware.RemoveTrailingSlash())
	e.Use(middleware.Recover())
	e.Use(middleware.RateLimiter(middleware.NewRateLimiterMemoryStore(
		rate.Limit(20),
	)))

	template.NewTemplateRenderer(e, "public/*.html")

	e.GET("/", func(c echo.Context) error {
		return c.Render(http.StatusOK, "index", nil)
	})

	e.GET("/get-info", func(c echo.Context) error {
		res := map[string]interface{}{
			"Name":       "Wene Alves",
			"Age":        22,
			"Profession": "Software Engineer",
			"Phone":      "+55 11 9 9999-9999",
			"Email":      "wenealves@pro.com.br",
		}

		return c.Render(http.StatusOK, "name_card", res)
	})

	e.Logger.Fatal(e.Start(":3000"))
}
