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

func NewWebHookHandler() []*solanum.SolaService {
	return []*solanum.SolaService{
	{
		Uri:        "/",
		Method:     http.MethodPatch,
		Handler:    Update_LambdaHandler,
		Middleware: indexMiddleware,
	},
	{
		Uri:        "/",
		Method:     http.MethodGet,
		Handler:    HealthCheckHandler,
		Middleware: indexMiddleware,
	},
	}
}



func HealthCheckHandler(ctx *gin.Context){
   
	
	ctx.JSON(http.StatusOK,gin.H{"code" : http.StatusOK,"content": "Health Check OK OK ."})
}
func indexMiddleware(ctx *gin.Context) {
	log.Println("Hello Index Middleware")
	
	ctx.Next()
}


func Update_LambdaHandler(ctx *gin.Context) {
	
	/* Required
	1. Repository name
	2. image_uri
	*/

	/* ADD fucntion
	   1. Update function options
	   2. OIDC로 credentials 로그인 하기 찾아보기 https://aws.github.io/aws-sdk-go-v2/docs/configuring-sdk/
	*/
   	var requestBody UpdateImageRequestBody
	   if err := ctx.BindJSON(&requestBody); err != nil {
		   // DO SOMETHING WITH THE ERROR
		 
		   fmt.Println("ERROR");
		   fmt.Println(err)  
		   ctx.JSON(http.StatusBadRequest,gin.H{"Success" : false,"content": "Bad Request"})
		   return
	   }
	
	Result,msg:=Update_lambda_image(requestBody.Func_name,requestBody.Image_uri)
	if Result ==true{
	ctx.JSON(http.StatusOK, gin.H{"Success" : Result,"content": msg})
	}
	if Result ==false{
      ctx.JSON(http.StatusInternalServerError,gin.H{"Success" : Result,"content": msg})
	}

}


func indexMiddleware3(ctx *gin.Context) {
	log.Println("Hello Index Middleware")
	
	ctx.Next()
}
