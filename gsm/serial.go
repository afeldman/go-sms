package gsm

import (
	"errors"
	"log"
	"strings"
	"time"

	"github.com/tarm/serial"
)

type GSMSerialModem struct {
	ComPort      string
	BaudRate     int
	Port         *serial.Port
	DeviceId     string
	Success      bool
	Transmission time.Time
}

func New(ComPort, DeviceId string, BaudRate int) (modem *GSMSerialModem) {
	modem = &GSMSerialModem{
		ComPort:      ComPort,
		BaudRate:     BaudRate,
		DeviceId:     DeviceId,
		Success:      false,
		Transmission: time.Unix(0, 0),
	}
	return modem
}

func (m *GSMSerialModem) Connect() (err error) {
	m.Port, err = serial.OpenPort(
		&serial.Config{
			Name: 		 m.ComPort,
			Baud:        m.BaudRate,
			ReadTimeout: time.Second
		}
	)

	if err == nil {
		m.initModem()
	}

	return err
}

func (m *GSMSerialModem) Write(message string) {
	m.Port.Flush()
	_, err := m.Port.Write([]byte(message))
	if err != nil {
		log.Fatal(err)
	}
}

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

func (m *GSMSerialModem) Send(command string, waitForOk bool) string {
	m.Write(command)

	if waitForOk {
		output, _ := m.Expect([]string{"OK\r\n", "ERROR\r\n"}) // we will not change api so errors are ignored for now
		return output
	} else {
		return m.Read(1)
	}
}

func (m *GSMSerialModem) initModem() {
	m.Send(ECHOOFF, true)
	m.Send(USEFULLERRORMESSAGES, true)
	m.Send(DISABLENOTIFICATIONS, true)
	m.Send(ENABLETEXTMODE, true)
}

func (m *GSMSerialModem) SMS(mobileno, message string) string {

	m.Write("AT+CMGS=\"" + mobileno + "\"\r")
	m.Read(3)

	return m.Send(message, true)
}
