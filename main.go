package main

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"net/http"
)

func main() {
	app := gin.Default()

	//template configuration
	app.LoadHTMLGlob("templates/*")

	app.GET("/", func(res *gin.Context) {
		res.JSON(200, gin.H{
			"message": "ping pong",
		})
	})

	app.GET("/ascii-json", func(res *gin.Context) {
		data := map[string]interface{}{
			"lang": "GO语言",
			"tag":  "<br>",
		}
		res.AsciiJSON(http.StatusOK, data)
	})

	//render html
	app.GET("/layout", func(res *gin.Context) {
		res.HTML(http.StatusOK, "layout.html", gin.H{
			"title": "Gin Page Layout",
		})
	})

	app.GET("/home", func(res *gin.Context) {
		res.HTML(http.StatusOK, "layout.html", gin.H{
			"title":   "GIN Home Page",
			"content": "about.html",
		})
	})

	app.GET("/about", func(res *gin.Context) {
		res.HTML(http.StatusOK, "layout.html", gin.H{
			"title":   "Gin Page About",
			"content": "about.html",
		})
	})

	//run server
	func() {
		err := app.Run(":8080")

		if err != nil {
			fmt.Println("Application Error :", err)
		}
	}()
}
