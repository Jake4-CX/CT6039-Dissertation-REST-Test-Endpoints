package controllers

import (
	"io"
	"strconv"
	"time"

	"github.com/gin-gonic/gin"
	log "github.com/sirupsen/logrus"
)

// Get Requests

func GetRequest(c *gin.Context) {
	c.JSON(200, gin.H{
		"message": "GET request received",
	})
}

func GetRequestErrorCode(c *gin.Context) {
	errorCode, err := strconv.Atoi(c.Param("errorCode"))

	if err != nil {
		log.Error("Error parsing error code: " + err.Error())
		errorCode = 400
	} else if errorCode < 100 || errorCode > 599 {
		log.Error("Error code out of range: " + strconv.Itoa(errorCode))
		errorCode = 400
	}

	log.Infof("GET request received with error code: %d", errorCode)

	c.JSON(errorCode, gin.H{
		"message":    "GET request received",
		"error_code": errorCode,
	})
}

func GetRequestDelayedResponse(c *gin.Context) {
	delay, err := strconv.Atoi(c.Param("delay"))
	errCode := 200

	if err != nil {
		log.Error("Error parsing delay: " + err.Error())
		delay = 0
		errCode = 400
	}

	log.Infof("GET request received with delay: %d", delay)

	time.Sleep(time.Duration(delay) * time.Millisecond)

	c.JSON(errCode, gin.H{
		"message": "GET request received",
		"delay":   delay,
	})
}

// Post Requests

func PostRequest(c *gin.Context) {

	data, err := io.ReadAll(c.Request.Body)
	if err != nil {
		log.Error("Error reading request body: " + err.Error())
		c.JSON(400, gin.H{
			"message": "Error reading request body",
		})
		return
	}

	dataStr := string(data)

	log.Info("POST request received with data: " + dataStr)

	c.JSON(200, gin.H{
		"message": "POST request received",
		"data":    dataStr,
	})
}

// Put Requests

func PutRequest(c *gin.Context) {

	data, err := io.ReadAll(c.Request.Body)
	if err != nil {
		log.Error("Error reading request body: " + err.Error())
		c.JSON(400, gin.H{
			"message": "Error reading request body",
		})
		return
	}

	dataStr := string(data)

	log.Info("PUT request received with data: " + dataStr)

	c.JSON(200, gin.H{
		"message": "PUT request received",
		"data":    dataStr,
	})
}

// Delete Requests

func DeleteRequest(c *gin.Context) {
	param := c.Param("param")

	c.JSON(200, gin.H{
		"message": "DELETE request received",
		"param":   param,
	})
}
