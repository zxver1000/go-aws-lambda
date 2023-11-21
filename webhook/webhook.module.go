package webhook

import (
	"log"

	"github.com/annuums/solanum"
	"github.com/gin-gonic/gin"
)

var WebHookModule *solanum.SolaModule

func NewWebHookModule(router *gin.RouterGroup, uri string) (*solanum.SolaModule, error) {
	if WebHookModule == nil {
		WebHookModule, _ = solanum.NewModule(router, uri)
		attachWebhookControllers()
	}

	return WebHookModule, nil
}

func attachWebhookControllers() {
	//* Attatching Controller Directly
	var (
		ctr solanum.Controller
		err error
	)
	ctr, err = NewWebHookController()

	if err != nil {
		log.Fatal(err)
	}
	// ctr2, _ := NewAnotherController()
	//	...

	WebHookModule.SetControllers(&ctr)
}
