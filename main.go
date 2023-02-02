package main

import "fmt"
import "net/http"

func defaultHandler(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "text/html;charset=utf-8")
	if r.URL.Path == "/" {
		fmt.Fprintln(w, "<h1>Hello,欢迎来到 goblog!</h1>")
	} else {
		w.WriteHeader(http.StatusNotFound)
		fmt.Fprint(w, "<h1>请求页面未找到 :(</h1>"+
			"<p>如有疑惑，请联系我们。</p>")
	}
}

func aboutHandler(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "text/html;charset=utf-8")
	fmt.Fprintln(w, "此博客是用以记录编程笔记，如您有反馈或建议，"+
		"请联系<a href=\\\"mailto:herrluk@example.com\\\">herrluk@example.com</a>")

}

func main() {
	router := http.NewServeMux()
	http.HandleFunc("/", defaultHandler)
	http.HandleFunc("/about", aboutHandler)
	err := http.ListenAndServe(":3000", router)
	if err != nil {
		//fmt.Fprintln(err)
	}
}
