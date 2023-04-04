package routers

import (
	"assinatura-api/configuration"
	"net/http"

	"github.com/aws/aws-sdk-go-v2/service/dynamodb"
	"github.com/gin-gonic/gin"
)

func ResponseOK(c *gin.Context, log configuration.Logfile) {
	c.IndentedJSON(http.StatusOK, "Servidor up")
}

func GetPlano(idPlano string, meses string, c *gin.Context, log configuration.Logfile, dynamoClient *dynamodb.Client) {

}
