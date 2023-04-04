package routers

import (
	"assinatura-api/configuration"
	"assinatura-api/query"
	"net/http"

	"github.com/aws/aws-sdk-go-v2/service/dynamodb"
	"github.com/gin-gonic/gin"
)

func ResponseOK(c *gin.Context, log configuration.Logfile) {
	c.IndentedJSON(http.StatusOK, "Servidor up")
}

func GetPlano(idPlano string, meses string, c *gin.Context, log configuration.Logfile, dynamoClient *dynamodb.Client) {
	//Vai la no banco pegar o valor do plano nessa quantidade de meses
	valor := query.SelectPunch(idPlano, meses, *dynamoClient, log)
	c.IndentedJSON(http.StatusOK, valor)
}
