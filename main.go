package main

import (
	"fmt"
	"github.com/gorilla/mux"
	"strings"
)
import "net/http"

func homeHandler(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "text/html;charset=utf-8")
	fmt.Fprintln(w, "<h1>Hello,欢迎来到 goblog!</h1>")
}

func aboutHandler(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "text/html;charset=utf-8")
	fmt.Fprintln(w, "此博客是用以记录编程笔记，如您有反馈或建议，"+
		"请联系<a href=\\\"mailto:herrluk@example.com\\\">herrluk@example.com</a>")

}

func notFoundHandler(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "text/html; charset=utf-8")
	w.WriteHeader(http.StatusNotFound)
	fmt.Fprint(w, "<h1>请求页面未找到 :(</h1><p>如有疑惑，请联系我们。</p>")
}

func articlesShowHandler(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	id := vars["id"]
	fmt.Fprintln(w, "文章 ID"+id)
}

func articlesIndexHandler(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintln(w, "访问文章列表")

}

func articlesStoreHandler(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintln(w, "创建新文章")
}

func forceHTMLMiddleware(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		// 1. 设置标头
		w.Header().Set("Content-Type", "text/html;charset=utf-8")
		// 2. 继续处理请求
		next.ServeHTTP(w, r)
	})
}

func removeTrailingSlash(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		if r.URL.Path != "/" {
			r.URL.Path = strings.TrimSuffix(r.URL.Path, "/")
		}
		next.ServeHTTP(w, r)
	})
}
func main() {
	router := mux.NewRouter()
	router.HandleFunc("/", homeHandler).Methods("GET").Name("home")
	router.HandleFunc("/about", aboutHandler).Methods("GET").Name("about")

	router.HandleFunc("/articles/{id:[0-9]+}", articlesShowHandler).
		Methods("GET").Name("articles.show")
	router.HandleFunc("/articles", articlesIndexHandler).
		Methods("GET").Name("articles.index")
	router.HandleFunc("/articles", articlesStoreHandler).
		Methods("POST").Name("articles.store")

	// 自定义 404 界面
	router.NotFoundHandler = http.HandlerFunc(notFoundHandler)

	// 中间件：强制内容类型为 HTML
	router.Use(forceHTMLMiddleware)

	// 通过命名获取 URL 示例
	homeURL, _ := router.Get("home").URL()
	fmt.Println("homeURL: ", homeURL)
	articleURL, _ := router.Get("articles.show").URL("id", "23")
	fmt.Println("articleURL: ", articleURL)
	err := http.ListenAndServe(":3000", removeTrailingSlash(router))
	if err != nil {
		//fmt.Fprintln(err)
	}
}
