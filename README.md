# cron.navjs.cn

## 说明

TimeCron是一个用于创建和管理定时任务的轻量级系统。它支持跨平台,可以在Linux、Windows、MacOS等操作系统上运行。

部署简单，无需数据库,只需一个json配置文件即可

您可以通过官方网站查看详细的内容  [cron.navjs.cn](http://cron.navjs.cn)


## 本地开发构建

您可以根据自己需要自行构建目标平台可执行文件,或者自行构建docker容器镜像

```bash
git clone https://github.com/xnkyn/timecron.git

go build
```

此项目要求golang 1.16及以上


## 技术栈

后端 golang
前端 vue3

## 使用过程中遇到问题

您可以通过 [npee社区](http://www.npee.cn) 参与交流

开机自启  [如何设置开机自启](https://www.npee.cn/d/20-jian-yi-jia-ru-kai-ji-qi-dong/2) 在这里我详细的讲解了在linux下的开机自启配置过程

## 内网穿透

在开发中需要公网访问本地项目,推荐使用 [frpee.com 免费内网穿透](http://frpee.com) 进行映射解决

## 开源说明

受到很多使用者反馈参与开发维护，现已完全免费开源,可用于任何场景

您可以提交pr参与项目开发或优化

使用本开源项目，请合法使用，禁止用于违法犯罪用途

![alt text](image/README/image.png)