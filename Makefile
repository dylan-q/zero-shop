


run-user-api: # 启动 user-api 服务
	@go run ./app/user/api/user.go  -f ./app/user/api/etc/user-api.yaml

run-user-rpc: # 启动 user-rpc 服务
	@go run ./app/user/rpc/user.go  -f ./app/user/rpc/etc/user.yaml

ctl-user-api: # 初始化user-api服务
	@goctl api go -api ./app/user/api/user.api -dir ./app/user/api --style=go_zero

ctl-user-rpc:
	@goctl rpc protoc ./app/user/rpc/user.proto --go_out=./app/user/rpc/ --go-grpc_out=./app/user/rpc/ --zrpc_out=./app/user/rpc --style=go_zero

ctl-category-model:
	@goctl model mysql datasource -url="root:root@tcp(127.0.0.1:3306)/zero-shop" -table="goods"  -dir="./genModel"  --style=go_zero

ctl-goods-rpc:
	@goctl rpc protoc ./app/goods/rpc/goods.proto --go_out=./app/goods/rpc/ --go-grpc_out=./app/goods/rpc/ --zrpc_out=./app/goods/rpc --style=go_zero

run-goods-rpc: # 启动 user-rpc 服务
	@go run ./app/goods/rpc/goods.go  -f ./app/goods/rpc/etc/goods.yaml
