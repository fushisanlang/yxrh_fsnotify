#!/bin/bash
#for simple control yxrh_fsnotify
#by https://github.com/fushisanlang/

WD="/usr/local/yxrh_fsnotify/"

pidfile=${WD}pid.file
start() {
    cd ${WD}
    if [[ -e conf/config.ini ]]; then
        if [[ ${2} = nosend ]]; then
            nohup ${WD}sbin/yxrh_fsnotify_nosend  > /dev/null 2>&1 &
        else
            nohup ${WD}sbin/yxrh_fsnotify  > /dev/null 2>&1 &
        fi
        pid=`ps -ef| grep "/usr/local/yxrh_fsnotify/sbin/yxrh_fsnotify" |grep -v grep |awk '{print $2}'`
        echo ${pid} >${pidfile}
        echo -e "\e[1;32mserver run as ${pid}\e[0m" 
    else
        echo -e "\e[1;31mserver start faild,please check your config file \e[0m"
    fi
}

stop() {
    cd ${WD}
        pid=`ps -ef| grep "/usr/local/yxrh_fsnotify/sbin/yxrh_fsnotify" |grep -v grep |awk '{print $2}'`
    kill -9 ${pid}
if [[ $? = 0 ]]; then
    echo -e "\e[1;32mserver is stoped\e[0m" 
    rm -f ${pidfile}
else
    echo -e "\e[1;31mSomething went wrong \e[0m"
fi
}

status() {
if [[ -e ${pidfile} ]]; then
        pid=`ps -ef| grep "/usr/local/yxrh_fsnotify/sbin/yxrh_fsnotify" |grep -v grep |awk '{print $2}'`
    echo -e "\e[1;32mserver is run as ${pid}\e[0m" 
else
    echo -e "\e[1;32mserver is stoped\e[0m" 
fi
}

#restart
restart () {
stop
start
}

shell_help() {
    echo "Usage: /etc/init.d/yxrh_fsnotify {start|stop|status|restart}"
}

case $1 in
        start )
        start;;
        stop )
        stop;;
        status )
        status ;;
        restart )
        restart ;;
        *) shell_help; break ;;
esac
