package router

import (
	"fmt"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/soifres/go_lessons/hillel/news"
	"github.com/soifres/go_lessons/hillel/wheather"
)

func indexHandler(c *gin.Context) {
	// c.String(http.StatusOK, "Hello world!")
	c.HTML(
		// Set the HTTP status to 200 (OK)
		http.StatusOK,
		// Use the index.html template
		"index.html",
		// Pass the data that the page uses (in this case, 'title')
		gin.H{
			"title": "Домашняя страница",
			"wht":   wheather.WheatherNow,
		},
	)
}

func collectHandler(c *gin.Context) {
	category := c.Param("category")

	items := news.Collect(category)

	// tmp := fmt.Sprintf("Search is initialized, collected %d items", items)
	// c.String(http.StatusOK, tmp)
	c.HTML(
		http.StatusOK,
		"index.html",
		gin.H{
			"title":   "Результат запроса",
			"count":   fmt.Sprintf("Найдено %d", items),
			"catlink": fmt.Sprintf("/result/%s", category),
			"wht":     wheather.WheatherNow,
			// "catlink": strconv.QuoteToASCII(fmt.Sprintf("/result/%s", category)),
		},
	)

}

func resultHandler(c *gin.Context) {
	category := c.Param("category")
	topics := news.Result(category)
	// c.JSON(http.StatusOK, topics)

	// type Message struct {
	// 	Name string
	// 	Body string
	// 	Time int64
	// }
	// var topics = []Message{
	// 	{"Alice1", "Hello1", 1294706395881547000},
	// 	{"Alice2", "Hello2", 1294706395881547000},
	// 	{"Alice3", "Hello3", 1294706395881547000},
	// }

	// jsResult, err := json.Marshal(topics)
	// if err != nil {
	// 	log.Fatalf("Сбой маршалинга JSON: %s", err)
	// 	return
	// }
	c.HTML(
		http.StatusOK,
		"index.html",
		gin.H{
			"title":  "Результат запроса",
			"topics": topics,
			"wht":    wheather.WheatherNow,
		},
	)

}
