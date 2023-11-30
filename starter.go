package main

import (
	"log"
	"webhook/webhook"

	"github.com/annuums/solanum"
	"github.com/joho/godotenv"
)

func main(){
	env_err := godotenv.Load(".env")
    if env_err != nil {
        log.Fatal("Error loading .env file")
    }
server := *solanum.NewSolanum(5050)

 var webhookModule solanum.Module
 webhookUri:="/"
 webhookModule, _ = webhook.NewWebHookModule(
	 server.GetGinEngine().Group(webhookUri),
	 webhookUri,
 )

 server.AddModule(&webhookModule)
 
   server.Run()
}