标签： 反向代理

# LVS

---

[toc]

## Linux Cluster

### 系统扩展的方式

> * scale up: 向上扩展/纵向扩展
> * scale out: 向外扩展/横向扩展

### 集群类型

> * LB:负载均衡集群 Load Banlancing
> * HA:高可用性集群 High Availability
> * HP:高性能集群	High Performa>ncing
> * 大规模并行处理平台:Hadoop(MapReduce)

### 系统可用性
	
> * Availability(可用性百分比)=平均无故障时间/平均无故障时间+平均故障修复时间{90% 95% 99% 99.9% 99.99%}
> * 系统衡量指标:(容量和性能要取折衷方案)
> * 可扩展性
> * 可用性
> * 容量(保证可接受性能的情况下,所能够提供的最大吞吐量)
> * 性能(强调系统的响应时间)  		
> * 系统运维:可用 --> 标准化 -->自动化
> * https://www.top500.org
> * 稳定压倒一切
> * 服务组件遵循一致标准	
> * 构建高可扩展性的重要原则:在系统内部尽量避免串行化和交互
> * GSLB:全局负载均衡,Global Service Load Balancing
> * SLB:服务负载均衡,Service Load Balancing
		
## LB集群需求

> * 负载均衡(Load Balance)集群提供了一种廉价、有效、透明的方法,来扩展网络设备和服务器的负载、贷款和吞吐量，同时加强了网络数据处理能力,提高了网络的灵活性和可用性;
> * 搭建负载均衡服务器的需求如下:
 - 把单台计算机无法承受的大规模并发访问或数据流量分担到多台节点设备上,分别进行处理,减少用户等待响应的时间,提升用户体验;
 - 单个重负载均衡的运算分担到多台节点设备上做并行处理,每个节点设备处理结束后,将结果汇总,返回给用户,系统处理能力得到了大幅度提高;
 - 7*24小时的服务保证,任意一个或多个有限后面节点设备宕机,不能影响业务

## LB集群的实现

### 硬件

> * F5 -- > BIG-IP
> * Citrix 	NetScaler
> * A10		A10
> * Array
> * Redware

### 软件

> * lvs,haproxy,nginx最为流行
 - lvs
 - haproxy
 - nginx
 - ats(apache软件基金会产品,apache traffic server)
 - perlbal
> * 基于工作的协议层次划分:
 - 传输层: lvs,haproxy(mode tcp)
 - 应用层: haproxy,nginx,ats,perbal

## LVS相关概念

> * LVS: Linux Virtual Service,即为Linux虚拟服务器,是一个虚拟的服务器集群系统,可以在UNIX/LINUX平台下实现负载均衡集群功能;该项目在1998年5月由章文嵩博士组织成立,是中国国内最早出现的自由软件项目之一;
> * LVS负载均衡调度技术是在Linux内核中实现的,因此被称之为Linux虚拟服务器(Linux Virtual Server);我们使用该软件配置LVS时候,不能直接配置内核中的ipvs,而需要使用ipvs的管理工具ipvsadm进行管理,但Keepalived软件可以直接管理ipvs,并不是通过ipvsadm管理ipvs;
> * LVS真正实现负载均衡调度的工具是IPVS,工作在linux内核层面;
> * LVS自带的IPVS管理工具是ipvsadm(用户控件的命令行工具,用于管理集群服务)
> * keepalived可以实现管理ipvs及对负载均衡器的高可用
> * ipvs工作在内核中netfilter INPUT钩子上,支持TCP,UDP,AH,EST,AH_EST,SCTP等诸多协议
 - ①根据请求报文的目标IP和PORT将其转发至后端主机集群中的某一台主机(根据挑选算法)
 - ②netfilter:
   PREROUTING --> INPUT
   PREROUTING --> FORWARD --> POSTROUTING
   OUTPUT --> POSTROUTING

## LVS集群术语

> * 调度器: director,dispatcher,balancer
> * RS: Real Server(真实节点)
> * Client IP: CIP(客户端IP)
> * Director Virtual IP(调度器虚拟IP)
> * Director IP: DIP(调度器IP)
> * Real Server IP: RIP(后端节点IP)

## lvs工作模式

### NAT

> lvs-nat(masquerade,地址伪装)
> * 多目标的DNAT(iptables): 它通过修改请求报文的目标IP地址(同时可能会修改目标端口)至挑选出来某RS的RIP地址实现转发
> * LVS NAT模式强行将由PRERROUTING链上发送到INPUT链上的数据包规则改写到POSTROUTING链上!
> * NAT模式注意要点:
 - ①RS应该和DIP都应该使用私网地址RS的网关要指向DIP
 - ②请求和响应报文都要经由director转发,极高负载的场景中,director可能会成为系统瓶颈
 - ③支持端口映射(请求和响应均经过调度器)
 - ④RS可以使用任意OS
 - ⑤RS的RIP和Director的DIP必须在同一网段
 - NAT技术将请求的报文(通过DNAT方式改写)和响应报文(通过SNAT方式改写),通过调度器地址重写后再转发给内部的服务器,报文返回时在改写成原先的用户请求的地址;

### DR

> lvs-dr(direct routing,直接路由,gateway)

> * dr通过修改请求目标的MAC地址进行转发
 - Director: DIP,VIP
 - RS: RIP, VIP
> * DIP和RIP要在同一物理网络中
> * DR模式注意要点:
 - ①保证前端路由器将目标IP为VIP的请求报文发送给director(静态绑定;arptables;修改RS主机内核的参数/linux特有)
 - ②RS的RIP可以使用私有地址;但也可以使用公网地址
 - ③RS跟Director必须在同一物理网络中
 - ④请求报文经由Director调度,但响应报文是由RS直接响应给Client;
 - ⑤不支持端口映射
 - ⑥RS可以是大多数OS(支持arp广播抑制等内核参数)
 - ⑦RS的网关不能指向DIP

### TUN

> lvs-tun(ip tunneling,ipip)

> * tun模式响应报文不能超过MTU(最大传输单元)大小
> * 不修改请求报文的ip首部,而是通过在原有的ip首部(cip<-->vip)之外,再封装一个首部(dip<-->rip)
> * TUN模式注意要点:
 - ①RIP、DIP、VIP全得是公网地址
 - ②RS的网关不能指向DIP
 - ③请求报文必须经由director调度,但响应报文必须不能经由director
 - ④不支持端口映射
 - ⑤RS的OS必须支持隧道功能

### FULLNAT

> * lvs-fullnat: director通过同时修改请求报文的源地址和目标地址进行转发
> * LVS当前应用主要采用DR和NAT模式,但这2中模式要求RealServer和LVS在同一个vlan中,导致部署成本过高;TUNNEL模式虽然可以跨vlan,但是RealServer上需要部署ipip隧道模块等,网络拓补上需要连通外网,较复杂,不宜运维
> * FULLNAT模式和NAT模式的区别是: Packet IN时,除了做DNAT,还做SNAT(用户ip->内网ip),从而实现LVS-RealServer间可以跨vlan通讯,RealServer只需要连接到内网;
> * FULLNAT将作为一种新的工作模式(同DR/NAT/TUNNEL),实现如下功能:
 - Packet IN时,目标ip更换为RealServer IP,源IP更换为内网local IP
 - Packet OUT时,目标ip更换为client ip,源ip更换为vip
 - local ip为一组内网IP地址,性能和NAT相比,正常转发性能下降<10%
> * FULLNAT模式注意要点:
 - ①VIP是公网地址,RIP和DIP是私网地址,二者无需在同一网络中
 - ②RS接受到的请求报文的源地址为DIP,因此要响应DIP;
 - ③请求报文和响应报文都必须经由Director
 - ④支持端口映射机制
 - ⑤RS可以使用任意OS

## http stateless

> * http协议是无状态的
> * session保持方案
 - ①source ip hash
 - ②cookie
> * session集群: 浪费资源,消耗性能
> * session服务器: memcache,redis等

## LVS调度方法

> * lvs scheduler: lvs调度算法
> * 静态方法: 仅根据算法本身进行调度
 - RR: round robin,轮询
 - WRR: wrighted round robin,加权轮询
 - SH: source hash: 实现session保持的机制,将来自同一个IP的请求始终调度至同一RS,有损集群的负载均衡的效果
 - DH: destination hash: 将对同一个目标的请求始终发往同一个RS
> * 动态方法: 根据算法及各RS的当前负载状态进行调度
 - LC: Least-Connection,最小连接数,Overhead=Active*256+Inactive
 - WLC: Weighted Least-Connection,加权最小连接数,Overhead=(Active*256+Inactive)/weight
 - SED: Shortest Expection Delay,最短期望延迟,Overhead=(Active+1)*256/weight
 - NQ: Never Queue,永不排队,NQ是基于SED算法的改进
 - LBLC: Locality-Based Least-Connection，即动态的LC算法，正向代理情形下的cache server调度
 - LBLCR: Locality-Based Least-Connection with Replicaion,带复制功能的LBLC算法
> * 基于局部性的最少链接LBLC和带复制的基于局部性最少链接LBLCR主要使用于Web Cache和Db Cache集群，但是我们很少这样用;普遍一致性哈希
> * 源地址散列调度SH和目标地址散列调度DH可以结合使用在防火墙集群中,它们可以保证整个系统的唯一出入口
> * 最短预期延时调度SED和不排队调度NQ主要是对处理时间相对较长的网络服务。

## ipvs集群服务

> * 支持协议: tcp,udp,ah,esp,ah_esp,sctp
> * 一个ipvs主机可以同时定义多个cluster service,协议: tcp,udp
> * 一个cluster service上至少应该有一个real server,定义时指明lvs的工作模式和调度算法

### lvs安装

> * ipvs模块工作在内核空间,用户空间不能直接操作,需要通过ipvsadm管理工具操作;
```
# 安装ipvsadm管理工具
$ rpm -qa ipvsadm
$ yum install ipvsadm -y

# 创建linux内核的软链接
$ ln -s /usr/src/kernel/`uname -r` /usr/src/linux

# 查看是否有ipvsadm组件
$ lsmod|grep ip_vs
``` 

### vip添加

> * 使用ip命令添加VIP
```
$ ip addr add 10.0.0.3/24 dev label eth0:0
```

### ipvsadm用法

> * 管理集群服务/Virtual Server
```
# ipvsadm添加/修改/删除Virtual Server服务
ipvsadm -A|E -t|u|f service-address [-s scheduler] [-p [timeout]] [-M netmask]
ipvsadm -D -t|u|f service-address

# -A: 添加Virtual Server
# -E: 修改指定的Virtual Server
# -D: 删除指定的Virtual Server

# service-address: Virtual Server地址
 -t: ip:port/tcp协议
 -u: ip:port/udp协议
 -f: FireWall Mark(防火墙标记/iptables)

# -s scheduler: 指定调度算法,默认为wlc

# 案例:
$ ipvsadm -A -t 10.0.0.3:80 -s wrr
```
> * 管理集群服务中的RS
```
ipvsadm -a|e -t|u|f service-address -r server-address [-g|i|m] [-w weight]
ipvsadm -d -t|u|f service-address -r server-address

# -a: 给Virtual Server添加Real Server节点
# -e: 给Virtual Server修改Real Server节点信息
# -d: 删除Virtual Server下面的Real Server节点

# service-address: Virtual Server/ip[:port]
# -r server-address: 指定real server地址

# -g: 工作模式/gateway dr
# -i: 工作模式/ipip,tun
# -m: 工作模式/masquerade,nat
# -w: 指定对应节点的权重

# 添加rs节点
$ ipvsadm -a -t 10.0.0.3:80 -r 10.0.0.7:80 -g -w 1

# 删除rs节点
$ ipvsadm -d -t 10.0.0.3:80 -r 10.0.0.8:80 -g -w 1
```
> * 清空和查看ipvs规则
```
# 清空原有的ipvs规则
$ ipvsadm -C

# 查看ipvs规则
ipvsadm -L|l [options]
 -n: numberic,基于数字格式显示地址和端口
 -c: connection,显示当前ipvs已经建立的链接
 --stats: 显示ipvs链接的统计数据
 --rate: 显示统计速率
 --exact: 显示精确值

# 查看添加的Virtual Server和Real Server信息
$ ipvsadm -Ln
```
> * 保存和重载
```
# 重载ipvs规则
$ ipvsadm -R < /etc/sysconfig/ipvsadm

# 保存ipvs规则到指定文件
ipvsadm -S [-n] > /etc/sysconfig/ipvsadm
```
> * 置零计数器
```
# 针对集群服务设置
ipvsadm -Z [-t|u|f service-address]
```

## LVS工作模式配置

### NAT配置

> * NAT模式配置前提:
 - RS网关指向DIP
 - DIP开启路由转发,
> * LVS环境配置: 
 - DIP: 192.168.20.1(RS网关指向此地址)
 - VIP: 172.16.100.9
 - RS1: 192.168.20.7
 - RS2: 192.168.20.8
> * ipvs集群服务配置
```
# 开启网卡间核心转发
$ sysctl -w net.ipv4.ip_forward = 1

# 添加集群服务地址
$ ipvsadm -A -t 172.16.100.9:80 -s rr
$ ipvsadm -L -n

# 添加real server地址, real server网关指向192.168.20.1
$ ipvsadm -a -t 172.16.100.9:80 -r 192.168.20.7 -m(伪装,模式为nat)
$ ipvsadm -a -t 172.16.100.9:80 -r 192.168.20.8 -m(伪装,模式为nat)
$ ipvsadm -L -n
```
> * 保存/重载规则
```
# 保存ipvs规则到指定文件
$ ipvsadm -S > /etc/sysconfig/ipvsadm  

# 清空ipvs规则
$ ipvsadm -C 
$ ipvsadm -L -n

# 导入已有的ipvs规则
$ ipvsadm -R < /etc/sysconfig/ipvsadm
```
> * 修改lvs调度算法
```
$ ipvsadm -E -t 172.16.100.9:80 -s sh(source IP hash)
```
> * 配置端口映射
```
# VIP:PORT --> RS:PORT
$ ipvsadm -e -t 172.16.100.9:80 -r 192.168.20.7:8080 -m
$ ipvsadm -e -t 172.16.100.9:80 -r 192.168.20.8:8080 -m
```

###  DR配置

> * DR模式两个内核参数设置(<font color="red">重点</font>)
> * `arp_ignore`: 是否响应别人的arp广播请求
 - 0: 使用任何地址响应,默认值
 - 1: 使用本地地址响应arp请求
> * `arp_announce`: 是否接受通告,记录通告并通告给别人
 - 0: 默认值,通告所有地址给其他主机
 - 1: 尽量不适用本地地址接受通告
 - 2: 使用最佳的本地地址接受通告
> * Director设置
```
# 配置Director的VIP
$ ifconfig eth0:0 172.16.100.10/32 broadcast 172.16.100.10 up
$ route add -host 172.16.100.10 dev eth0:0
```
> * 配置Real Server
```
# arp抑制参数设置
$ echo 1 > /proc/sys/net/ipv4/conf/all/arp_ignore
$ echo 1 > /proc/sys/net/ipv4/conf/eth0/arp_ignore
$ echo 2 > /proc/sys/net/ipv4/conf/all/arp_announce
$ echo 2 > /proc/sys/net/ipv4/conf/eth0/arp_announce

# 配置Real Server的VIP(绑定在lo上)
$ ifconfig lo:0 172.16.100.10/32 broadcast 172.16.100.10 up
$ route add -host 172.16.100.10 dev lo:0
```
> * 配置ipv规则
```
$ ipvsadm -C
$ ipvsadm -A -t 172.16.100.10:80 -s rr
$ ipvsadm -a -t 172.16.100.10:80 -r 172.16.100.21 -g(gateway,DR模式)
$ ipvsadm -a -t 172.16.100.10:80 -r 172.16.100.22 -g(gateway,DR模式)
```

## ipvs基于防火墙标记转发

### FWM转发介绍

> * ipvsadm -A|E -t|u|f service-address [-s scheduler] [-p timeout]] [-M netmask]
```
-f, --fwmark-service integer
    Use a firewall-mark, an integer value greater than zero, to denote a virtual service instead of an address, port  and  protocol (UDP  or  TCP).  The marking of packets with a firewall-mark is configured using the -m|--mark option to iptables(8). It can be used to build a virtual service associated with the same  real  servers,  covering  multiple  IP  address,  port  and  protocol triplets. If IPv6 addresses are used, the -6 option must be used.
    Using firewall-mark virtual services provides a convenient method of grouping together different IP addresses, ports and protocols into a single virtual service. This is useful for both simplifying configuration if a large number of virtual services are required and grouping persistence across what would otherwise be multiple virtual services.
```
> * iptabels/netfilter数据包流向
```
PREROUTING --> INPUT
PREROUTING --> FORWARD --> POSTROUTING
OUTPUT --> POSTROUTING
```
> * ipvs工作在INPUT链上

### FWM转发配置

> * 将ssh,http,https服务的数据包进入mangle表的PREROUTING链时打上标记mark
```
# 使用iptables在mangle表的PREROUTING链上添加规则
$ iptables -t mangle -A PREROUTING -d $vip -p $protocol -m multiport --dports 22,80,443 -j MARK -set-mark $sum

# $vip: vip地址
# $protocol: 协议
# $sum: 标记值
```
> * ipvsadm定义集群服务和添加规则
```
# Director上配置vip
$ ifconfig eth0:0 172.16.100.10/32 broadcast 172.16.100.10 up
$ route add -host 172.16.100.10 dev eth0:0
$ ipvsadm -C
$ ipvsadm -A -f 10 -s rr

# Real Server在lo上绑定VIP地址，并且修改内核参数做ARP广播抑制
$ ifconfig lo:0 172.16.100.10/32 broadcast 172.16.100.10 up
$ route add -host 172.16.100.10 dev lo:0
$ echo 1 >/proc/sys/net/ipv4/conf/all/arp_ignore
$ echo 1 >/proc/sys/net/ipv4/conf/eth0/arp_ignore 
$ echo 2 >/proc/sys/net/ipv4/conf/all/arp_announce
$ echo 2 >/proc/sys/net/ipv4/conf/eth0/arp_announce

# Director上添加配置RS
$ ipvsadm -a -f 10 -r 172.16.100.21 -g
$ ipvsadm -a -f 10 -r 172.16.100.22 -g
```
> * ipvsadm基于FWM定义集群的功用: 将共享一组RS的集群服务统一进行调度

## lvs集群session保持

> * lvs集群的session保持机制
 - session绑定: lvs sh算法,针对某一特定服务
 - session复制: 易造成集群session混乱
 - session服务器: 使用redis实现会话共享
> * lvs persistence: lvs的持久链接
 - 功能: 无论ipvs使用何种调度方法,都能实现将来自同一个client的请求始终定向至第一次调度时挑选的RS
 - 持久连接模板: 独立于算法,source ip rs timer
> * 对于共享同一组的RS的服务器,需要进行统一绑定
> * 持久链接的实现方式
 - 每端口持久: PPC,单服务持久调度
 - 每FWM持久: PFWMC,单FWM持久调度,PORT AFFINITY
 - 每客户端持久: PCC,单客户端持久调度; director会将用户的任何请求都识别为集群服务,并向RS进行调度,TCP: 1-65535,UDP: 1-65535,0代表所有端口

## LVS集群HA

> * SPOF: Single Point of Failure/单点故障
> * Director: 高可用集群,通过keepalived实现
> * RealServer: 让director对其做健康检查,并且根据检测的结果自动完成添加或移除等管理功能
> * LVS健康检查
```
# 1.基于协议层次检查
  ip: icmp
  传输层: 检测端口的开放状态/port
  应用层: 请求获取关键性的资源/http method

# 2.检查频度,定义检查的时间间隔,连续多次超时则认为故障,下线

# 3.状态判断
 下线: ok --> failure --> failure --> failure
 上线: failure --> ok --> ok

# 4.back server
  设置备用服务器
```

## LVS健康检查脚本

> * 开发LVS检查检测脚本
```
#!/bin/bash

fwm=6
sorry_server=127.0.0.1
rs=('172.16.100.21' '172.16.100.22')
rw=('1' '2')
type='-g'
chkloop=3
rsstatus=(0 0)
logfile=/var/log/ipvs_health_check.log

addrs() {
    ipvsadm -a -f $fwm -r $1 $type -w $2
    [ $? -eq 0 ] && return 0 || return 1
}

delrs() {
    ipvsadm -d -f $fwm -r $1
    [ $? -eq 0 ] && return 0 || return 1
}
			
chkrs() { 
    local i=1
    while [ $i -le $chkloop ]; do
        if curl --connect-timeout 1 -s http://$1/.health.html | grep "OK" &> /dev/null; then
            return 0
        fi
        let i++
        sleep 1
    done
    return 1
}

initstatus() {
    for host in `seq 0 $[${#rs[@]}-1]`; do 
        if chkrs ${rs[$host]}; then
            if [ ${rsstatus[$host]} -eq 0 ]; then
                rsstatus[$host]=1
            fi
        else
            if [ ${rsstatus[$host]} -eq 1 ]; then
                rsstatus[$host]=0
            fi
        fi
    done
}
			
initstatus
while :; do
    for host in `seq 0 $[${#rs[@]}-1]`; do 
        if chkrs ${rs[$host]}; then
            if [ ${rsstatus[$host]} -eq 0 ]; then
                addrs ${rs[$host]} ${rw[$host]}
	        [ $? -eq 0 ] && rsstatus[$host]=1
            fi
        else
            if [ ${rsstatus[$host]} -eq 1 ]; then
                delrs ${rs[$host]} ${rw[$host]}
                [ $? -eq 0 ] && rsstatus[$host]=0
            fi
        fi
    done
sleep 5
done
```

## DR类型director部署脚本

```
#!/bin/bash

vip=172.16.100.33
rip=('172.16.100.8' '172.16.100.9')
weight=('1' '2')
port=80
scheduler=rr
ipvstype='-g'

case $1 in
    start)
        iptables -F -t filter
        ipvsadm -C
				
        ifconfig eth0:0 $vip broadcast $vip netmask 255.255.255.255 up
        route add -host $vip dev eth0:0
        echo 1 > /proc/sys/net/ipv4/ip_forward

        ipvsadm -A -t $vip:$port -s $scheduler
        [ $? -eq 0 ] && echo "ipvs service $vip:$port added."  || exit 2
        for i in `seq 0 $[${#rip[@]}-1]`; do
            ipvsadm -a -t $vip:$port -r ${rip[$i]} $ipvstype -w ${weight[$i]}
            [ $? -eq 0 ] && echo "RS ${rip[$i]} added."
        done
        touch /var/lock/subsys/ipvs
        ;;
    stop)
        echo 0 > /proc/sys/net/ipv4/ip_forward
        ipvsadm -C
        ifconfig eth0:0 down
        rm -f /var/lock/subsys/ipvs
        echo "ipvs stopped."
        ;;
    status)
        if [ -f /var/lock/subsys/ipvs ]; then
            echo "ipvs is running."
            ipvsadm -L -n
        else
            echo "ipvs is stopped."
        fi
        ;;
    *)
        echo "Usage: `basename $0` {start|stop|status}"
        exit 3
        ;;
esac
```

## DR类型RS部署脚本

```
#!/bin/bash

vip=172.16.100.33
interface="lo:0"

case $1 in
    start)
        echo 1 > /proc/sys/net/ipv4/conf/all/arp_ignore
        echo 1 > /proc/sys/net/ipv4/conf/lo/arp_ignore
        echo 2 > /proc/sys/net/ipv4/conf/all/arp_announce
        echo 2 > /proc/sys/net/ipv4/conf/lo/arp_announce
        ifconfig $interface $vip broadcast $vip netmask 255.255.255.255 up
        route add -host $vip dev $interface
        ;;
    stop)
        echo 0 > /proc/sys/net/ipv4/conf/all/arp_ignore
        echo 0 > /proc/sys/net/ipv4/conf/lo/arp_ignore
        echo 0 > /proc/sys/net/ipv4/conf/all/arp_announce
        echo 0 > /proc/sys/net/ipv4/conf/lo/arp_announce
        ifconfig $interface down
        ;;
    status)
        if ifconfig lo:0 |grep $vip &> /dev/null; then
            echo "ipvs is running."
        else
            echo "ipvs is stopped."
        fi
        ;;
    *)
        echo "Usage: `basename $0` {start|stop|status}"
        exit 1
esac
```

## LVS+Keepalived实现负载均衡高可用(DR)

### RS节点配置

> * 添加VIP
```
$ ip addr add 10.0.0.3/32 dev lo label lo:0
```
> * 添加主机路由
```
$ route add -host 10.0.0.3 dev lo:0
```
> * ARP广播抑制
```
$ echo "1" > /proc/sys/net/ipv4/conf/lo/arp_ignore
$ echo "1" > /proc/sys/net/ipv4/conf/all/arp_ignore
$ echo "2" > /proc/sys/net/ipv4/conf/lo/arp_announce
$ echo "2" > /proc/sys/net/ipv4/conf/all/arp_announce
```

### Keepalived配置

> * Director安装好LVS服务,手工测试ok,安装keepalived高可用服务
> * 清空LVS原有的负载均衡配置
```
$ ipvsadm -C
```
> * master负载均衡keepalived配置
```
$ cat /etc/keepalived/keepalived.conf
! Configuration File for keepalived

global_defs {
    router_id lb01
    vrrp_mcast_group4 224.0.1.18
}

vrrp_instance VI_1 {
    state MASTER
    interface eth0
    virtual_router_id 55
    priority 150
    advert_int 1
    authentication {
        auth_type PASS
        auth_pass 1111
    }
    virtual_ipaddress {
      10.0.0.3/24 dev eth0 label eth0:1
    }
}

vrrp_instance VI_2 {
    state BACKUP
    interface eth0
    virtual_router_id 56
    priority 50
    advert_int 1
    authentication {
        auth_type PASS
        auth_pass 1112
    }
    virtual_ipaddress {
      10.0.0.4/24 dev eth0 label eth0:2

    }
}

#ipvsadm -A -t 10.0.1.3:80 -s rr -p 300
virtual_server 10.0.0.3 80 {
    delay_loop 6          
    lb_algo wrr                
    lb_kind DR                
    nat_mask 255.255.255.0
    #persistence_timeout 10
    protocol TCP            
#ipvsadm -a -t 10.0.1.3:80 -r 10.0.0.7:80 -g 
    real_server 10.0.0.7 80 {
        weight 1              
        TCP_CHECK {
            connect_timeout 8       
            nb_get_retry 3
            delay_before_retry 3
            connect_port 80
        }
    }
    real_server 10.0.0.8 80 {
        weight 1              
        TCP_CHECK {
            connect_timeout 8       
            nb_get_retry 3
            delay_before_retry 3
            connect_port 80
        }
    }
} 
```
> * BACKUP负载均衡keepalived配置
```
$ cat /etc/keepalived/keepalived.conf
! Configuration File for keepalived

global_defs {
    router_id lb02
    vrrp_mcast_group4 224.0.1.18
}

vrrp_instance VI_1 {
    state BACKUP
    interface eth0
    virtual_router_id 55
    priority 100
    advert_int 1
    authentication {
        auth_type PASS
        auth_pass 1111
    }
    virtual_ipaddress {
        10.0.0.3/24 dev eth0 label eth0:1
    }
}

vrrp_instance VI_2 {
    state MASTER
    interface eth0
    virtual_router_id 56
    priority 150
    advert_int 1
    authentication {
        auth_type PASS
        auth_pass 1112
    }
    virtual_ipaddress {
        10.0.0.4/24 dev eth0 label eth0:2
    }
}

# ipvsadm -A -t 10.0.1.3:80 -s rr -p 300
virtual_server 10.0.0.3 80 {
    delay_loop 6          
    lb_algo wrr                
    lb_kind DR                
    nat_mask 255.255.255.0
    persistence_timeout 50     
    protocol TCP             
# ipvsadm -a -t 10.0.1.3:80 -r 10.0.0.7:80 -g 
    real_server 10.0.0.7 80 {
        weight 1              
        TCP_CHECK {
            connect_timeout 8       
            nb_get_retry 3
            delay_before_retry 3
            connect_port 80
        }
    }
    real_server 10.0.0.8 80 {
        weight 1              
        TCP_CHECK {
            connect_timeout 8       
            nb_get_retry 3
            delay_before_retry 3
            connect_port 80
        }
    }
}
```

### HTTP健康检查配置

> * 上文中keepalived配置文件中的rs节点健康检测是基于TCP;
> * 下面实现基于HTTP的健康检测方法
```
! Configuration File for keepalived
global_defs {
    router_id LVS_DEVEL
    vrrp_mcast_group4 224.0.1.118
}

vrrp_script chk_mt {
    script "[[ -f /etc/keepalived/down ]] && exit 1 || exit 0"
    interval 1
    weight -20
}

vrrp_instance VI_1 {
    state MASTER
    interface eno16777736
    virtual_router_id 144
    priority 100
    advert_int 1
    authentication {
        auth_type PASS
        auth_pass 84ae57f7f4f6
    }

    virtual_ipaddress {
        172.16.100.88/16 dev eno16777736 label eno16777736:1
    }
				
    track_script {
        chk_mt
    }
				
    notify_master "/etc/keepalived/notify.sh master"
    notify_backup "/etc/keepalived/notify.sh backup"
    notify_fault "/etc/keepalived/notify.sh fault"
}

virtual_server 172.16.100.88 80 {
    delay_loop 6
    lb_algo wrr
    lb_kind DR
    nat_mask 255.255.0.0
    protocol TCP
    sorry_server 127.0.0.1 80

    real_server 172.16.100.6 80 {
        weight 1
        HTTP_GET {
            url {
                path /
                status_code 200 
            }
            connect_timeout 3
            nb_get_retry 3
            delay_before_retry 3
        }
    }
    real_server 172.16.100.69 80 {
        weight 2
        HTTP_GET {
            url {
                path /
                status_code 200 
            }
            connect_timeout 3
            nb_get_retry 3
            delay_before_retry 3
        }
    }
}
```

### 通知脚本

```
#!/bin/bash

vip=172.16.100.88
contact='root@localhost'

notify() {
    mailsubject="`hostname` to be $1: $vip floating"
    mailbody="`date '+%F %H:%M:%S'`: vrrp transition, `hostname` changed to be $1"
    echo $mailbody | mail -s "$mailsubject" $contact
}

case "$1" in
    master)
        notify master
        exit 0
        ;;
    backup)
        notify backup
        exit 0
        ;;
    fault)
        notify fault
        exit 0
        ;;
    *)
        echo 'Usage: `basename $0` {master|backup|fault}'
        exit 1
        ;;
esac
```