package main

import "fmt"
import "net/http"

func handlerFunc(w http.ResponseWriter, r *http.Request) {

	if r.URL.Path == "/" {
		fmt.Fprintln(w, "<h1>Hello, 欢迎来到 goblog ！</h1>")
	} else if r.URL.Path == "/about" {
		fmt.Fprintln(w, "此博客是用以记录编程笔记，如您有反馈或建议，"+
			"请联系<a href=\\\"mailto:herrluk@example.com\\\">herrluk@example.com</a>")
	} else {
		fmt.Fprint(w, "<h1>请求页面未找到 :(</h1>"+
			"<p>如有疑惑，请联系我们。</p>")
	}

}

func main() {
	http.HandleFunc("/", handlerFunc)
	err := http.ListenAndServe(":3000", nil)
	if err != nil {
		//fmt.Fprintln(err)
	}
}
