package aws_session

import (
	"fmt"
	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/aws/credentials"
	"github.com/aws/aws-sdk-go/aws/session"
)

func CreateSession(region string, accesskey string, secretaccesskey string, token string) (*session.Session){
	sess , err := session.NewSession(&aws.Config{
		Region: aws.String(region),
		Credentials: credentials.NewStaticCredentials(accesskey, secretaccesskey, token),
	})

	if err != nil{
		fmt.Println("Create Session Error!!")
	} else{
		fmt.Println("Return sess")
	}
	//sess := session.Must(session.NewSessionWithOptions(session.Options{
	//	SharedConfigState: session.SharedConfigEnable,
	//}))

	return sess
}
