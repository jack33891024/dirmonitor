标签： 反向代理

# NGINX

---

[toc]

## 集群简介

> * 简单地说,集群就是指一组(若干个)相互独立的计算机,利用高速通信网络组成的一个较大的计算机服务系统,每个集群节点(即集群中的每台计算机)都是运行各自服务的独立服务器;这些服务器之间可以彼此通信,协同向用户提供应用程序、系统资源和数据,并以单一系统的模式加以管理,当用户客户机请求集群系统时,集群给用户的感觉就像一个单一独立的服务器,而实际上用户请求的是一组集群服务器;
> * 打开谷歌、百度的页面,看起来好简单,也许你觉得用几分钟就可以制作出相似的网页,而实际上,这个页面的背后是有成千上万服务器集群系统工作的结果;而这么多服务器维护和管理,以及相互协调工作也许就是运维人员未来工作的职责;
> * 若要用一句话来描述集群,即一堆服务器合作做同一件事,这些机器可能需要整个技术团队架构、涉及和统一协调管理;这些机器可以分布在同一机房,也可以分布在全国全球各个地方的多个机房;

## 为什么要使用集群

### 高性能(performance)

> * 一些国家重要的计算密集型应用(如天气预报、核试验模拟等),需要计算机有很强的运算处理能力;以全世界现有的技术,即使是大型机,其计算能力也是有限的,很难单独完成此任务;因为计算时间可能会相当长,也许几天,甚至几年或更久;因此,对于这列复杂的计算业务,便使用了计算机集群技术,几种几十上百台,甚至成千上万台计算机进行计算;
> * 大家耳熟能详的大型网站谷歌、百度、淘宝等,都不是几台大型机可以构建的,都是上万台服务器组成的高性能集群,分布于不同的地点;
> * 假如你配一个LNMP环境,每次只需要服务10个并发请求,那么单台服务器一定会比多个服务器集群要快;只有当并发或总请求数量超过单台服务器的承受能力时,服务器集群才会体现出优势;
> * 下图是某大型计算机设备的超强硬件配置:
![image_1d3qgkg071c4e3t41g8a1u8680sp.png-149.7kB][1]

### 价格有效性(Cost-effectiveness)

> * 通常一套系统集群架构,只需要几台或数十台服务器主机即可;与动则价值上百万的专用超级计算机相比便宜了很多;在达到同样性能需求的条件下,采用计算机集群架构比采用同等运算能力的大型计算机具有更高的性价比;
> * 早期的淘宝、支付宝的数据库等核心系统就是用的上百万的小型机服务器;后因使用维护成本太高以及扩展设备费用成几何级数翻倍,甚至发展为扩展瓶颈,人员维护也十分困难,最终使用PC服务器集群替换,比如,把数据库系统从小型机结合Oracle数据库迁移到MySQL开源数据库结合PC服务器上来;不但成本下降了,扩展和维护也更容易了;

### 可伸缩性(Scalablility)

> * 当服务负载、压力增长时,针对集群系统进行较简单的扩展即可满足需求,且不会降低服务质量;
> * 通常情况下,硬件设备若想扩展性能能力,不得不购买增加新的CPU和存储器设备,如果加不上去了,就不得不购买更高性能的服务器,就拿我们现有的服务器来讲,可以增加的设备总是有限的;如果采用集群技术,则只需要将新的单个服务器加入现有集群架构中即可,从访问的客户角度来看,系统服务无论是连续性还是性能上都机会没有变化,系统在不知不觉中完成了升级,加大了访问能力,轻松地实现了扩展;集群系统中节点数目可以增长到几千甚至上万个,其伸缩性远超单台超级计算机;

### 高可用性(Availability)

> * 单一的计算机系统总会面临设备损毁的问题,例如:CPU、内存、主板、电源、硬盘等,只要一个部件坏掉,这个计算机系统就可能会宕机,无法正常提供服务;在集群系统中,尽管部分硬件和软件也还是会发生故障,但整个系统的服务可以是每天24*7可用的;
> * 集群架构技术,可以使得系统在若干硬件设备故障发生时扔可以继续工作,这样就将系统的停机时间减少到了最小;集群系统在提高系统可靠性的同时,也大大减小了系统故障带来的业务损失,目前几乎100%的互联网的网站都要求7*24小时提供服务;

### 透明性(Transparency)

> * 多台独立计算机组成的松耦合集群系统构成一个虚拟服务器;用户或客户端程序访问集群系统时,就像访问一台高性能、高可用的服务器一样,集群中一部分服务器的上线下线不会中断整个系统服务,这对用户也是透明的;

### 可管理型(Manageability)

> * 整个系统可能在物理上很大,但其实容易管理,就像管理一个单一映像系统一样;在理性状况下,软硬件模块的插入能做到即插即用(Plug&Play);

### 可编程性(Programmability)

> * 在集群系统上,容易开发及修改各类应用程序;

## 集群的分类

### 集群的常见分类

> * 根据功能和结构划分:
 - 负载均衡集群(Load balancing clusters),简称LBC或者LB
 - 高可用性集群(High-availability clusters),简称HAC
 - 高性能计算集群(High-performance clusters),简称HPC
 - 网格计算(Grid computing)

### 集群种类介绍

> * 负载均衡集群:
 - 1、负载均衡集群为企业提供了更为实用、性价比更高的系统架构解决方案;负载均衡集群可以把很多客户集中访问请求负载压力尽可能平均地分摊在计算机集群中处理;客户访问请求负载通常包括应用程序处理负载和网络流量负载;这样的系统非常适合使用同一组应用程序为大量用户提供服务的模式,每个节点都可以承担一定的访问请求负载压力,并且可以实现访问请求在各个节点之间动态分配,以实现负载均衡;
 - 2、负载均衡集群运行时,一般是通过一个或多个前端负载均衡器将客户访问请求分发到后端的一组服务器上,从而达到整个系统的高性能和高可用性;一般高可用性集群和负载均衡集群会使用类似的技术,或同时具有高可用性与负载均衡的特点;
 - 3、负载均衡集群典型的开源软件包括: LVS、Nginx、Haproxy等,架构图如下:
![image_1d3qh208qq0n5eo13ri6rbtuv1j.png-38.1kB][2]
 - 负载均衡集群的作用:
   分担用户访问请求及数据流量(负载均衡)
   保持业务连续性,即7*24小时服务(高可用)
   用于与web业务及数据库从库等服务器的业务
> * 高可用集群:
 - 1、一般是指在集群中任意一个节点失效的情况下,该节点上所有任务会自动转移到其他正常的节点上;此处过程并不影响这个集群的运行,不同的业务会有若干秒的切换时间,db业务明显长于web业务的切换时间;
 - 2、当集群中的一个节点系统发生故障时,运行着的集群服务会迅速做出反应,将该系统的服务分配到集群中其他正在工作的系统上运行;考虑到计算机硬件和软件的容错性,高可用性集群的主要目的是使集群的整体服务尽可能可用;如果高可用性集群中的主节点发生了故障,那么这段时间内将由备节点代替它,备节点通常是主节点的镜像,当它代替主节点时,它可以完成接管主节点(包括IP地址及其他资源)的服务于;因此使集群系统环境对于用户来说是一致的,即不会影响用户的访问;
 - 3、高可用性集群使服务器系统的运行速度和响应速度会尽可能地快;它们经常利用在多台机器上运行的冗余节点和服务来相互跟踪,如果某个节点失败,它的替补者将在几秒钟或更短时间内接管它的职责;因此对用户而言,集群里的任意一台机器宕机,业务都不会受影响(理论情况下);
 - 4、高可用集群常用的开源软件包括Keepalived、Heartbeat等,架构图如下:
![image_1d3qhaljstoo1k7fgbklu1137f20.png-53.3kB][3]
 - 5、高可用性集群的作用:
   当一台机器宕机时,另外一台机器接管宕机的机器的IP资源,提供服务;
   常用于不易实现负载均衡的应用,比如负载均衡器,主数据库,主存储对之间;
> * 高性能计算集群:
 - 性能计算集群也称为并行计算;通常高性能计算集群涉及为集群开发的并行应用程序,已解决复杂的科学问题(天气预报、石油勘测、核反应模拟等);高性能计算集群对外就好像一台超级计算机,这种超级计算机内部由数十甚至上万个独立服务器组成,并且在公共消息传递层上进行通信以运行并行应用程序;在工作中实际是把任务切成蛋糕,然后下发到集群节点计算,计算后返回结果,然后继续领新任务计算,如此往复;

## 常见的集群软硬件介绍及选型

### 企业运维常见的集群软硬件产品

> * 互联网企业常用的开源软件: Nginx,LVS,Haproxy,Keepalived,Heartbeat;
> * 互联网企业常用的商业机器硬件有: F5,Netscaler,Radware,A10等,工作模式相当于haproxy的工作模式,下图就是F5硬件负载均衡:
![image_1d3qhn02m1nbhvr32rmku0vt62t.png-54.4kB][4]
> * 淘宝,赶集网,新浪等公司曾今使用过Netscaler负载均衡产品;集群硬件Netscaler的产品如下图:
![image_1d3qhp4ab1spb17dd15t91kjcokf3a.png-54.5kB][5]

### 对于集群软硬件产品如何选型

> * 当企业业务重要,技术力量又薄弱,并且希望出钱购买产品及获取更好的服务时,可以选择硬件负载均衡产品,如F5、Netscaler、Radware等,此类公司多为传统的大型非互联网企业,如银行、金融、证券、宝马、奔驰等;
> * 对于门户网站来说,大多会使用软件及硬件产品来分担单一产品的风险,如淘宝、腾讯、新浪等;融资了的企业会购买硬件产品,如赶集等网站;
> * 中小型互联网企业,由于起步阶段无利润可转或者利润很低,会希望通过使用开源免费的方案来解决问题,因此会雇佣专门的运维人员进行维护,例如:51CTO等;
> * 相比较而言,商业的负载均衡产品成本高,性能好,更稳定,缺点是不能二次开发;开源的负载均衡软件对运维人员的能力要求较高,如果是运维及开发能力强,那么开源软件的负载均衡是不错的选择,目前的互联网行业更偏向使用开源的负载均衡软件;

### 如何选择开源软件产品
	
> * 中小型互联网公司网站并发访问和总访问量不是很大的情况下,首选Nginx负载均衡,理由是Nginx负载均衡配置简单、使用方便、安全稳定社区活跃,使用的人逐渐增多,流行趋势好,另外一个实现负载均衡的类似产品为Haproxy(支持L4和L7负载,同样优秀,但社区不如Nginx活跃);
> * 如果要考虑Nginx负载均衡的高可用功能,建议首选Keepalived软件,理由是安装、配置简单、使用方便，安全稳定;
> * 如果是大型互联网公司,负载均衡产品可以使用LVS+Keepalived在前端做四层转发(一般是主备或主主,如果需要扩展可以使用DNS或前端使用OSPF),后端使用Nginx或者Haproxy做7层转发(可以扩展到百台),再后面是应用服务器,如果是数据库和存储的负载均衡和高可用,建议选择LVS+Heartbeat,LVS支持tcp转发且dr模式效率很高,Heartbeat可以配合drbd,不但可以进行VIP的切换,还可以支持块设备级别的数据同步以及资源服务的管理;

## Nginx负载均衡集群介绍

### 搭建负载均衡服务的需求

> * 负载均衡(Load Balance)集群提供了一种廉价、有效、透明的方法,来扩展网络设备和服务器的负载、带宽和吞吐量,同时加强了网络数据处理能力,提高了网络的灵活性和可用性;
> * 搭建负载均衡服务器的需求如下:
 - 把单台计算机无法承受的大规模并发访问或数据流量分担到多台节点设备上,分别进行处理,减少用户等待响应的时间,提升用户体验;
 - 单个重负载均衡的运算分担到多台节点设备上做并行处理,每个节点设备处理结束后,将结果汇总返回给用户,系统处理能力得到了大幅度提高;
 - 7*24小时的服务保证,任意一个或多个有限后面节点设备宕机,不能影响业务;

### 反向代理和负载均衡概念简介

> * 严格的说,Nginx仅仅是作为Nginx Proxy反向代理使用的,因为这个反向代理功能表现的效果是负载均衡集群的效果,所以本文称之为Nginx负载均衡;那么反向代理和负载均衡有什么区别呢？
> * 普通负载均衡软件,例如LVS,其实现的功能只是对请求数据包的转发(也可能会改写数据包)、传递,其中DR模式明显的特征是从负载均衡下面的节点服务器来看,接受到的请求还是来自访问负载均衡的客户端的真是用户,而反向代理就不一样了,反向代理接受访问用户的请求后,会代理用户重新发起请求代理下的节点服务器,最后把数据返回给客户端用户,在节点服务器看来,访问的节点服务器的客户端用户就是反向代理服务器了,而非真实的网站访问用户;
> * LVS等负载均衡是转发用户请求的数据包,而Nginx反向代理是接受用户的请求然后重新发起请求去请求其后面的节点服务器;

### 实现Nginx负载均衡的组件

> * 实现Nginx负载均衡的组件主要有两个,如下:

|Nginx http功能模块|模块说明|
|:---|:---|
|ngx_http_proxy_module|proxy代理模块,用于把请求抛给服务器节点或upstream服务器池|
|ngx_http_upstream_module|负载均衡模块,可以实现网站的负载均衡功能及节点的健康检查|

## Nginx负载均衡环境

> * 所有用户的请求统一发送到Nginx负载均衡器,然后负载均衡器根据调度算法来请求web01和web02

### 软硬件准备

> * 硬件准备: 4台VM虚拟机,两台负载均衡,两台做RS,如下:

|HOSTNANE|IP|说明|
|:---|:---|:---|
|lb01|10.0.0.5|Nginx主负载均衡器|
|lb02|10.0.0.6|Nginx辅负载均衡器|
|web01|10.0.0.8|web01服务器|
|web02|10.0.0.7|web02服务器|
> * 软件准备
```
# 4台VM系统环境如下
[root@lb01 ~]# cat /etc/redhat-release 
CentOS release 6.7 (Final)
[root@lb01 ~]# uname -r
2.6.32-573.el6.x86_64
[root@lb01 ~]# uname -m
x86_64

# 需要的软件(Nginx-1.6.3)
[root@lb01 tools]# wget http://nginx.org/download/nginx-1.6.3.tar.gz
[root@lb01 tools]# ls
nginx-1.6.3.tar.gz
```

### 安装Nginx

> * 4台VM上安装Nginx
> * 安装需要的依赖软件包
```
yum install openssl openssl-devel pcre pcre-devel -y
rpm -qa openssl openssl-devel pcre pcre-devel
```
> * 编译安装Nginx
```
useradd -u 888 -s /sbin/nologin -M www
mkdir /server/tools -p
cd /server/tools/
wget http://nginx.org/download/nginx-1.6.3.tar.gz
tar xf nginx-1.6.3.tar.gz
cd nginx-1.6.3
./configure \
--prefix=/application/nginx-1.6.3 \
--user=www \
--group=www \
--with-http_ssl_module \
--with-http_stub_status_module
make
make install 
ln -s /application/nginx-1.6.3/ /application/nginx
echo "/application/nginx/sbin/nginx" >>/etc/rc.local
```

### 配置RS节点

> * Nginx web01和web02配置如下:
```
$ cat nginx.conf
worker_processes  1;
events {
  worker_connections  1024;
}
http {
  include       mime.types;
  default_type  application/octet-stream;
  sendfile        on;
  keepalive_timeout  65;
  log_format  main  '$remote_addr - $remote_user [$time_local] "$request" '
                    '$status $body_bytes_sent "$http_referer" '
                    '"$http_user_agent" "$http_x_forwarded_for"';
  server {
      listen       80;
      server_name  bbs.xxx.org;
      location / {
          root   html/bbs;
          index  index.html index.htm;
      }
      access_log logs/access_bbs.log main;
      }
  server {
      listen       80;
      server_name  www.xxx.org;
      location / {
          root   html/www;
          index  index.html index.htm;
      }
      access_log logs/access_www.log main;
      }
  }
```
> * 启动Nginx
```
# 创建站点目录
mkdir /application/nginx/html/{www,bbs}
# 检查nginx配置文件语法
/application/nginx/sbin/nginx -t
# 启动nginx
application/nginx/sbin/nginx 
netstat -lntup|grep nginx 
```
> * 创建测试文件
```
# web01
for dir in {www,bbs};do echo "`ifconfig eth1|awk -F "[ :]+" 'NR==2{print $4}'` $dir" >/application/nginx/html/$dir/index.html;done

# web02
for dir in {www,bbs};do echo "`ifconfig eth1|awk -F "[ :]+" 'NR==2{print $4}'` $dir" >/application/nginx/html/$dir/index.html;done
```

### 负载均衡简易实现

> * 实践服务器: lb01节点操作,仅用到主负载均衡节点
> * 负载均衡配置: 代理www.xxx.org服务,RS节点为web01和web02;lb01的nginx配置文件如下:
```
[root@lb01 conf]# cat nginx.conf
worker_processes  1;
events {
  worker_connections  1024;
}
http {
  include       mime.types;
  default_type  application/octet-stream;
  sendfile        on;
  keepalive_timeout  65;

  upstream www_server_pools {
      server 10.0.0.7:80 weight=1;
      server 10.0.0.8:80 weight=1;
  }

  server {
      listen       80;
      server_name  www.xxx.org;
      location / {
      proxy_pass http://www_server_pools;
      }
  }
}

[root@lb01 conf]# /application/nginx/sbin/nginx -t
[root@lb01 conf]# /application/nginx/sbin/nginx 
[root@lb01 conf]# netstat -lnutp|grep nginx
tcp        0      0 0.0.0.0:80                  0.0.0.0:*                   LISTEN      32519/nginx
```
> * 负载均衡测试结果: 轮询
```
# 配置host解析
[root@lb01 conf]# echo "10.0.0.5 www.xxx.org" >>/etc/hosts

# 模拟访问
[root@lb01 conf]# curl www.xxx.org
10.0.0.7 bbs
[root@lb01 conf]# curl www.xxx.org
10.0.0.8 bbs
[root@lb01 conf]# curl www.xxx.org
10.0.0.7 bbs
[root@lb01 conf]# curl www.xxx.org
10.0.0.8 bbs
```
> * 宕掉任意一个web节点,测试结果:
```
# 模拟web01宕机来测试结果,可以看到请求被分配到了web02上
[root@web01 conf]# /application/nginx/sbin/nginx -s stop 
[root@lb01 conf]# curl www.xxx.org
10.0.0.7 bbs
[root@lb01 conf]# curl www.xxx.org
10.0.0.7 bbs
[root@lb01 conf]# curl www.xxx.org
10.0.0.7 bbs
[root@lb01 conf]# curl www.xxx.org
10.0.0.7 bbs
```
> * 宕掉所有的web节点,测试结果:
```
# 模拟web01和web02宕机
[root@web01 conf]# /application/nginx/sbin/nginx -s stop
[root@web02 conf]# /application/nginx/sbin/nginx -s stop

# 可以看出所有web节点宕机,Nginx会向用户报502错误 
[root@lb01 conf]# curl -I www.xxx.org 
HTTP/1.1 502 Bad Gateway
Server: nginx/1.6.3
Date: Mon, 30 May 2016 23:14:32 GMT
Content-Type: text/html
Content-Length: 172
Connection: keep-alive
```
> * RS节点恢复正常时,返回结果正常

## Nginx负载均衡核心组件

### nginx upstream模块

#### upstream模块介绍

> * Nginx的负载均衡功能依赖于`ngx_http_upstream_module`模块,所支持的代理方式包括`proxy_pass`、`fastcgi_pass`、`memcached_pass`等,新版Nginx软件支持的方式有所增加;本次实验使用proxy_pass代理方式;
> * `ngx_http_upstream_module`模块允许Nginx定义一组或多组节点服务器组,使用时可以通过proxy_pass代理方式把网站的请求发送到事先定义好的对应Upstream组的名字上,具体写法为`proxy_pass http://www_server_pools`,其中www_server_pools就是一个UPstream节点服务器组名字;
> * ngx_http_upstream_module模块官方地址：http://nginx.org/en/docs/http/ngx_http_upstream_module.html

#### upstream模块语法

> * upstream简单配置案例如下
```
# upstream是关键字,www_server_pools为一个upstream集群组的名字,自定义,调用时使用这个名字
upstream www_server_poos {
    # server是关键字,使用域名或IP,端口不指定默认80,weight表示权重,数字越大被分配的请求越多
    server 10.0.0.7:80 weight=5;
    server 10.0.0.8:80 weight=10; # 注意结尾有分号 
}
```
> * upstream较为完整的配置案例如下
```
upstream bbs_server_tools {
    server 10.0.0.7;
    server 10.0.0.8:80 weight=1 max_fails=1 fail_timeout=10s;
    # 备份服务器,上面指定的服务器都不可访问时会启用,backup的用法和haproxy中用法一样
    server 10.0.0.9:80 weight=1 max_fails=2 fail_timeout=20s backup;
    server 10.0.0.10:80 weight=1 max_fails=2 fail_timeout=20s backup;   
}
```
> * 使用域名及socket的upstream配置案例如下
```
upstream backend {
    server backend1.example.com weight=5;
    # 域名加端口,转发到后端指定的端口上
    server backend2.example.com:8080 
    server unix:/var/run/docker.sock; # 指定socket文件
}
```
> * 如果是两台web服务器做高可用,常规方案就是需要Keepalived配合,这里使用的Nginx的backup参数通过负载均衡功能就能实现web服务器集群,对于企业应用来说,能做集群就不做高可用

#### upstream模块说明

> * upstream模块的内容应放于nginx.conf配置的http{}标签内,其默认调度节点算法是wrr(weighted round_robin,即权重轮询)
> * 下表为upstream模块内部server标签部分参数说明:

|upstream模块内参数|参数说明|
|:---|:---|
|server 10.0.0.8:80|负载均衡后的RS配置,可以是ip或域名,如果端口不写默认为80,高并发场景下,IP可以换成域名,通过DNS做负载均衡|
|weight=1|代表服务器的权重,默认值为1,权重数字越大表示 接受的请求比例越大|
|max_fails=1|Nginx尝试连接后端主机失败的次数,这个数字是配合proxy_next_upstream,fastcgi_next_upstream,memcached_next_upstream这三个参数自定义的状态码时,会将这个请求转发给正常的后端服务器,例如:404,502,503;max_fails的默认值为1;企业场景下建议2~3次,如京东1次,蓝汛10次,根据业务需求去配置|
|backup|热备配置(RS节点的高可用),当前面激活的RS都失败后会自动启用热备RS,这标志着这个服务器作为备份服务器,若主服务器全部宕机了,就会向它转发请求,注意:当负载调度算法为ip_hash时,后端服务器在负载均衡调度中的状态不能是weight和backup|
|fail_timeout=10s|在max_fials定义的失败次数后,距离下次检查的间隔时间,默认是10s;如果max_fails是5,它就检查5次,如果5次都是502,那么它就会根据fail_timeout的值,等待10s再去检查,还是只检查一次,如果持续502,在不重新加载nginx配置的情况下,每个10s都只检查一次,常规业务2~3秒比较合理,如京东3秒,蓝汛3秒,可以根据业务需求去配置;|
|down|这个标志着服务器永远不可用,这个参数可配合ip_hash使用|
> * 以下是一个较为完整upstream标签
```
upstream backend {
    server backend1.example.com weight=5;
    # 检查次数等于5的时候,5次连续检测失败后,间隔10s在重新检查
    server 127.0.0.1:8080 max_fails=5 fail_timeout=10s;
    server unix:/tmp/backend3;
    server backup1.example.com:8080 backup;    #热备机设置
}
```
> * 如果Nginx代理cache服务,可能需要使用hash算法,此时若宕机,可以通过设置down参数确保客户端用户按照当前的hash算法访问,配置如下
```
upstream backend {
    ip_hash;
    server backend1.example.com;
    server backend2.example.com;
    server backend3.example.com down; # 注意此处的down
    server backend4.example.com;
}
```
> * 下面是haproxy负载均衡器server标签的配置示例
```
# 开启对后端服务器的健康检测,通过GET /test/index.php来判断后端服务器的健康情况
server php_server1 10.12.25.68:80 cookie 1 check inter 2000 rise 3 fall 3 weight 2
server php_server1 10.12.25.72:80 cookie 2 check inter 2000 rise 3 fall 3 weight 1
server php_server1 10.12.25.79:80 cookie 3 check inter 1500 rise 3 fall 3 backup

# 上述命令说明:
 weight: 调节服务器的请求分配权重
 check: 开启对该服务器健康检查
 inter: 设置连续两次的健康检查间隔时间,默认毫秒,默认值为1200
 rise: 指定多少次连续成功的健康检查后,即可认定该服务器处于可用状态
 fall: 指定多少次不成功的健康检查后,即认为该服务器为宕机状态,默认值3
 maxconn: 指定可被发送到该服务器的最大并发连接数
```

#### upstream模块调度算法

> 调度算法一般分为两类,第一类为静态调度算法,第二类为动态调度算法;
> * 静态调度算法: 即负载均衡根据自身设定的规则进行分配,不需要考虑后端节点服务器的情况,例如:rr,wrr,ip_hash等都属于静态调度算法;
> * 动态调度算法: 即负载均衡器会根据后端节点的当前状态来决定是否分发请求,例如:连接数少的优先获得请求,响应时间短的优先获得请求,例如:least_conn,fair等都属于动态调度算法;

> 下面介绍一些常见的调度算法
> * rr轮询(默认调度算法,静态调度算法): 按客户端请求顺序把客户端的请求分配到不同的后端节点服务器,类似LVS的rr算法,如果后端节点服务器宕机(默认情况下nginx只检测80端口),宕机的服务器会被自动从节点服务器池中剔除,使客户端的用户访问不受影响,新的请求会分配给正常的服务器;
> * wrr(权重轮询,静态调度算法): 在rr轮询算法的基础上加上权重,即为权重轮询算法;当使用该算法时,权重和用户访问成正比,权重值越大,被转发的请求也就越多;可以根据服务器的配置和性能指定权重值大小,有效解决新旧服务器性能不均带来的请求分配问题,以下为一个使用权重轮询的例子:
```
# 后端服务器192.168.1.2的配置为: E55202 CPU,8GB 内存.
# 后端服务器192.168.1.3的配置为: Xeon(TM)2.80GHz2,4GB 内存
# 假设希望在有30个请求到达前端时,其中20个请求交给192.16.1.3处理,剩余10个请求交给192.168.1.2处理,就可做如下配置:
upstream wrr_lb {
    server 192.168.1.2 weight=1;
    server 192.168.1.3 weight=2;
 }
```
> * ip_hash(静态调度算法): 每个请求按客户端ip的hash结果分配,当新的请求到达时,先将其客户端IP通过哈希算法哈希出一个值,在随后的客户端请求中,客户IP的哈希值只要相同,就会被分配至同一台服务器;该调度算法可以解决动态网页的session共享问题,但有时会导致请求分配不均,即无法保证1:1的负载均衡,因为在国内大多数公司都是NAT上网模式,多个客户端会对应一个外部IP,所以这些客户端都会被分配到统一节点服务器上,从而导致请求分配不均;LVS负载均衡的-p参数,keepalived配置里的persistence_timeout 50参数都类似这里Nginx里的ip_hash参数,其功能均为解决动态网页的session共享问题;
```
# 以下是一个简单的ip_hash算法配置案例:
  upstream ip_hash_lb {
    ip_hash;
    server 192.168.1.2:80;
    server 192.168.1.3:8080;
  }

  upstream backend {
    ip_hash;
    server backupend1.example.com;
    server backupend2.example.com;
    server backupend3.example.com down;
    server backupend4.example.com;
  }
# 注意: 当负载调度算法为ip_hash时,后端服务器在负载均衡调度中的状态不能有weight和backup,即使有也不会生效;
```
> * fair(动态调度算法): 此算法会根据后端节点服务器的响应时间来分配请求,响应时短的优先分配;这是更加智能的调度算法,此种算法可以根据页面大小和加载时间长短智能地进行负载均衡,也就是根据后端服务器的响应时间来分配请求,响应时间短的优先分配;Nginx本身是不支持fair调度算法的;如果需要使用这种调度算法,需要下载Nginx的相关模块upstream_fair,示例如下:
```
upstream fair_lb {
    server 192.168.1.2;
    server 192.168.1.3;
    fair;
}
```
> * least_conn(动态调度算法): least_conn算法会根据后端节点的连接数来决定分配情况,哪个机器连接数少就分发;
> * url_hash算法: 和ip_hash类似,这里是根据访问URL的hash结果来分配请求的,让每个URL定向到同一个后端服务器,后端服务器为缓存服务器时效果显著;在upstream中加入hash语句,server语句中不能写入weight等其他参数,hash_method使用的是hash算法;url_hash按访问URL的hash结果来分配请求,使每个URL定向到同一个后端服务器,可以进一步提高后端缓存服务器的命中率;Nginx本身是不支持url_hash的,如果需要使用这种调度算法,必须安装Nginx的hash模块软件包;url_hash(web缓存节点)和ip_hash(会话保持)类似,配置如下:
```
upstream url_hash_lb {
    server squid1:3128;
    server squid2:3128;
    hash $request_uri;  # hash对象: 请求的URL
    hash_method crc32;  # hash算法: crc32
}
```
> * 一致性HASH算法: 一致性hash算法一般用于后端业务为缓存服务(squid,memcached)的场景,通过将用户请求的URI或指定字符串进行计算,然后调度到后端的服务器上,此后任何用户查找同一个URI或者指定字符串都会被调度到这一台服务器上,因此后端的每个节点缓存的内容都是不同的,一致性hash算法可以解决后端某个或几个节点宕机后,缓存的数据动荡最小,一致性hash算法知识比较复杂,这里仅仅给出配置示例:
```
http {
    upstream test{
        consistent_hash $request_uri;
        server 127.0.0.1:9001 id=1001 weight=3;
        server 127.0.0.1:9002 id=1002 weight=10;
        server 127.0.0.1:9003 id=1003 weight=20;
    }
}

# 虽然Nginx本身不支持一致性hash算法,但是Nginx的分之tengine支持
# 一致性hash模块使用简介: http://tengine.taobao.org/document_cn/http_upstream_consistent_hash_cn.html
```

### http_proxy_module模块

> * proxy_pass指令属于`ngx_http_proxy_module`模块,此模块可以将请求转发到另一台服务器,在实际的反向代理工作中,会通过location功能匹配指定的URI,然后把接受到的符合匹配URI的请求通过proxy_pass抛给定义好的upstream节点池;
> * 该指令的官方地址如下: http://nginx.org/en/docs/http/ngx_http_proxy_module.html#proxy_pass
> * 下面是proxy_pass的使用案例:
```
# 1、将匹配URI为name的请求抛给http://127.0.0.1/remote/
location /name/ {
    proxy_pass http://127.0.0.1/remote
}

# 2、将匹配URI为some/path的请求抛给http://127.0.0.1
location /some/path {
    proxy_pass http://127.0.0.1
}

# 3、将匹配URI为那么的请求应用指定的rewrite规则,然后抛给http://127.0.0.1.
location /name/ {
    rewrite /name/([^/]+) users?name=$1 break;
    proxy_pass http://127.0.0.1;
}
```
> * http proxy模块参数
 - nginx的代理功能是通过http_proxy模块来实现的,默认在安装nginx时已经安装了http proxy模块,因此可直接使用http proxy模块,下面详细解释模块中每个选项代表的含义:
 
|http proxy模块相关参数|参数说明|
|:---|:---|
|proxy_set_header|设置http请求header项传给后端服务器节点,例如:可实现让代理后端的服务器节点获取访问客户端用户的真实IP地址,而不是负载均衡器的地址|
|client_body_buffer_size|用于指定客户端请求body缓冲区大小|
|proxy_connect_timeout|表示反向代理与后端节点服务器连接的超时时间,即发起握手等候响应的超时时间|
|proxy_send_timeout|表示代理后端服务器的数据回传时间,即在规定时间内后端服务器必须传完所有数据,否则nginx将断开这个连接|
|proxy_read_timeout|设置nginx代理的后端服务器获取信息的时间,表示连接建立成功后,nginx等待后端服务器的响应时间,其实是nginx已经进入后端的排队之中等候处理的时间|
|proxy_buffer_size|设置缓冲区大小,默认该缓冲区大小等于指定proxy_buffers设置的大小|
|proxy_buffers|设置缓冲区的数量和大小,nginx从代理的后端服务器获取响应信息,会放置到缓冲区|
|proxy_busy_buffers_size|用于设置系统很忙时可以使用的proxy_buffers大小,官方推荐的大小为proxy_buffers*2|
|proxy_temp_file_write_size|指定proxy缓存临时文件的大小|

## Nginx负载均衡配置实践

### 基于域名的虚拟主机配置

> * nginx web服务器节点信息如下

|主机名|IP地址|角色说明|
|:---|:---|:---|
|web01|10.0.0.8|nginx web01服务器|
|web02|10.0.0.7|nginx web02服务器|
> * 配置基于域名的虚拟主机
```
# 两台节点一致
$ cat /application/nginx/conf/nginx.conf
worker_processes 1;
events {
    worker_connections 1024;
}
http {
    include mime.types;
    default_type application/octet-stream;
    sendfile on;
    keepalive_timeout 65;
    log_format main '$remote_addr - $remote_user [$time_local] "$request" '
                    '$status $body_bytes_sent "$http_referer" '
                '   "$http_user_agent" "$http_x_forwarded_for"';
    server {
        listen       80;
        server_name  bbs.xxx.org;
        location / {
            root   html/bbs;
            index  index.html index.htm;
        }
        access_log logs/access_bbs.log main;
    }

    server {
        listen       80;
        server_name  www.xxx.org;
        location / {
            root   html/www;
            index  index.html index.htm;
        }
        access_log logs/access_www.log main;
    }
}
```
> * 创建站点目录及对应的测试文件
```
[root@web01 ~]# for n in {bbs,www};do echo "10.0.0.8 $n.xxx.org" >/application/nginx/html/$n/index.html;done
[root@web02 ~]# for n in {bbs,www};do echo "10.0.0.7 $n.xxx.org" >/application/nginx/html/$n/index.html;done

[root@web01 ~]# cat /application/nginx/html/{bbs,www}/index.html
10.0.0.8 bbs.xxx.org
10.0.0.8 www.xxx.org

[root@web02 conf]# cat /application/nginx/html/{bbs,www}/index.html
10.0.0.7 bbs.xxx.org
10.0.0.7 www.xxx.org

$ /application/nginx/sbin/nginx -t #检查语法
$ /application/nginx/sbin/nginx -s reload #重新加载

[root@web01 ~]# curl {www,bbs}.xxx.org
10.0.0.8 www.xxx.org
10.0.0.8 bbs.xxx.org

[root@web02 conf]# curl {www,bbs}.xxx.org
10.0.0.7 www.xxx.org
10.0.0.7 bbs.xxx.org
```

### nginx负载均衡反向代理配置

> * 利用upstream定义一组www服务器池
```
# 先定义一个名字为www_server_pools的服务器池,里面有2台web服务器
upstream www_server_pools {
    server 10.0.0.7:80 weight=1;
    server 10.0.0.8:80 weight=1;
}
```
> * 配置www服务的虚拟主机server负载代理
```
server {
    listen 80;
    server_name www.xxx.org;
    location / {
        proxy_pass http://www_server_pools;
    }
}
```
> * 实际的配置内容如下
```
[root@lb01 conf]# cat nginx.conf
worker_processes 1;
events {
    worker_connections 1024;
}
http {
    include mime.types;
    default_type application/octet-stream;
    sendfile on;
    keepalive_timeout 65;

    upstream www_server_pools {
        server 10.0.0.7:80 weight=1;
        server 10.0.0.8:80 weight=1;
    }

    server {
        listen 80;
        server_name www.xxx.org;
        location / {
            proxy_pass http://www_server_pools;
        }
    }
}
```
> * 配置hosts文件解析,然后重新加载服务,访问测试:
```
root@lb01 conf]# tail -1 /etc/hosts
10.0.0.5 www.xxx.org bbs.xxx.org

[root@lb01 conf]# /application/nginx/sbin/nginx -s reload

[root@lb01 conf]# curl www.xxx.org
10.0.0.7 bbs.xxx.org

[root@lb01 conf]# curl www.xxx.org
10.0.0.8 bbs.xxx.org

# 提示: 
  默认调度算法是weight round-robin,即权重轮询算法;
  upstream仅仅是定义服务器池,并不会处理用户的请求,必须要有其他方式将请求转发给这个服务器池才行;
  虽然定义的是www服务器池,但这个服务器池也可以作为BBS等业务的服务器池,因为节点服务器的虚拟主机都是根据访问的主机头字段区分的;
```

### 反向代理多虚拟主机节点服务器

> * 上面的代理结果不对,究其原因是当用户访问域名时确实是携带了www.xxx.org主机头请求Nginx反向代理服务器,但是反向代理服务器向下面的节点重新发起请求时,默认并没有在请求头里告诉节点服务器要找哪台虚拟主机,所以web节点服务器接收到请求后发现没有主机头信息,因此就把节点服务器的第一个虚拟主机发给了反向代理(而节点上的第一个虚拟主机放置的是故意这样放置的BBS);
> * 解决这个问题的办法,就是当反向代理向后重新发起请求时,要携带主机头信息,以明确告诉节点服务器要找哪个虚拟主机;具体的配置很简单,就是在Nginx代理www服务虚拟主机配置里增加如下一行配置即可:
```
proxy_set_header Host $host;
```
> * 在代理向后端服务器发送的http请求中加入host字段信息,用于当后端服务器配置有多个虚拟主机时,可以识别代理的是那个虚拟主机,这是节点服务器多虚拟主机时的关键配置;整个Nginx代理配置为:
```
[root@lb01 conf]# cat nginx.conf
worker_processes 1;
events {
    worker_connections 1024;
}
http {
    include mime.types;
    default_type application/octet-stream;
    sendfile on;
    keepalive_timeout 65;

    upstream www_server_pools {
        server 10.0.0.7:80 weight=1;
        server 10.0.0.8:80 weight=1;
    }

    server {
        listen       80;
  		server_name  www.xxx.org;
  		location / {
  			proxy_pass http://www_server_pools;
  			proxy_set_header Host  $host; 
 		}
    }
}

# 修改后的测试结果如下:
[root@lb01 conf]# /application/nginx/sbin/nginx -s reload
[root@lb01 conf]# curl www.xxx.org
10.0.0.7 www.xxx.org
[root@lb01 conf]# curl www.xxx.org
10.0.0.8 www.xxx.org
```

### 反向代理的节点服务器记录用户IP

> * web01节点服务器对应的www虚拟主机的访问日志的第一个字段并不是客户端的IP,而是反向代理服务器本身的IP,最后一个字段也是一个“-”,解决方法如下:
```
# 在反向代理请求后端节点服务器的请求头中增加获取的客户端IP的字段信息,然后后端节点可以通过程序或者相关的配置接收X-Forwarded-For传来的用户真实IP的信息;
proxy_set_header X-Forwarded-For $remote_addr;
```

> * 在代理向后端服务器发送的http请求头中加入X-Forwarded-For字段信息,用于后端服务器程序、日志等接收记录真实用户的IP,而不是代理服务器的IP;
> * 特别注意,虽然反向代理这块已经配好了,但是节点服务器(web服务器)需要的访问日志如果要记录用户的真实IP,还必须进行日志格式配置,这样才能把代理传过来的X-Forwarded-For头信息记录下来,如果希望第一行显示,可以替换到第一行的$remote_addr变量;具体配置如下:
```
# Nginx配置日志:
log_format main '$remote_addr - $remote_user [$time_local] "$request" '
                '$status $body_bytes_sent "$http_referer" '
                '"$http_user_agent" "$http_x_forwarded_for"';

# Apache配置日志:
LogFormat "\"%{X-Forwarded-For}i\" %l %u %t \"%r\" %>s %b \"%{Referer}i\" \"%{User-Agent}i\"" common
```
> * 解决这个问题的整个Nginx代理配置为:
```
[root@lb01 conf]# cat nginx.conf
worker_processes 1;
events {
	worker_connections 1024;
}
http {
    include mime.types;
    default_type application/octet-stream;
    sendfile on;
    keepalive_timeout 65;

    upstream www_server_pools {
        server 10.0.0.7:80 weight=1;
        server 10.0.0.8:80 weight=1;
    }

    server {
        listen 80;
        server_name www.etiantain.org;
        location / {
            proxy_pass http://www_server_pools;
            proxy_set_header Host $host;
            proxy_set_header X-Forwarded-For $remote_addr;
        }
    }
}
```

### nginx反向代理相关参数说明

|nginx反向代理重要参数|解释说明|
|:---|:---|
|proxy_pass http://www_server_pools|通过proxy_pass功能把用户的请求转向到反向代理定义的upstream|
|proxy_set_header Host \$host|在代理向后端服务器发送的http请求头中加入host字段信息,用于当后端服务器配置有多个虚拟主机时,可以识别代理的是哪个虚拟主机;这是节点服务器多虚拟主机的关键配置|
|proxy_set_header X-Forwarded-For \$remote_addr|在代理向后端服务器发送的http请求头中加入X-Forwarded_For字段信息,用于后端服务器程序,日志等接受记录真实用户的IP,而不是代理服务器的IP;这是反向代理时,节点服务器获取用户真实IP的必要功能配置|
> * 除了具有多虚拟主机代理以及节点服务器记录真实用户IP的功能外,Nginx软件还提供了相当多的作为反向代理和后端节点服务器对话的相关控制参数;由于参数众多,最好把这些参数放到一个配置文件里,然后用include方式包含到虚拟主机配置里,效果如下:
```
[root@lb01 conf]# cat nginx.conf
worker_processes 1;
error_log logs/error.log;
events {
  worker_connections 1024;
 }
http {
    include       mime.types;
    default_type  application/octet-stream;
    sendfile        on;
    keepalive_timeout  65;
    upstream www_server_pools {
        server 10.0.0.8:80  weight=1;
        server 10.0.0.7:80  weight=2;
    }
		
    server {
        listen      10.0.0.3:80;
        server_name  www.xxx.org;
        location / {
            proxy_pass http://www_server_pools;
            include proxy.conf; #这就是包含的配置，具体内容如下
        }
    }
}
```
> * 可以把参数写成一个文件,使用include包含,看起来更简洁规范:
```
[root@lb01 conf]# cat proxy.conf
proxy_set_header Host $host;
proxy_set_header X-Forwarded-For $remote_addr;
proxy_connect_timeout 60;
proxy_send_timeout 60;
proxy_read_timeout 60;
proxy_buffer_size 4k;
proxy_buffers 4 32k;
proxy_busy_buffers_size 64k;
proxy_temp_file_write_size 64k;
```

## 根据URL中的目录地址实现代理转发

> * 通过Nginx实现动静分离,即通过Nginx反向代理配置规则实现让动态资源和静态资源及其他业务分别由不同的服务器解析,以解决网站性能,安全,用户体验等重要问题;
![image_1d3ssbs0816qdga81pt9681h1a9.png-117.8kB][7]
> * 当用户请求www.xxx.org/upload/xx地址时,实现由upload上传服务器池处理请求;
```
# upload_pools为上传服务器池,有一个服务器,地址为10.0.0.8端口为80
upstream upload_pools {
    server 10.0.0.8:80 weight=1;
}
```
> * 当用户请求www.xxx.org/static/xx地址时,实现由静态服务器池处理请求;
```
static_pools为静态服务器池,有一个服务器,地址为10.0.0.7端口为80
upstream static_pools {
    server 10.0.0.7:80 weight=2;
}
```
> * 除此之外,对于其他访问请求,全部由默认的动态服务器池处理请求;
```
# default_pools为默认的服务器池,即动态服务器池,地址为10.0.0.7端口为8080
upstream default_pools {
    server 10.0.0.7:8080 weight=2;
}
```
> * 下面利用location或者if语句把不同的URI(路径)请求,分给不同的服务器池处理,具体配置如下:
```
# 方案1: 以location方案实现,
# 将符合static的请求交给静态服务器池static_pools,配置如下:
location /static/ {
    proxy_pass http://static_pools;
    include proxy.conf;
}

# 将符合upload的请求交给上传服务器池upload_pools，配置如下：
location /upload/ {
    proxy_pass http://upload_pools;
include proxy.conf;
}

# 不符合上述规则的请求，默认全部交给动态服务器池default_pools，配置如下：
location / {
    proxy_pass http://default_pools;
include proxy.conf;
}

# 方案2: 以if语句实现
if ($request_uri ~* "^/static/(.*)$")
{
    proxy_pass http://static_pools/$1;
}
if ($request_uri ~* "^/upload/(.*)$")
{
    proxy_pass http://upload_pools/$1;
}
location / {
    proxy_pass http://default_pools;
    include proxy.conf;
}
```
> * nginx反向代理的实际配置文件如下:
```
[root@lb01 conf]# cat nginx.conf
worker_processes 1;
events {
    worker_connections 1024;
}
http {
    include mime.types;
    default_type application/octet-stream;
    sendfile on;
    keepalive_timeout 65;

    upstream static_pools {
        server 10.0.0.7:80 weight=1;
    }
    upstream upload_pools {
        server 10.0.0.8:80 weight=1;
    }
    upstream default_pools {
        server 10.0.0.7:8080 weight=1;
    }

    server {
  		listen       80;
  		server_name  www.xxx.org;
  		location / {
  			proxy_pass http://default_pools;
  			include proxy.conf;
  	    }
  		location /static/ {
  			proxy_pass http://static_pools;
  			include proxy.conf;
 		}
 	 	location /upload/ {
 			 proxy_pass http://upload_pools;
  			include proxy.conf;
 		}
	}
}
```
> * web01配置文件如下:
```
[root@web01 ~]# cat /application/nginx/conf/nginx.conf
worker_processes 1;
events {
	worker_connections 1024;
}
http {
    include mime.types;
    default_type application/octet-stream;
    sendfile on;
    keepalive_timeout 65;
    log_format main '$remote_addr - $remote_user [$time_local] "$request" '
		            '$status $body_bytes_sent "$http_referer" '
					'"$http_user_agent" "$http_x_forwarded_for"';
    server {
        listen       80;
        server_name  bbs.xxx.org;
        location / {
         	root   html/bbs;
      		index  index.html index.htm;
  	    }
  		access_log logs/access_bbs.log main;
    }

    server {
  		listen       80;
  		server_name  www.xxx.org;
  		location / {
      		root   html/www;
      		index  index.html index.htm;
 		}
 		access_log logs/access_www.log main;
  	}
}

[root@web01 ~]# tree /application/nginx/html/www/
/application/nginx/html/www/
├── index.html
└── upload
		└── index.html
[root@web01 ~]# cat /application/nginx/html/www/upload/index.html
upload page
```
> * web02配置文件如下:
```
[root@web02 ~]# cat /application/nginx/conf/nginx.conf
worker_processes 1;
events {
    worker_connections 1024;
}
http {
    include mime.types;
    default_type application/octet-stream;
    sendfile on;
    keepalive_timeout 65;
    log_format main '$remote_addr - $remote_user [$time_local] "$request" '
		            '$status $body_bytes_sent "$http_referer" '
		            '"$http_user_agent" "$http_x_forwarded_for"';

    server {
        listen       80;
        server_name  bbs.xxx.org;
        location / {
            root   html/bbs;
            index  index.html index.htm;
		}
		access_log logs/access_bbs.log main;
	}

    server {	
        listen       10.0.0.7:80;
        server_name  www.xxx.org;
        location / {
            root   html/www;
            index  index.html index.htm;
		}
		access_log logs/access_www.log main;
 	}

    server {
		listen       10.0.0.7:8080;
        server_name  www.xxx.org;
        location / {
            root   html/www8080;
            index  index.html index.htm;
  	    }
  		access_log logs/access_www.log main;
    }
}


[root@web02 ~]# tree /application/nginx/html/www
/application/nginx/html/www
├── index.html
└── static		
   └── index.html
   
[root@web02 ~]# cat /application/nginx/html/www/static/index.html
static page

[root@web02 ~]# tree /application/nginx/html/www8080/
/application/nginx/html/www8080/
└── index.html

[root@web02 ~]# cat /application/nginx/html/www8080/index.html
default page
```

> * 测试结果如下
```
[root@web01 ~]# tail -1 /etc/hosts
10.0.0.5 www.xxx.org bbs.xxx.org
[root@lb02 ~]# curl www.xxx.org
default page
[root@lb02 ~]# curl www.xxx.org/static/
static page
[root@lb02 ~]# curl www.xxx.org/upload/
upload page
```
> * 根据URL目录地址转发的应用场景
 - 根据HTTP的URL进行转发的应用情况,被称为第7层(应用层)的负载均衡,而LVS的负载均衡一般用于TCP等的转发,因此被称为第4层(传输层)的负载均衡;
 - 在企业中,有时希望只用一个域名对外提供服务,不希望使用多个域名对应同一个产品业务,此时就需要在代理服务器上通过配置规则,使得匹配不同规则的请求会交给不同的服务器池处理;这类业务有:
   1、业务的域名没有拆分或者不希望拆分,但希望实现动静分离、多业务分离;
   2、不同的客户端设备(例如：手机和pc端)使用同一个域名访问同一个业务网站就需要根据规则将不同设备的用户请求交给后端不同的服务器处理,以便得到最佳用户体验;

## 根据user_agent转发实践

> * 企业中,为了让不同的客户端设备访问有更好的体验,需要在后端架设不同服务器来满足不同的客户端访问,例如:移动客户端访问网站,就需要部署单独的移动服务器及程序,体验才能更好,而且移动端还分为苹果、安卓、iPad等,在传统情况下,一般用下面的办法解决这个问题;
> * 1、常规4层负载均衡架构下
 - 可以使用不同的域名来实现这个需求,例如人为分配好让移动端访问wap.xxx.org, pc客户端用户访问www.xxx.org,通过不同域名来引导用户到指定的后端服务器;此解决方案最大问题就是不同客户度的用户要记住对应的域名!这样一来就会导致用户体验不是很好;
> * 2、第7层负责均衡解决方案	
 - 在第7层负责均衡架构下,就可以不需要人为拆分域名了,对外只需要用一个域名,例如www.xxx.org,然后通过获取用户请求中的设备信息($http_user_agent获取)根据这些信息转给后端合适的服务器处理,这个方案最大好处就是不需要让用户记忆多个域名了,用户只需要记住主网站地址www.xxx.org,剩下的由网站服务器处理,这样的思路大大提升了用户访问体验,这是当前企业网站非常常用的解决方案;
> * 根据客户端设备(user_agent)转发请求实践
```
# 这里还是使用static_pools、upload_pools、作为本次实验的后端服务器池,下面先根据计算机客户端浏览器的不同设置对应的匹配规则

location / {
   # 如果请求的浏览器为微软IE浏览器(MSIE),则让请求有static_pools池处理
  if ($http_user_agent  ~*  "MSIE")
    {
      proxy_pass  http://static_pools; 
    }

# 如果请求的浏览器为谷歌浏览器(Chrome),则让请求有upload_pools池处理
  if ($http_user_agent  ~*  "Chrome")
    {
      proxy_pass  http://upload_pools;
    }

# 剩余默认
  proxy_pass  http://default_pools;
  include proxy.conf;
}
```
> * 实际中完整的配置文件内容如下
```
[root@lb01 conf]# cat nginx.conf
worker_processes 1;
error_log logs/error.log;
events {
	worker_connections 1024;
}
http {
	include mime.types;
    default_type application/octet-stream;
    sendfile on;
    keepalive_timeout 65;

    upstream static_pools {
   		server 10.0.0.7:80  weight=2;
    }

    upstream upload_pools {
   		server 10.0.0.8:80  weight=1;
    }
		
    upstream default_pools {
        server 10.0.0.7:8080  weight=2;
    }
		
    server {
        listen 80;
        server_name www.xxx.org;
 		location / {
 		 	if ($http_user_agent  ~*  "MSIE")
            {
                proxy_pass  http://static_pools; 
            }

  			if ($http_user_agent  ~*  "Chrome")
            {
                proxy_pass  http://upload_pools;
            }

  			proxy_pass  http://default_pools;
  		    include proxy.conf;
	 	}
	}
}
```
> * 除了针对浏览器外,上述"$http_user_agent"变量也可以针对移动端,比如安卓、苹果、Ipad设备进行匹配,去请求指定的服务器,具体配置如下:
```
location / {
  if ($http_user_agent ~* "android")
    {
      proxy_pass  http://android_pools; #<==这是android服务器池，需要提前定义upstream。
    }

  if ($http_user_agent ~* "iphone")
    {
      proxy_pass  http://iphone_pools; #<==这是iphone服务器池，需要提前定义upstream。
     }

  proxy_pass  http://pc_pools;
  include proxy.conf;
 }

# 完整配置文件为
[root@lb01 conf]# cat nginx.conf
worker_processes 1;
events {
	worker_connections 1024;
}
http {
    include mime.types;
    default_type application/octet-stream;
    sendfile on;
    keepalive_timeout 65;

    upstream android_pools {
        server 10.0.0.7:80  weight=1; ##apache
    }

    upstream iphone_pools {
        server 10.0.0.7:8080  weight=1; ##apache
    }
		
    upstream pc_pools {
        server 10.0.0.8:80  weight=1; #nginx
    }
		
    server {
        listen       80;
        server_name   blog.xxx.org;
        location / {
            if ($http_user_agent ~* "android")
            {
                proxy_pass  http://android_pools; 
            }

            if ($http_user_agent ~* "iphone")
            {
                proxy_pass  http://iphone_pools;
            }

            proxy_pass  http://pc_pools;
            include proxy.conf;
        }  
    }
}
```

## 根据文件扩展名实现代理转发

> * 除了根据URI路径及user_agent转发外,还可以实现根据文件扩展名进行转发
> * 以下是相关server配置
```
# 先看看location方法的匹配规则,如下:
location ~ .*.(gif|jpg|jpeg|png|bmp|swf|css|js)$ {
    proxy_pass http://static_pools;
    include proxy.conf
}

# 下面是if语句方法的匹配规则
if ($request_uri ~ "..(php|php5)$")
{
    proxy_pass http://php_server_pools;
}

if ($request_uri ~ "..(jsp|jsp|do|do)$")
{
    proxy_pass http://java_server_pools;
}

# 根据扩展名转发的应用场景: 可根据扩展名实现资源动静分离访问,如图片、视频等请求静态服务器池,PHP、JSP等请求动态服务器池;示例代码如下:
location ~ ..(gif|jpg|jpeg|png|bmp|swf|css|js)$ {
    proxy_pass http://static_pools;
    include proxy.conf;
}

location ~ ..(php|php3|php5)$ {
    proxy_pass http://dynamic_pools;
    include proxy.conf;
}

# 在开发无法通过程序实现动静分离的时候,运维可以根据资源实体进行动静分离,而不依赖于开发,具体实现策略先把后端服务器分成不同的组;注意每组服务器的程序都是相同的,因为开发没有把程序拆开,分组后,在前端代理服务器上通过讲解过的路径、扩展名进行规则匹配,从而实现动静分离;
```

## Nginx负载均衡检测节点状态

> * 淘宝技术团队开发了一个Tengine(Nginx的分支)模块`nginx_upstream_check_module`,用于提供主动式后端服务器健康检查;通过它可以检测后端realserver的健康状态,如果后端realserver不可用,则所有的请求就不会转发到该节点上;
> * 安装nginx_upstream_check_module模块:
```
# 查看nginx的编译参数
root@lb01 conf]# /application/nginx/sbin/nginx -V 
nginx version: nginx/1.6.3
built by gcc 4.4.7 20120313 (Red Hat 4.4.7-16) (GCC)
TLS SNI support enabled
configure arguments: --prefix=/application/nginx-1.6.3 --user=www --group=www --with-http_ssl_module --with-http_stub_status_module

# 下载补丁包
[root@lb01 conf]# cd /server/tools/
[root@lb01 tools]# wget https://codeload.github.com/yaoweibin/nginx_upstream_check_module/zip/master 
[root@lb01 tools]# unzip master

# 对源程序打补丁
[root@lb01 nginx-1.6.3]# cd nginx-1.6.3
[root@lb01 nginx-1.6.3]# patch -p1 < ../nginx_upstream_check_module-master/check_1.5.12+.patch 

# 重新编译参数,加上--add-module=../nginx_upstream_check_module-master/这个模块
[root@lb01 nginx-1.6.3]# ./configure \
--prefix=/application/nginx-1.6.3 \
--user=www \
--group=www \
--with-http_ssl_module \
--with-http_stub_status_module \
--add-module=../nginx_upstream_check_module-master/ 

# 将编译的参数写入内核,如果是新装的Nginx则继续执行make install,如果给已经安装的Nginx系统打监控补丁就不用执行make install
[root@lb01 nginx-1.6.3]# make 

# 备份之前的启动脚本
[root@lb01 nginx-1.6.3]# mv /application/nginx/sbin/nginx{,.ori} 

# 将打过补丁的nginx二进制程序复制到/application/nginx/sbin/目录下
[root@lb01 nginx-1.6.3]# cp ./objs/nginx /application/nginx/sbin/ 

# 查看重新编译参数
[root@lb01 nginx-1.6.3]# /application/nginx/sbin/nginx -V 
nginx version: nginx/1.6.3
built by gcc 4.4.7 20120313 (Red Hat 4.4.7-16) (GCC)
TLS SNI support enabled
configure arguments: --prefix=/application/nginx-1.6.3 --user=www --group=www --with-http_ssl_module --with-http_stub_status_module --add-module=../nginx_upstream_check_module-master/
```
> * 配置Nginx健康检查,如下:
```
[root@lb01 conf]# cat nginx.conf
worker_processes 1;
events {
    worker_connections 1024;
}
http {
    include mime.types;
    default_type application/octet-stream;
    sendfile on;
    keepalive_timeout 65;

    upstream static_pools {
        server 10.0.0.7:80 weight=1;
        check interval=3000 rise=2 fall=5 timeout=1000 type=http;#健康检查参数
    }

    upstream upload_pools {
        server 10.0.0.8:80 weight=1;
        check interval=3000 rise=2 fall=5 timeout=1000 type=http;#健康检查参数
    }

    upstream default_pools {
        server 10.0.0.7:8080 weight=1;
        check interval=3000 rise=2 fall=5 timeout=1000 type=http;#健康检查参数
		}

    server {
        listen 80;
        server_name www.xxx.org;
        location / {
            proxy_pass http://default_pools;
            include proxy.conf;
        }

        location /static/ {
            proxy_pass http://static_pools;
            include proxy.conf;
        }

        location /upload/ {
            proxy_pass http://upload_pools;
            include proxy.conf;
        }
			
        location /status { #添加健康检查标签
            check_status;
            access_log off;
	    }
    }
}

## 注意此处必须重启Nginx,不能重新加载 
[root@lb01 conf]# /application/nginx/sbin/nginx -s stop
[root@lb01 conf]# /application/nginx/sbin/nginx 

# check interval=3000 rise=2 fall=5 timeout=1000 type=http;
上面配置的意思是,对static_pools这个负载均衡条目中的所有节点,每隔3秒检测一次,请求2次正常则标记realserver状态为up,如果检测5次都失败,则标记realserver的状态为down,超时时间为1秒,检查的协议是http
```

	
## proxy_next_upstream参数补充

> * 当Nginx接收到后端服务器返回proxy_next_upstream参数定义的状态码时,会将这个请求转发给正常工作的后端服务器,例如500、502、503、504,此参数可以提升用户的访问体验,具体配置如下:
```
server {
    listen 80;
    server_name www.xxx.org;
    location / {
        proxy_pass http://static_pools;
        proxy_next_upstream error timeout invalid_header http_500 http_502 http_503 http_504;
    }
    include proxy.conf;
}
```
 

  [1]: http://static.zybuluo.com/yujianfeng/j7bfay8etcwi0t3j2ny7woyf/image_1d3qgkg071c4e3t41g8a1u8680sp.png
  [2]: http://static.zybuluo.com/yujianfeng/eyzvslypx0juwkxgsn7e2idz/image_1d3qh208qq0n5eo13ri6rbtuv1j.png
  [3]: http://static.zybuluo.com/yujianfeng/cgpa8lfjc9ueqnj8vn68gja8/image_1d3qhaljstoo1k7fgbklu1137f20.png
  [4]: http://static.zybuluo.com/yujianfeng/z13ijr3n5t81imuny1kudsd0/image_1d3qhn02m1nbhvr32rmku0vt62t.png
  [5]: http://static.zybuluo.com/yujianfeng/vg8raew4uy04f4e7hq71h2su/image_1d3qhp4ab1spb17dd15t91kjcokf3a.png
  [6]: http://static.zybuluo.com/yujianfeng/vj5q67v1h02x321piqprpo33/image_1d3qie5rl13tp11405ka13tvili3n.png
  [7]: http://static.zybuluo.com/yujianfeng/fkbunsgcbp43qffri44o774v/image_1d3ssbs0816qdga81pt9681h1a9.png