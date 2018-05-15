## 概述
    关于go的各种各样项目。
## 目录
#### commonTools
    封装的一些go的工具。
#### designPatternsDemo
    根据对设计模式的理解，用go实现的Demo。
#### goServer
[xiaohe.info](http://xiaohe.info/) 的代码总集，以下描述各个部分的作用：
* chatServer:聊天服务器，包括了`webSocketServer` `webServer` `rpcServer`三种通信方式。
* fileServer:文件服务器，支持分片上传。
* goClient:聊天服务器的测试客户端。
* goServer:一个业务服务器，支持用户注册登录等操作，与`chatServer` `fileServer` `httpServer` `webSite` 共同支撑了`xiaohe.info`的个人网站。
* goServerModel：model库。
* httpServer：一个静态页面的服务器。
* webSite：`xiaohe.info`的静态网页文件。
#### taskManager
    任务管理器，一个自用的定时邮件提醒的项目。
#### test
* 【Go并发编程实战 第2版】的一些笔记测试Demo。
* 其他有意义的测试项目。
#### vendor
	引用项目。