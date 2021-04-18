package main

import (
	"newsfeeder/httpd/handler"
	newsfeed "newsfeeder/platform"

	"github.com/gin-gonic/gin"
)

func main() {
	r := gin.Default()
	feed := newsfeed.New()

	r.GET("/ping", handler.PingGet("King Kong"))
	r.GET("/newsfeed", handler.NewsfeedGet(feed))
	r.POST("/newsfeed", handler.NewsfeedPost(feed))

	r.Run() // listen and serve on 0.0.0.0:8080 (for windows "localhost:8080")

	// feed := newsfeed.New()

	// fmt.Println(feed)

	// feed.Add(newsfeed.Item{
	// 	"Hello", "How ya ' doing mate?",
	// })

	// fmt.Println(feed)
}
