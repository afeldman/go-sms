package main

import (
	"fmt"
	"os"

	"github.com/afeldman/go-sms/smsconfig"
	"github.com/afeldman/go-sms/smsroutes"
	"github.com/gin-gonic/gin"
)

/** create a sms rest gateway
*
 */
func main() {
	// current dir
	workingdir, err := os.Getwd()
	if err != nil {
		panic(err)
	}

	fmt.Println(workingdir)

	// load config from workdir
	smsconfig.LoadConfig(workingdir)

	// start rest interface
	router := gin.Default()

	// route settings
	api := router.Group("/api")
	{
		v1 := api.Group("/v1")
		{
			// handler functions
			v1.POST("/sms", smsroutes.SMSHandler)
		}
	}
	// no rout no handler
	router.NoRoute(smsroutes.NoHandler)

	// server start
	router.Run(smsconfig.SMSConfiguration.ServerAddress + ":" + smsconfig.SMSConfiguration.ServerPort)
}
