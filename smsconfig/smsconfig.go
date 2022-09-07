package smsconfig

import (
	"fmt"
	"log"

	"github.com/afeldman/go-sms/modem"
	"github.com/spf13/viper"
)

type ModemType int

const (
	SERIAL ModemType = iota
	ADB
)

type SMSDevice struct {
	DeviceType ModemType
	COMPort    string
	Baudrate   int
	DeviceId   string
}

type SMSConfig struct {
	ServerAddress  string
	ServerPort     string
	Retries        int
	Buffersize     int
	BufferLow      int
	MSGTimeOut     int
	MSGTimeOutLong int
	Devices        []SMSDevice
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

	for _, device := range SMSConfiguration.Devices {
		if device.DeviceType == SERIAL {
			gsm_modem := modem.NewSerialModem(device.COMPort, device.DeviceId, device.Baudrate)
			modem.GSMModem = append(modem.GSMModem, gsm_modem)
		}
		if device.DeviceType == ADB {
			modem_list, err := modem.ModemList()
			if err != nil {
				panic(err.Error())
			}
			for _, gms_modem := range modem_list {
				modem.GSMModem = append(modem.GSMModem, gms_modem)
			}
		}
	}

}
