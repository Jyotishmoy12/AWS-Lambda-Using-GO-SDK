package main

import (
	"context"
	"fmt"
	"github.com/aws/aws-lambda-go/lambda"
	"github.com/aws/aws-lambda-go/events"
)

// request defines the input structure for our lambda function
// the payload will be json
type Request struct {
	Name string `json:"name"`
}
// response defines the output structure for our lambda function
type Response struct {
	Message string `json:"message"`
}

// Handler function now accepts a context.Contect and a event for S3
func Handler(ctx context.Context, s3Event events.S3Event){
	fmt.Println("S3 Event: ", s3Event)
	// and s3 event can contain multple records
	// so we need to loop through them
	for _, record :=range s3Event.Records{
		bucket:= record.S3.Bucket.Name
		key:= record.S3.Object.Key
		// kif the information to cloudwatch
		fmt.Printf("File upload detected!\n")
		fmt.Printf("Bucket: %s\n", bucket)
		fmt.Printf("Key: %s\n", key)
	}
}

// Handler is our function that lambda will execute
// func Handler(req Request)(Response, error){
// 	if req.Name == ""{
// 		return Response{
// 			Message: "Hello, World",
// 		}, nil
// 	}
// 	return Response{
// 		Message: fmt.Sprintf("Hello, %s", req.Name),
// 	}, nil
// }

func main(){
	// lambda.Start(Handler) function takes our handler and wraps
	// it in a function that will be called by the lambda runtime
	lambda.Start(Handler)
}