标签： 反向代理

# HAPROXY

---

[toc]

## haproxy基础概述

> 架构体系中cache: 存储形式为key-value键值对,算法基于`取模法`扩展,`一致性hash算法/consistent hashing`;

> LB: Load Balance(负载均衡),包括4层和7层两种
> * tcp: 代表产品如lvs,haproxy,nginx
> * http: haproxy,nginx,ats,apache

> proxy: 代理http,分为正向代理和反向代理两种;
> * iptables的SNAT规则做服务器的网关代理;
> * 代理作用: web缓存(加速),反向代理,内容路由(根据流量及内容类型等将请求转发至特定服务器),在代理服务器上添加Via;
> * 缓存的作用: 减少冗余内容传输,节省带宽,缓解网络瓶颈;降低了对原始服务器的请求压力;降低了传输延迟;
> * haproxy只是http协议的反向代理,不提供缓存功能,但额外支持对tcp层基于tcp通信的应用做负载均衡;

### haproxy简介

> HAProxy提供高可用性、负载均衡以及基于TCP和HTTP应用的代理,支持虚拟主机,它是免费、快速并且可靠的一种解决方案;

> HAProxy特别适用于那些负载特大的web站点,这些站点通常又需要会话保持或七层处理;

> HAProxy运行在时下的硬件上,完全可以支持数以万计的并发连接并且它的运行模式使得它可以很简单安全的整合进您当前的架构中, 同时可以保护你的web服务器不被暴露到网络上;

> HAProxy实现了一种事件驱动,单一进程模型,此模型支持非常大的并发连接数;多进程或多线程模型受内存限制,系统调度器限制以及无处不在的锁限制,很少能处理数千并发连接;事件驱动模型因为在有更好的资源和事件管理的用户端(UserSpace)实现所有这些任务,所以没有这些问题;此模型的弊端是在多核系统上,这些程序通常扩展性较差这就是为什么他们必须进行优化以使每个CPU时间片(Cycle)做更多的工作;

> HAProxy内部的数据结构: 弹性二叉树;
> * 评判数据结构复杂度: `Big O`, `红黑树: O(logN)`;

> HAProxy版本特性
> * 1.4提供较好的弹性: 衍生于1.2版本,并提供了额外的新特性,其中大多数是期待已久的;
```
客户端侧的长连接(client-side keep-alive)
TCP加速(TCP speedups)
响应池(response buffering)
RDP(Remote Desktop Protocol/3389)协议
基于源的粘性(source-based stickiness)
更好的统计数据接口(a much better stats interfaces)
更详细的健康状态检测机制(more verbose health checks)
基于流量的健康评估机制(traffic-based health)
支持HTTP认证
服务器管理命令行接口(server management from the CLI)
基于ACL的持久性(ACL-based persistence)
日志分析器
```
> * 1.3内容交换和超强负载: 衍生于1.2版本,并提供了额外的新特性;
```
内容交换(content switching): 基于任何请求标准挑选服务器池;
ACL: 编写内容交换规则;
负载均衡算法(load-balancing algorithms): 更多的算法支持;
内容探测(content inspection): 阻止非授权协议;
透明代理(transparent proxy): 在Linux系统上允许使用客户端IP直接连入服务器；
内核TCP拼接(kernel TCP splicing): 无copy方式在客户端和服务端之间转发数据以实现数G级别的数据速率；
分层设计(layered design): 分别实现套接字、TCP、HTTP处理以提供更好的健壮性、更快的处理机制及便捷的演进能力;
快速、公平调度器(fast and fair scheduler): 为某些任务指定优先级可实现理好的QoS;
会话速率限制(session rate limiting): 适用于托管环境;
```

> 支持的平台及OS
> * x86、x86_64、Alpha、SPARC、MIPS及PARISC平台上的Linux 2.4
> * x86、x86_64、ARM (ixp425)及PPC64平台上的Linux2.6；
> * UltraSPARC 2和3上的Sloaris 8/9；
> * Opteron和UltraSPARC平台上的Solaris 10；
> * x86平台上的FreeBSD 4.1-8；
> * i386, amd64, macppc, alpha, sparc64和VAX平台上的OpenBSD 3.1-current；

> 若要获得最高性能,需要在Linux 2.6或打了epoll补丁的Linux 2.4上运行haproxy 1.2.5以上的版本;
> * haproxy 1.1默认使用的polling系统为select(),其处理的文件数达数千个时性能便会急剧下降;
> * 1.2和1.3版本默认的为poll(),在有些操作系统上可会也会有性能方面的问题,但在Solaris上表现相当不错;
> * HAProxy 1.3在Linux 2.6及打了epoll补丁的Linux 2.4上默认使用epoll,在FreeBSD上使用kqueue,这两种机制在任何负载上都能提供恒定的性能表现;

> 在较新版本的Linux 2.6(>=2.6.27.19)上,HAProxy还能够使用splice()系统调用在接口间无复制地转发任何数据,这甚至可以达到10Gbps的性能;

> 基于以上事实,在x86或x86_64平台上,要获取最好性能的负载均衡器,建议按顺序考虑以下方案
> * Linux 2.6.32及之后版本上运行HAProxy 1.4;
> * 打了epoll补丁的Linux 2.4上运行HAProxy 1.4;
> * FreeBSD上运行HAProxy 1.4;
> * Solaris 10上运行HAProxy 1.4;

> haproxy性能:
> * HAProxy借助于OS上几种常见的技术来实现性能的最大化;
> * 单进程、事件驱动模型显著降低了上下文切换的开销及内存占用;
> * O(1)事件检查器(event checker)允许其在高并发连接中对任何连接的任何事件实现即时探测;
> * 在任何可用的情况下,单缓冲(single buffering)机制能以不复制任何数据的方式完成读写操作,这会节约大量的CPU时钟周期及内存带宽;
> * 借助于Linux 2.6(>= 2.6.27.19)上的splice()系统调用,HAProxy可以实现零复制转发(Zero-copy forwarding),在Linux 3.5及以上的OS中还可以实现零复制启动(zero-starting);
> * 内存分配器在固定大小的内存池中可实现即时内存分配,这能够显著减少创建一个会话的时长;
> * 树型存储: 侧重于使用作者多年前开发的弹性二叉树,实现了以O(log(N))的低开销来保持计时器命令、保持运行队列命令及管理轮询及最少连接队列;
> * 优化的HTTP首部分析: 优化的首部分析功能避免了在HTTP首部分析过程中重读任何内存区域;
> * 精心地降低了昂贵的系统调用,大部分工作都在用户空间完成,如时间读取、缓冲聚合及文件描述符的启用和禁用等;
> * 所有的这些细微之处的优化实现了在中等规模负载之上依然有着相当低的CPU负载,甚至于在非常高的负载场景中,5%的用户空间占用率和95%的系统空间占用率也是非常普遍的现象,这意味着HAProxy进程消耗比系统空间消耗低20倍以上;因此对OS进行性能调优是非常重要的;即使用户空间的占用率提高一倍,其CPU占用率也仅为10%,这也解释了为何7层处理对性能影响有限这一现象;由此,在高端系统上HAProxy的7层性能可轻易超过硬件负载均衡设备;

> 在生产环境中,在7层处理上使用HAProxy作为昂贵的高端硬件负载均衡设备故障时的紧急解决方案也时常可见;
> * 硬件负载均衡设备在“报文”级别处理请求,这在支持跨报文请求(request across multiple packets)有着较高的难度,并且它们不缓冲任何数据,因此有着较长的响应时间;
> * 对应的软件负载均衡设备使用TCP缓冲,可建立极长的请求,且有着较大的响应时间;

> 可以从三个因素来评估负载均衡器的性能:
> * 会话率
> * 会话并发能力
> * 数据率

## haproxy配置选项

> 配置文件格式
> * haproy的配置处理3类主要参数来源
```
# 最优先处理命令行参数
# "global"配置段,用于设定全局配置参数
# proxy相关配置段,如"defaults","listen","front","backend"
```
> 时间格式
> * 一些包含了值得参数表示时间,如超时时长,这些值一般以毫秒为单位,但也可以使用其它的时间单位后缀
```
us: 微秒(microseconds),即1/1000000秒;
ms: 毫秒(milliseconds),即1/1000秒;
s: 秒(seconds);
m: 分钟(minutes);
h:小时(hours);
d: 天(days);
```

> 配置示例
> * 下面的例子配置了一个监听在所有接口的80端口上HTTP proxy服务,它转发所有的请求至后端监听在127.0.0.1:8000上的"server"
```
global
    daemon
    maxconn 25600
defaults
    mode http
    timeout connect 5000ms
    timeout client 50000ms
    timeout server 50000ms
frontend http-in
    bind *:80
    default_backend servers
backend servers
    server server1 127.0.0.1:8000 maxconn 32
```
> haproxy配置项包括
> * 全局配置/global
> * 代理配置/defaults,listen,frontend,backend


## haproxy全局配置项

> 全局配置项: `global`配置中的参数为进程级别的参数,通常与其运行的OS相关

> 进程管理及安全相关的参数
> * `chroot <dir>`: 修改haproxy的工作目录至指定的目录并在放弃权限之前执行chroot()操作,可以提升haproxy的安全级别,但要确保指定的目录为空目录且任何用户均不能有写权限;
> * `daemon`: 让haproxy以守护进程的方式工作于后台,其等同于"D”选项的功能,当然也可以在命令行中以"db”选项将其禁用;
> * `gid <number>`: 以指定的GID运行haproxy,建议使用专用于运行haproxy的GID,以免因权限问题带来风险;
> * `group <group name>`: 同gid,不过指定的组名;
> * `log  <address> <facility> [max level [min level]]`: 定义全局的syslog服务器,最多可以定义两个;
> * `logsendhostname [<string>]`: 在syslog信息的首部添加当前主机名,可以为"string”指定的名称,也可以缺省使用当前主机名;
> * `nbproc <number>`: 指定启动的haproxy进程的个数,只能用于守护进程模式的haproxy;默认启动一个,一般只在单进程仅能打开少数文件描述符的场景中才使用多进程模式;
> * `pidfile: <file>`: 指定haproxy进程运行的PID文件存放路径
> * `uid <number>`: 以指定的UID身份运行haproxy进程;
> * `ulimit-n <number>`: 设定每进程所能够打开的最大文件描述符数目,默认情况下其会自动进行计算,因此不推荐修改此选项;
> * `user <user name>`: 同uid,但使用的是用户名;
> * `node <node name>`: 定义当前节点的名称,用于HA场景中多haproxy进程共享同一个IP地址时;
> * `description <string>`: 当前实例的描述信息;

> 性能调整相关参数
> * `maxconn <number>`: 设定每个haproxy进程所接受的最大并发连接数,其等同于命令行选项"n","ulimit -n"自动计算的结果正是参照此参数设定的;
> * `maxpipes <number>`: haproxy使用pipe完成基于内核的tcp报文重组,此选项则用于设定每进程所允许使用的最大pipe个数;每个pipe会打开两个文件描述符,因此"ulimit -n"自动计算时会根据需要调大此值;默认为maxconn/4,其通常会显得过大;
> * `noepoll`: 在Linux系统上禁用epoll机制;
> * `nokqueue`: 在BSD系统上禁用kqueue机制;
> * `nopoll`: 禁用poll机制;
> * `nosepoll`: 在Linux禁用启发式epoll机制;
> * `nosplice`: 禁止在Linux套接字上使用内核tcp重组,这会导致更多的recv/send系统调用;不过在Linux 2.6.2528系列的内核上,tcp重组功能有bug存在;
> * `spreadchecks <0..50, in percent>`: 在haproxy后端有着众多服务器的场景中,在精确的时间间隔后统一对众服务器进行健康状况检查可能会带来意外问题;此选项用于将其检查的时间间隔长度上增加或减小一定的随机时长;
> * `tune.bufsize <number>`: 设定buffer的大小,同样的内存条件下,较小的值可以让haproxy有能力接受更多的并发连接,较大的值可以让某些应用程序使用较大的cookie信息;默认为16384,其可以在编译时修改,不过强烈建议使用默认值;
> * `tune.chksize <number>`: 设定检查缓冲区的大小,单位为字节;更大的值有助于在较大的页面中完成基于字符串或模式的文本查找,但也会占用更多的系统资源,不建议修改;
> * `tune.maxaccept <number>`: 设定haproxy进程内核调度运行时一次性可以接受的连接的个数,较大的值可以带来较大的吞吐率,默认在单进程模式下为100,多进程模式下为8,设定为-1可以禁止此限制;不建议修改;
> * `tune.maxpollevents <number>`: 设定一次系统调用可以处理的事件最大数,默认值取决于OS;其值小于200时可节约带宽,但会略微增大网络延迟,而大于200时会降低延迟,但会稍稍增加网络带宽的占用量;
> * `tune.maxrewrite <number>`: 设定为首部重写或追加而预留的缓冲空间,建议使用1024左右的大小;在需要使用更大的空间时,haproxy会自动增加其值;
> * `tune.rcvbuf.client <number>`/`tune.rcvbuf.server <number>`: 设定内核套接字中服务端或客户端接收缓冲的大小,单位为字节;强烈推荐使用默认值;
> * `tune.sndbuf.client <number>`/`tune.sndbuf.server <number>`: 设定内核套接字中服务端或客户端f缓冲的大小,单位为字节;强烈推荐使用默认值;

> Debug相关的参数
> * debug
> * quiet

## haproxy代理配置项

> defaults
> * `defaults <name>`: "defaults"段用于为所有其它配置段提供默认参数,这配置默认配置参数可由下一个"defaults"所重新设定;

> frontend
> * `frontend <name>`: "frontend"段用于定义一系列监听的套接字,这些套接字可接受客户端请求并与之建立连接;

> backend
> * `backend  <name>`: "backend"段用于定义一系列"后端"服务器,代理将会将对应客户端的请求转发至这些服务器;

> listen
> * `listen <name>`: "listen"段通过关联"前端"和"后端"定义了一个完整的代理,通常只对TCP流量有用;
> * 所有代理的名称只能使用`大写字母,小写字母,数字,-(中线),_(下划线),.(点号)和:(冒号)`;此外ACL名称会区分字母大小写;

> bind
> * 用法1: `bind [<address>]:<port_range> [, ...]`
> * 用法2: `bind [<address>]:<port_range> [, ...] interface <interface>` 
> * 此指令仅能用于frontend和listen区段,用于定义一个或几个监听的套接字;
```
<address>: 可选选项,其可以为主机名,IPv4地址,IPv6地址或*;省略此选项,将其指定为*或0.0.0.0时,将监听当前系统的所有IPv4地址;

<port_range>: 可以是一个特定的TCP端口,也可是一个端口范围(如5005-5010),代理服务器将通过指定的端口来接收客户端请求;需要注意的是,每组监听的套接字<address:port>在同一个实例上只能使用一次,而且小于1024的端口需要有特定权限的用户才能使用,这可能需要通过uid参数来定义;

<interface>: 指定物理接口的名称,仅能在Linux系统上使用;其不能使用接口别名,而仅能使用物理接口名称,而且只有管理有权限指定绑定的物理接口;

# 实例:
frontend 
    bind *:80
    bind *:8080
    ...
```

> mode
> * `mode { tcp|http|health }`: 设定实例的运行模式或协议;当实现内容交换时,前端和后端必须工作于同一种模式(一般说来都是HTTP模式),否则将无法启动实例;
> * `tcp`: 实例运行于纯TCP模式,在客户端和服务器端之间将建立一个全双工的连接,且不会对7层报文做任何类型的检查;此为默认模式,通常用于SSL,SSH,SMTP等应用;
> * `http`: 实例运行于HTTP模式,客户端请求在转发至后端服务器之前将被深度分析,所有不与RFC格式兼容的请求都会被拒绝;
> * `health`: 实例工作于health模式,其对入站请求仅响应"OK"信息并关闭连接,且不会记录任何日志信息;此模式将用于响应外部组件的健康状态检查请求;目前业讲,此模式已经废弃,因为tcp或http模式中的monitor关键字可完成类似功能;

> hash-type
> * `hash-type <method>`: 定义用于将hash码映射至后端服务器的方法;其不能用于frontend区段;可用方法有map-based和consistent,在大多数场景下推荐使用默认的map-based方法;
> * `map-based`: hash表是一个包含了所有在线服务器的静态数组;其hash值将会非常平滑,会将权重考虑在列,但其为静态方法,对在线服务器的权重进行调整将不会生效,这意味着其不支持慢速启动;此外,挑选服务器是根据其在数组中的位置进行的,因此,当一台服务器宕机或添加了一台新的服务器时,大多数连接将会被重新派发至一个与此前不同的服务器上,对于缓存服务器的工作场景来说,此方法不甚适用;
> * `consistent`: hash表是一个由各服务器填充而成的树状结构;基于hash键在hash树中查找相应的服务器时,最近的服务器将被选中;此方法是动态的,支持在运行时修改服务器权重,因此兼容慢速启动的特性;添加一个新的服务器时,仅会对一小部分请求产生影响,因此,尤其适用于后端服务器为cache的场景;不过此算法不甚平滑,派发至各服务器的请求未必能达到理想的均衡效果,因此可能需要不时的调整服务器的权重以获得更好的均衡性;

> log
> * `log global`
> * `log <address> <facility> [<level> [<minlevel>]]`: 为每个实例启用事件和流量日志,因此可用于所有区段;每个实例最多可以指定两个log参数,不过如果使用了"log global"且"global"段已经定了两个log参数时,多余了log参数将被忽略;
> * `global`: 当前实例的日志系统参数同"global"段中的定义时,将使用此格式;每个实例仅能定义一次"log global"语句,且其没有任何额外参数;
> * `<address>`: 定义日志发往的位置,其格式之一可以为`<IPv4_address:PORT>`,其中的port为UDP协议端口,默认为514;格式之二为Unix套接字文件路径,但需要留心chroot应用及用户的读写权限;
> * `<facility>`: 可以为syslog系统的标准facility之一;
> * `<level>`: 定义日志级别,即输出信息过滤器,默认为所有信息;指定级别时,所有等于或高于此级别的日志信息将会被发送;

> maxconn
> * `maxconn <conns>`: 设定一个前端的最大并发连接数,因此其不能用于backend区段;对于大型站点来说,可以尽可能提高此值以便让haproxy管理连接队列,从而避免无法应答用户请求;当然此最大值不能超出"global"段中的定义;此外,需要留心的是,haproxy会为每个连接维持两个缓冲,每个缓冲的大小为8KB,再加上其它的数据,每个连接将大约占用17KB的RAM空间;这意味着经过适当优化后,有着1GB的可用RAM空间时将能维护40000-50000并发连接;
> * 如果为<conns>指定了一个过大值,极端场景下,其最终占据的空间可能会超出当前主机的可用内存,这可能会带来意想不到的结果;因此,将其设定了一个可接受值方为明智决定;其默认为2000;

> default_backend
> * `default_backend <backend>`: 在没有匹配的"use_backend"规则时为实例指定使用的默认后端,因此其不可应用于backend区段;在"frontend"和"backend"之间进行内容交换时,通常使用"use-backend"定义其匹配规则;而没有被规则匹配到的请求将由此参数指定的后端接收;
> * `<backend>`: 指定使用的后端的名称;
> * 示例
```
use_backend     dynamic  if  url_dyn
use_backend     static   if  url_css url_img extension_img
default_backend dynamic
```

> server
> * `server <name> <address>[:port] [param*]`: 为后端声明一个server,因此不能用于defaults和frontend区段;
> * `<name>`: 为此服务器指定的内部名称,其将出现在日志及警告信息中;如果设定了"http-send-server-name",它还将被添加至发往此服务器的请求首部中;
> * `<address>`: 此服务器的的IPv4地址,也支持使用可解析的主机名,只不过在启动时需要解析主机名至相应的IPv4地址;
> * `[:port]`: 指定将连接请求所发往的此服务器时的目标端口,其为可选项;未设定时,将使用客户端请求时的同一端口;
> * `[param*]`: 为此服务器设定的一系参数;其可用的参数非常多,具体请参考官方文档中的说明,下面仅说明几个常用的参数;
```
###服务器或默认服务器参数###
  backup: 设定为备用服务器,仅在负载均衡场景中的其它server均不可用于启用此server;
  check: 启动对此server执行健康状态检查,其可以借助于额外的其它参数完成更精细的设定,如:
    inter <delay>: 设定健康状态检查的时间间隔,单位为毫秒,默认为2000;也可以使用fastinter和downinter来根据服务器端状态优化此时间延迟;
    rise <count>: 设定健康状态检查中,某离线的server从离线状态转换至正常状态需要成功检查的次数;
    fall <count>: 确认server从正常状态转换为不可用状态需要检查的次数;
  cookie <value>: 为指定server设定cookie值,此处指定的值将在请求入站时被检查,第一次为此值挑选的server将在后续的请求中被选中,其目的在于实现持久连接的功能;

###基于浏览器cookie实现session sticky###
backend websrvs
    balance 	roundrobin
    cookie SERVERID insert nocache indirect
    server web1 172.16.100.68:80 check weight 1 cookie websrv1
    server web2 172.16.100.69:80 check weight 3 cookie websrv2
						
###要点###
    1.每个server有自己惟一的cookie标识;
    2.在backend中定义为用户请求调度完成后操纵其cookie;

maxconn <maxconn>: 指定此服务器接受的最大并发连接数;如果发往此服务器的连接数目高于此处指定的值,其将被放置于请求队列,以等待其它连接被释放;
maxqueue <maxqueue>: 设定请求队列的最大长度;
observe <mode>: 通过观察服务器的通信状况来判定其健康状态,默认为禁用,其支持的类型有"layer4"和"layer7","layer7"仅能用于http代理场景;
redir <prefix>: 启用重定向功能,将发往此服务器的GET和HEAD请求均以302状态码响应;需要注意的是,在prefix后面不能使用/,且不能使用相对地址,以免造成循环;
 例如: server srv1 172.16.100.6:80 redir http://imageserver.magedu.com check
weight <weight>:权重,默认为1,最大值为256,0表示不参与负载均衡;

###后端http服务时的健康状态的检查方法###
  option httpchk
  option httpchk <uri>
  option httpchk <method> <uri>

option httpchk <method> <uri> <version>: 不能用于frontend段

backend https_relay
    mode tcp
    option httpchk OPTIONS * HTTP/1.1\r\nHost:\ www.magedu.com
    server apache1 192.168.1.1:443 check port 80

###使用案例###
    server first  172.16.100.7:1080 cookie first  check inter 1000
    server second 172.16.100.8:1080 cookie second check inter 1000
```

> balance
> * 负载均衡算法配置
```
balance
balance <algorithm> [ <arguments> ]

# 定义负载均衡算法,可用于"defaults","listen"和"backend";
# <algorithm>用于在负载均衡场景中挑选一个server,其仅应用于持久信息不可用的条件下或需要将一个连接重新派发至另一个服务器时;支持的算法有:

roundrobin: 基于权重进行轮叫,在服务器的处理时间保持均匀分布时,这是最平衡,最公平的算法;此算法是动态的,这表示其权重可以在运行时进行调整,不过,在设计上,每个后端服务器仅能最多接受4128个连接;

static-rr: 基于权重进行轮叫,与roundrobin类似,但是为静态方法,在运行时调整其服务器权重不会生效;不过其在后端服务器连接数上没有限制;

leastconn: 新的连接请求被派发至具有最少连接数目的后端服务器;在有着较长时间会话的场景中推荐使用此算法,如LDAP,SQL等,其并不太适用于较短会话的应用层协议,如HTTP;此算法是动态的,可以在运行时调整其权重;

source: 将请求的源地址进行hash运算,并由后端服务器的权重总数相除后派发至某匹配的服务器;这可以使得同一个客户端IP的请求始终被派发至某特定的服务器;不过,当服务器权重总数发生变化时,如某服务器宕机或添加了新的服务器,许多客户端的请求可能会被派发至与此前请求不同的服务器;常用于负载均衡无cookie功能的基于TCP的协议;其默认为静态,不过也可以使用hash-type修改此特性;

uri: 对URI的左半部分("问题"标记之前的部分)或整个URI进行hash运算,并由服务器的总权重相除后派发至某匹配的服务器;这可以使得对同一个URI的请求总是被派发至某特定的服务器,除非服务器的权重总数发生了变化;此算法常用于代理缓存或反病毒代理以提高缓存的命中率;需要注意的是,此算法仅应用于HTTP后端服务器场景;其默认为静态算法,不过也可以使用hash-type修改此特性;

url_param: 通过<argument>为URL指定的参数在每个HTTP GET请求中将会被检索;如果找到了指定的参数且其通过等于号"="被赋予了一个值,那么此值将被执行hash运算并被服务器的总权重相除后派发至某匹配的服务器;此算法可以通过追踪请求中的用户标识进而确保同一个用户ID的请求将被送往同一个特定的服务器,除非服务器的总权重发生了变化;如果某请求中没有出现指定的参数或其没有有效值,则使用轮叫算法对相应请求进行调度;此算法默认为静态的,不过其也可以使用hash-type修改此特性;

hdr(<name>): 对于每个HTTP请求,通过<name>指定的HTTP首部将会被检索;如果相应的首部没有出现或其没有有效值,则使用轮叫算法对相应请求进行调度;其有一个可选选项"use_domain_only",可在指定检索类似Host类的首部时仅计算域名部分(比如通过www.magedu.com来说,仅计算magedu字符串的hash值)以降低hash算法的运算量;此算法默认为静态的,不过其也可以使用hash-type修改此特性;
```

> capture request header
> * `capture request header <name> len <length>`: 捕获并记录指定的请求首部最近一次出现时的第一个值,仅能用于"frontend"和"listen"区段;捕获的首部值使用花括号{}括起来后添加进日志中;如果需要捕获多个首部值,它们将以指定的次序出现在日志文件中,并以竖线"|"作为分隔符;不存在的首部记录为空字符串,最常需要捕获的首部包括在虚拟主机环境中使用的"Host",上传请求首部中的"Content-length",快速区别真实用户和网络机器人的"User-agent",以及代理环境中记录真实请求来源的"X-Forward-For";
> * `<name>`: 要捕获的首部的名称,此名称不区分字符大小写,但建议与它们出现在首部中的格式相同,比如大写首字母;需要注意的是,记录在日志中的是首部对应的值,而非首部名称;
> * `<length>`: 指定记录首部值时所记录的精确长度,超出的部分将会被忽略;
> * 可以捕获的请求首部的个数没有限制,但每个捕获最多只能记录64个字符;为了保证同一个frontend中日志格式的统一性,首部捕获仅能在frontend中定义;

> capture response header
> * `capture response header <name> len <length>`: 捕获并记录响应首部,其格式和要点同请求首部;

> stats
> * stats: haproxy管理页面,stats参数列表如下:
> * stats enable
```
# 启用基于程序编译时默认设置的统计报告,不能用于"frontend"区段;只要没有另外的其它设定,它们就会使用如下的配置:
stats uri   : /haproxy?stats
stats realm : "HAProxy\ Statistics"
stats auth  : no authentication
stats scope : no restriction
# 尽管"stats enable"一条就能够启用统计报告,但还是建议设定其它所有的参数,以免其依赖于默认设定而带来非期后果
```
> * stats hide-version
```
# 启用统计报告并隐藏HAProxy版本报告,不能用于"frontend"区段;默认情况下,统计页面会显示一些有用信息,包括HAProxy的版本号,然而,向所有人公开HAProxy的精确版本号是非常有风险的,因为它能帮助恶意用户快速定位版本的缺陷和漏洞;尽管"stats hide-version"一条就能够启用统计报告,但还是建议设定其它所有的参数,以免其依赖于默认设定而带来非期后果;具体请参照"stats enable"一节的说明;
```
> * stats realm
```
stats realm <realm>
# 启用统计报告并高精认证领域,不能用于"frontend"区段;haproxy在读取realm时会将其视作一个单词,因此,中间的任何空白字符都必须使用反斜线进行转义;此参数仅在与"stats auth"配置使用时有意义;

<realm>: 实现HTTP基本认证时显示在浏览器中的领域名称,用于提示用户输入一个用户名和密码;
# 尽管"stats realm"一条就能够启用统计报告,但还是建议设定其它所有的参数,以免其依赖于默认设定而带来非期后果;具体请参照"stats enable"一节的说明;
```
> * stats scope
```
stats scope { <name> | "." }
# 启用统计报告并限定报告的区段,不能用于"frontend"区段;当指定此语句时,统计报告将仅显示其列举出区段的报告信息,所有其它区段的信息将被隐藏;如果需要显示多个区段的统计报告,此语句可以定义多次;需要注意的是,区段名称检测仅仅是以字符串比较的方式进行,它不会真检测指定的区段是否真正存在;
<name>:可以是一个"listen","frontend"或"backend"区段的名称,而"."则表示stats scope语句所定义的当前区段;
# 尽管"stats scope"一条就能够启用统计报告,但还是建议设定其它所有的参数,以免其依赖于默认设定而带来非期后果;
# 下面是一个配置案例:
backend private_monitoring
    stats enable
    stats uri     /haproxyadmin?stats
    stats refresh 10s
```
> * stats auth
```
stats auth <user>:<passwd>
# 启用带认证的统计报告功能并授权一个用户帐号,其不能用于"frontend"区段;
<user>:授权进行访问的用户名;
<passwd>:此用户的访问密码,明文格式;
# 此语句将基于默认设定启用统计报告功能,并仅允许其定义的用户访问,其也可以定义多次以授权多个用户帐号;可以结合"stats realm"参数在提示用户认证时给出一个领域说明信息;在使用非法用户访问统计功能时,其将会响应一个"401 Forbidden"页面;其认证方式为HTTP Basic认证,密码传输会以明文方式进行,因此,配置文件中也使用明文方式存储以说明其非保密信息故此不能相同于其它关键性帐号的密码;
# 尽管"stats auth"一条就能够启用统计报告,但还是建议设定其它所有的参数,以免其依赖于默认设定而带来非期后果;
```
> * stats admin
```
stats admin { if | unless } <cond>
# 在指定的条件满足时启用统计报告页面的管理级别功能,它允许通过web接口启用或禁用服务器,不过,基于安全的角度考虑,统计报告页面应该尽可能为只读的;此外,如果启用了HAProxy的多进程模式,启用此管理级别将有可能导致异常行为;
# 目前来说,POST请求方法被限制于仅能使用缓冲区减去保留部分之外的空间,因此,服务器列表不能过长,否则,此请求将无法正常工作;因此,建议一次仅调整少数几个服务器;下面是两个案例,第一个限制了仅能在本机打开报告页面时启用管理级别功能,第二个定义了仅允许通过认证的用户使用管理级别功能;
```

> * 启用haproxy的stats:
```
listen statistics
    bind *:9090
    stats enable
    stats hide-version
    #stats scope .
    stats uri /haproxyadmin?stats
    stats realm "haproxy\ statistics"
    stats auth admin:password
    stats admin if TRUE		#启用管理功能
```

> option httplog
> * `option httplog [ clf ]`: 启用记录HTTP请求,会话状态和计时器的功能;
> * `clf`: 使用CLF格式来代替HAProxy默认的HTTP格式,通常在使用仅支持CLF格式的特定日志分析器时才需要使用此格式;
> * 默认情况下,日志输入格式非常简陋,因为其仅包括源地址,目标地址和实例名称,而"option httplog"参数将会使得日志格式变得丰富许多,其通常包括但不限于HTTP请求,连接计时器,会话状态,连接数,捕获的首部及cookie,"frontend","backend"及服务器名称,当然也包括源地址和端口号等;

> option logasap
> * `no option logasap`: 启用或禁用提前将HTTP请求记入日志,不能用于"backend"区段;
> * 默认情况下,HTTP请求是在请求结束时进行记录以便能将其整体传输时长和字节数记入日志,由此,传较大的对象时,其记入日志的时长可能会略有延迟;"option logasap"参数能够在服务器发送complete首部时即时记录日志,只不过,此时将不记录整体传输时长和字节数;此情形下,捕获"Content-Length"响应首部来记录传输的字节数是一个较好选择;
> * 配置示例如下:
```
listen http_proxy 0.0.0.0:80
    mode http
    option httplog
    option logasap
    log 172.16.100.9 local2
```

> option forwardfor
> * `option forwardfor [ except <network> ] [ header <name> ] [ if-none ]`: 允许在发往服务器的请求首部中插入"X-Forwarded-For"首部;
> * `<network>`: 可选参数,当指定时,源地址为匹配至此网络中的请求都禁用此功能;
> * `<name>`: 可选参数,可使用一个自定义的首部,如"X-Client"来替代"X-Forwarded-For";有些独特的web服务器的确需要用于一个独特的首部;
> * `if-none`: 仅在此首部不存在时才将其添加至请求报文文档中;
> * HAProxy工作于反向代理模式,其发往服务器的请求中的客户端IP均为HAProxy主机的地址而非真正客户端的地址,这会使得服务器端的日志信息记录不了真正的请求来源,"X-Forwarded-For"首部则可用于解决此问题;HAProxy可以向每个发往服务器的请求上添加此首部,并以客户端IP为其value;
> * 需要注意的是,HAProxy工作于隧道模式,其仅检查每一个连接的第一个请求,因此,仅第一个请求报文被附加此首部;如果想为每一个请求都附加此首部,请确保同时使用了"option httpclose","option forceclose"和"option http-server-close"几个option;
> * 配置示例如下:
```
frontend www
    mode http
    option forwardfor except 127.0.0.1
```

> errorfile
> * `errorfile <code> <file>`: 在用户请求不存在的页面时,返回一个页面文件给客户端而非由haproxy生成的错误代码;可用于所有段中;
> * `<code>`: 指定对HTTP的哪些状态码返回指定的页面;这里可用的状态码有200,400,403,408,500,502,503和504;
> * `<file>`: 指定用于响应的页面文件;
> * 配置示例如下:
```
errorfile 400 /etc/haproxy/errorpages/400badreq.http
errorfile 403 /etc/haproxy/errorpages/403forbid.http
errorfile 503 /etc/haproxy/errorpages/503sorry.http
```

> errorloc/errorloc302 
> * `errorloc <code> <url>`
> * `errorloc302 <code> <url>`
> * 请求错误时,返回一个HTTP重定向至某URL的信息;可用于所有配置段中;
> * `<code>`: 指定对HTTP的哪些状态码返回指定的页面;这里可用的状态码有200,400,403,408,500,502,503和504;
> * `<url>`: Location首部中指定的页面位置的具体路径,可以是在当前服务器上的页面的相对路径,也可以使用绝对路径;需要注意的是,如果URI自身错误时产生某特定状态码信息的话,有可能会导致循环定向;
> * 需要留意的是,这两个关键字都会返回302状态码,这将使得客户端使用同样的HTTP方法获取指定的URL,对于非GET方法的场景(如POST)来说会产生问题,因为返回客户的URL是不允许使用GET以外的其它方法的;如果的确有这种问题,可以使用errorloc303来返回303状态码给客户端;

> errorloc303
> * `errorloc303 <code> <url>`: 请求错误时,返回一个HTTP重定向至某URL的信息给客户端;可用于所有配置段中;
> * `<code>`: 指定对HTTP的哪些状态码返回指定的页面;这里可用的状态码有400,403,408,500,502,503和504;
> * `<url>`: Location首部中指定的页面位置的具体路径,可以是在当前服务器上的页面的相对路径,也可以使用绝对路径;需要注意的是,如果URI自身错误时产生某特定状态码信息的话,有可能会导致循环定向;
> * 配置示例如下:
```
backend webserver
    server 172.16.100.6 172.16.100.6:80 check maxconn 3000 cookie srv01
    server 172.16.100.7 172.16.100.7:80 check maxconn 3000 cookie srv02
    errorfile 403 /etc/haproxy/errorpages/sorry.htm
    errorfile 503 /etc/haproxy/errorpages/sorry.htm
```

## haproxy负载均衡算法

> 关键字: `balance`,指明调度算法
> * 动态: 权重可动态调整
> * 静态: 调整权重不会实时生效

> `roundrobin`: 轮询,动态算法,每个后端主机最多支持4128个连接;

> `staic-rr`: 轮询,静态算法,每个后端主机支持的数量无上限;

> `leastconn`: 根据后端主机的负载数量进行调度: 仅适用于长连接的会话,动态;

> `source`: 将请求的源地址进行hash运算,并由后端服务器的权重总数相除;
 
> `hash-type`: `map-based`: 取模法,静态; `consistent`: 一致性哈希法,动态;

> `uri`: 对URI的左半部分或整个URI进行hash运算,并由服务器的总权重相除

> `url_param`: 根据url中指定的参数的值进行调度: 把值做hash计算,并除以权重;

> `hdr(<name>)`: 根据请求报文中指定的header(如user_agent,referer,hostname)进行调度,把指定的header的值做hash计算;


## haproxy ACL控制

> haproxy的ACL用于实现基于请求报文的首部,响应报文的内容或其它的环境状态信息来做出转发决策,这大大增强了其配置弹性;

> 配置法则通常分为两步,首先去定义ACL,即定义一个测试条件,而后在条件得到满足时执行某特定的动作,如阻止请求或转发至某特定的后端;

> 定义ACL的语法格式如下
```
# acl <aclname> <criterion> [flags] [operator] <value> ...
 
<aclname>: ACL名称,区分字符大小写,且其只能包含大小写字母,数字,-(连接线),_(下划线),.(点号)和:(冒号);haproxy中,acl可以重名,这可以把多个测试条件定义为一个共同的acl;

<criterion>: 测试标准,即对什么信息发起测试;测试方式可以由[flags]指定的标志进行调整;而有些测试标准也可以需要为其在<value>之前指定一个操作符[operator];

[flags]: 目前haproxy的acl支持的标志位有3个:
    -i:不区分<value>中模式字符的大小写;
    -f:从指定的文件中加载模式;
    --:标志符的强制结束标记,在模式中的字符串像标记符时使用;

<value>: acl测试条件支持的值有以下四类:
    整数或整数范围: 如1024:65535表示从1024至65535;仅支持使用正整数(如出现类似小数的标识,其为通常为版本测试),且支持使用的操作符有5个,分别为eq,ge,gt,le和lt;
    字符串: 支持使用"-i"以忽略字符大小写,支持使用"\"进行转义;如果在模式首部出现了-i,可以在其之前使用"--"标志位;
    正则表达式: 其机制类同字符串匹配;
    IP地址及网络地址

# 同一个acl中可以指定多个测试条件,这些测试条件需要由逻辑操作符指定其关系;
# 条件间的组合测试关系有三种:"与"(默认即为与操作),"或"(使用"||"操作符)以及"非"(使用"!"操作符)
```
> 常用的测试标准(criteria)
```
be_sess_rate <integer>
be_sess_rate(backend) <integer>
# 用于测试指定的backend上会话创建的速率(即每秒创建的会话数)是否满足指定的条件;常用于在指定backend上的会话速率过高时将用户请求转发至另外的backend,或用于阻止攻击行为;

# 配置示例如下:
backend dynamic
    mode http
    acl being_scanned be_sess_rate gt 50
    redirect location /error_pages/denied.html if being_scanned

fe_sess_rate <integer>
fe_sess_rate(frontend) <integer>
# 用于测试指定的frontend(或当前frontend)上的会话创建速率是否满足指定的条件;常用于为frontend指定一个合理的会话创建速率的上限以防止服务被滥用;
# 例如下面的例子限定入站邮件速率不能大于50封/秒,所有在此指定范围之外的请求都将被延时50毫秒;
frontend mail
    bind :25
    mode tcp
    maxconn 500
    acl too_fast fe_sess_rate ge 50
    tcp-request inspect-delay 50ms
    tcp-request content accept if ! too_fast
    tcp-request content accept if WAIT_END

hdr <string>
hdr(header) <string>
# 用于测试请求报文中的所有首部或指定首部是否满足指定的条件;指定首部时,其名称不区分大小写,且在括号"()"中不能有任何多余的空白字符;测试服务器端的响应报文时可以使用shdr();
# 例如下面的例子用于测试首部Connection的值是否为close:
hdr(Connection) -i close

method <string>
# 测试HTTP请求报文中使用的方法;

path_beg <string>
# 用于测试请求的URL是否以<string>指定的模式开头;
# 下面的例子用于测试URL是否以/static,/images,/javascript或/stylesheets头:
acl url_static  path_beg  -i /static /images /javascript /stylesheets

path_end <string>
# 用于测试请求的URL是否以<string>指定的模式结尾;
# 下面的例子用户测试URL是否以jpg,gif,png,css或js结尾:
acl url_static    path_end  -i .jpg .gif .png .css .js

hdr_beg <string>
# 用于测试请求报文的指定首部的开头部分是否符合<string>指定的模式;
# 下面的例子用记录测试请求是否为提供静态内容的主机img,video,download或ftp:
acl host_static hdr_beg(host) -i img. video. download. ftp.

hdr_end <string>
# 用于测试请求报文的指定首部的结尾部分是否符合<string>指定的模式;

url_beg <string>
# 用于测试url的开头部分是否符合<string>指定的模式

url_end <string>
# 用于测试url的结尾部分是否符合<string>指定的模式

path_reg
url_reg
```

## haproxy配置示例

> haproxy基础服务器配置示例
```
#---------------------------------------------------------------------
# Global settings
#---------------------------------------------------------------------
global
    log         127.0.0.1 local2
    chroot      /var/lib/haproxy
    pidfile     /var/run/haproxy.pid
    maxconn     4000
    user        haproxy
    group       haproxy
    daemon

defaults
    mode                    http
    log                     global
    option                  httplog
    option                  dontlognull
    option http-server-close
    option forwardfor       except 127.0.0.0/8
    option                  redispatch
    retries                 3
    timeout http-request    10s
    timeout queue           1m
    timeout connect         10s
    timeout client          1m
    timeout server          1m
    timeout http-keep-alive 10s
    timeout check           10s
    maxconn                 30000

listen stats
    mode http
    bind 0.0.0.0:1080
    stats enable
    stats hide-version
    stats uri     /haproxyadmin?stats
    stats realm   Haproxy\ Statistics
    stats auth    admin:admin
    stats admin if TRUE

frontend http-in
    bind *:80
    mode http
    log global
    option httpclose
    option logasap
    option dontlognull
    capture request  header Host len 20
    capture request  header Referer len 60
    default_backend servers

frontend healthcheck
    bind :1099
    mode http
    option httpclose
    option forwardfor
    default_backend servers

backend servers
    balance roundrobin
    server websrv1 192.168.10.11:80 check maxconn 2000
    server websrv2 192.168.10.12:80 check maxconn 2000
```

> haproxy负载均衡MySQL服务的配置示例
```
#---------------------------------------------------------------------
# Global settings
#---------------------------------------------------------------------
global
    log         127.0.0.1 local2
    chroot      /var/lib/haproxy
    pidfile     /var/run/haproxy.pid
    maxconn     4000
    user        haproxy
    group       haproxy
    daemon

defaults
    mode                    tcp
    log                     global
    option                  httplog
    option                  dontlognull
    retries                 3
    timeout http-request    10s
    timeout queue           1m
    timeout connect         10s
    timeout client          1m
    timeout server          1m
    timeout http-keep-alive 10s
    timeout check           10s
    maxconn                 600

listen stats
    mode http
    bind 0.0.0.0:1080
    stats enable
    stats hide-version
    stats uri     /haproxyadmin?stats
    stats realm   Haproxy\ Statistics
    stats auth    admin:admin
    stats admin if TRUE

frontend mysql
    bind *:3306
    mode tcp
    log global
    default_backend mysqlservers

backend mysqlservers
    balance leastconn
    server dbsrv1 192.168.10.11:3306 check port 3306 intval 2 rise 1 fall 2 maxconn 300
    server dbsrv2 192.168.10.12:3306 check port 3306 intval 2 rise 1 fall 2 maxconn 300
```

> * haproxy动静分离配置示例
```
global
    log         127.0.0.1 local2
    chroot      /var/lib/haproxy
    pidfile     /var/run/haproxy.pid
    maxconn     4000
    user        haproxy
    group       haproxy
    daemon
    # turn on stats unix socket
    stats socket /var/lib/haproxy/stats

defaults
    mode                    http
    log                     global
    option                  httplog
    option                  dontlognull
    option http-server-close
    option forwardfor       except 127.0.0.0/8
    option                  redispatch
    retries                 3
    timeout http-request    10s
    timeout queue           1m
    timeout connect         10s
    timeout client          1m
    timeout server          1m
    timeout http-keep-alive 10s
    timeout check           10s
    maxconn                 30000

listen stats
    mode http
    bind 0.0.0.0:1080
    stats enable
    stats hide-version
    stats uri     /haproxyadmin?stats
    stats realm   Haproxy\ Statistics
    stats auth    admin:admin
    stats admin if TRUE

frontend http-in
    bind *:80
    mode http
    log global
    option httpclose
    option logasap
    option dontlognull
    capture request  header Host len 20
    capture request  header Referer len 60
    acl url_static       path_beg       -i /static /images /javascript /stylesheets
    acl url_static       path_end       -i .jpg .jpeg .gif .png .css .js

    use_backend static_servers          if url_static
    default_backend dynamic_servers

backend static_servers
    balance roundrobin
    server imgsrv1 172.16.200.7:80 check maxconn 6000
    server imgsrv2 172.16.200.8:80 check maxconn 6000

backend dynamic_servers
    cookie srv insert nocache
    balance roundrobin
    server websrv1 172.16.200.7:80 check maxconn 1000 cookie websrv1
    server websrv2 172.16.200.8:80 check maxconn 1000 cookie websrv2
    server websrv3 172.16.200.9:80 check maxconn 1000 cookie websrv3
```