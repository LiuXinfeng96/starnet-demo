# starnet-demo

## docker 一键部署

### 在线环境准备

- docker

### 离线环境依赖

- docker离线安装包
- mysql: 8.0 镜像包
- nginx: 1.23.3 镜像包
- starnet-demo 镜像包

### demo部署包目录结构

```shell
├── conf
│   ├── masterchain_sdk_config.yaml
│   ├── satellitechain_sdk_config.yaml
│   └── system_config.yaml
├── crypto-config
│   ├── exec-chain
│   │   ├── ca
│   │   └── user
│   ├── master-chain
│   │   ├── ca
│   │   └── user
│   ├── wx-org1.chainmaker.org
│   │   ├── ca
│   │   ├── node
│   │   └── user
│   ├── wx-org2.chainmaker.org
│   │   ├── ca
│   │   ├── node
│   │   └── user
│   ├── wx-org3.chainmaker.org
│   │   ├── ca
│   │   ├── node
│   │   └── user
│   └── wx-org4.chainmaker.org
│       ├── ca
│       ├── node
│       └── user
├── deploy.sh
├── DockerFile
├── go.mod
├── go.sum
├── log
│   └── default.log
├── README.md
├── src
│   ├── configs
│   │   ├── config.go
│   │   └── config_test.go
│   ├── contract
│   │   ├── convert_params.go
│   │   ├── exec.go
│   │   └── master.go
│   ├── db
│   │   ├── block_chain_identity.go
│   │   ├── common.go
│   │   ├── comm_state.go
│   │   ├── constellation.go
│   │   ├── control.go
│   │   ├── db.go
│   │   ├── debris.go
│   │   ├── fault.go
│   │   ├── instruction.go
│   │   ├── login_log.go
│   │   ├── net_state.go
│   │   ├── operation.go
│   │   ├── orbit.go
│   │   ├── satellite.go
│   │   └── user.go
│   ├── handlers
│   │   ├── common.go
│   │   ├── comm_state.go
│   │   ├── constellation.go
│   │   ├── control.go
│   │   ├── debris.go
│   │   ├── fault.go
│   │   ├── instruction.go
│   │   ├── login.go
│   │   ├── monitor.go
│   │   ├── net_state.go
│   │   ├── operation.go
│   │   ├── orbit.go
│   │   └── satellite_state.go
│   ├── loggers
│   │   ├── gin_logger.go
│   │   ├── gorm_logger.go
│   │   └── logger.go
│   ├── main.go
│   ├── models
│   │   ├── contract.go
│   │   ├── req.go
│   │   └── resp.go
│   ├── routers
│   │   └── routers.go
│   └── services
│       ├── block_chain_client.go
│       ├── contract.go
│       ├── curd.go
│       ├── server.go
│       └── token.go
└── web
    ├── conf.d
    │   └── nginx.conf
    └── resources
        ├── assets
        └── index.html
```

### 部署前必要配置修改

#### 数据库配置

配置目录：conf/system_config.yaml

配置内容：

```yaml
db_config:
  user: root
  password: 123456
  ip: 172.16.2.225
  port: 13306
  dbname: chainmaker
```

- user：数据库用户名
- password：数据库密码
- ip：数据库IP地址
- port：数据库端口
- dbname：数据库名称

#### 区块链配置

配置目录：conf/system_config.yaml

配置内容：

```yaml
block_chain_config:
  - user_name: master
    user_pwd: 1234
    sdk_config_path: ../conf/masterchain_sdk_config.yaml
  - user_name: exec
    user_pwd: 1234
    sdk_config_path: ../conf/satellitechain_sdk_config.yaml
master_contract: 
  name: masterC
  chain_id: chain_dcwjj
exec_contract: 
  name: execC
  chain_id: chain_dcwjc
```

- block_chain_config：区块链SDK配置

  user_name：内置绑定sdk证书的用户名称

  user_pwd：内置绑定sdk证书的用户密码

  sdk_config_path：sdk配置文件路径

- master_contract：主链合约名称和链ID配置

  name：合约名称

  chain_id：链ID

- exec_contract：星座链合约名称和链ID配置

  name：合约名称

  chain_id：链ID

#### 区块链SDK配置文件

配置目录：conf/masterchain(satellitechain)_sdk_config.yaml

配置内容：

```yaml
chain_client:
  # 链ID
  chain_id: "chain_dcwjj"
  # 组织ID
  org_id: "69vkhyyorg.cm-3yt9l7uir5"
  # 客户端用户私钥路径
#  user_key_file_path: "../crypto-config/wx-org1.chainmaker.org/user/client1/client1.tls.key"
  user_key_file_path: "../crypto-config/master-chain/user/client1/client1.tls.key"
  # 客户端用户证书路径
#  user_crt_file_path: "../crypto-config/wx-org1.chainmaker.org/user/client1/client1.tls.crt"
  user_crt_file_path: "../crypto-config/master-chain/user/client1/client1.tls.crt"
  # 客户端用户交易签名私钥路径(若未设置，将使用user_key_file_path)
#  user_sign_key_file_path: "../crypto-config/wx-org1.chainmaker.org/user/client1/client1.sign.key"
  user_sign_key_file_path: "../crypto-config/master-chain/user/client1/client1.sign.key"
  # 客户端用户交易签名证书路径(若未设置，将使用user_crt_file_path)
#  user_sign_crt_file_path: "../crypto-config/wx-org1.chainmaker.org/user/client1/client1.sign.crt"
  user_sign_crt_file_path: "../crypto-config/master-chain/user/client1/client1.sign.crt"
  # 同步交易结果模式下，轮询获取交易结果时的最大轮询次数，删除此项或设为<=0则使用默认值 10
  retry_limit: 10
  # 同步交易结果模式下，每次轮询交易结果时的等待时间，单位：ms 删除此项或设为<=0则使用默认值 500
  retry_interval: 500
  # 当前签名证书的别名。当设置此配置项时，chain client 对象将自动检查链上是否已添加此别名，如果没有则自动上链此证书别名，
  # 并且后续所有交易都会使用别名，别名可降低交易体大小。若为空则不启用。
#  alias: mycert1
  nodes:
    - # 节点地址，格式为：IP:端口:连接数
      node_addr: "172.16.2.64:30494"
      # 节点连接数
      conn_cnt: 10
      # RPC连接是否启用双向TLS认证
      enable_tls: true
      # 信任证书池路径
      trust_root_paths:
#        - "../crypto-config/wx-org1.chainmaker.org/ca"
        - "../crypto-config/master-chain/ca"
      # TLS hostname
      tls_host_name: "consensus1-69vkhyyorg.cm-3yt9l7uir5"
  archive:
    # 数据归档链外存储相关配置
    type: "mysql"
    dest: "root:123456:localhost:3306"
    secret_key: xxx
  rpc_client:
    max_receive_message_size: 16 # grpc客户端接收消息时，允许单条message大小的最大值(MB)
    max_send_message_size: 16 # grpc客户端发送消息时，允许单条message大小的最大值(MB)
  pkcs11:
    enabled: false # pkcs11 is not used by default
    library: /usr/local/lib64/pkcs11/libupkcs11.so # path to the .so file of pkcs11 interface
    label: HSM # label for the slot to be used
    password: 11111111 # password to logon the HSM(Hardware security module)
    session_cache_size: 10 # size of HSM session cache, default to 10
    hash: "SHA256" # hash algorithm used to compute SKI
```

- chain_id：区块链ID
- org_id：链上组织ID
- user_key_file_path：客户端通信TLS私钥路径
- user_crt_file_path：客户端通信TLS证书路径
- user_sign_key_file_path：客户端用户交易签名私钥路径
- user_sign_crt_file_path：客户端用户交易签名证书路径
- node_addr：连接区块链节点的地址
- trust_root_paths：TLS双向认证客户端认证节点服务端证书池路径
- tls_host_name：区块链节点通信TLS证书中的域名地址

#### Nginx配置文件

配置目录：web/conf.d/nginx.conf

配置内容：

```yaml
server {
    listen       80;
    listen  [::]:80;
    server_name  _;
    #access_log  /var/log/nginx/host.access.log  main;
    location / {
        root   /usr/share/nginx/resources/;
        index  index.html index.htm;
    }
    location /satellitebc/ {
      proxy_pass http://172.16.2.225:8086/;
    }
    #error_page  404              /404.html;
    # redirect server error pages to the static page /50x.html
    #
    error_page   500 502 503 504  /50x.html;
    location = /50x.html {
        root   /usr/share/nginx/html;
    }
}
```

- proxy_pass：starnet-demo服务的地址和端口

### 使用部署脚本启动

在项目主目录下，运行部署脚本：

```shell
$ ./deploy.sh
```











