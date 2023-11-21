package main

import (
	"fmt"
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

 fmt.Println("hihi")

 var webhookModule solanum.Module
 webhookUri:="/webhook"
 webhookModule, _ = webhook.NewWebHookModule(
	 server.GetGinEngine().Group(webhookUri),
	 webhookUri,
 )

 server.AddModule(&webhookModule)

   server.Run()
}