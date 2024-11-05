#!/bin/bash
# MC_TOOL=mc.exe
# # 备份源数据目录
# DB_DIR=/root/mysql/data/databackup

# # MinIO URL
# # 桶名称
# BUCKET=backup

# # 目标备份目录,使用当前日期时间
# DEST_DIR=$(date +%Y-%m-%d-%H-%M-%S)-allsql

# # 数据库备份
# $MC_TOOL mirror --overwrite $DB_DIR myminio/$BUCKET/database/$DEST_DIR


mc config host add minio http://ip:9000 13845678 13845678
#看官方文档,设置别名方式可以添加多个s3服务,也可以配置地址
mc alias set minio http://ip:9000 123456 123456
# 列出桶中对象
mc ls minio/static

# 将服务端对象全复制到本地 后面不需要带static目录,复制会自带目录名
mc cp --recursive minio/static E:/minio

# 应该使用这个--recursive递归目录 --newer-than 48h变化的文件时间
mc cp --recursive --newer-than 48h minio/static G:/minio

# // 根据更新时间,增量同步

LAST_TIME=`date -d "24 hours ago" +%Y-%m-%dT%H:%M:%SZ`

mc find --newer-than=$LAST_TIME minio/mybucket | mc cp --recursive /local/data  # 这个不行,会报错



## 同步网站数据到本地
mc cp --recursive --newer-than 240h minio/backup G:\minio


## 把本地文件推到minio服务端

# // 这个是新的 这个是同步本地文件到minio服务的,不要用这个
mc mirror --newer-than 24h minio/static E:/minio/static

#  加上--overwrite参数,强行把本地的覆盖,不然会报错 (这个可以用)
mc mirror --overwrite --newer-than 24h minio/static E:/minio/static

# 寻找100小时更新的文件
mc find --newer-than="100h" minio/static


## 具体内容参考https://min.io/docs/minio/linux/reference/minio-mc/mc-mirror.html,有详细使用教程