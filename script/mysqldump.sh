#!/bin/bash

mysqldump -uroot -123456 db > /var/lib/mysql/databackup/db.sql
RET=$?
echo $RET