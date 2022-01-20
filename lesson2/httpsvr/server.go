/*
 * @Author: zhangniannian
 * @Date: 2022-01-20 21:06:49
 * @LastEditors: zhangniannian
 * @LastEditTime: 2022-01-20 21:12:19
 * @Description: 请填写简介
 */
package main

import (
	"fmt"
	"net/http"
)

func myHandler(w http.ResponseWriter, r *http.Request) {
	fmt.Println("header = ", r.Header) // 请求头  输出：header =  map[Accept-Encoding:[gzip] User-Agent:[Go-http-client/1.1]]
	w.Write([]byte("hello http"))      //给客户端回复数据
}

func main() {
	http.HandleFunc("/hello", myHandler) // 注册处理函数
	http.ListenAndServe("127.0.0.1:8000", nil)
}
