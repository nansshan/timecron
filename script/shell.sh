#!/bin/bash  
  
# 设置URL  
# url="http://baidu.com"  
  
# # 发送HTTP GET请求并获取响应  
# response=$(curl -s "$url")  
  
# # 打印响应结果  
# echo "$response"

# echo "哈哈哈哈哈" 
should_run=true
counter=0

while [ $counter -lt 10 ] && [ $should_run = true ]; do
   echo "当前时间:$(date)"
   counter=$((counter+1))
   sleep 1
done

echo "脚本执行结束"