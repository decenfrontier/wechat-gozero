# wechat-gozero

![Go](https://img.shields.io/badge/Go-1.18-blue.svg)
![go-zero](https://img.shields.io/badge/go_zero-1.3.3-blue.svg) 
![goctl](https://img.shields.io/badge/goctl-1.3.5-blue.svg)
![star](https://img.shields.io/github/stars/wslynn/wechat-gozero?style=social)


## 1 项目简介
本项目的目标是模仿微信UI开发一个IM系统, 实现消息单聊, 群聊, 在线推送, 离线拉取等功能

为了提升开发和重构效率，后端使用go-zero微服务框架。

为了一套代码支持全平台和便捷的状态管理，前端使用Flutter+GetX开发。

都是比较新的技术栈, 如果对您有帮助, 请点个star支持一下

免责声明: 本代码仅供技术交流，请勿用于商业用途

现已实现单聊功能(加好友目前需要通过post请求后端接口实现, 前端页面暂未完成)
![784ea66737aaaad8d8e0cb84efdc7b63.gif](https://img.gejiba.com/images/784ea66737aaaad8d8e0cb84efdc7b63.gif)


## 2 相关开源地址
后端开源地址：https://github.com/wslynn/wechat-gozero

前端开源地址：https://github.com/wslynn/wechat_flutter


## 3 架构图及文档
### (1) 架构图
[![XbgESH.png](https://s1.ax1x.com/2022/06/17/XbgESH.png)](https://imgtu.com/i/XbgESH)
### (2) 文档
[语雀在线文档](https://www.yuque.com/docs/share/77c846d2-51f8-4a25-8330-fa036a8a4cbe)
### (3) API接口文档
[ApiPost接口文档](https://console-docs.apipost.cn/preview/6c245af8bcc075c4/42820335d3df842c)

## 4 快速开始
### (1) 下载依赖包
> $ go mod tidy
### (2) 启动环境
> $ docker-compose up -d
### (3) kafka配置
创建topic, partition, replica
[![Xq4Rds.png](https://s1.ax1x.com/2022/06/17/Xq4Rds.png)](https://imgtu.com/i/Xq4Rds)
### (4) 运行各个微服务
> $ modd

## 5 常见问题及解决方案
### (1) elasticsearcj容器报错nested: AccessDeniedException[/usr/share/elasticsearch/data/nodes]
解决: 
> $ chmod 777 data/elasticsearch/**

## 6 接下来的开发计划

- 搜索用户
- 查看用户信息
- 获取好友列表
- 发起群聊
- 修改用户信息
- 修改密码
- 上传图片和视频到云存储
- 保存用户在线状态

空闲时间会继续更新, 直至实现微信的大部分功能, 也欢迎各位大佬PR~