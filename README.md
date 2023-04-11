# wwbtg

> 这是一个模仿php电商项目的go微服务框，会尽可能保持项目目录结构的不变，中间的一些php组件会被go技术栈替换。

原框架由下图所示
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

## 功能列表

### Common类模块
- [x] 加载框架基础依赖库
  - [ ] json处理组件替换
  - [ ] rpc调用组件替换
  - [ ] db连接管理组件替换
- [x] 实现不同模块（UI、Server）服务启动环境
  - [ ] http请求连接管理组件`php-fmp > multiplexer(net/http)`
- [ ] 实现静态资源服务

### UI类模块
- [x] 提供基本路由配置
- [x] 提供接口
  - [x] 实现一个`invoke()`
  - [x] 默认请求返回格式`json`
  - [ ] 支持其他返回格式

### Server类模块
- [ ] 实现服务注册配置