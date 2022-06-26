# 项目结构Review



## 项目说明

一个用户组管理的微服务

* 用户组的增删改查
* 用户组与用户查询，绑定与解绑



## 结构总览

```bash
.
├── api
│   ├── goapi<grpc框架自动生成的代码桩，可不理会...>
│   │   └── janus-biz-groups
│   │       └── service
│   │           └── v1
│   │               ├── group.pb.go
│   │               ├── group.transport_grpc.pb.go
│   │               └── group.transport.pb.go
│   └── proto
│       └── janus-biz-groups
│           └── service
│               └── v1
│                   ├── group.proto
│                   └── group.transport.proto
├── cmd
│   └── server
│       └── main.go
├── config
│   └── config.yaml
├── internal
│   ├── domain
│   │   ├── entity
│   │   │   ├── group.go
│   │   │   └── member.go
│   │   ├── repository
│   │   │   ├── group.go
│   │   │   └── repository.go
│   │   ├── service
│   │   │   ├── group.go
│   │   │   ├── helper
│   │   │   │   ├── dependency.go
│   │   │   │   ├── dependency_test.go
│   │   │   │   ├── group.go
│   │   │   │   └── group_test.go
│   │   │   └── service.go
│   │   └── valueobject
│   │       └── group.go
│   ├── infrastructure
│   │   ├── component
│   │   │   ├── ent.go
│   │   │   └── grpc.go
│   │   ├── config
│   │   │   ├── config.pb.go
│   │   │   ├── config.proto
│   │   │   └── reader.go
│   │   └── constructor
│   │       ├── provider.go
│   │       ├── wire_gen.go
│   │       └── wire.go
│   ├── interface
│   │   ├── datasource
│   │   │   ├── datasource.go
│   │   │   ├── group.go
│   │   │   └── sqldb
│   │   │       ├── ent<orm框架自动生成的代码，可不理会>
│   │   │       ├── group.go
│   │   │       ├── helper.go
│   │   │       └── sqldb.go
│   │   └── transport
│   │       ├── group.go
│   │       ├── helper
│   │       │   └── group.go
│   │       ├── interceptor.go
│   │       └── transport.go
│   └── usecase
│       ├── group.go
│       ├── helper.go
│       └── usecase.go
└── unittest
    ├── client.go
    ├── group_test.go
    └── mock.go

39 directories, 75 files
```



## 结构说明

### api

存放proto buffer与使用grpc工具编译protobuf生成的代码桩



### cmd

项目启动的入口

#### - server

表示是作为一个服务端启动，main.go是启动项目的入口文件



### config

存放启动项目的配置文件



### internal

存放项目源码（go的机制是internal文件夹内的东西无法它所在的目录以外的代码调用，用来保护源码防止被其他人恶意依赖....）



#### - domain

领域层，存放实体，值对象，仓储定义，服务等

本层仅能依赖go自身的标准库

##### - entity

存放实体定义，本项目的实体为用户组 Group，有id，名称，描述，父级的组的id，用户组成员id共五个属性

##### - repository

存放仓储的接口定义，对应的实现的位置是`internal.interface.datasource`

存放以下定义：

* 查找用户组
* 查找用户组并返回符合条件的记录总数
* 创建用户组
* 更新用户组
* 移除（软删除）用户组
* 按单个用户成员id查找用户组
* 查找用户组成员
* 添加用户组成员
* 移除用户组成员

##### - service

领域服务

本项目中的领域服务有三个内容

* 查询特定id的组是否存在
* 查询group实体中的父级组id是否存在
* 检查组之间是否存在循环依赖的关系，如a的父级组是b，b的父级组是a这种情况

##### - valueobject

值对象，定义如下内容

* 用户组查询器
* 错误定义:用户组不存在的



#### - usecase

用例层，按照业务需求组合领域层的内容

本层仅能依赖go标准库与domain层的内容

定义了用户组与用户组成员的增删改查的用例



#### - interface

适配器层

* datasource：存储的实现
* transport：系统的输入

##### - datasource

数据源，是[repository](#- repository)的实现

###### - sqldb

定义与关系型数据库的交互方法

###### - cache

定义与缓存的交互方法

###### - rpcclient

定义与其它微服务的交互方法

###### - mq

定义与消息队列的交互方法



##### - transport

实现grpc编译后的代码桩的接口中的方法，用来转换系统外部的输入与输出

输入的内容经过transport层处理，进入到usecase层，usecase层处理完毕后再回到transport层然后输出给外部调用者



#### - infrastructure

基础设施层，主要做以下工作：

* 初始化与数据库的连接
* 初始化grpc服务
* 依赖注入的构建
* 读取配置

##### - component

负责初始化一些基础设施的组件，比如与数据库的连接，启动一个服务等

##### - constructor

应用实例化的依赖注入工作

##### - config

负责定义配置文件的形状与读取配置文件的内容





