package main

import (
	"aws_test/aws_web"
)

const ()

func main() {
	//sess := aws_session.CreateSession(Region, WcgAccessKeyId, WcgSecretAccessKey, Token);
	server := aws_web.CreateWebServer()
	server.Run()
}
