#!/bin/sh

# 杀掉modd, 及其运行的所有子进程
batchkill() {
    # 根据程序的ppid获取程序的pid
	pids=`ps -ef|awk '{if($3=='$1'){print $2} }'`;
	#杀掉父程序的pid，防止子程序被杀掉后开启新的子程序
    echo "killing $1"
    kill -9 $1
    #如果获得了pid，则以已获得的pid作为ppid继续进行查找
	if [ -n "$pids" ]; then
		for pid in $pids
		do
		batchkill $pid
		done
	fi
}

pid_modd=`pgrep -f modd | head -n 1`
echo $pid_modd
batchkill $pid_modd
echo "done"