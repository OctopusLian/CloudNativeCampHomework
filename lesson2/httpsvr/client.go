/*
 * @Author: zhangniannian
 * @Date: 2022-01-20 21:06:42
 * @LastEditors: zhangniannian
 * @LastEditTime: 2022-01-20 21:12:36
 * @Description: 请填写简介
 */
package main

import (
	"fmt"
	"io"
	"net/http"
)

func main() {
	resp, err := http.Get("http://127.0.0.1:8000/hello")
	if err != nil {
		fmt.Println("Get err:", err)
		return
	}
	defer resp.Body.Close()

	fmt.Println("Header = ", resp.Header) // 响应头部

	buf := make([]byte, 4096) // 定义切片缓冲区，存读到的内容
	var result string
	// 获取服务器发送的数据包内容
	for {
		n, err := resp.Body.Read(buf) // 读body中的内容。
		if n == 0 {
			fmt.Println("--Read finish!")
			break
		}
		if err != nil && err != io.EOF {
			fmt.Println("resp.Body.Read err:", err)
			return
		}
		result += string(buf[:n]) // 累加读到的数据内容
	}
	// 打印从body中读到的所有内容
	fmt.Println("result = ", result) //result =  hello http

}
