package smsconfig

import (
	"fmt"
	"log"

	"github.com/spf13/viper"
)

type ModemType int

const (
	SERIAL ModemType = iota
	ADB
)

type SMSADBDevice struct {
	DeviceId string
}

type SMSSerialDevice struct {
	COMPort  string
	Baudrate int
	DeviceId string
}

type SMSConfig struct {
	ServerAddress  string
	ServerPort     string
	Retries        int
	Buffersize     int
	BufferLow      int
	MSGTimeOut     int
	MSGTimeOutLong int
	Devices        []interface{}
}

var SMSConfiguration *SMSConfig = &SMSConfig{
	ServerAddress:  "0.0.0.0",
	ServerPort:     "1712",
	Retries:        5,
	Buffersize:     10,
	BufferLow:      4,
	MSGTimeOut:     10,
	MSGTimeOutLong: 20,
	Devices:        nil,
}

func LoadConfig(workingdir string) {
	viper.SetConfigFile(workingdir + "/config/smsconfig.toml")
	viper.SetConfigType("toml")

	if err := viper.ReadInConfig(); err != nil {
		log.Fatal(err)
	}

	if err := viper.Unmarshal(SMSConfiguration); err != nil {
		log.Fatal(err)
	}
	fmt.Printf("%#v\n", SMSConfiguration)

}
