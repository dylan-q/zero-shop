## 基于go-zero的微服务项目 zero-shop

```bash
https://github.com/zeromicro/go-zero
```

- 开发环境 go1.20.4 + Win11
- 开发工具 goland

### 开发步骤

- 新建业务表，根据业务表的数据生成 model
- 新建 proto 文件，根据proto文件生成`rpc`服务
- 新建业务 api 文件，根据 api 文件生成`api`服务
- 修改配置，让api 服务连接 rpc 服务