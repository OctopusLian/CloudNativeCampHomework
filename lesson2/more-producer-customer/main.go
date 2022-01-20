/*
 * @Author: zhangniannian
 * @Date: 2022-01-20 20:46:02
 * @LastEditors: zhangniannian
 * @LastEditTime: 2022-01-20 20:52:43
 * @Description: 请填写简介
 */
package main

import "fmt"

//生产者
func Producer(id int, ch chan int) {
	for i := 1; i <= 20; i++ {
		fmt.Printf("id: %d, send: %d\n", id, i)
		ch <- i //生成者往队列塞数据
	}
	close(ch) //塞完后关闭通道

}

//消费者
func Consumer(id int, ch chan int, done chan bool) {
	for {
		value, ok := <-ch //消费者从队列中拿数据
		if ok {
			//有数据，打印
			fmt.Printf("id: %d, recv: %d\n", id, value)
		} else {
			//没有数据，打印closed
			fmt.Printf("id: %d,closed\n", id)
			break
		}
	}
	done <- true
}

func main() {
	ch := make(chan int, 10) //创建一个长度为10，元素类型为int的队列
	done := make(chan bool)  //创建一个通知停止的通道
	coNum := 2
	prNum := 2

	for i := 1; i <= coNum; i++ {
		go Consumer(i, ch, done)
	}
	for i := 1; i <= prNum; i++ {
		go Producer(i, ch)
	}

	for i := 1; i <= coNum; i++ {
		<-done
	}

}
