package solanum

import (
	"log"
	"net/http"

	_ "github.com/annuums/solanum"
	"github.com/gin-gonic/gin"
)

func NewWebHookHandler() *SolaService {
	return &SolaService{
		Uri:        "/",
		Method:     http.MethodPost,
		Handler:    indexHandler3,
		Middleware: indexMiddleware3,
	}
}




func indexHandler3(ctx *gin.Context) {
	ctx.JSON(200, "WebHook Handler.")
}

func indexMiddleware3(ctx *gin.Context) {
	log.Println("Hello Index Middleware")
	ctx.Next()
}
