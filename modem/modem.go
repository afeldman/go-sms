package modem

var GSMModem []interface{}

type GSM interface {
	Connect() error
	initModem()
	Expect(possibilities []string) (string, error)
	Write(message string) bool
	Read(message_size int) string
	Send(command string, waitForOk bool) string
	SMS(mobileno, message string) string
}
