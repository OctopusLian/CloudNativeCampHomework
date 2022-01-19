/*
 * @Author: zhangniannian
 * @Date: 2022-01-19 21:12:46
 * @LastEditors: zhangniannian
 * @LastEditTime: 2022-01-19 21:14:50
 * @Description: 请填写简介
 */
package main

import "fmt"

func main() {
	strArray := []string{"I", "am", "stupid", "and", "weak"}
	for index, value := range strArray {
		if value == "stupid" {
			strArray[index] = "smart"
		}
		if value == "weak" {
			strArray[index] = "strong"
		}
	}

	fmt.Println(strArray) //[I am smart and strong]
}
