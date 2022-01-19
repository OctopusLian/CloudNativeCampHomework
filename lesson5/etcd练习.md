<!--
 * @Author: zhangniannian
 * @Date: 2022-01-19 20:58:09
 * @LastEditors: zhangniannian
 * @LastEditTime: 2022-01-19 21:00:34
 * @Description: 请填写简介
-->
## etcd练习  

查看集群成员状态
```
etcdctl member list --write-out=table
+------------------+---------+---------+-----------------------+-----------------------+------------+
| ID | STATUS | NAME | PEER ADDRS | CLIENT ADDRS | IS LEARNER |
+------------------+---------+---------+-----------------------+-----------------------+------------+
| 8e9e05c52164694d | started | default | http://localhost:2380 | http://localhost:2379 | false |
+------------------+---------+---------+-----------------------+-----------------------+------------+
```

基本的数据读写操作  

- 写入数据  

```
etcdctl --endpoints=localhost:12379 put /a b
OK
```

- 读取数据  

```
etcdctl --endpoints=localhost:12379 get /a
/a
b
```

- 按key的前缀查询数据  

```
etcdctl --endpoints=localhost:12379 get --prefix /
```

- 只显示键值  

```
etcdctl --endpoints=localhost:12379 get --prefix / --keys-only --debug
```

- 修改值  

```
etcdctl put x 1
```

- 查询最新值  

```
etcdctl get x
x
1
```

- 查询历史版本值  

```
etcdctl get x --rev=2
x
0
```