package main

//import (
//	yb "Test/control"
//	handler "Test/handler"
//	"log"
//	"net/http"
//)
//
//func main() {
//	http.HandleFunc("/", yb.Tes)       // 设置访问的路由
//	err := http.ListenAndServe(":9090", nil) // 设置监听的端口
//	if err != nil {
//		log.Fatal("ListenAndServe: ", err)
//	}
//}

import (
	"Test/handler"
)

func main() {
	handler.Init()
	handler.BusinessHandleW()
}
