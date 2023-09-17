#!/bin/bash

HOST="127.0.0.1"
PORT=3306
USERNAME="root"
PASSWORD="#Duong1091"
DB="api_greet"

if [ -f "./api_greet.sql" ];then
mysql -h${HOST} -P${PORT} -u${USERNAME} ${DB} -p${PASSWORD} -e "source ./api_greet.sql"
mv ./api_greet.sql ./api_greet.sql.tmp
else
echo "数据库文件不存在或已导入"
fi
