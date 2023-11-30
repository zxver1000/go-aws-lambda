package main

import (
	"webhook/webhook"

	"github.com/aws/aws-lambda-go/events"
	"github.com/aws/aws-lambda-go/lambda"
	ginadapter "github.com/awslabs/aws-lambda-go-api-proxy/gin"
	"github.com/gin-gonic/gin"

	"github.com/annuums/solanum"
)
var ginLambda *ginadapter.GinLambda

func handler(request events.APIGatewayProxyRequest) (events.APIGatewayProxyResponse, error) {
	if ginLambda == nil {
		/*
        env_err := godotenv.Load(".env")
    	if env_err != nil {
        log.Fatal("Error loading .env file")
    	}
		*/
		
		server := *solanum.NewSolanum(8080)
		server.GetGinEngine().Use(gin.Logger(), gin.Recovery())

		solanum.NewHelloWorldHandler()

		var webhookModule solanum.Module
		webhookUri:="/"
		webhookModule, _ = webhook.NewWebHookModule(
			server.GetGinEngine().Group(webhookUri),
			webhookUri,
		)
	   
		server.AddModule(&webhookModule)
		server.InitModules()

		ginLambda = ginadapter.New(server.GetGinEngine())
	}

	return ginLambda.Proxy(request)
}


func main(){
	
   lambda.Start(handler)
}