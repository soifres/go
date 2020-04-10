package main

import (
	"github.com/soifres/go_lessons/hillel/news"
	"github.com/soifres/go_lessons/hillel/router"
)

func main() {
	r := router.New()
	a := news.New()

	// Запуск бесконечного цикла
	go a.Serve()

	// Запуск web сервера
	r.Run()
}

// http://localhost:8080/search/health
// http://localhost:8080/result/health
