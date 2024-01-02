<p align="center">
    ACS ApiCallerStats
</p>

## 简介

ACS是一款使用go和vue3开发的用于记录API访问记录、网页浏览记录等的开源项目。它提供了一个简单易用的后台管理页面，可以查看调用折线图、饼图等，帮助您更好地了解API的使用情况以及发现一些异常。

- [GitHub](https://github.com/linkaias/acs): https://github.com/linkaias/acs.git

- [Gitee](https://gitee.com/tolinkai/acs): https://gitee.com/tolinkai/acs.git

- [在线预览](http://acsdemo.uiucode.com) 用户名：admin 密码：123456

- [文档](https://funcs.gitbook.io/acs_docs/)
  
- 小睹为快

<p align="center">
    <img width="900" src="https://818816519-files.gitbook.io/~/files/v0/b/gitbook-x-prod.appspot.com/o/spaces%2Fv9sLeXVmUlZt17fbW2U4%2Fuploads%2FGuZjYaC38YbOFZ4K19Wn%2Fimage.png?alt=media">
</p>

<p align="center">
    <img width="900" src="https://818816519-files.gitbook.io/~/files/v0/b/gitbook-x-prod.appspot.com/o/spaces%2Fv9sLeXVmUlZt17fbW2U4%2Fuploads%2FUdttMnLKLduN0fXk1McW%2Fimage.png?alt=media">
</p>

<p align="center">
    <img width="900" src="https://818816519-files.gitbook.io/~/files/v0/b/gitbook-x-prod.appspot.com/o/spaces%2Fv9sLeXVmUlZt17fbW2U4%2Fuploads%2F9Ma5TGw8wp0qWubHRem9%2Fimage.png?alt=media">
</p>

<p align="center">
    <img width="900" src="https://818816519-files.gitbook.io/~/files/v0/b/gitbook-x-prod.appspot.com/o/spaces%2Fv9sLeXVmUlZt17fbW2U4%2Fuploads%2Fm8KkmfShNRmu1W6eCJhJ%2Fimage.png?alt=media">
</p>


## 主要功能
1. 记录 API 访问记录、网页浏览记录等
2. 提供后台管理页面，查看调用折线图、饼图等
3. 支持Restful风格API和GRPC两种通信模式，快速接入您的系统。
4. 支持多种数据库记录方式。（目前仅支持mysql，其他数据库引擎开发中...）

## 使用场景
* 场景1: 小明在自己的项目中大量调用一些第三方的API服务，为了方便查看调用记录和调用频率等问题。
* 场景2: 小李有一个自己的博客网站，想观察一下自己的博客实际的访问状态，访问趋势等。
* 场景3: 小红的项目给别人提供了接口服务，为了统计别人调用自己接口的调用量和调用趋势

## 技术选型
1. 前端：用基于vue3的Element-Plus构建基础页面。
2. 后端：采用高效的go语言开发，使用Gin快速搭建基础restful风格API。
3. 鉴权：采用JWT鉴权。
4. 数据库：采用MySql>5.7版本,数据库引擎 innoDBimportant，使用gorm实现对数据库的基本操作。Clickhouse和Monge待实现。
5. 配置文件：使用viper实现.env格式的配置文件。
6. 日志：使用logrous和lfshook实现日志记录和分隔。
7. 数据通信：支持GRPC和HTTP两种通信方式。


## 原生安装

```bash
# 克隆项目
git clone https://github.com/linkaias/acs.git
# 或者
git clone https://gitee.com/tolinkai/acs.git

# 进入项目目录
cd acs

# 根据文档修改配置文件
vim .env

# 安装依赖
go mod tidy

# 启动后端服务
go run main.go

# 启动前端服务
cd web/acs_admin
# 安装依赖
npm install
# 启动服务
npm run dev

# 浏览器访问 
http://localhost:9528

```
## Docker安装

```bash
# 克隆项目
git clone https://github.com/linkaias/acs.git
# 或者
git clone https://gitee.com/tolinkai/acs.git

cd acs

# 构建镜像
docker build -t acs:1.1 .

# 启动容器 说明：.env为项目配置文件， var 为日志目录。 8080为后台管理界面端口，9528为http接口服务端口，9529为grpc服务端口。
# 注意：--network 建议和你的mysql容器网络一致。如果mysql没有使用docker安装，可以删除--network参数
docker run -d -it --network=[yourNetworkName] -v ./.env:/run/.env -v ./var:/run/var  -p 8080:80 -p 9528:9528 -p 9529:9529 --name acs acs:1.1 /bin/sh


# 打开浏览器，打开http://127.0.0.1:8080，即可访问后台登陆界面。

```

详情参考安装文档：https://funcs.gitbook.io/acs_docs/kai-shi-kai-shi/install


## Online Demo

[在线体验](http://acsdemo.uiucode.com)
http://acsdemo.uiucode.com

用户名：admin 密码：123456



**❤️ 创作不易，点个Star便是我更新的动力！您也可以分享好友圈子以支持此项目。**



## License

[MIT](https://github.com/linkaias/acs/blob/master/LICENSE)

Copyright (c) 2024-present LinKai
