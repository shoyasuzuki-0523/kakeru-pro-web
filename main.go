package main

import "github.com/gin-gonic/gin"

type Meta struct {
	Title string
}

func main() {
	r := gin.Default()
	r.LoadHTMLGlob("templates/*.html")

	r.GET("/", func(ctx *gin.Context) {
		meta := Meta{Title: "Home"}
		ctx.HTML(200, "index.html", gin.H{
			"meta": meta,
			"data": "This is Home.",
		})
	})

	r.GET("/about", func(ctx *gin.Context) {
		meta := Meta{Title: "About"}
		ctx.HTML(200, "index.html", gin.H{
			"meta": meta,
			"data": "This is About.",
		})
	})

	r.Run()
}
