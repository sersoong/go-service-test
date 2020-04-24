package web

import (
	"fmt"

	"github.com/gin-gonic/gin"
)

//StartWeb start web service
func StartWeb() {
	fmt.Println("start a web service")
	r := gin.Default()
	r.GET("/", handleRoot)
	r.Run("127.0.0.1:8880")
}

func handleRoot(c *gin.Context) {
	c.JSON(200, "helloworld")
}
