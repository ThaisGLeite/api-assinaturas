package routers

import (
	"assinatura-api/configuration"
	"net/http"

	"github.com/gin-gonic/gin"
)

func ResponseOK(c *gin.Context, app configuration.Logfile) {
	c.IndentedJSON(http.StatusOK, "Servidor up")
}
