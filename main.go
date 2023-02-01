package main

import "fmt"
import "net/http"

func handlerFunc(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "text/html;charset=utf-8")
	if r.URL.Path == "/" {
		fmt.Fprintln(w, "<h1>Hello, 欢迎来到 goblog ！</h1>")
	} else if r.URL.Path == "/about" {
		fmt.Fprintln(w, "此博客是用以记录编程笔记，如您有反馈或建议，"+
			"请联系<a href=\\\"mailto:herrluk@example.com\\\">herrluk@example.com</a>")
	} else {
		w.WriteHeader(http.StatusNotFound)
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
