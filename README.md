# PeckerGo 脚手架工具

为 [peckergo](http://git.ti-ding.com/wangji/peckergo) 开发的脚手架工具

### 安装

```
// 编辑安装脚手架工具
go install 

```

### 命令

##### init

进入 `peckergo` 工程目录, 初始化 `peckergo` 工程 (将工程内的 `peckergo` 都替换成项目名字)

```
pecker init
```

##### model

创建 model (`model_name` 为小写下划线格式)
为项目生成空 `model` 代码， `CRUD` 相关后端 `controller`、`router` 代码, 以及前端 `router`, `store`, `view` 代码

```
pecker model {model_name}
```