package routers

import (
	"assinatura-api/configuration"
	"assinatura-api/models"
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
	valor := query.SelectValorPlano(idPlano, meses, *dynamoClient, log)
	c.IndentedJSON(http.StatusOK, valor)
}

func GetAssinatura(nome string, sobrenome string, c *gin.Context, log configuration.Logfile, dynamoClient *dynamodb.Client) {
	//Vai la no banco pegar o plano de assinatura do cliente
	validade := query.SelectValidadePlano(nome, sobrenome, *dynamoClient, log)
	if validade == "" {
		c.IndentedJSON(http.StatusNotFound, "Cliente nao encontrado")
		return
	}
	//Checar se o plano esta valido
	status := configuration.CheckValidadeAssinante(validade, log)
	c.IndentedJSON(http.StatusOK, status)
}

func PostAssinante(assinante models.Assinante, c *gin.Context, log configuration.Logfile, dynamoClient *dynamodb.Client) {
	//Vai la no banco gravar o plano de assinatura do cliente
	query.InsertAssinante(assinante, *dynamoClient, log)
	c.IndentedJSON(http.StatusOK, "Assinante criado com sucesso")
}

func GetAssinantes(c *gin.Context, log configuration.Logfile, dynamoClient *dynamodb.Client) {
	//Vai la no banco pegar o nome de todos os assinantes
	assinantes := query.SelectAssinantes(*dynamoClient, log)
	assinantesValidos := make([]models.Assinante, 0)
	//Checar qual desses assinantes esta com o plano valido
	for _, assinante := range assinantes {
		validade := assinante.Validade
		status := configuration.CheckValidadeAssinante(validade, log)
		if status == "valido" {
			assinantesValidos = append(assinantesValidos, assinante)
		}
	}
	c.IndentedJSON(http.StatusOK, assinantesValidos)
}
