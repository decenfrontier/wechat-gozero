# wechat-gozero

![Go](https://img.shields.io/badge/Go-1.18-blue.svg)
![go-zero](https://img.shields.io/badge/go_zero-1.3.3-blue.svg) 
![goctl](https://img.shields.io/badge/goctl-1.3.5-blue.svg)
![star](https://img.shields.io/github/stars/wslynn/wechat-gozero?style=social)


## 1 项目简介
本项目的目标是利用websocket, kafka等技术, 模仿微信UI开发一个IM系统, 实现消息单聊, 群聊, 在线推送, 离线拉取等功能

为了提升开发和重构效率，后端使用go-zero微服务框架。

为了一套代码支持全平台和便捷的状态管理，前端使用Flutter+GetX开发。

都是比较新的技术栈, 网上参考资料不多, 一步步做到现在不容易, 如果对您有帮助, 请点个star支持一下 

有时间会继续更新, 直至实现微信的大部分功能, 也欢迎各位大佬PR~


## 2 相关开源地址
后端开源地址：https://github.com/wslynn/wechat-gozero

前端开源地址：https://github.com/wslynn/wechat_flutter


## 3 架构图及文档
[语雀在线文档](https://www.yuque.com/docs/share/77c846d2-51f8-4a25-8330-fa036a8a4cbe)

[![XbPYqA.png](https://s1.ax1x.com/2022/06/16/XbPYqA.png)](https://imgtu.com/i/XbPYqA)

## 4 常见问题及解决方案
### (1) elasticsearcj容器报错nested: AccessDeniedException[/usr/share/elasticsearch/data/nodes]
解决: 
> chmod 777 data/elasticsearch/**
