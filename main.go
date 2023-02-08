package main

import (
	"fmt"
	"goBlogByMyself/app/http/middlewares"
	"goBlogByMyself/bootstrap"
	"net/http"
)

func main() {

	bootstrap.SetupDB()
	router := bootstrap.SetupRoute()

	// 通过命名获取 URL 示例
	homeURL, _ := router.Get("home").URL()
	fmt.Println("homeURL: ", homeURL)
	articleURL, _ := router.Get("articles.show").URL("id", "23")
	fmt.Println("articleURL: ", articleURL)

	//
	err := http.ListenAndServe(":3000", middlewares.RemoveTrailingSlash(router))
	if err != nil {
		fmt.Println(err)
	}
}
