package smsroutes

import (
	"fmt"
	"log"
	"net/http"

	"github.com/afeldman/go-sms/modem"
	"github.com/gin-gonic/gin"
)

type sms struct {
	Mobile  string `json:"no"`
	Message string `json:"msg"`
}

type sms_response struct {
	Number  string `json:"no"`
	Message string `json:"msg"`
}

func SMSHandler(c *gin.Context) {
	sms_request := sms{}

	response := sms_response{
		Number:  "404",
		Message: ""}

	err := c.BindJSON(&sms_request)
	if err != nil {
		response.Message = err.Error()
		log.Println(response)
		c.JSON(http.StatusBadGateway, response)
		return
	}

	log.Println(fmt.Sprintf("sms: %+v", sms_request))

	response.Message = modem.SendSMS(sms_request.Mobile, sms_request.Message)
	response.Number = sms_request.Mobile

	log.Println(response)

	c.JSON(200, response)
}

func NoHandler(c *gin.Context) {
	c.JSON(404, gin.H{"code": http.StatusNotFound, "message": "Page not found"})
}
