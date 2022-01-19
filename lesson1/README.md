<!--
 * @Author: zhangniannian
 * @Date: 2022-01-19 20:41:16
 * @LastEditors: zhangniannian
 * @LastEditTime: 2022-01-19 20:42:37
 * @Description: 请填写简介
-->
## 课后练习1.1  

- 安装 Go • 安装 IDE 并安装 Go 语言插件  
- 安装 IDE 并安装 Go 语言插件  
- 编写一个小程序  
```
给定一个字符串数组
["I","am","stupid","and","weak"]
用 for 循环遍历该数组并修改为
["I","am","smart","and","strong"]
```

## 课后练习1.2  

- 基于 Channel 编写一个简单的单线程生产者消费者模型  

队列：队列长度10，队列元素类型为 int  
生产者：每1秒往队列中放入一个类型为 int 的元素，队列满时生产者可以阻塞  
消费者：每一秒从队列中获取一个元素并打印，队列为空时消费者阻塞  