package gsm

const (
	NEWLINE              = "\r\n"
	ECHOOFF              = "ATE0" + NEWLINE
	USEFULLERRORMESSAGES = "AT+CMEE=1" + NEWLINE
	DISABLENOTIFICATIONS = "AT+WIND=0" + NEWLINE
	ENABLETEXTMODE       = "AT+CMGF=1" + NEWLINE
)

type GSM interface {
	Connect() error
	initModem()
	Expect(possibilities []string) (string, error)
	Write(message string) bool
	Read(message_size int) string
	Send(command string, waitForOk bool) string
	SMS(mobileno, message string) string
}
