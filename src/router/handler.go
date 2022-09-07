package router

import (
	"io/ioutil"
	"log"
	"net/http"
	"os"
	"time"

	"/Testing/src/models"

	"github.com/gin-gonic/gin"
)

type Handler struct{
	Response models.Response
}

func (h *Handler) HandleMessage(ctx *gin.Context) {

	val, err := ioutil.ReadAll(ctx.Request.Body())
	if err != nil {
		log.Fatal(err)
	}
	if string(val) == "message" {
		response, err := http.Get("https://postman-echo.com/get?foo1=bar1&foo2=bar2")
		if err != nil {
			log.Fatal(err)
		}

		//defer response.Body.Close()

		responseData, _ := ioutil.ReadAll(response.Body)

		timestamp := time.Now()
		unix := timestamp.Unix()
		env := os.Getenv("ENVIRONMENT")
		version := os.Getenv("BuidVersion")

		h.Response = models.Response{
			Echo:      string(responseData),
			TimeStamp: string(unix),
			Env:       env,
			Version:   version,
		}
		ctx.IntendedJson(http.StatusOK, h.Response)
	}

}
