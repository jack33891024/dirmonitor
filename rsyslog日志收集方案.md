标签： rsyslog

# rsyslog日志收集方案

---

[toc]

## rsyslog
	
> * 日志:历史日志
 - 历史事件:
   时间,事件
   日志级别: 事件的关键性程度,loglevel
> * 系统日志服务:
 - syslog:CentOS 5
   syslogd:system log,记录用户空间日志;
   klogd:kernel log,记录内核空间日志;
 - rsyslog:CentOS 6及以上
   syslogd:
   klogd:
> * rsyslog: 多线程, 支持传输协议: UDP,TCP,SSL,TLS,RELP;
 - MySQL,PGSQL,Oracle实现日志存储;
 - 强大的过滤器,可实现过滤日志信息中任何部分;
 - 自定义输出格式
 - elasticsearch,logstash,kibana = elk
> * 日志收集方: 
 - facility: 设施,从功能或程序上对日志进行分类;如: auth,authpriv,cron,daemon,kern,lpr,mail,mark,news,security,user,uucp,local0-local7,syslog	
 - priority: 级别,debug,info,notice,warn(warning),err(error),crit(critical),alert,emerg(panic)
 - 指定级别:
   *:所有级别
   none:没有级别
   priority:此级别及更高级别的日志信息
   =priority:此级别
 - 系统细致定义: facility.priority 	/var/log/message
> * 程序环境:
 - 主程序: rsyslogd
 - 配置文件: /etc/rsyslog.conf
 - 服务脚本: /etc/rc.d/init.d/rsyslog
> * rsyslog配置文件: rsyslog.conf
```
RULES:
	facility.priority	target
			
	target:
        文件路径:记录于指定的日志文件中,通常应该在/var/log目录下;文件路径前的"-"表示异步写入;
        用户:将日志通知给指定用户
            *:所有用户
	    日志服务器:@host
            host:必须要监听在tcp或udp协议514端口上提供服务;
	    管道:|COMMAND
		
    文件记录的日志的格式:
        时间产生的日期时间	主机	进程(pid):	事件内容

        有些日志记录二进制格式: /var/logwtmp,/var/log/btmp
            /var/log/wtmp:当前系统上成功登录的日志;
            # last
        /var/log/btmp:当前系统上失败的登录尝试;
            # lastb
			# lastlog命令: 显示当前系统上每一个用户最近一次的登录时间;
```

## rsyslog配置
	
> * 服务端配置: 日志收集端
```
# 开启指定模块和监听套接字(TCP或UDP任选):
    /etc/rsyslog.conf:
    # Provides UDP syslog reception
    $ModLoad imudp
    $UDPServerRun 514

    # Provides TCP syslog reception
    $ModLoad imtcp
    $InputTCPServerRun 514
    重载rsyslog
```
> * 客户端配置: 日志发送端
```
/etc/rsyslog.conf:
    facility.priority	@RSYSLOG_SERVER_IP
    重载rsyslog
```

## 配置使用基于mysql存储日志信息
		
> 准备MySQL服务器,创建用户,授权对Syslog数据库的全部访问权限
```
$ yum install mysql-server mysql -y 
$ service mysqld start
mysql> grant all privileges on Syslog.* to syslog@'172.16.%.%' identified by 'syslogpass';
mysql> flush privileges;
$ mysql -usyslog -psyslogpass -h172.16.100.10

# 注意:172.16.100.10为mysql服务端所在IP地址,/etc/my.cnf配置中加入skip_name_resolv = on;
```
			
> 安装rsyslog-mysql程序包
```
$ yum list all rsyslog-mysql
$ yum install rsyslog-mysql -y
```
	
> 导入rsyslog-mysql依赖的数据库
```
$ rpm -ql rsyslog-mysql
$ mysql -usyslog -psyslogpass -h172.16.100.10 </usr/share/doc/rsyslog-mysql-VERSION/createDB.sql
```
	
> 配置rsyslog使用ommysql模块
```
/etc/rsyslog.conf:
#### MODULES ####
$ModLoad ommysql

#### RULES ####
facility.priority 	:ommysql:DBHOST,DB,DBUSER,USERPASS
```
	
> 重启rsyslog服务
```
$ service rsyslog restart
```
		
> 安装loganalyzer
```
# 1.配置webserver,支持php:
$ yum instal httpd php php-mysql php-gd -y
$ service httpd start

# 2.配置loganalyzer程序:
$ cp -r loganalyzer-3.6.5/src /var/www/html/loganalyzer
$ cp loganalyzer-3.6.5/contrib/*.sh /var/www/html/loganalyzer
$ cd /var/www/html/loganalyzer
$ chmod +x *.sh
$ ./configure.sh
$ ./secure.sh
$ chmod 666 config.php
```

> 浏览器安装loganalyzer程序: http://httpd_ip/loganalyzer/install.php