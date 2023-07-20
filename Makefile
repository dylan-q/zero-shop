
run-user-api: # 启动 user-api 服务
	@go run ./app/user/api/user.go  -f ./app/user/api/etc/user-api.yaml

run-user-rpc: # 启动 user-rpc 服务
	@go run ./app/user/rpc/user.go  -f ./app/user/rpc/etc/user.yaml

ctl-user-api: # 初始化user-api服务
	@goctl api go -api ./app/user/api/user.api -dir ./app/user/api --style=go_zero

ctl-user-rpc:
	@goctl rpc protoc ./app/user/rpc/user.proto --go_out=./app/user/rpc/ --go-grpc_out=./app/user/rpc/ --zrpc_out=./app/user/rpc --style=go_zero

ctl-user-model:
	@goctl model mysql datasource -url="root:root@tcp(127.0.0.1:3306)/zero-shop" -table="user"  -dir="./genModel"  --style=go_zero
