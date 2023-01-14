# 介绍

Docker提供了Docker Engine API与Docker daemon交互，Docker提供了Go和Python版本的SDK。本质上Docker Engine API是一个RESTful的接口，可以通过HTTP请求操作。

本项目使用Go语言调用Docker的api接口，并将服务注册到DBus中，提供给其他程序调用。

项目会作为集成环境管理软件Docker管理的后端部分。

项目会使用到deepin社区开发的golib库，引用部分项目未使用go mod进行管理，所该项目也暂时不使用go mod

```go
# 取消go mod管理
export GO111MODULE=off
```

# 参考

[Dockore](https://github.com/HsOjo/Dockore) 该项目是使用Python调用了Docker的api接口，参考该项目减少了很多开发时间，如果觉得软件不错可以给大佬点个star

[go-lib](https://github.com/linuxdeepin/go-lib)  Deepin GoLang Library是一个包含许多有用的go例程的库，例如glib, gettext, archive, graphic等。
