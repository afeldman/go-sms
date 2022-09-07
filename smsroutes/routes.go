package smsroutes

import (
	"net/http"

	"github.com/afeldman/go-sms/modem"
	"github.com/gin-gonic/gin"
)

type sms struct {
	mobile  string `json:"no"`
	message string `json:"msg"`
}

type sms_response struct {
	number  string `json:"no"`
	message string `json:"msg"`
}

func SMSHandler(c *gin.Context) {
	var SMS sms

	response := sms_response{
		number:  "404",
		message: ""}

	if err := c.BindJSON(&SMS); err != nil {
		response.message = err.Error()
		c.JSON(http.StatusBadGateway, response)
	}

	response.message = modem.SendSMS(SMS.mobile, SMS.message)
	response.number = SMS.mobile

	c.JSON(200, response)
}

func NoHandler(c *gin.Context) {
	c.JSON(404, gin.H{"code": http.StatusNotFound, "message": "Page not found"})
}
