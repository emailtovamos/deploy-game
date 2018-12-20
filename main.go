package main

import (
	"github.com/gin-gonic/gin"
	// "net/http"
	// "fmt"
	"strconv"
)

var router *gin.Engine
var bestScore = 999999.0

func main() {
	router = gin.Default()
	// router.LoadHTMLGlob("index.html")
	initializeRoutes()

	router.Run(":8080")
}

func initializeRoutes() {
	// router.POST("/setbs", handleSet)
	router.GET("/setbs/:hs", handleSet)
	router.GET("/getbs", handleGet)
}

func handleGet(c *gin.Context) {
	// message, _ := c.GetQuery("m")
	bsString := strconv.FormatFloat(bestScore, 'e', -1, 64)
	// c.String(http.StatusOK, bsString)
	c.JSONP(200, gin.H{
		"hs": bsString,
	})
}

func handleSet(c *gin.Context) {
	// message, _ := c.GetQuery("hs")
	// // c.String(http.StatusOK, string(bestScore))
	// bs, _ := strconv.ParseFloat(message, 64)
	// bestScore := bs
	// fmt.Println(bestScore)
	// c.JSONP(200, gin.H{
	// 	"hs": message,
	// })

	hs := c.Param("hs")
	bestScore, _ = strconv.ParseFloat(hs, 64)

	c.JSONP(200, gin.H{
		"hs": hs,
	})
}
