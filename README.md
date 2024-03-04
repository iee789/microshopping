


#### 创建初始化服务

```
# 货币服务
mkdir currencyservice 
cd currencyservice
go mod init currencyservice

# 商品分类服务
mkdir productcatalogservice
cd productcatalogservice
go mod init productcatalogservice

# 购物车服务
mkdir cartservice
cd cartservice
go mod init cartservice

# 广告服务
mkdir adservice
cd adservice
go mod init adservice

# 邮件服务
mkdir emailservice
cd emailservice
go mod init emailservice

# 付款服务
mkdir paymentservice
cd paymentservice
go mod init paymentservice

# 配送服务
mkdir shippingservice
cd shippingservice
go mod init shippingservice

# 推荐服务
mkdir recommendationservice
cd recommendationservice
go mod init recommendationservice

# 结算服务
mkdir checkoutservice
cd checkoutservice
go mod init checkoutservice

# 前端服务
mkdir frontend
cd frontend
go mod init frontend

# 启动微服务也应该是这个顺序，因为存在依赖关系
```

#### 使用工作空间

```
go work init
# 再对各个目录分别进行
go work use adservice
go work use currencyservice
go work use productcatalogservice
go work use cartservice
go work use emailservice
go work use paymentservice
go work use shippingservice
go work use recommendationservice
go work use checkoutservice
go work use frontend
# 同时每个目录再go mod tidy
```

#### protobuf相关库的操作

安装库

```
#先get再install，否则可能会出现问题
go get google.golang.org/protobuf/cmd/protoc-gen-go
go install google.golang.org/protobuf/cmd/protoc-gen-go

go get google.golang.org/grpc/cmd/protoc-gen-go-grpc
go install google.golang.org/grpc/cmd/protoc-gen-go-grpc
```

切换到具体服务目录中，比如currencyservice中执行如下命令得到文件，其他目录也一样：

```
protoc --go_out=./ proto/*.proto
protoc --go-grpc_out=require_unimplemented_servers=false:./ proto/*.proto
```

#### 服务注册到consul

```
# 启动consul时同时开启ip权限
consul agent -dev -client 0.0.0.0 -ui
# 浏览器中http://localhost:8500:ui/dc1/services查看情况

# api包： github.com/hashicorp/consul/api
```

