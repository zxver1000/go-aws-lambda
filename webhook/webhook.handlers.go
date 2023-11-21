package webhook

import (
	"log"
	"net/http"

	"fmt"

	"github.com/annuums/solanum"
	"github.com/gin-gonic/gin"
)
type UpdateImageRequestBody struct {
    Func_name string `json:"function_name" binding:"required"`
	Image_uri string `json:"image_uri" binding:"required"`
}

func NewWebHookHandler() *solanum.SolaService {
	return &solanum.SolaService{
		Uri:        "/",
		Method:     http.MethodPatch,
		Handler:    Update_LambdaHandler,
		Middleware: indexMiddleware3,
	}
}




func Update_LambdaHandler(ctx *gin.Context) {
	
	/* Required
	1. Repository name
	2. image_uri
	*/
   	var requestBody UpdateImageRequestBody

	
	   if err := ctx.BindJSON(&requestBody); err != nil {
		   // DO SOMETHING WITH THE ERROR
		 
		   fmt.Println("ERROR");
		   fmt.Println(err)  
		   ctx.JSON(404,"BAD REQEUST")
		   return
	   }
	

	result:=Update_lambda_image(requestBody.Func_name,requestBody.Image_uri)
	if result ==200{
	ctx.JSON(200, "Your Function Udpate complete.")
	}
	if result >=400{
      ctx.JSON(500,"Server error")
	}

}

func indexMiddleware3(ctx *gin.Context) {
	log.Println("Hello Index Middleware")
	ctx.Next()
}
