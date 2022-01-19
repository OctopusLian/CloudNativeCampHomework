<!--
 * @Author: zhangniannian
 * @Date: 2022-01-19 21:06:09
 * @LastEditors: zhangniannian
 * @LastEditTime: 2022-01-19 21:07:10
 * @Description: 请填写简介
-->
## 作业1  

除了将 httpServer 应用优雅的运行在 Kubernetes 之上，我们还应该考虑如何将服务发布给对内和对外的调用方。
来尝试用 Service, Ingress 将你的服务发布给集群外部的调用方吧  
在第一部分的基础上提供更加完备的部署 spec，包括（不限于）  

- Service  
- Ingress  

可以考虑的细节
- 如何确保整个应用的高可用
- 如何通过证书保证 httpServer 的通讯安全
- 结合[模块8的课后作业](../lesson8/README.md)