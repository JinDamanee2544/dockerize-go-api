package main

import (
	"github.com/JinDamanee2544/go-post-api/post"
	"github.com/joho/godotenv"
	"github.com/labstack/echo/v4"
)

func main() {
	godotenv.Load()

	e := echo.New()

	post.ConnectDB()

	e.GET("/posts", post.GetPosts)

	e.POST("/posts", post.CreatePost)

	e.GET("*", func(c echo.Context) error {
		return c.String(200, "Hello, World!")
	})

	e.Logger.Fatal(e.Start(":8080"))
}
