
mytarget: #make mytarget s=order e=api
	@echo $s-$e

run:
	@go run ./app/$s/api/$s.go  -f ./app/$s/api/etc/$s-api.yaml

run-user-api: # 启动 user-api 服务
	@go run ./app/user/api/user.go  -f ./app/user/api/etc/user-api.yaml

run-user-rpc: # 启动 user-rpc 服务
	@go run ./app/user/rpc/user.go  -f ./app/user/rpc/etc/user.yaml

ctl-user-api: # 初始化user-api服务
	@goctl api go -api ./app/user/api/user.api -dir ./app/user/api --style=go_zero

ctl-user-rpc:
	@goctl rpc protoc ./app/user/rpc/user.proto --go_out=./app/user/rpc/ --go-grpc_out=./app/user/rpc/ --zrpc_out=./app/user/rpc --style=go_zero

ctl-model:
	@goctl model mysql datasource -url="root:root@tcp(127.0.0.1:3306)/zero-shop" -table="order"  -dir="./genModel"  --style=go_zero

ctl-goods-api: # 初始化user-api服务
	@goctl api go -api ./app/goods/api/goods.api -dir ./app/goods/api --style=go_zero

ctl-goods-rpc:
	@goctl rpc protoc ./app/goods/rpc/goods.proto --go_out=./app/goods/rpc/ --go-grpc_out=./app/goods/rpc/ --zrpc_out=./app/goods/rpc --style=go_zero

run-goods-api: # 启动 goods-api 服务
	@go run ./app/goods/api/goods.go  -f ./app/goods/api/etc/goods-api.yaml

run-goods-rpc: # 启动 goods-rpc 服务
	@go run ./app/goods/rpc/goods.go  -f ./app/goods/rpc/etc/goods.yaml

ctl-order-rpc: #初始化订单服务
	@goctl rpc protoc ./app/order/rpc/order.proto --go_out=./app/order/rpc/ --go-grpc_out=./app/order/rpc/ --zrpc_out=./app/order/rpc --style=go_zero

run-order-rpc: # 启动 order-rpc 服务
	@go run ./app/order/rpc/order.go  -f ./app/order/rpc/etc/order.yaml