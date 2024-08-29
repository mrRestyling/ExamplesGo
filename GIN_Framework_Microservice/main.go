package main

import "github.com/gin-gonic/gin"

type Links struct {
	ID    uint64 `json:"id"`
	Title string `json:"title"`
	URL   string `json:"url"`
	Score uint64 `json:"score"`
}

var links = []Links{
	{ID: 1, Title: "Alpha", URL: "http://alpha", Score: 100},
	{ID: 2, Title: "Beta", URL: "http://beta", Score: 150},
}

func Example(c *gin.Context) {
	c.JSON(200, links)
}

func main() {

	route := gin.Default()
	route.GET("/example", Example)
	route.Run()

}
