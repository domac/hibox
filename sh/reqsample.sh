#!/usr/bin/env bash

logfile=req.log
servers=( http://localhost:9090/test http://localhost:9090/test2 http://localhost:9090/test )

echo "profile request is running"
echo "-----profile request is running-----" > $logfile
for srv in ${servers[@]}; do
    curl -s -I $srv | awk '/HTTP/{if($2==200) print suc; else print fail }' \
    suc="$srv handle success" fail="$srv handle fail" >> req.log &
done
