# wwbtg

> 这是一个模仿php电商项目的go微服务框，会尽可能保持项目目录结构的不变，中间的一些php组件会被go技术栈替换。

1. 原框架代码组织架构如下图所示
```text
      +----------------------------------------------------------+
      |                         GW                               |
      +--------------------------|-------------------------------+
      |                        Nginx                             |
      +--------------------------|-------------------------------+
      |                        <App>                             |
      |                          |                               |
      |                          |                               |
      +---------- UI ------------+------ Service ----------------+
      |  uiorder  |  uisale   |       |      |      |     |      |
      +-----------------------+ order | sale | item | ugc | user |
      |         uilib         |       |      |      |     |      |
      +----------------------------------------------------------+
      |           wwbtg       util        gconfig                |
      +----------------------------------------------------------+
```

2. 服务监控系统
```text
log file -> filebeat -> kafka -> logstash -> es -> grafana
```

3. 数据管理组织架构
```text
+------------------------------ relation database ------------------------------+
|        [ mysql ]                                                              |
|            |                                                                  |
|         <binlog>                                                              |
|            |                                                                  |
|        [ kafka ]                                                              |
|            |                                                                  |
|            ---->>> [ RealTime Data WorkHouse ] (ALI - hologress)              |
|            |                                                                  |
|            ---->>> [ ETL ] >>> [ Data WorkHouse ] (AWS - redshift)            |
|            |                                                                  |
|            ---->>> [ RealTime Stream Compute ] (ALI - flink) ------+          |
|                                                                    |          |
+------------------------------ no relation db ----------------------|----------+
|                                                                    v          |
|        [ redis ]                                              [ dynamodb ]    |
+-------------------------------------------------------------------------------+
```

## 功能列表

### Common类模块
- [x] 加载框架基础依赖库
  - [ ] json处理组件替换
  - [x] rpc调用组件替换
    - [x] net/rpc实现
    - [ ] grpc实现
  - [x] db连接管理组件替换
    - [x] go/mysql驱动连接池管理
- [x] 实现不同模块（UI、Server）服务启动环境
  - [x] http请求连接管理组件`php-fmp > multiplexer(net/http)`
- [ ] 实现静态资源服务

### UI类模块
> 以`uibook`模块为例

- [x] 提供基本路由配置
- [x] 提供接口
  - [x] 实现一个`invoke()`
  - [x] 默认请求返回格式`json`
  - [ ] 支持其他返回格式

### Server类模块
> 以`user`模块为例

- [x] 实现服务注册配置
- [x] 功能接口层（组合多个数据实体对外提供的最小单位）
    - [x] 默认rpc请求格式`json`
- [x] 数据访问层
    - [x] 关系数据库实现
    - [ ] `nosql`数据库实现

### 网关服务
- [ ] 服务注册、发现
- [ ] 服务链路追踪
- [ ] 基于日志的服务监控系统

### 消息中间件
- [ ] `kafka`

### 数据库中间件
- [ ] `binlog`
