package routes

import (
	"github.com/MarceloPetrucio/go-scalar-api-reference"
	"github.com/gin-gonic/gin"
	"net/http"
)

func DocsRoutes(r *gin.Engine) {
	r.GET("/docs", func(c *gin.Context) {
		htmlContent, err := scalar.ApiReferenceHTML(&scalar.Options{
			SpecURL: "./docs/swagger.json",
			CustomOptions: scalar.CustomOptions{
				PageTitle: "ZorpPay API Documentation",
			},
			DarkMode: true,
		})

		if err != nil {
			c.String(http.StatusInternalServerError, err.Error())
		}

		c.Data(http.StatusOK, "text/html; charset=utf-8", []byte(htmlContent))
	})
}
