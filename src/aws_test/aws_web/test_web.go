package aws_web

import (
	"aws_test/aws_data_struct"
	"aws_test/aws_ec2_create"
	"aws_test/aws_ec2_delete"
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
		v1.GET("/playerinfo", playerinfo)
		v1.GET("/adduser", adduser)
		v1.POST("/adduser_post", adduser_post)

		v1.GET("/deluser", deluser)
		v1.POST("/deluser_post", deluser_post)
	}
	return r
}

func deluser_post(context *gin.Context) {
	aws_ec2_delete.Delete_DeepLearning_Plyaer_Instances(MySvc)
	playerinfo(context)
}

func deluser(context *gin.Context) {
	context.HTML(
		http.StatusOK,
		"deluser.html",
		gin.H{
			"temp": aws_data_struct.Player,
		})

}

func playerinfo(context *gin.Context) {

	context.HTML(
		http.StatusOK,
		"playerinfo.html",
		gin.H{
			"temp": aws_data_struct.Player,
		})

}

func adduser_post(context *gin.Context) {
	fmt.Println("AddUser_Post func")
	TeamName := context.PostForm("teamname")
	Type := context.PostForm("type")

	p := aws_data_struct.PlayerEc2Status{}
	p.TeamName = TeamName
	p.Type = Type

	//인스턴스 만들어서 할차례
	aws_ec2_create.Create_DeepLearning_Player_Instances(MySvc, &p)
	aws_data_struct.Player = append(aws_data_struct.Player, p)

	playerinfo(context)
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
