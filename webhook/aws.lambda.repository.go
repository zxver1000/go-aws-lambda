package webhook

import (
	"context"
	"encoding/json"
	"errors"
	"fmt"
	"log"
	"os"
	"sync"

	_ "github.com/aws/smithy-go"

	"github.com/aws/aws-sdk-go-v2/config"
	"github.com/aws/aws-sdk-go-v2/service/lambda"
	"github.com/aws/smithy-go"
)

type AWS_Lambda struct {
	lambdaClient *lambda.Client
}

var singleInstance *AWS_Lambda
var once sync.Once
func Lambda_client_getInstance() *AWS_Lambda {
	if singleInstance == nil {
		once.Do(func() {
			fmt.Println("[Once] create lambda single instance")
			cfg, err := config.LoadDefaultConfig(context.TODO(), config.WithSharedCredentialsFiles(
				[]string{"../env"},
				), config.WithRegion(os.Getenv("region")))
			if err != nil {
				fmt.Println("Couldn't load default configuration. Have you set up your AWS account?")
				fmt.Println(err)
				return
			}
			lambdaClient_ := lambda.NewFromConfig(cfg)

			singleInstance = &AWS_Lambda{
				lambdaClient: lambdaClient_,
			}
		})
	} else {
		fmt.Println("[Once] lambda single instance already created-1")
	}
	return singleInstance
}



func Update_lambda_image(Func_name string,Image_uri string) (bool,string){

	/* 예외 처리 추가할것 
	   1. ECR에 이미지가 있는지 확인 -> 없으면  Bad Request
	   2. lambda fucntion 이 없는 경우  
	   3. 업데이트가 안될경우
	   4. lambda함수 자체 변경	   
	*/

	aws_instance:=Lambda_client_getInstance()
	
	var oe *smithy.OperationError

	update_lambda,err3:=aws_instance.lambdaClient.UpdateFunctionCode(context.TODO(),&lambda.UpdateFunctionCodeInput{
        FunctionName:&Func_name,
		DryRun:false,
		ImageUri:&Image_uri,
	})
	if err3!=nil{
		log.Println("Lambda Udpate error :",err3)
		//idx:=strings.Index("StatusCode",err3)
		//fmt.Println(idx)
		if errors.As(err3, &oe) {
			log.Printf("failed to call service: %s, operation: %s, error: %v", oe.Service(), oe.Operation(), oe.Unwrap())
			ErrorString := fmt.Sprintf("%+v", oe.Unwrap())
			//idx:=strings.Indexof(ErrorString,"Statuscode")
			return false,ErrorString
		}
	}
	
    jsonUpdate_lambda, err5 := json.Marshal(update_lambda)
   fmt.Println(err5)
     
	fmt.Println(jsonUpdate_lambda)
	
	

 return true,"Update Success"

}

func GetLambdaFucntion(Func_name string){

    aws_instance:=Lambda_client_getInstance()
	result1,err2:=aws_instance.lambdaClient.GetFunction(context.TODO(),&lambda.GetFunctionInput{
		FunctionName:&Func_name,
	})
	if err2 != nil {
		fmt.Printf("Couldn't list functions for your account. Here's whyhihi: %v\n", err2)
		return
	}
	e, err3 := json.Marshal(result1)
    if err3 != nil {
        fmt.Println(err3)
        return
    }
	fmt.Println("----------------")
	
	fmt.Println(string(e))

	fmt.Println("----------------")

}