package webhook

import (
	"github.com/annuums/solanum"
)
var WebHookController *solanum.SolaController

func NewWebHookController() (*solanum.SolaController, error) {
	if WebHookController == nil {
		WebHookController, _ = solanum.NewController()
		addWebhookHandlers()
	}

	return WebHookController, nil
}

func addWebhookHandlers() {
	WebHookHandler := NewWebHookHandler()
    
	// anotherHandler := NewHelloWorldHandler()
	//* ...
	WebHookController.AddHandler(WebHookHandler...)
	
}
