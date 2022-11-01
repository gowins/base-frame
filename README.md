```
├── common                              -----  公共包
│  ├── env                              ----- 环境变量
│  │  └── env.go
│  │  └── env_test.go
│  ├── constant
│  │  └── error.go                      ----- 错误码
│  ├── model                            ----- gorm默认生成目录
│  └── repository                       ----- gorm默认生成目录
├── go.mod
├── go.sum
├── README.md
└── service                             ----- 服务
   ├── bar-single                       ----- 单体服务
   │  ├── Dockerfile
   │  ├── etc                           ----- 配置文件
   │  │  └── user.yaml
   │  ├── internal                      ----- 私有包
   │  │  ├── config                     ----- 配置
   │  │  │  └── config.go
   │  │  ├── handler                    ----- 控制器
   │  │  │  ├── example
   │  │  │  │  └── example.go
   │  │  │  ├── middleware              ----- 中间件
   │  │  │  │  ├── auth_middle.go
   │  │  │  │  └── context_middle.go
   │  │  │  └── routes.go               ----- 路由
   │  │  ├── healthy_checker            ----- 健康检测
   │  │  │  └── healthy_checkout.go
   │  │  ├── logic                      ----- 业务逻辑
   │  │  │  └── getuserlogic.go
   │  └── user-api.go                   ----- 入口文件(可执行文件)
   └── foo-multi                        ----- 包含多个服务
      ├── ctlcmd                        ----- 命令行服务
      │  ├── ctl.go
      │  ├── Dockerfile
      │  └── manifests
      │     └── deploy.yaml
      ├── etc                           ----- 配置文件
      │  └── user-api.yaml
      ├── grpccmd                       ----- grpc服务
      │  ├── Dokcerfile
      │  ├── manifests
      │  │  └── deploy.yaml
      │  ├── user-rpc.go
      │  └── user.proto
      └── internal
         ├── config
         │  └── config.go
         ├── handler
         │  ├── foo
         │  │  └── v1
         │  │     └── loginhandler.go
         │  └── routes.go
         ├── logic
         │  └── foo
         │     └── v1
         │        └── loginlogic.go
         ├── rpc_client
         │  └── rpc_client.go
         ├── rpc_server
         │  └── rpc_server.go
         ├── types
         │  └── types.go
         └── user                       ----- protoc生成的代码
            ├── user.pb.go
            ├── user.proto
            └── user_grpc.pb.go
```
