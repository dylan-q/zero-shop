ifeq ($e,api)
	CONF=$s-api
	STR = @goctl api go -api ./app/$s/$e/$s.api -dir ./app/$s/$e --style=go_zero --home=./tpl
else
	CONF=$s
	STR = @goctl rpc protoc ./app/$s/$e/$s.proto --go_out=./app/$s/$e/ --go-grpc_out=./app/$s/$e/ --zrpc_out=./app/$s/$e --style=go_zero
endif
run: # 启动order的api服务 make run s=order e=api   启动order的rpc服务  make run s=order e=rpc
	@go run ./app/$s/$e/$s.go  -f ./app/$s/$e/etc/$(CONF).yaml
ctl: #生成根据api和proto文件生成对应的服务 make ctl s=order e=api  make ctl s=order e=rpc
	$(STR)
#make ctl-model m=abc abc 是表名
ctl-model:
	@goctl model mysql datasource -url="root:root@tcp(127.0.0.1:3306)/zero-shop" -table="$m"  -dir="./genModel"  --style=go_zero

