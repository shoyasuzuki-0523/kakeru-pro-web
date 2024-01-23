package main

import "github.com/gin-gonic/gin"

func main() {
	r := gin.Default()
	r.LoadHTMLGlob("templates/*.html")

	r.GET("/", func(ctx *gin.Context) {
		ctx.HTML(200, "index.html", gin.H{
			"data": "This is Home.",
		})
	})

	r.GET("/about", func(ctx *gin.Context) {
		ctx.HTML(200, "index.html", gin.H{
			"data": "This is About.",
		})
	})

	r.Run()
}
