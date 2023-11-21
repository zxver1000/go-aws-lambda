package webhook

import (
	"context"
	"fmt"
	"os"
	"sync"

	"encoding/json"

	"github.com/aws/aws-sdk-go-v2/config"
	"github.com/aws/aws-sdk-go-v2/service/lambda"
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



func Update_lambda_image(Func_name string,Image_uri string) int64{

	/* 예외 처리 추가할것 
	   1. ECR에 이미지가 있는지 확인 -> 없으면  Bad Request
	   2. lambda fucntion 이 없는 경우  
	   3. 업데이트가 안될경우
	   4. lambda함수 자체 변경	   
	*/

	aws_instance:=Lambda_client_getInstance()
	fmt.Println(aws_instance)
//	fun_name:="ko"

//	image:="455883942660.dkr.ecr.ap-northeast-2.amazonaws.com/koko:0.1"
	update_lambda,err3:=aws_instance.lambdaClient.UpdateFunctionCode(context.TODO(),&lambda.UpdateFunctionCodeInput{
        FunctionName:&Func_name,
		DryRun:false,
		ImageUri:&Image_uri,
	})
	if err3!=nil{
		fmt.Printf("error update.: %v\n", err3)
		return 404
	}
	/*
	json_lambda, err5 := json.Marshal(update_lambda)
    if err5 != nil {
        fmt.Println(err5)
        return 404
    }
	
	fmt.Println(json_lambda.function_name)

	*/

 return 200

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