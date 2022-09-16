package smsroutes

import (
	"fmt"
	"log"
	"net/http"

	"github.com/afeldman/go-sms/modem"
	"github.com/gin-gonic/gin"
)

// the rest routes
type SMSRequest struct {
	Mobile  string `json:"no"`  // module number for device
	Message string `json:"msg"` // message to device
}

// answer to rest client
type sms_response struct {
	Number  string `json:"no"`  // number of target
	Message string `json:"msg"` // answer message
}

// handle the sms
func SMSHandler(c *gin.Context) {
	sms_request := SMSRequest{}

	response := sms_response{
		Number:  "404",
		Message: ""}

	// read the request
	err := c.BindJSON(&sms_request)
	if err != nil {
		response.Message = err.Error()
		log.Println(response)
		c.JSON(http.StatusBadGateway, response)
		return
	}

	log.Printf(fmt.Sprintf("sms: %+v", sms_request) + "\n")

	// send the sms
	response.Message = modem.SendSMS(sms_request.Mobile, sms_request.Message)
	response.Number = sms_request.Mobile

	log.Println(response)

	// return all ok data return
	c.JSON(200, response)
}

// no endpoint found
func NoHandler(c *gin.Context) {
	c.JSON(404, gin.H{"code": http.StatusNotFound, "message": "Page not found"})
}
