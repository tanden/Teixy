#!/bin/bash
if [ $# -eq 0 ] ; then 
    echo 'usage: ./create_sql.sh <sql_file_name>'
    exit 1
fi

timestamp=`date +%Y%m%d%H%M`
up=./sql/${timestamp}_${1}.up.sql
down=./sql/${timestamp}_${1}.down.sql

touch ${up}
touch ${down}

# 実行権限を付与
chmod 755 ${up}
chmod 755 ${down}

# 念の為ファイルがあるかどうかチェック
if [ -f ${up} ] && [ -f ${down} ] ; then
    echo 'create sql file: success!'
    exit 0
else
    echo 'create sql file: fail...'
    exit 1
fi

# 空の.sqlのファイルを作成する
# example 
# $ ./create_sql.sh create_table
# $ ls ./sql
# 201902091202_create_table.up.sql
# 201902091202_create_table.down.sql