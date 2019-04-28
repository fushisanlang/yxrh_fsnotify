## 文件监控工具

* 环境说明

该工具源码需要golang语言支持，开发版本为
```
go version go1.11.2 linux/amd64
centos 6.9 x64
```

为了方便使用，已将所有源码进行了编译，可以直接在linux系统下使用。

另：本工具建议结合微信企业号进行使用。或通过zabbix等监控工具进行联动。

* 程序目录结构及说明

```shell
yxrh_fsnotify
├── conf
│   └── config.ini.example //配置文件示例
├── log
│   └── fsnotify.log  //监控生成的日志文件，用于记录被监控目录的文件操作情况
├── sbin
│   ├── yxrh_fsnotify //主程序，用于监控指定目录的文件操作情况
│   ├── yxrh_fsnotify_nosend //主程序，没有发送错误信息的功能
│   └── yxrh_sendmail //用于发送微信消息的程序，可以单独使用
└── src //程序源码
    ├── yxrh_fsnotify.go
    ├── yxrh_fsnotify_nosend.go    
    └── yxrh_sendmail.go
```

* 安装说明

```bash
#以下操作需要使用root用户进行操作
wget  https://raw.githubusercontent.com/fushisanlang/own_note/master/go/yxrh_fsnotify/yxrh_fsnotify.tar.gz -O /usr/local/yxrh_fsnotify.tar.gz
cd /usr/local/
tar zxvf /usr/local/yxrh_fsnotify.tar.gz
cd /usr/local/yxrh_fsnotify/src/install
sh install.sh

#配置
cd /usr/local/yxrh_fsnotify/conf
mv config.ini.example config.ini
```

* 配置文件说明

```bash
[topicArr]        #配置头
ipaddr = 192.168.1.1  #本机ip，用于在发送消息时区别主机
corpid = xxx   #企业微信号的id
corpsecret = xxxxxx  #企业微信号的secret
appid = xxx #企业微信号中应用的appid
dirname1 = /tmp/home #想要监控的目录的绝对路径
dirname2 = /usr/local/sbin #第二个路径，理论支持若干路径，但是建议只监控几个比较外层的目录
dirname3 = /usr/local/nginx/conf #不建议监控日志目录，因为经常写入的目录会不停的发送报警
```

* 使用方法

```shell
/etc/init.d/yxrh_fsnotify start #启动
/etc/init.d/yxrh_fsnotify start nosend #以无微信报警方式启动
/etc/init.d/yxrh_fsnotify stop #关闭
/etc/init.d/yxrh_fsnotify restart #重新启动
/etc/init.d/yxrh_fsnotify status #查看状态
```

* 日志相关

日志写在 `/usr/local/yxrh_fsnotify/log/fsnotify.log` 中。

因为没有定期删除机制，需要手工将历史日志进行删除。

但是因为生产中，需要实时监控，有报警就要及时响应并清理日志，所以日志文件在正常情况下不会很大。

日志存在的意义在于，当报警过多时可以通过日志查看信息。因为微信的消息是通过互联网方式发送的，前后顺序可能会有颠倒，对问题溯源不是很友好。
