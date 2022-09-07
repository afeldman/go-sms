package main

import (
	"fmt"
	"os"

	"github.com/afeldman/go-sms/smsconfig"
)

func main() {
	workingdir, err := os.Getwd()
	if err != nil {
		panic(err)
	}

	fmt.Println(workingdir)

	smsconfig.LoadConfig(workingdir)

	fmt.Println(smsconfig.SMSConfiguration.ServerPort)

}
