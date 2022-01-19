<!--
 * @Author: zhangniannian
 * @Date: 2022-01-19 20:44:15
 * @LastEditors: zhangniannian
 * @LastEditTime: 2022-01-19 20:52:23
 * @Description: 请填写简介
-->
## Namespace 练习  

- 在新 network namespace 执行 sleep 指令：  
```
unshare -fn sleep 60
```

- 查看进程信息  

```
ps -ef|grep sleep
root 32882 4935 0 10:00 pts/0 00:00:00 unshare -fn sleep 60
root 32883 32882 0 10:00 pts/0 00:00:00 sleep 60
```

- 查看网络 Namespace  

```
lsns -t net
4026532508 net 2 32882 root unassigned unshare
```

- 进入改进程所在 Namespace 查看网络配置，与主机不一致  

```
nsenter -t 32882 -n ip a
1: lo: <LOOPBACK> mtu 65536 qdisc noop state DOWN group default qlen 1000
link/loopback 00:00:00:00:00:00 brd 00:00:00:00:00:00
```

## CPU 子系统练习  

- 在 cgroup cpu 子系统目录中创建目录结构  

```
cd /sys/fs/cgroup/cpu
mkdir cpudemo
cd cpudemo
```

- 运行 busyloop  

- 执行 top 查看 CPU 使用情况，CPU 占用 200%  

- 通过 cgroup 限制 cpu  

cd /sys/fs/cgroup/cpu/cpudemo  

- 把进程添加到 cgroup 进程配置组  

```
echo ps -ef|grep busyloop|grep -v grep|awk '{print $2}' > cgroup.procs
```

- 设置 cpuquota  

```
echo 10000 > cpu.cfs_quota_us
```

- 执行 top 查看 CPU 使用情况，CPU 占用变为10%  

## 课后练习3.1  

- Memory 子系统练习  
- 在 cgroup memory 子系统目录中创建目录结构  

```
cd /sys/fs/cgroup/memory
mkdir memorydemo
cd memorydemo
```

- 运行 malloc（在linux机器make build）  
- 查看内存使用情况  

```
watch 'ps -aux|grep malloc|grep -v grep‘
```

- 通过 cgroup 限制 memory  

- 把进程添加到cgroup进程配置组  

```
echo ps -ef|grep malloc |grep -v grep|awk '{print $2}' > cgroup.procs
```

- 设置 memory.limit_in_bytes  

```
echo 104960000 > memory.limit_in_bytes
```

- 等待进程被 oom kill  

## OverlayFS 文件系统练习  

```
$ mkdir upper lower merged work
$ echo "from lower" > lower/in_lower.txt
$ echo "from upper" > upper/in_upper.txt
$ echo "from lower" > lower/in_both.txt
$ echo "from upper" > upper/in_both.txt
$ sudo mount -t overlay overlay -o lowerdir=`pwd`/lower,upperdir=`pwd`/upper,workdir=`pwd`/work `pwd`/merged
$ cat merged/in_both.txt
$ delete merged/in_both.txt
$ delete merged/in_lower.txt
$ delete merged/in_upper.txt
```

## 课后练习 3.2  

- 构建本地镜像  
- 编写 Dockerfile 将练习 2.2 编写的 httpserver 容器化  
- 请思考有哪些最佳实践可以引入到 Dockerfile 中来  
- 将镜像推送至 Docker 官方镜像仓库
- 通过 Docker 命令本地启动 httpserver  
- 通过 nsenter 进入容器查看 IP 配置  

