# ws_chat即时通信系统

![Go](https://img.shields.io/badge/Go-1.18-blue.svg)
![go-zero](https://img.shields.io/badge/go_zero-1.3.3-blue.svg)
![star](https://img.shields.io/github/stars/wslynn/ws_chat?style=social)


## 1 项目简介

个人出于兴趣做的开源项目，目标是实现一个类似于微信的即时通信软件!

项目正在稳定开发中, 敬请期待...

为了一套代码支持全平台，前端使用Flutter+GetX开发。

为了稳定可扩展，后端使用了集成了众多工程最佳实践的go-zero微服务框架。


## 2 相关开源地址
后端开源地址：https://github.com/wslynn/ws_chat

前端开源地址：https://github.com/wslynn/ws_chat_flutter


## 3 架构图及文档
[语雀在线文档](https://www.yuque.com/docs/share/77c846d2-51f8-4a25-8330-fa036a8a4cbe)

[![OZycv9.png](https://s1.ax1x.com/2022/05/05/OZycv9.png)](https://imgtu.com/i/OZycv9)


## 4 modd (hot reload)
### (1) 下载modd
> go install github.com/cortesi/modd/cmd/modd
### (2) 写配置文件modd.conf
```conf
# user
app/user/rpc/**/*.go {
    prep: go build -o ./bin/user-rpc  -v app/user/rpc/user.go
    daemon +sigkill: ./bin/user-rpc -f app/user/rpc/etc/user.yaml
}
app/user/api/**/*.go {
    prep: go build -o ./bin/user-api  -v app/user/api/user.go
    daemon +sigkill: ./bin/user-api -f app/user/api/etc/user.yaml
}
```
### (3) 运行modd
> root@tencent:~/code/ws_chat (master) # modd