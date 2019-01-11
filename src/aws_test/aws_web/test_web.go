package aws_web

import (
	"aws_test/aws_data_struct"
	"aws_test/aws_ec2_status"
	"aws_test/aws_session"
	"fmt"
	"github.com/aws/aws-sdk-go/aws/session"
	"github.com/aws/aws-sdk-go/service/ec2"
	"github.com/gin-gonic/gin"
	"net/http"
)

var MySession *session.Session
var MySvc *ec2.EC2

func CreateWebServer() (server *gin.Engine) {
	r := gin.Default()
	//r.Use(static.Serve("/", static.LocalFile("./src/aws_test/aws_web/views", true)))
	r.LoadHTMLGlob("templates/*")

	r.GET("/", rootPage)
	r.POST("/checkkey", checkKey)

	v1 := r.Group("/v1")
	{
		v1.GET("/", index)
		v1.GET("/adduser", adduser)
		v1.GET("/playerinfo", playerinfo)
		v1.POST("/adduser_post", adduser_post)
	}
	return r
}

func playerinfo(context *gin.Context) {

}

func adduser_post(context *gin.Context) {
	TeamName := context.PostForm("teamname")
	Type := context.PostForm("type")

	p := aws_data_struct.PlayerEc2Status{}
	p.TeamName = TeamName
	p.Type = Type

	//인스턴스 만들어서 할차례
}

func adduser(context *gin.Context) {

	context.HTML(
		http.StatusOK,
		"adduser.html",
		gin.H{})

}

func checkKey(context *gin.Context) {
	AKey := context.PostForm("AccessKey")
	SKey := context.PostForm("SecretAccessKey")
	RKey := context.PostForm("Region")

	fmt.Println("Key Check : ", AKey, SKey, RKey)
	sess, _ := aws_session.CreateSession(RKey, AKey, SKey, "")
	MySession = sess
	svc := ec2.New(sess)
	fmt.Println("111")
	_, err := aws_ec2_status.Describe(svc)
	if err != nil {
		context.HTML(
			http.StatusOK,
			"makesession.html",
			gin.H{
				"message": "Key error",
			})
	} else {
		MySvc = svc
		index(context)
	}
}

func rootPage(context *gin.Context) {

	context.HTML(
		http.StatusOK,
		"makesession.html",
		gin.H{})
}

func index(context *gin.Context) {
	desc, _ := aws_ec2_status.Describe(MySvc)

	context.HTML(
		http.StatusOK,
		"index.html",
		gin.H{
			"title": "Home Page",
			"Page":  desc,
		})
}

func create(context *gin.Context) {

	test_1 := context.PostForm("nickname")
	test_2 := context.PostForm("job")
	fmt.Println(test_1, " ", test_2)

	context.HTML(
		http.StatusOK,
		"createInstances.html",
		gin.H{
			"title": "Home Page",
		})
}
