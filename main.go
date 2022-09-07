package main

import (
	"fmt"
	"os"

	"github.com/afeldman/go-sms/modem"
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

	for _, device := range modem.GSMModem {
		switch v := device.(type) {
		case modem.GSMADBModem:
			fmt.Println("ADB device")

		case modem.GSMSerialModem:
			fmt.Println("Serial device")
		default:
			panic("modem type not found")
		}
	}

}
