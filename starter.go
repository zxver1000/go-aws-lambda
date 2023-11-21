package main

import (
	"fmt"

	"webhook"

	"github.com/annuums/solanum"
)

func main(){

	server := *solanum.NewSolanum(5050)

 fmt.Println("hihi")

 var webhookModule solanum.Module
 webhookUri:="/webhook"
 webhookModule, _ = webhook.NewWebHookModule(
	 server.GetGinEngine().Group(webhookUri),
	 webhookUri,
 )

 server.AddModule(&webhookModule)


 server.AddModule(&helloWorldModule)
   server.Run()
}