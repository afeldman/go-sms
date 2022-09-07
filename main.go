package main

import (
	"fmt"
	"os"

	"github.com/afeldman/go-sms/smsconfig"
	"github.com/afeldman/go-sms/smsroutes"
	"github.com/gin-gonic/gin"
)

func main() {
	workingdir, err := os.Getwd()
	if err != nil {
		panic(err)
	}

	fmt.Println(workingdir)

	smsconfig.LoadConfig(workingdir)

	router := gin.Default()

	api := router.Group("/api")
	{
		v1 := api.Group("/v1")
		{
			v1.POST("/sms", smsroutes.SMSHandler)
		}
	}

	router.NoRoute(smsroutes.NoHandler)

	router.Run(smsconfig.SMSConfiguration.ServerAddress + ":" + smsconfig.SMSConfiguration.ServerPort)
}
