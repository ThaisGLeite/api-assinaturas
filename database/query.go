package query

import (
	"assinatura-api/configuration"
	"context"

	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/feature/dynamodb/attributevalue"
	"github.com/aws/aws-sdk-go-v2/feature/dynamodb/expression"
	"github.com/aws/aws-sdk-go-v2/service/dynamodb"
)

// Pega o valor do plano com base no nome e duraçao
func SelectPunch(idPlano string, meses string, dynamoClient dynamodb.Client, log configuration.Logfile) float32 {
	var valorPlano float32

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
		err = attributevalue.UnmarshalMap(i, &valorPlano)
		configuration.Check(err, log)
	}

	return valorPlano
}
