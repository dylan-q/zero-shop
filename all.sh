#!/usr/bin bash
servers=("user" "order" "goods" "pay")

# 使用for循环遍历列表中的元素
for svr in "${servers[@]}"
do
  start  go run ./app/"$svr"/rpc/"$svr".go  -f ./app/"$svr"/rpc/etc/"$svr".yaml
  start  go run ./app/"$svr"/api/"$svr".go  -f ./app/"$svr"/api/etc/"$svr"-api.yaml
done