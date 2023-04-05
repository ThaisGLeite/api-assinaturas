package query

import (
	"assinatura-api/configuration"
	"assinatura-api/models"
	"context"

	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/feature/dynamodb/attributevalue"
	"github.com/aws/aws-sdk-go-v2/feature/dynamodb/expression"
	"github.com/aws/aws-sdk-go-v2/service/dynamodb"
)

// Pega o valor do plano com base no nome e duraçao
func SelectValorPlano(idPlano string, meses string, dynamoClient dynamodb.Client, log configuration.Logfile) float32 {
	var plano models.Plano

	//Monta a query com o nome do plano e a duraç~ao em meses
	query := expression.And(
		expression.Name("Plano").Equal(expression.Value(idPlano)),
		expression.Name("Duracao").Equal(expression.Value(meses)),
	)

	proj := expression.NamesList(expression.Name("Valor"))

	expr, err := expression.NewBuilder().WithFilter(query).WithProjection(proj).Build()
	configuration.Check(err, log)

	params := &dynamodb.ScanInput{
		ExpressionAttributeNames:  expr.Names(),
		ExpressionAttributeValues: expr.Values(),
		FilterExpression:          expr.Filter(),
		ProjectionExpression:      expr.Projection(),
		TableName:                 aws.String("PlanoAssinatura"),
	}

	// Make the DynamoDB Query API call
	result, err := dynamoClient.Scan(context.TODO(), params)
	configuration.Check(err, log)

	for _, i := range result.Items {
		err = attributevalue.UnmarshalMap(i, &plano)
		configuration.Check(err, log)
	}

	return plano.Valor
}

// Pega a data de validade do plano com base no nome do cliente
func SelectValidadePlano(nome string, sobrenome string, dynamoClient dynamodb.Client, log configuration.Logfile) string {
	//Monta a query com o nome do cliente
	query := expression.And(
		expression.Name("Nome").Equal(expression.Value(nome)),
		expression.Name("SobreNome").Equal(expression.Value(sobrenome)),
	)

	proj := expression.NamesList(expression.Name("Validade"))

	expr, err := expression.NewBuilder().WithFilter(query).WithProjection(proj).Build()
	configuration.Check(err, log)

	params := &dynamodb.ScanInput{
		ExpressionAttributeNames:  expr.Names(),
		ExpressionAttributeValues: expr.Values(),
		FilterExpression:          expr.Filter(),
		ProjectionExpression:      expr.Projection(),
		TableName:                 aws.String("Assinantes"),
	}

	// Make the DynamoDB Query API call
	result, err := dynamoClient.Scan(context.TODO(), params)
	configuration.Check(err, log)
	var assinatura models.Assinante

	for _, i := range result.Items {
		err = attributevalue.UnmarshalMap(i, &assinatura)
		configuration.Check(err, log)
	}

	return assinatura.Validade
}

// Cadastro de um novo cliente
func InsertAssinante(assinante models.Assinante, dynamoClient dynamodb.Client, log configuration.Logfile) {
	item, err := attributevalue.MarshalMap(assinante)
	configuration.Check(err, log)

	input := &dynamodb.PutItemInput{
		Item:      item,
		TableName: aws.String("Assinantes"),
	}

	_, err = dynamoClient.PutItem(context.TODO(), input)
	configuration.Check(err, log)
}

func SelectAssinantes(dynamoClient dynamodb.Client, log configuration.Logfile) []models.Assinante {
	//Setup das informações da query no banco
	input := &dynamodb.ScanInput{
		TableName: aws.String("Assinantes"),
	}
	result, err := dynamoClient.Scan(context.Background(), input)
	configuration.Check(err, log)
	var assinantes []models.Assinante
	err = attributevalue.UnmarshalListOfMaps(result.Items, &assinantes)
	configuration.Check(err, log)
	return assinantes
}
