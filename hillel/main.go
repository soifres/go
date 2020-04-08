package main

import (
	"go_lessons/hillel/news"
	"go_lessons/hillel/router"
)

func main() {
	r := router.New()
	a := news.New()

	go a.Serve()
	r.Run()
}
