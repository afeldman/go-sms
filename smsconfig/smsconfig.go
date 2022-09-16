package smsconfig

// configuration information
import (
	"fmt"
	"log"

	"github.com/afeldman/go-sms/modem"
	"github.com/spf13/viper"
)

// modem type
type ModemType int

/* two available modem types for the beginning
* 1. serial (USB-Stick)
* 2. Android developer interface
 */
const (
	SERIAL ModemType = iota
	ADB
)

// sms device is set as
type SMSDevice struct {
	DeviceType ModemType // device type
	COMPort    string    // comport. not necessary for adb
	Baudrate   int       // baudrate. not needed for adb
	DeviceId   string    // device id for sending data
}

// config object
type SMSConfig struct {
	ServerAddress  string      // server address
	ServerPort     string      // server port
	Retries        int         // how many retries after fail. serial
	Buffersize     int         // buffer size for serial
	BufferLow      int         // buffer low for serial
	MSGTimeOut     int         // message time out serial
	MSGTimeOutLong int         // message time out for long term
	Devices        []SMSDevice //devices
}

// default server settings
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

//load the config utilizing viper
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

	// build devices based on device type
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
