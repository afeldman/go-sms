package modem

import (
	"log"
	"strconv"
)

// GSMModem are interfaces :D
var GSMModem []interface{}

// interfrace function for all modems
type GSM interface {
	Connect() error
	initModem()
	Expect(possibilities []string) (string, error)
	Write(message string) bool
	Read(message_size int) string
	Send(command string, waitForOk bool) string
	SMS(mobileno, message string) string // send sms function
}

/* send the sms.
 * given a message and a number
 */
func SendSMS(number, message string) string {

	log.Println(number)
	log.Println(message)

	// send on all devices. this has to be changed.
	// using an intelligent queue
	for _, device := range GSMModem {
		switch v := device.(type) {
		case GSMADBModem:
			log.Println("ADB device")
			log.Println("device id: " + v.DeviceId)
			log.Println("android version: " + strconv.Itoa(v.AndroidVersion))
			return v.SMS(number, message)
		case GSMSerialModem:
			log.Println("Serial device")
			log.Println(v.DeviceId)
			return v.SMS(number, message)
		default:
			panic("modem type not found")
		}
	}
	return ""
}
