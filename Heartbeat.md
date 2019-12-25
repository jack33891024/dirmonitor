标签: 高可用软件

# Heartbeat高可用实践

---

[toc]

## Heartbeat介绍
	
> heartbeat可以将资源(IP及程序服务资源)从一台已经故障的计算机快速转移到另一台正常运转的机器上继续提供服务;

> heartbeat和keepalived相似,heartbeat可以实现failover(故障转移)功能,但不能实现对后端的健康检查;

> heartbeat官方地址: http://linux-ha.org/wiki/Main_Page

## Heartbeat工作原理

> 通过修改heartbeat软件的配置文件,可以指定哪一台Hearbeat服务器作为主服务器,则另一台将自动成为热备服务器;然后在热备服务器上配置Heartbeat守护程序来监听来自主服务器的心跳消息;如果热备服务器在指定时间内未监听到来自主服务器的心跳,就会启动故障转移程序,并取得主服务器上的相关资源服务的所有权,接替主服务器继续不间断的提供服务,从而达到资源及服务高可用性的目的;此种方式属于主备模式;

> heartbeat还支持主主模式,即两台服务器互为主备,这时它们之间会相互发送报文来告诉对方自己当前的状态,如果在指定时间内未收到对方发送的心跳报文,那么一方就会认为对方失效或者宕机了;这时每个运行正常的主机就会启动自身的资源接管模块来接管运行在对方主机上的资源或者服务,继续为用户提供服务;一般情况下可以较好的实现一台主机故障后,企业业务仍能不间断的持续运行;注意: 所谓的业务不间断,在故障转移期间也是需要切换时间的,heatbeat的切换时间一般是在5-20秒左右;

> 和Keepalived服务一样,heartbeat高可用是服务器级别的,不是服务级别的;

> heartbeat切换的常见条件: 服务器宕机、Heartbeat服务本身故障、心跳连接故障;

> 服务故障不会导致切换,可以通过服务器宕机把heartbeat服务停掉;	

## Hearbeat心跳连接

> 要部署heartbeat服务,至少需要两台主机来完成;

> 要实现高可用服务,两台主机之间是通过以下几种方式实现互相通信、监测;
> * 串行电缆,所谓的串口(首选,缺点距离不能太远);
> * 一根以太网电缆两网卡直连(推荐);
> * 以太网电缆,通过交换机等网络设备连接(次选);但是增加了交换机的故障点,同时,线路不是专用心跳线,容易受其他数据传输的影响,导致心跳报文发送问题;

> 高可用服务器上对Heartbeat软件会利用这条心跳线来检查对端的机器是否存活,进而决定是否做故障转移,资源切换来保证业务的连续性;

> 如果条件允许,以上的连接可同时使用,来加大保险系数防止裂脑问题发生;

## 裂脑介绍

> 由于两台高可用服务器对之间在指定时间内,无法互相检测对方心跳而各自启动故障转移功能,取的了资源及服务的所有权,而此时的两台高可用服务器对都还活着并正常运行,这样就会导致同一个IP或服务在两端同时启动而发生冲突的严重问题,最严重的是两台主机占用同一个VIP地址,当用户写入数据时可能会分别写入到两端,这样可能会导致服务器两端的数据不一致或造成数据丢失,这种情况就称之为裂脑,也有的人称其为分区集群或大脑垂直分割,英文为split brain;

## 裂脑发生原因

> 高可用服务器对之间心跳线路故障,导致无法正常通信(硬件故障);
> * 1、心跳线坏了(断了、老化等);
> * 2、网卡及相关驱动问题,IP配置及冲突问题;
> * 3、心跳线间连接的设备故障(网卡及交换机);
> * 4、仲裁的机器出问题;

> 高可用服务器对上开启了防火墙阻挡了心跳消息传输(软件故障);

> 其他服务配置不当等原因,如心跳方式不同,心跳广播冲突、软件BUG等;

> 另外的高可用软件Keepalived配置里如果virtual_router_id参数,两端配置不一致,也会导致脑裂问题发生;

## 防止脑裂的措施

> 发生裂脑时,对业务的影响是极其严重的,有时甚至是致命的;如两台高可用服务器对之间发生脑裂,导致互相争用同一IP资源,就如同我们在局域网内常见的IP地址冲突一样,两个机器就会有一个或者两个都不正常,影响用户正常访问服务器;如果是应用在数据库或者存储服务这种极重要的高可用上,那就可能会导致用户发布的数据间断的写在两台不同服务器上的恶果,最终数据恢复极其困难或难以恢复;

> 实际生产环境中,我们可以从以下几个方面来防止裂脑问题的发生;
> * 同时使用串行电缆和以太网电缆连接,同时用两条心跳线,这样一条线坏了,另一个还是好的,依然能传递心跳消息(网卡设备和网线设备);
> * 检测到脑裂时强行关闭一个心跳节点(互联网场景应用不多,银行等业务应用较多);(这个功能需要特殊设备支持,如Stonith、fence)相当于程序上备节点发现心跳线故障,发送关机命令到主节点;
> * 做好对裂脑时的监控报警(如邮件或手机短信等),在问题发生时人为第一时间介入仲裁,降低损失,百度监控有上行和下行,和人工交互的过程;当然在实施高可用方案时,要根据业务实际需求确定是否能容忍这样的损失;对于一般的网站业务,这个损失是可控的;
> * 启用磁盘锁;正在服务一方锁住共享磁盘,"裂脑"发生时,让对方完全"抢不走"共享磁盘资源;但使用磁盘锁也会有一个不小的问题,如果占用共享盘的一方不主动"解锁",另一方就永远得不到共享磁盘;现实中假如服务节点突然死机或崩溃,就不可能执行解锁命令;后备节点也就接管不了共享资源和应用服务;于是有人在HA中设计了"智能"锁;即,正在服务的一方只在心跳线全部断开(察觉不到对端)时才会启用磁盘锁;平时就不上锁了;
> * 报警报在服务器接管之前,给人员处理留足够时间(1分钟内报警,但是服务器此时没有接管,而是5分钟接管,接管的时间较长;数据不会丢,导致用户无法写数据);
> * 报警后,不直接自动服务器接管,而是人为人员控制接管;
> * 增加仲裁机制,确定谁该获得资源;以下有几点思路:
```
1、增加一个仲裁机制;例如设置参考IP(如网关IP),当心跳线完全断开时,2个节点各自ping一个参考IP,不通则表明断点就出在本段,不仅心跳线、还有对外服务的本地网络链路断了,这样就主动放弃竞争,让能ping通参考IP的一端去接管服务;ping不通参考IP的一方可以自我重启,可以彻底释放有可能还占用的那些共享资源(heartbeat也有此功能);

2、通过第三方软件仲裁谁该获得资源;
```
> * 如何开发程序判断裂脑:
```
#1、只要备节点出现VIP就报警(a、主机宕机,备机接管；b、主机没宕机,裂脑了),不管哪个情况,人工查看;
#2、严谨判断,备机出现VIP,并且主机及服务还活着,裂脑了(依赖报警);
```


## Heartbeat消息类型

> Heartbeat高可用软件在工作过程中,一般来说,有三种消息类型,具体为: 
> * 心跳消息: 心跳消息为约150字节的数据包,可能为单播、广播或多播的方式,控制心跳频率及出现故障要等待多久进行故障切换;
> * 集群转换消息(ip-request和ip-request-resp): 当主服务器恢复在线状态时,通过ip-request消息要求备机释放主服务器失败时备服务器取得的资源,然后备份服务器关闭释放主服务器失败时取得的资源及服务；备服务器释放主服务器失败时取得的资源及服务后,就会通过ip-request-resp消息通知主服务器它不在拥有该资源及服务,主服务器收到来自备服务器的ip-request-resp消息通知后,启动失败时释放的资源及服务,并开始提供正常的访问服务;
> * 重传消息: rexmit-request控制重传心跳请求,此消息不太重要；以上心跳控制消息都使用UDP协议发送到/etc/ha.d/ha/cf文件指定的任意端口,或指定的多播地址;

## Heartbeat IP地址接管和故障转移

> Heartbeat是通过IP地址接管和ARP广播进行故障转移的;

> ARP广播: 在主服务器故障时,备节点接管资源后,会立即强制更新所有客户端本地的ARP表(即清除客户端本地缓存的失败服务器的vip地址和mac地址的解析记录);确保客户端和新的服务器对话;

## 管理IP与虚拟IP

> 真实IP,又被称为管理IP,一般是配置在物理网卡上的实际IP,在负载均衡及高可用环境中,管理IP是不对外提供用户访问服务的,仅作管理服务器使用,如SSH可以通过这个管理IP来连接服务器;

> VIP是虚拟IP,ifconfig命令配置的称为别名IP(CentOS 7没有),ip命令配置的称为辅助IP;实际上就是heartbeat临时绑定在物理网卡上的别名IP(heartbeat3以上也采用了辅助IP),如eth0:x,x为0-255的任意数字,你可以在一块网卡上绑定多个别名;在实际生产环境中,需要在DNS配置中把网站域名地址解析到这个VIP地址上,由这个VIP对用户提供服务;

> 这样做的好处是当提供服务的服务器宕机后,在接管的服务器上直接会自动配置上同样的VIP提供服务,如果是使用管理IP的话,来回迁移就难以做到,而且,管理IP迁移走了;我们就只能去机房连接服务器了;

> VIP的实质就是确保两台服务器各有一个管理IP不动,就是随时可以连上机器,然后,增加绑定其他的IP,这样就算VIP转移走了,也不至于服务器本身连不上;


## Hearbeat的生产应用场景

> web服务器对之间使用高可用(尽量使用keepalived);

> 数据库主库和从库之间使用heartbeat高可用;

> 存储服务器之间使用heartbeat实现高可用;

> 分布式存储系统配合heartbeat的高可用;

## Heartbeat脚本默认目录和配置文件路径

> 启动脚本: `/etc/init.d/heartbeat`

> 资源控制目录: `/etc/ha.d/resource.d/`(控制资源的脚本,可以被heartbeat直接调用)

> heartbeat的默认配置文件目录为/etc/ha.d;heartbeat常用的配置文件有三个,分别为ha.cf、authkey、haresource;

|配置名称|作用|说明|
|:---|:---|:---|
|ha.cf|heartbeat参数配置文件|在这里配置heartbeat的一些基本参数|
|authkey|heartbeat认证文件|高可用服务器对之间根据对端的authkey,对对端进行认证;|
|haresource|heartbeat资源配置文件|如配置IP资源及脚本程序等|

## heartbeat服务主机规划
  
> 下表可以适用于对外提供服务的业务(例如web服务):

|主机名|接口|IP地址|作用|
|:---|:---|:---|:---|
|mysql-master|eth0|10.0.0.64|外网管理IP,用于WAN数据转发|
|mysql-master|eth1|172.16.1.64|内网管理IP,用于LAN数据转发|
|mysql-master|eth2(内网)|10.0.10.64|用于服务器间心跳连接(直连)|
|mysql-master|VIP|172.16.1.164|用于程序提供对外服务|
|mysql-backup|eth0|10.0.0.65|外网管理IP,用于WAN数据转发|
|mysql-backup|eth1|172.16.1.65|内网管理IP,用于LAN数据转发|
|mysql-backup|eth2(内网)|10.0.10.65|用于服务器间的心跳连接|
|mysql-backup|VIP|172.16.1.165|用于程序对外提供服务|

> 下表适合于不对外提供服务的业务(例如: MySQL)

|主机名|接口|IP地址|作用|
|:---|:---|:---|:---|
|mysql-master|eth0|禁用|数据库不用公网地址|
|mysql-master|eth1|172.16.1.64|内网管理IP,用于LAN数据转发|
|mysql-master|eth2(内网)|10.0.10.64|用于服务器间心跳连接(heartbeat直连)|
|mysql-master|VIP|172.16.1.164|用于程序提供对外服务|
|mysql-backup|eth0|禁用|数据库不用公网地址|
|mysql-backup|eth1|172.16.1.65|内网管理IP,用于LAN数据转发|
|mysql-backup|eth2(内网)|10.0.10.65|用于服务器间心跳连接(heartbeat直连)|
|mysql-backup|VIP|172.16.1.165|用于程序对外提供服务|

> 提示: 需要另外添加一块eth2的内网卡用于做心跳线

> 服务器的网卡、hosts文件、心跳线如下: 
```
[root@mysql-master ~]# ip add
1: lo: <LOOPBACK,UP,LOWER_UP> mtu 65536 qdisc noqueue state UNKNOWN 
    link/loopback 00:00:00:00:00:00 brd 00:00:00:00:00:00
    inet 127.0.0.1/8 scope host lo
    inet6 ::1/128 scope host 
       valid_lft forever preferred_lft forever
2: eth0: <BROADCAST,MULTICAST,UP,LOWER_UP> mtu 1500 qdisc pfifo_fast state UP qlen 1000
    link/ether 00:0c:29:fb:c4:3a brd ff:ff:ff:ff:ff:ff
    inet 10.0.0.64/24 brd 10.0.0.255 scope global eth0 #WAN数据转发IP,本次环境不会使用到
    inet6 fe80::20c:29ff:fefb:c43a/64 scope link 
       valid_lft forever preferred_lft forever
3: eth1: <BROADCAST,MULTICAST,UP,LOWER_UP> mtu 1500 qdisc pfifo_fast state UP qlen 1000
    link/ether 00:0c:29:fb:c4:44 brd ff:ff:ff:ff:ff:ff
    inet 172.16.1.64/24 brd 1716.1.255 scope global eth1 #LAN数据转发IP
    inet6 fe80::20c:29ff:fefb:c444/64 scope link 
       valid_lft forever preferred_lft forever
4: eth2: <BROADCAST,MULTICAST,UP,LOWER_UP> mtu 1500 qdisc pfifo_fast state UP qlen 1000
    link/ether 00:0c:29:fb:c4:4e brd ff:ff:ff:ff:ff:ff
    inet 10.0.10.64/24 brd 10.0.10.255 scope global eth2 #心跳线IP
    inet6 fe80::20c:29ff:fefb:c44e/64 scope link 
       valid_lft forever preferred_lft forever

[root@mysql-backup ~]# ip add
1: lo: <LOOPBACK,UP,LOWER_UP> mtu 65536 qdisc noqueue state UNKNOWN 
    link/loopback 00:00:00:00:00:00 brd 00:00:00:00:00:00
    inet 127.0.0.1/8 scope host lo
    inet6 ::1/128 scope host 
       valid_lft forever preferred_lft forever
2: eth0: <BROADCAST,MULTICAST,UP,LOWER_UP> mtu 1500 qdisc pfifo_fast state UP qlen 1000
    link/ether 00:0c:29:2d:33:92 brd ff:ff:ff:ff:ff:ff
    inet 10.0.0.65/24 brd 10.0.0.255 scope global eth0
    inet6 fe80::20c:29ff:fe2d:3392/64 scope link 
       valid_lft forever preferred_lft forever
3: eth1: <BROADCAST,MULTICAST,UP,LOWER_UP> mtu 1500 qdisc pfifo_fast state UP qlen 1000
    link/ether 00:0c:29:2d:33:9c brd ff:ff:ff:ff:ff:ff
    inet 172.16.1.65/24 brd 1716.1.255 scope global eth1  #LAN数据转发IP
    inet6 fe80::20c:29ff:fe2d:339c/64 scope link 
       valid_lft forever preferred_lft forever
4: eth2: <BROADCAST,MULTICAST,UP,LOWER_UP> mtu 1500 qdisc pfifo_fast state UP qlen 1000
    link/ether 00:0c:29:2d:33:a6 brd ff:ff:ff:ff:ff:ff
    inet 10.0.10.65/24 brd 10.0.10.255 scope global eth2 #心跳线IP
    inet6 fe80::20c:29ff:fe2d:33a6/64 scope link 
       valid_lft forever preferred_lft forever

[root@mysql-master ~]# tail -3 /etc/hosts
172.16.1.64 mysql-master
172.16.1.65 msyql-bakcup
172.16.1.66 msyql-slave

[root@mysql-backup ~]# tail -3 /etc/hosts
1716.1.64 mysql-master
1716.1.65 msyql-bakcup
1716.1.66 msyql-slave

[root@mysql-master ~]# route add -host 10.0.10.65 dev eth2 #配置心跳线直连通信,可以不配置
[root@mysql-backup ~]# route add -host 10.0.10.64 dev eth2 #配置心跳线直连通信,可以不配置
```

## 双机安装heartbeat

> 修改yum源并安装
```
$ mv /etc/yum.repos.d/CentOS-Base.repo /etc/yum.repos.d/CentOS-Base.repo.backup
$ wget -O /etc/yum.repos.d/CentOS-Base.repo http://mirrors.aliyun.com/repo/Centos-6.repo
$ wget -O /etc/yum.repos.d/epel.repo http://mirrors.aliyun.com/repo/epel-6.repo
$ yum install heartbeat* -y
```

## 配置heartbeat模板配置文件

> 两台机器操作一样
```
$ ls /etc/ha.d/
$ ls -l /usr/share/doc/heartbeat-3.0.4/ # heartbeat的模板配置目录
$ cd /usr/share/doc/heartbeat-3.0.4/

# 拷贝3个文件到/etc/ha.d/下
$ cp -a authkeys ha.cf haresources /etc/ha.d/
$ ls -l /etc/ha.d/
authkeys ha.cf harc haresources rc.d README.config resource.d shellfuncs
```

## ha.cf、authkeys及haresource文件配置及说明

> * README文件说明
```
[root@mysql-master ~]# cat /etc/ha.d/README.config 
You need three configuration files to make heartbeat happy,
and they all go in this directory.

They are:
        ha.cf           Main configuration file
        haresources     Resource configuration file
        authkeys        Authentication information

These first two may be readable by everyone, but the authkeys file must not be.

The good news is that sample versions of these files may be found in
the documentation directory (providing you installed the documentation).

If you installed heartbeat using rpm packages then
this command will show you where they are on your system:
                rpm -q heartbeat -d

If you installed heartbeat using Debian packages then
the documentation should be located in /usr/share/doc/heartbeat	
```
> * 查看rpm方式安装软件的帮助文档信息
```
[root@mysql-master ~]# rpm -qd heartbeat 
/usr/share/doc/heartbeat-3.0.4/AUTHORS
/usr/share/doc/heartbeat-3.0.4/COPYING
/usr/share/doc/heartbeat-3.0.4/COPYING.LGPL
/usr/share/doc/heartbeat-3.0.4/ChangeLog
/usr/share/doc/heartbeat-3.0.4/README
/usr/share/doc/heartbeat-3.0.4/apphbd.cf
/usr/share/doc/heartbeat-3.0.4/authkeys
/usr/share/doc/heartbeat-3.0.4/ha.cf
/usr/share/doc/heartbeat-3.0.4/haresources
/usr/share/man/man1/cl_status.1.gz
/usr/share/man/man1/hb_addnode.1.gz
/usr/share/man/man1/hb_delnode.1.gz
/usr/share/man/man1/hb_standby.1.gz
/usr/share/man/man1/hb_takeover.1.gz
/usr/share/man/man5/authkeys.5.gz
/usr/share/man/man5/ha.cf.5.gz
/usr/share/man/man8/apphbd.8.gz
/usr/share/man/man8/heartbeat.8.gz
```

> ha.cf文件说明

|参数|说明|
|:---|:---|
|debugfile /var/log/ha-debug|heartbeat的调试日志存放位置|
|logfile /var/log/ha-log|heartbeat的日志存放位置|
|logfacility local1|在syslog服务中配置通过local1设备接收日志|
|keepalive  2|指定心跳间隔时间为2秒(即每2秒中在eth1上发送一次广播)|
|deadtime  30|指定若备节点在30秒内没有收到主节点的心跳信号,则立即接管主节点的服务资源;|
|warntime  10|指定心跳延迟的时间为10秒;当10秒中内备节点不能接收到主节点的心跳信号时,就会往日志中写入一个警告日志,但是不会切换服务;|
|initdead  120|指定在heartbeat首次运行后,需要等待120秒才启动主服务器的任何资源;该选项用于解决这种情况产生的时间间隔取值至少为deadtime的两倍,单机启动时会遇到vip绑定很慢,为正常现象,该值设置的长的原因;|
|#bcast    eth1|指明心跳使用以太网广播方式在eth1接口上进行广播;如使用两个实际网络来传送心跳#bcast  eth0  eth1|
|mcast eth2 225.0.0.0 694 1 0|设置多播通信使用的端口,694为默认使用的端口|
|auto_failback on|#用来定义主节点恢复后,是否将服务自动切回;|
|node     mysql-master|主节点主机名|
|node     mysql-backup|备节点主机名|
|crm      no|是否开启Cluster Resource Manager(集群资源管理)功能;|

> ha.cf配置文件配置内容
```
cat /etc/ha.d/ha.cf 
debugfile /var/log/ha-debug
logfile /var/log/ha-log
logfacility     local1

keepalive 2
deadtime 30
warntime 10
initdead 60

#bcast  eth1
mcast eth2 225.0.0.64 694 1 0

auto_failback on
node    mysql-master
node    mysql-backup
crm     no
```

> authkeys文件说明:
```
[root@mysql-master ~]# cat /etc/ha.d/authkeys 
#
#       Authentication file.  Must be mode 600 #权限必须是600
#
#
#       Must have exactly one auth directive at the front.
#       auth    send authentication using this method-id
#
#       Then, list the method and key that go with that method-id
#
#       Available methods: crc sha1, md5.  Crc doesn't need/want a key.
#
#       You normally only have one authentication method-id listed in this file
#
#       Put more than one to make a smooth transition when changing auth
#       methods and/or keys.
#
#
#       sha1 is believed to be the "best", md5 next best.
#
#       crc adds no security, except from packet corruption.
#               Use only on physically secure networks.
#
#auth 1
#1 crc
#2 sha1 HI!
#3 md5 Hello!
#######################################

# 生成authkeys的密码
$ echo heartbeat|sha1sum 
1c81b56e0878d42e95064acfe07613edd5c6c29d  -

# 把authkeys文件的权限改成600
$ chmod 600 /etc/ha.d/authkeys 

# 把authkeys的内容修改成如下
$ cat /etc/ha.d/authkeys   
auth 1
1 sha1 1c81b56e0878d42e95064acfe07613edd5c6c29d
```

> haresource文件说明
```
[root@mysql-master ~]# cat /etc/ha.d/haresources    
mysql-master IPaddr::172.16.1.164/24/eth1
mysql-backup IPaddr::172.16.1.165/24/eth1

# mysql-master为主机名,表示初始状态会在mysql-maste绑定vip 172.16.1.164;
# Ipaddr为heartbeat配置IP的默认脚本,双冒号后面的IP等都是脚本的参数;
# 172.16.1.164/24/eth1为集群对外服务的VIP,初始启动在mysql-master上,#24为子网掩码,eth1为ip绑定的实际物理网卡,为heartbeat提供对外服务的通信接口;
# 172.16.1.165/24/eth1为集群对外服务的VIP,初始启动在mysql-backup上,#24为子网掩码,eth1为ip绑定的实际物理网卡,为heartbeat提供对外服务的通信接口;

# 用heartbeat自带资源启VIP的实际过程例子如下:
[root@mysql-master ~]# /etc/ha.d/resource.d/IPaddr 172.16.1.164/24/eth1 start
INFO: Adding inet address 172.16.1.164/24 to device eth1
INFO: Bringing device eth1 up
INFO: /usr/libexec/heartbeat/send_arp -i 200 -r 5 -p /var/run/resource-agents/send_arp-172.16.1.164 eth1 172.16.1.164 auto not_used not_used
INFO:  Success
INFO:  Success

[root@mysql-master ~]# ip add|grep 172.16.1.164
    inet 172.16.1.164/24 scope global eth1
```

## 启动heartbeat查看配置结果

> 启动hearbeat查看vip信息
```
[root@mysql-master ~]# /etc/init.d/heartbeat start
[root@mysql-master ~]# ip add|grep 10.0.0.16[45]
    inet 172.16.1.164/24 scope global eth1
[root@mysql-backup ~]# /etc/init.d/heartbeat start
[root@mysql-backup ~]# ip add|grep 10.0.0.16[45]
    inet 172.16.1.165/24 scope global eth1
```
> 模拟主宕机查看VIP情况
```
[root@mysql-master ~]# /etc/init.d/heartbeat stop
Stopping High-Availability services: Done.
[root@mysql-backup ~]# ip add|grep 10.0.0.16[45]
    inet 172.16.1.165/24 scope global eth1
    inet 172.16.1.164/24 scope global secondary eth1

[root@mysql-master ~]# lsof -i :694
COMMAND    PID USER   FD   TYPE DEVICE SIZE/OFF NODE NAME
heartbeat 8085 root    7u  IPv4  18978      0t0  UDP 225.0.0.64:ha-cluster 
heartbeat 8086 root    7u  IPv4  18978      0t0  UDP 225.0.0.64:ha-cluster

# 发现故障问题,第一时间看日志报错
# 主动释放资源
[root@mysql-master ~]# /usr/share/heartbeat/hb_standby 
/usr/share/heartbeat/hb_standby --help
usage:
/usr/share/heartbeat/hb_standby [all|foreign|local|failback]    # 默认接管所有,all接管所有,local接管本机控制的资源

# 主动接管资源
[root@mysql-master ~]# /usr/share/heartbeat/hb_takeover 
########################################
# 主通过释放资源查看日志信息理解过程
[root@mysql-master ~]# cat /var/log/ha-log  
Jul 05 08:23:16 mysql-master heartbeat: [12285]: info: mysql-master wants to go standby [all]
Jul 05 08:23:17 mysql-master heartbeat: [12285]: info: standby: mysql-backup can take our all resources
Jul 05 08:23:17 mysql-master heartbeat: [14599]: info: give up all HA resources (standby).
ResourceManager(default)[14612]:        2016/07/05_08:23:17 info: Releasing resource group: mysql-master IPaddr::172.16.1.164/24/eth1
ResourceManager(default)[14612]:        2016/07/05_08:23:17 info: Running /etc/ha.d/resource.d/IPaddr 172.16.1.164/24/eth1 stop
IPaddr(IPaddr_172.16.1.164)[14675]:       2016/07/05_08:23:17 INFO: IP status = ok, IP_CIP=
/usr/lib/ocf/resource.d//heartbeat/IPaddr(IPaddr_172.16.1.164)[14649]:    2016/07/05_08:23:17 INFO:  Success
ResourceManager(default)[14739]:        2016/07/05_08:23:17 info: Releasing resource group: mysql-backup IPaddr::172.16.1.165/24/eth1
ResourceManager(default)[14739]:        2016/07/05_08:23:17 info: Running /etc/ha.d/resource.d/IPaddr 172.16.1.165/24/eth1 stop
IPaddr(IPaddr_172.16.1.165)[14802]:       2016/07/05_08:23:17 INFO: IP status = no, IP_CIP=
/usr/lib/ocf/resource.d//heartbeat/IPaddr(IPaddr_172.16.1.165)[14776]:    2016/07/05_08:23:17 INFO:  Success
Jul 05 08:23:17 mysql-master heartbeat: [14599]: info: all HA resource release completed (standby).
Jul 05 08:23:17 mysql-master heartbeat: [12285]: info: Local standby process completed [all].
Jul 05 08:23:18 mysql-master heartbeat: [12285]: WARN: 1 lost packet(s) for [mysql-backup] [885:887]
Jul 05 08:23:18 mysql-master heartbeat: [12285]: info: remote resource transition completed.
Jul 05 08:23:18 mysql-master heartbeat: [12285]: info: No pkts missing from mysql-backup!
Jul 05 08:23:18 mysql-master heartbeat: [12285]: info: Other node completed standby takeover of all resources.

# 通过在备上查看主释放资源后的日志来理解heartbeat的工作原理
[root@mysql-backup ~]# cat /var/log/ha-log  
Jul 05 08:23:19 mysql-backup heartbeat: [13172]: info: mysql-master wants to go standby [all]
Jul 05 08:23:20 mysql-backup heartbeat: [13172]: info: standby: acquire [all] resources from mysql-master
Jul 05 08:23:20 mysql-backup heartbeat: [14738]: info: acquire all HA resources (standby).
ResourceManager(default)[14751]:        2016/07/05_08:23:20 info: Acquiring resource group: mysql-master IPaddr::172.16.1.164/24/eth1
/usr/lib/ocf/resource.d//heartbeat/IPaddr(IPaddr_172.16.1.164)[14779]:    2016/07/05_08:23:20 INFO:  Resource is stopped
ResourceManager(default)[14751]:        2016/07/05_08:23:20 info: Running /etc/ha.d/resource.d/IPaddr 172.16.1.164/24/eth1 start
IPaddr(IPaddr_172.16.1.164)[14904]:       2016/07/05_08:23:20 INFO: Adding inet address 172.16.1.164/24 to device eth1
IPaddr(IPaddr_172.16.1.164)[14904]:       2016/07/05_08:23:20 INFO: Bringing device eth1 up
IPaddr(IPaddr_172.16.1.164)[14904]:       2016/07/05_08:23:20 INFO: /usr/libexec/heartbeat/send_arp -i 200 -r 5 -p /var/run/resource-agents/send_arp-172.16.1.164 eth1 172.16.1.164 auto not_used not_used
/usr/lib/ocf/resource.d//heartbeat/IPaddr(IPaddr_172.16.1.164)[14878]:    2016/07/05_08:23:20 INFO:  Success
ResourceManager(default)[14985]:        2016/07/05_08:23:20 info: Acquiring resource group: mysql-backup IPaddr::172.16.1.165/24/eth1
/usr/lib/ocf/resource.d//heartbeat/IPaddr(IPaddr_172.16.1.165)[15013]:    2016/07/05_08:23:20 INFO:  Running OK
Jul 05 08:23:20 mysql-backup heartbeat: [14738]: info: all HA resource acquisition completed (standby).
Jul 05 08:23:20 mysql-backup heartbeat: [13172]: info: Standby resource acquisition done [all].
Jul 05 08:23:21 mysql-backup heartbeat: [13172]: info: remote resource transition completed.
```

## HeartBeat实现WEB服务高可用

> 在mysql-master和mysql-backup上安装httpd服务
```
$ yum install httpd -y
$ /etc/init.d/httpd start
$ lsof -i :80
```
> 在mysql-master和mysql-backup上创建index文件
```
[root@mysql-master ~]# echo 10.0.0.64 >/var/www/html/index.html
[root@mysql-backup ~]# echo 10.0.0.65 >/var/www/html/index.html
```
> 查看服务状态
```
[root@mysql-master ~]# curl 10.0.0.64
10.0.0.64
[root@mysql-backup ~]# curl 10.0.0.65
10.0.0.65
```
> 配置web服务高可用,通过vip漂移实现高可用
> * 在浏览器上用vip访问,宕掉主节点(挂起虚拟机),服务切换过程需要30秒左右可以正常提供访问;
> * 两边不起httpd,将启动资源的权利交给heartbeat控制
```
[root@mysql-master ~]# cat /etc/ha.d/haresources 
mysql-master IPaddr::172.16.1.164/24/eth0 httpd
#mysql-master IPaddr::172.16.1.164/24/eth1
mysql-backup IPaddr::172.16.1.165/24/eth1 

[root@mysql-backup ~]# cat /etc/ha.d/haresources 
mysql-master IPaddr::172.16.1.164/24/eth0 httpd
#mysql-master IPaddr::172.16.1.164/24/eth1 
mysql-backup IPaddr::172.16.1.165/24/eth1

[root@mysql-master ~]# /etc/init.d/heartbeat start
[root@mysql-master ~]# lsof -i :80
COMMAND   PID   USER   FD   TYPE DEVICE SIZE/OFF NODE NAME
httpd   20168   root    8u  IPv6  31604      0t0  TCP *:http (LISTEN)
httpd   20170 apache    8u  IPv6  31604      0t0  TCP *:http (LISTEN)
httpd   20171 apache    8u  IPv6  31604      0t0  TCP *:http (LISTEN)
httpd   20172 apache    8u  IPv6  31604      0t0  TCP *:http (LISTEN)
httpd   20173 apache    8u  IPv6  31604      0t0  TCP *:http (LISTEN)
httpd   20174 apache    8u  IPv6  31604      0t0  TCP *:http (LISTEN)
httpd   20175 apache    8u  IPv6  31604      0t0  TCP *:http (LISTEN)
httpd   20176 apache    8u  IPv6  31604      0t0  TCP *:http (LISTEN)
httpd   20177 apache    8u  IPv6  31604      0t0  TCP *:http (LISTEN)

[root@mysql-backup ~]# /etc/init.d/heartbeat start
[root@mysql-backup ~]# lsof -i :80

# 在主上模拟服务器宕机(挂起),在备上查看httpd的状态,整个切换过程需要30秒以上;
[root@mysql-backup ~]# lsof -i :80
COMMAND   PID   USER   FD   TYPE DEVICE SIZE/OFF NODE NAME
httpd   20143   root    8u  IPv6  32613      0t0  TCP *:http (LISTEN)
httpd   20164 apache    8u  IPv6  32613      0t0  TCP *:http (LISTEN)
httpd   20165 apache    8u  IPv6  32613      0t0  TCP *:http (LISTEN)
httpd   20166 apache    8u  IPv6  32613      0t0  TCP *:http (LISTEN)
httpd   20167 apache    8u  IPv6  32613      0t0  TCP *:http (LISTEN)
httpd   20168 apache    8u  IPv6  32613      0t0  TCP *:http (LISTEN)
httpd   20169 apache    8u  IPv6  32613      0t0  TCP *:http (LISTEN)
httpd   20170 apache    8u  IPv6  32613      0t0  TCP *:http (LISTEN)
httpd   20171 apache    8u  IPv6  32613      0t0  TCP *:http (LISTEN)
```

## haresources的配置参数详解

> 在两台机器上分别拷贝httpd脚本到/etc/ha.d/resource.d/下,并确保具备可执行权限;
```
$ cp /etc/init.d/httpd /etc/ha.d/resource.d
```
> 如果不执行拷贝也OK,heartbeat可以通过自定义的脚本找到/etc/init.d/httpd

> heartbeat控制脚本的要求:
> * ①脚本路径要放入/etc/init.d/或/etc/ha.d/resource.d/;
> * ②脚本执行需要以/etc/init.d/httpd start/stop方式;
> * ③脚本要具备可执行权限(chmod +x);
> * ④/etc/init.d/httpd名字和mysql-master IPaddr::172.16.1.164/24/eth0 httpd相同,提示: `/etc/ha.d/resource.d`/为heartbeat的默认脚本文件目录
> * haresources文件注解
```
mysql IPaddr::172.16.1.164/24/eth0 drbddisk::data Filesystem::/dev/drbd0::/data::ext3 rsdata

# mysql <===为主机名,表示初始状态会在mysql绑定 IP 172.16.1.164
# IPaddr <===为heartbeat配置IP的默认脚本,其后的IP等都是其参数;
# 172.16.1.164/24/eth0 <===为集群对外服务的VIP,初始启动在mysql上,24为子网掩码,eth0为ip绑定的实际物理网卡,为heartbeat提供对外服务的通信接口;等价于下面
# /etc/ha.d/resource.d/IPaddr 172.16.1.164/24/eth0 stop/start #<===启VIP,清ARP缓存
# drbddisk::data #<===启动drbd data资源,这里相当于执行/etc/ha.d/resource.d/drbddisk data stop/start
# Filesystem::/dev/drbd0::/data::ext3 #<===drbd分区挂载到/data目录,这里相当于执行/etc/ha.d/resource.d/FIlesystem /dev/drbd0 /data ext3 stop/start
# rsdata #<===启动NFS/MFS服务脚本;这里相当于执行/etc/init.d/rsdata stop/start或/etc/ha.d/resource.d/rsdata start/stop

# 资源生效顺序:从左到右
    mysql drbddisk::data Filesystem::/dev/drbd0::/data::ext3 rsdata IPaddr::172.16.1.164/24/eth0
    1.设置drbd,drbddisk::data
    2.挂载/data到/dev/drbd0, Filesystem::/dev/drbd0::/data::ext3
    3.NFS/MFS服务配置,rsdata(不要开机启动)
    4.启动VIP IPaddr::172.16.1.164/24/eth0

# 这样用户的体验最好;
  一旦heartbeat无法控制资源的启动,heartbeat会采取极端的措施,例如重启系统,来释放没法管理的资源,因此,被管理的资源必须是不能开机启动的;可以写脚本来强制控制资源的处理,来防止heartbeat重启;
  先起VIP,可能没有数据,用户体验不好！！！
```

## 生产场景heartbeat调用资源的应用

> 在实际的工作场景中,heartbeat调用资源最常见的两种方式: 
> * ①heartbeat可以仅控制vip资源的漂移,不负责服务资源的启动及停止,本次实验的httpd服务就可以这样做;<===适合web服务;
> * ②heartbeat即控制vip资源的漂移,同时又控制服务资源启动及停止,本次实验的httpd服务例子既是,ip和服务都要切换;<===适合数据服务(数据库和存储)只能一端写入

> VIP正常,httpd服务宕了;这个时候不会做高可用切换;写个简单的脚本定时或守护进程判断httpd服务,如果有问题,则停止heartbeat,主动使线上的业务转移到另一台;

> 两端服务能同时起,那最好不要交给heartbeat,对于某些服务,不能两端同时起,heatbeat可以控制服务的启动;

## ha高可用httpd案例结论:

> ①日志很重要,不管是heartbeat,所有服务的日志都很重要;有问题时多查看相关日志;

> ②httpd的高可用还可以两边都处于启动状态,即httpd不需要交给ha起,而是默认状态就先启动运行;

> ③这个httpd高可用性配置在生产环境中用的很少,但是它却是生产环境需求的一个初级模型;如: heartbeat+brbd+mysql实现数据库高可用性配置,heartbeat+active/active+nfs/mfs实现存储高可用性配置;

## heartbeat和Keepalived应用场景区别

> ①对于一般的web、db、负载均衡(nginx、Haproxy)等,heartbeat和Keepalived都可以实现;

> ②lvs负载均衡最好和Keepalived结合,虽然heartbeat也可以调用带有ipvsadm命令的脚本来启动和停止lvs负载均衡,但是heartbeat本身并没有对下面节点rs的健康检查功能,heartbeat的这个缺陷可以通过ldirecetord插件来弥补;

> ③需要数据同步(配合drbd)的高可用业务最好用heartbeat,例如: mysql双主多从,NFS/MFS存储,它们的特点是需要数据同步,这样的业务最好用heartbeat;因为heartbeat自带了drbd的脚本,可以利用强大的drbd同步软件配合实现同步;如果你解决了数据同步可以不用drbd,例如: 共享存储或者inotify+rsync(sersync+rsync),那么就可以考虑Keepalived;

> ④运维人员对哪个更熟悉就用哪个,其实就是要你能控制维护你部署的服务;目前,总的来说还是使用Keepalived软件的多一些;

## heartbeat服务生产环境下维护要点

> 在我们每天的实战运维工作中,当有新项目上线或者VIP更改需求时,可能会进行添加修改服务VIP的操作,下面就是heartbeat+Haproxy/nginx高可用负载均衡的生产环境下的维护方法: 
> * 所有配置放到SVN,更改后提交SVN,对比,推送到正式环境;
> * 常见的情况就是修改配置文件,我们知道配置文件有3个,ha.cf、authkeys、haresource;
```
①在修改配置前执行/etc/init.d/heartbeat stop或/usr/lib64/heartbeat/hb_standby(此命令最好)把本机业务推送到备节点工作,当确认备节点正常工作后,开始修改本地的配置,修改好后可以在执行/etc/init.d/heartbeat stop/start把资源服务接管回来;记得在把业务推到备节点时及修改配置接管回服务时都要立即查看服务是否正常工作,特别是所有的VIP(新的旧的)是否启动OK,URL地址是不是能打开,这个检查过程可以写成脚本放heartbeat服务启动脚本的参数里等;

②先修改好一端的配置,然后同步到另一端,此时,准备好如下的命令操作: 
   /etc/init.d/heartbeat stop
   /etc/init.d/heartbeat stat
   ifconfig|egrep "ip1|ip2"
   wget url
   准备好后,拷贝粘贴同时执行上面3条命令,执行完毕后查看ip是否OK,如果5秒内IP不OK,则需要回滚配置或再次推到备节点;

# 方法一: 如果要回滚,将主服务覆盖备份的配置文件重启
/bin/cp haresource.bak hareasource.new
/etc/init.d/heartbeat stop
/etc/init.d/heartbeat start

# 方法二: 把业务推到备节点,然后在认真检查主节点配置 /usr/lib64/heartbeat/hb_standby
```

## 高可用切换的特别说明

> ①高可用服务的切换一般用于主故障备用自动切换接管,快速顶替故障机提供服务;

> ②备接管后的善后工作等,最好有人工处理解决！

> ③不管准备多么完善,监控多么智能,一般都不会自动切回主库;

> ④而是人工控制,因为这个回切是可控的,有时间准备;

> ⑤而一开的主挂了,备用顶替这个死突然的、不可控的;

> ⑥重要数据的业务是不能来回自动切换的,即auto_failback off应该是off状态;

> ⑦有关主的heartbeat是不是开机要自启动这个要具体业务具体分析,特别是控制资源的场景;