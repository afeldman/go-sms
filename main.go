package main

import (
	"fmt"
	"os"
	"strconv"

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

	fmt.Println(len(modem.GSMModem))

	for _, device := range modem.GSMModem {
		switch v := device.(type) {
		case modem.GSMADBModem:
			fmt.Println("ADB device")
			fmt.Println("device id: " + v.DeviceId)
			fmt.Println("android version: " + strconv.Itoa(v.AndroidVersion))

		case modem.GSMSerialModem:
			fmt.Println("Serial device")
			fmt.Println(v.DeviceId)
		default:
			panic("modem type not found")
		}
	}

}
