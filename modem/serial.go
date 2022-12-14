package modem

import (
	"errors"
	"log"
	"strings"
	"time"

	"github.com/tarm/serial"
)

// serial modem constant functions calls
const (
	NEWLINE              = "\r\n"
	ECHOOFF              = "ATE0" + NEWLINE
	USEFULLERRORMESSAGES = "AT+CMEE=1" + NEWLINE
	DISABLENOTIFICATIONS = "AT+WIND=0" + NEWLINE
	ENABLETEXTMODE       = "AT+CMGF=1" + NEWLINE
)

// serial modem
type GSMSerialModem struct {
	ComPort  string
	BaudRate int
	Port     *serial.Port
	DeviceId string
}

// new modem object
func NewSerialModem(ComPort, DeviceId string, BaudRate int) (modem *GSMSerialModem) {
	modem = &GSMSerialModem{
		ComPort:  ComPort,
		BaudRate: BaudRate,
		DeviceId: DeviceId,
	}
	return modem
}

// connect to the modem utilizing the serial port
func (m *GSMSerialModem) Connect() (err error) {
	m.Port, err = serial.OpenPort(
		&serial.Config{
			Name:        m.ComPort,
			Baud:        m.BaudRate,
			ReadTimeout: time.Second})

	if err == nil {
		m.initModem()
	}

	return err
}

// write a message to serial bus
func (m *GSMSerialModem) Write(message string) {
	m.Port.Flush()
	_, err := m.Port.Write([]byte(message))
	if err != nil {
		log.Fatal(err)
	}
}

// read a message vom serial
func (m *GSMSerialModem) Read(message_size int) string {
	var output string = ""
	buf := make([]byte, message_size)
	for i := 0; i < message_size; i++ {
		// ignoring error as EOF raises error on Linux
		c, _ := m.Port.Read(buf)
		if c > 0 {
			output = string(buf[:c])
		}
	}

	return output
}

// expect output on bus
func (m *GSMSerialModem) Expect(possibilities []string) (string, error) {
	readMax := 0
	for _, possibility := range possibilities {
		length := len(possibility)
		if length > readMax {
			readMax = length
		}
	}

	readMax = readMax + 2 // we need offset for \r\n sent by modem

	var status string = ""
	buf := make([]byte, readMax)

	for i := 0; i < readMax; i++ {
		// ignoring error as EOF raises error on Linux
		n, _ := m.Port.Read(buf)
		if n > 0 {
			status = string(buf[:n])

			for _, possibility := range possibilities {
				if strings.HasSuffix(status, possibility) {
					return status, nil
				}
			}
		}
	}

	return status, errors.New("match not found")
}

// send message and may wait for a success
func (m *GSMSerialModem) Send(command string, waitForOk bool) string {
	m.Write(command)

	if waitForOk {
		output, _ := m.Expect([]string{"OK\r\n", "ERROR\r\n"}) // we will not change api so errors are ignored for now
		return output
	} else {
		return m.Read(1)
	}
}

// initialize modem
func (m *GSMSerialModem) initModem() {
	m.Send(ECHOOFF, true)
	m.Send(USEFULLERRORMESSAGES, true)
	m.Send(DISABLENOTIFICATIONS, true)
	m.Send(ENABLETEXTMODE, true)
}

// send a sms
func (m *GSMSerialModem) SMS(mobileno, message string) string {

	m.Write("AT+CMGS=\"" + mobileno + "\"\r")
	m.Read(3)

	return m.Send(message, true)
}
