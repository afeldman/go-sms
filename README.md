# go-sms

Have to use my mobile android phone to send sms.

## Prerequisites
install go >1.18
download android tools

[Linux:](https://dl.google.com/android/repository/platform-tools-latest-linux.zip)
[macos:](https://dl.google.com/android/repository/platform-tools-latest-darwin.zip)
[Windows:](https://dl.google.com/android/repository/platform-tools-latest-windows.zip)

And you have to have an Android device with USB Debugging enabled. (Source: Enable developer options and USB debugging)

## adb 
install adb in the path
adb kill-server
bring your mobile into debug mode
connect with usb
enable usb debugging on the mobile
adb start-server