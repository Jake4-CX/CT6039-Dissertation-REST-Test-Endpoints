package main

import (
	"os"

	"github.com/Jake4-CX/CT6039-Dissertation-REST-Test-Endpoints/cmd/controllers"
	"github.com/Jake4-CX/CT6039-Dissertation-REST-Test-Endpoints/pkg/initializers"
	"github.com/gin-gonic/gin"
	log "github.com/sirupsen/logrus"
)

func main() {

	initializers.LoadEnvVariables()

	router := gin.Default()

	router.GET("/GET-request", controllers.GetRequest)
	router.GET("/GET-request/error-code/:errorCode", controllers.GetRequestErrorCode)
	router.GET("/GET-request/delayed-response/:delay", controllers.GetRequestDelayedResponse)

	router.POST("/POST-request", controllers.PostRequest)
	router.PUT("/PUT-request", controllers.PutRequest)
	router.DELETE("/DELETE-request/:param", controllers.DeleteRequest)

	log.Fatal((router.Run("0.0.0.0:" + os.Getenv("REST_PORT"))))

}