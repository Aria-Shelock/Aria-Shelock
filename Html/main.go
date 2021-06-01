package main

import "github.com/gin-gonic/gin"

func main() {
	app := gin.Default()
	app.LoadHTMLFiles("./home.html")
	app.GET("/niu", func(c *gin.Context) {
		c.String(200, "新年快乐")
	})
	app.GET("/home", func(c *gin.Context) {
		c.HTML(200, "home.html", gin.H{})
	})
	app.Run()
}
