package main

import (
	"assinatura-api/configuration"
	"assinatura-api/driver"
	"assinatura-api/models"
	"assinatura-api/routers"
	"log"
	"net/http"
	"os"

	"github.com/apex/gateway"
	"github.com/aws/aws-sdk-go-v2/service/dynamodb"
	"github.com/gin-gonic/gin"
)

var (
	dynamoClient *dynamodb.Client
	logs         configuration.Logfile
)

func inLambda() bool {
	if lambdaTaskRoot := os.Getenv("LAMBDA_TASK_ROOT"); lambdaTaskRoot != "" {
		return true
	}
	return false
}

func setupRouter() *gin.Engine {

	appRouter := gin.New()
	appRouter.GET("/", func(ctx *gin.Context) {
		logs.InfoLogger.Println("Servidor Ok")
		routers.ResponseOK(ctx, logs)
	})

	appRouter.GET("/planos/:id/:mes", func(ctx *gin.Context) {
		//Pegar o nome do plano e o mes solicitado no banco
		idPlano := ctx.Param("id")
		meses := ctx.Param("mes")

		routers.GetPlano(idPlano, meses, ctx, logs, dynamoClient)
	})

	appRouter.GET("/assinatura/:nome/:sobrenome", func(ctx *gin.Context) {
		//Pegar o plano de assinatura do cliente
		nome := ctx.Param("nome")
		sobrenome := ctx.Param("sobrenome")
		routers.GetAssinatura(nome, sobrenome, ctx, logs, dynamoClient)
	})

	appRouter.GET("/assinaturas", func(ctx *gin.Context) {
		//Pegar o nome de todos os assinantes
		routers.GetAssinantes(ctx, logs, dynamoClient)
	})

	appRouter.POST("/novaassinatura", func(ctx *gin.Context) {
		//Criar um novo assinante
		var assinante models.Assinante
		err := ctx.BindJSON(&assinante)
		configuration.Check(err, logs)
		routers.PostAssinante(assinante, ctx, logs, dynamoClient)
	})

	return appRouter
}

// Para compilar o binario do sistema usamos:
//
//	GOOS=linux GOARCH=amd64 CGO_ENABLED=0 go build -o assinatura-api .
//
// para criar o zip do projeto comando: zip lambda.zip assinatura-api
func main() {
	InfoLogger := log.New(os.Stdout, " ", log.LstdFlags|log.Lshortfile)
	ErrorLogger := log.New(os.Stdout, " ", log.LstdFlags|log.Lshortfile)

	logs.InfoLogger = *InfoLogger
	logs.ErrorLogger = *ErrorLogger
	var err error
	// chamada de função para a criação da sessao de login com o banco
	dynamoClient, err = driver.ConfigAws()
	//chamada da função para revificar o erro retornado
	configuration.Check(err, logs)

	if inLambda() {

		log.Fatal(gateway.ListenAndServe(":8080", setupRouter()))
	} else {

		log.Fatal(http.ListenAndServe(":8080", setupRouter()))
	}
}
