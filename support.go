package main

import (
	"fmt"
	"log"
	"net/http"
	"os"
	"strconv"
)

var api API

func main() {
	http.HandleFunc("/", func(res http.ResponseWriter, req *http.Request) {
		res.Header().Set("Access-Control-Allow-Origin", "*")
		res.Write([]byte("Hello"))
	})
	http.HandleFunc("/book", func(res http.ResponseWriter, req *http.Request) {
		res.Header().Set("Access-Control-Allow-Origin", "*")
		res.Header().Set("Content-Type", "text/json; charset=utf-8")
		req.ParseForm()
		locator := req.Form.Get("locator")

		api.GetBook(locator, res)
	})
	http.HandleFunc("/book/page", func(res http.ResponseWriter, req *http.Request) {
		res.Header().Set("Access-Control-Allow-Origin", "*")
		req.ParseForm()
		locator := req.Form.Get("locator")
		page := req.Form.Get("page")

		api.GetBookPage(locator, page, res)
	})
	http.HandleFunc("/pack", func(res http.ResponseWriter, req *http.Request) {
		res.Header().Set("Access-Control-Allow-Origin", "*")
		res.Header().Set("Content-Type", "text/json; charset=utf-8")
		req.ParseForm()
		path := req.Form.Get("path")
		dst := req.Form.Get("dst")

		api.MakeErisFromFolder(path, dst)
	})
	port := 3455
	if len(os.Args) > 1 {
		port, _ = strconv.Atoi(os.Args[1])
	}
	log.Fatal(http.ListenAndServe(fmt.Sprintf("127.0.0.1:%d", port), nil))
}

//func mainx() {
//	server.HandleFunc("a", func() {
//		fmt.Println(1)
//	})
//	for {
//		var v string
//		fmt.Scan(&v)
//		server.Handle("a")
//	}
//}
