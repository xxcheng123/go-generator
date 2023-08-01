## `GO` 开发脚手架 *go-generator*
### 说明
`CLD` 分层
```tree
|-- config.yaml
|-- controllers
|-- dao        
|-- go.mod     
|-- go.sum     
|-- log.log
|-- logger
|-- logic
|-- main.go
|-- models
|-- pkg
|-- readme.md
|-- routers
`-- settings
```
### 启动顺序
1. 加载配置信息
2. 初始化配置日志系统
3. 加载数据库
   - *MySQL*
   - *Redis*
4. 注册路由
   - 中间件注册
5. 优雅关机

### 相关第三方库
- *gin*
- *viper*
- *zap*