package routers

import (
	"assinatura-api/configuration"
	"assinatura-api/models"
	"assinatura-api/query"
	"net/http"
	"time"

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

	status := "valido"

	date, err := time.Parse("2006-01-02", validade)
	configuration.Check(err, log)

	agora := time.Now().UTC()

	if date.Before(agora) {
		status = "invalido"
	}

	c.IndentedJSON(http.StatusOK, status)
}

func PostAssinante(nome string, sobrenome string, plano string, validade string, c *gin.Context, log configuration.Logfile, dynamoClient *dynamodb.Client) {
	//Criar um assinante novo
	assinante := models.Assinante{
		Nome:       nome,
		SobreNome:  sobrenome,
		Plano:      plano,
		Validade:   validade,
		DataInicio: time.Now().UTC().Format("2006-01-02"),
	}

	//Vai la no banco gravar o plano de assinatura do cliente
	query.InsertAssinante(assinante, *dynamoClient, log)
	c.IndentedJSON(http.StatusOK, "Assinante criado com sucesso")
}
