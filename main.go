package main

import (
	"io"
	"net/http"

	"github.com/gin-gonic/gin"
)

func main() {
	r := gin.Default()

	r.Any("/", func(c *gin.Context) {
		body, _ := io.ReadAll(c.Request.Body)
		defer c.Request.Body.Close()

		c.String(http.StatusOK, "%s", string(body))
	})

	r.Run(":3000")
}
