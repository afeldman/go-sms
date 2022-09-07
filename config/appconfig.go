package smsconfig

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
