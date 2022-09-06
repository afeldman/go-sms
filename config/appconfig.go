package smsconfig

type SMSDevice struct {
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
	Devices        []SMSDevice
}

var SMSConfiguration *SMSConfig = &SMSConfig{
	ServerAddress:  "0.0.0.0",
	ServerPort:     "2510",
	Retries:        5,
	Buffersize:     10,
	BufferLow:      4,
	MSGTimeOut:     10,
	MSGTimeOutLong: 20,
	Devices: []SMSDevice{
		{
			COMPort:  "/dev/ttyUSB0",
			Baudrate: 115200,
			DeviceId: "7a2f01ed-d84c-48c9-aa8d-4b3a25191a44",
		},
	},
}
