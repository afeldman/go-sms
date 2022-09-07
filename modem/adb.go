package modem

import (
	"bufio"
	"errors"
	"fmt"
	"log"
	"os/exec"
	"strconv"
	"strings"
)

const (
	adb_program string = "adb"
	a8_command  string = "service call isms 7 i32 0 s16 \"com.android.mms.service\" s16 \"[receiver]\" s16 \"null\" s16 \"[msg]\" s16 \"null\" s16 \"null\""
	a11_command string = "service call isms 6 i32 0 s16 \"com.android.mms.service\" s16 \"null\" s16 \"[receiver]\" s16 \"null\" s16 \"[msg]\" s16 \"null\" s16 \"null\" s16 \"null\" s16 \"null\" s16 \"null\""
)

type GSMADBModem struct {
	DeviceId       string
	AndroidVersion int
}

func ModemList() ([]GSMADBModem, error) {
	dev := make([]GSMADBModem, 0)

	log.Println("create modem list")

	if !commandExists(adb_program) {
		return dev, errors.New("adb not found")
	} else {
		log.Println("adb found")
	}

	devices, dev_err := adb_device_ids()
	if dev_err != nil {
		return dev, errors.New("no device available")
	}

	for _, device := range devices {

		version, version_error := adbversion(device)
		if version_error != nil {
			fmt.Println(version_error.Error())
			return dev, version_error
		}

		dev = append(dev, *NewADBModem(device, version))
	}

	return dev, nil
}

func NewADBModem(DeviceId string, version int) *GSMADBModem {

	return &GSMADBModem{DeviceId: DeviceId, AndroidVersion: version}
}

func (m *GSMADBModem) Connnect() error {
	return nil
}

func (m *GSMADBModem) initModem() {
	fmt.Println("no init")
}

func (m *GSMADBModem) Expect(possibilities []string) (string, error) {
	return "", nil
}

func (m *GSMADBModem) Write(message string) bool {
	return true
}

func (m *GSMADBModem) Read(message_size int) string {
	return ""
}

func (m *GSMADBModem) Send(command string, waitForOk bool) string {
	return ""
}

func (m *GSMADBModem) SMS(mobileno, message string) string {
	var command string
	if m.AndroidVersion < 11 {
		command = a8_command
	} else {
		command = a11_command
	}

	command = strings.ReplaceAll(
		strings.ReplaceAll(command, "[receiver]", mobileno), "[msg]", message)

	fmt.Println(command)

	stdout, err := exec.Command(
		adb_program,
		"-s",
		m.DeviceId,
		"shell",
		command).Output()
	if err != nil {
		fmt.Println(err.Error())
		return ""
	}

	var lines []string

	sc := bufio.NewScanner(strings.NewReader(string(stdout)))
	for sc.Scan() {
		lines = append(lines, sc.Text())
	}

	return lines[0]
}

func adb_device_ids() ([]string, error) {
	stdout, deviceerr := exec.Command(
		adb_program,
		"devices").Output()
	if deviceerr != nil {
		fmt.Println(deviceerr.Error())
		return []string{}, deviceerr
	}

	var lines, devices []string

	sc := bufio.NewScanner(strings.NewReader(string(stdout)))
	for sc.Scan() {
		lines = append(lines, sc.Text())
	}

	devices_tmp := make([]string, len(lines)-1)

	copy(devices_tmp, lines[1:])

	for _, device := range devices_tmp {
		dsc := strings.Split(device, "\t")[0]
		devices = append(devices, dsc)
	}

	return devices, nil

}

func adbversion(deviceid string) (int, error) {
	stdout, versionerr := exec.Command(
		adb_program,
		"-s",
		deviceid,
		"shell",
		"getprop",
		"ro.build.version.release").Output()
	if versionerr != nil {
		fmt.Println(versionerr.Error())
		return -1, versionerr
	}

	var version_str string

	s := bufio.NewScanner(strings.NewReader(string(stdout)))
	for s.Scan() {
		version_str = s.Text()
	}

	version, err := strconv.Atoi(version_str)
	if err != nil {
		fmt.Println(err.Error())
		return -1, err
	}

	return version, nil
}

func commandExists(cmd string) bool {
	_, err := exec.LookPath(cmd)
	return err == nil
}
