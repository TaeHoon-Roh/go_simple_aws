package aws_web

import (
	"aws_test/aws_ec2_status"
	"fmt"
	"github.com/aws/aws-sdk-go/aws/session"
	"github.com/aws/aws-sdk-go/service/ec2"
	"github.com/gin-gonic/gin"
	"net/http"
)

const (
	//MyAws
	MyAccessKeyId = "AKIAJU2KVGEPDAHX5Q4Q"
	MySecretAccessKey = "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"
	//WcgAws
	WcgAccessKeyId = "AKIAJQKXLNE67JVFRMMA"
	WcgSecretAccessKey = "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"


	Region = "ap-northeast-2"
	Token = ""
	InstanceId = "i-093effb4c7399d644"
)

var MySession *session.Session
func CreateWebServer(sess *session.Session) (server *gin.Engine){
	MySession = sess
	r := gin.Default()


	//r.Use(static.Serve("/", static.LocalFile("./src/aws_test/aws_web/views", true)))
	r.LoadHTMLGlob("templates/*")

	r.GET("/", index)

	v1 := r.Group("/v1")
	{
		v1.GET("/desc", desc)
		v1.POST("/create", create)
	}
	return r
}

func index(context *gin.Context) {
	svc := ec2.New(MySession)
	desc := aws_ec2_status.Describe(svc)

	context.HTML(
		http.StatusOK,
		"index.html",
		gin.H{
			"title":"Home Page",
			"Page":desc,

		})
}

func desc(context *gin.Context) {

	context.HTML(
		http.StatusOK,
		"adduser.html",
		gin.H{
			"title":"Home Page",
		})
}

func create(context *gin.Context) {

	test_1 := context.PostForm("nickname")
	test_2 := context.PostForm("job")
	fmt.Println(test_1, " " , test_2)
	context.HTML(
		http.StatusOK,
		"createInstances.html",
		gin.H{
			"title":"Home Page",
		})
	//context.JSON(http.StatusCreated, gin.H{
	//	"message":"create Instances",
	//})
}
