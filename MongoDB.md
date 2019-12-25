标签: MongoDB

# MongoDB

------

[toc]

## Mongod的常用选项

```
$ mongod --help
	--fork	{true|false}	#mongod是否运行在后台
	--bing_ip	IP			#指定监听的地址
	--port	PORT			#指定监听的端口
	
	--maxCons	NUM			#并发最大连接数
	--logpath	PATH		#指定日志路径
	--httpinterface			#启用http统计接口
	
	--auth					#启用数据库用户认证
	
	调试相关：
	--cpu					#阶段性显示cpu和iowait的利用率
	--sysinfo				#打印系统相关信息
```

## MongoDB CRUD

> [MongoDB CRUD][1]

## MongoDB索引

> 索引类型: B+ Tree、hash、空间索引、全文索引

> MongoDB索引类型:
> * 单字段索引
> * 组合索引(多字段索引)
> * 多键索引
> * 空间索引
> * 文本索引
> * hash索引

### 索引操作

> 生成数据
```
# 创建并进入数据库
use testdb

# for循环生成数据
for (i=1;i<=10000;i++) db.students.insert({name:"student"+i,age:(i%120),address:"#89 Wenhua Road,Zhengzhou,china"})
db.students.find().count()
```
> 创建索引
```
db.students.ensureIndex({name: 1})
```

> * 查看索引
```
db.students.getIndexes()
```

> * 删除索引
```
db.students.dropIndex("name_1")
```

> 创建唯一键索引
```
db.students.ensureIndex({name: 1},{unique: true})
```

> MongoDB与索引相关的方法
```
db.mycoll.ensureIndex(field[,options])
    name、unique、dropDups、sparse
db.mycoll.dropIndex(index_name)
db.mycoll.dropIndexes()
db.mycoll.getIndexes()
db.mycoll.reIndex()
```
		
> 生成索引后查询
```
db.students.find({name:"student5000"})
db.students.find({name:"student5000"}).explain()
db.students.find({name:{$gt:"student5000"}})
db.students.find({name:{$gt:"student5000"}}).explain()
```

## MongoDB复制

> MongoDB实现复制的两种类型:
> * master/slave: mongodb已废弃
> * replica set: 复制集、副本集

> MongoDB复制集说明
```
# 服务于同一数据集的多个mongodb实例
# 一个复制集只能有一个主节点、可以有多个从节点,主节点负责读写操作、从节点只负责读操作
# 主节点将数据修改操作保存至oplog中
# arbiter：仲裁者
```

> 工作特性
> * 至少三个,且应该为奇数个节点,可以使用arbiter来参与选举;
> * heartbeat(2s),自动失效转移(通过选择方式实现)

> 复制集中的节点分类:
> * 0优先级的节点: 冷备节点,不会被选举成为主节点,但可以参与选举;
> * 被隐藏的从节点: 首先是一个0优先级的从节点,且对客户端不可见,也可以参与选举;
> * 延迟复制的从节点: 首先是一个0优先级的从节点,且复制时间落后于主节点一个固定时长;
> * arbiter:仲裁节点
		
### MongoDB的复制架构

> 主节点的数据修改操作保存至oplog中
> * oplog是一个大小固定的文件,存储在local数据库
> * local: 存放了副本集的所有元数据和oplog,用于存储oplog的是一个名为oplog.rs的collection;
   oplog.rs的大小依赖于OS及文件系统,但可以自定义其大小oplogSize

> 主从节点通过heartbeat通信,判断状态
> * 初始同步(initial sync)
> * 回滚后追赶(post-rollback catch-up)
> * 切分块迁移(sharding chunk migrations)

> Mongo的数据同步类型
> * 初始同步
> * 节点没有任何数据时
> * 节点丢失副本复制历史
> * 复制

> 初始同步的步骤
> * 1、克隆所有数据库
> * 2、应用数据集的所有改变,复制oplog并应用于本地
> * 3、为所有collection构建索引

> 副本集的重新选举的影响条件
> * 心跳信息
> * 优先级
> * optime
> * 网络连接
> * 网络分区

> 选举机制
> * 触发选举的事件
> * 新副本集初始化时
> * 从节点联系不到主节点
> * 主节点"下台"时: 主节点收到stepDown()命令时; 某从节点有更高的优先级且已经满足成主节点其他所有条件; 主节点无法联系到副本集的"多数方"
				
### MongoDB副本集实现步骤
		
> 准备3个节点: 做好集群hosts文件解析
> * mongodb-node1	192.168.56.13	主节点
> * mongodb-node2	192.168.56.14	从节点
> * mongodb-node3	192.168.56.15	从节点

> 安装配置服务: 所有节点都执行,安装可以参考前文
```
$ yum install mongodb-org-server-2.6.4-1.x86_64.rpm mongodb-org-shell-2.6.4-1.x86_64.rpm mongodb-org-tools-2.6.4-1.x86_64.rpm -y
$ mkdir -pv /mongodb/data
$ chown -R mongod.mongod /mongodb
$ vim /etc/mongod.conf	#修改为以下选项
dbpath=/mongodb/data
bind_ip=$IP
httpinterface=true
rest=true
replSet=testSet
replIndexPrefetch=_id_only
$ service mongod start
```
	
> 配置主节点mongo shell
```
$ mongo --host mongodb-node1
> rs.help()	#查看副本集操作帮助
> rs.status()	#查看副本集状态
> rs.initiate()		#初始化
> rs.add("mongodb-node2:27017")		#添加从节点,默认端口可以不用写
> rs.add("mongodb-node3:27017")		#STARUP2:追赶主节点状态,追赶完成后显示SECONDARY
> rs.isMaster()		#查询自己是否为Master
> rs.conf()		#显示当前副本集的配置信息
> rs.stepDown() 	#放弃主节点强制成为从节点,主节点由当前的从节点根据选举法则产生;
> rs.printReplicationInfo()	#显示当前副本集的oplog状态信息
```
		
> 配置从节点mongo shell可读
```
$ mongo --host mongodb-node2
> rs.status()
> rs.slaveOk()	#执行此步才可以执行查询操作
> db.mycoll.findOne()	#执行查询操作
# 从节点不允许写入数据
```
		
> 节点下线后health状态为0,正常为1,查看方法
```
> rs.stats() # 查看health字段的值
```
				
> 修改从节点优先级: 必须主节点执行
```
PRIMARY> cfg=rs.conf()
PRIMARY> cfg.members[ID].priority=2 	#ID使用rs.stats()命令查看
PRIMARY> rs.reconfig(cfg)
```
		
> 添加仲裁节点: 必须主节点执行
```
PRIMARY> rs.addArb("ARBITER_HOSTNAMT:PORT")	#初始化添加仲裁节点
PRIMARY> rs.conf()	#查看状态
```
	
> 查看从节点状态信息
```
PRIMARY> rs.printSlaveReplicationInfo()
```

## MongoDB分片概述

> CPU、Memory、IO遇到性能瓶颈,MongoDB需要sharding Diagram;

> MySQL sharding方案: Gizzard,HiveDB,MySQL Proxy + HSACLE,Hiberbate Shard,Pyshards

> MongoDB sharding架构中的角色:
> * mogos: Router,架构统一入口,需要高可用
> * config server: 元数据服务器
> * shard: 数据节点,也称为mongod实例

> sharding类型
> * 基于范围切片: 支持顺序排序的索引,如Btree索引
> * 基于列表切片: list
> * 基于hash切片: hash

> sharding分片根据业务决定: 写离散,读集中;
			
### MongoDB实现sharding集群
		
> * 准备4个节点: 确保时间同步

|角色|主机名|IP地址|注释|
|:---|:---|:---|:---|
|mongos|mongodb-mongos|192.168.56.7|生产实现HA|
|config server|mongodb-config|192.168.56.8|生产实现HA|
|shard|mongodb-shard1|192.168.56.9|
|shard|mongodb-shard2|192.168.56.10|

> 安装配置mongodb: 所有节点执行
```
$ yum install mongodb-org-server-2.6.4-1.x86_64.rpm mongodb-org-shell-2.6.4-1.x86_64.rpm mongodb-org-tools-2.6.4-1.x86_64.rpm -y 
$ mkdir -pv /mongodb/data
$ chown -R mongod.mongod /mongodb  # install -o mongod -g mongod -d /mongodb/data 
```

> config server配置
```
$ vim /etc/mongod.conf
dbpath=/mongodb/data
bind_ip=192.168.56.8
httpinterface=true
rest=true
configsrv=true

$ service mongod start
$ ss -tnlp|grep 27019
```
	
> shard配置
```
$ mkdir -pv /mongodb/data
$ chown -R mongod.mongod /mongodb/data/
$ vim /etc/mongod.conf
dbpath=/mongodb/data
bind_ip=[192.168.56.9|192.168.56.10]
auth=true	#根据实际情况开启用户认证
httpinterface=true
rest=true

$ service mongod start
$ ss -tnlp|grep 27017
```
			
> * mongos配置
```
$ yum install mongodb-org-mongos-2.6.4-1.x86_64.rpm -y
$ mongos --configdb=192.168.56.8:27019 --fork --logpath=/var/log/mongodb/mongos.log
$ ss -tnlp|grep 27017
$ mongo --host 192.168.56.7
> sh.help()		#查看shard相关帮助
> sh.status()	#查看shard状态

# 添加sharding节点
> sh.addShard("192.168.56.9:27017")
> sh.addShard("192.168.56.10:27017")

# 激活启用sharding的database
> sh.enableShareding("testdb")	

# 对testdb库中students集合的age字段做范围分片
> sh.shardCollection("students",{"age": 1})	

> use admin
> db.runCommand("listShards")	#列出集群中的shards,需要切换到admin库中
> db.printShardingStatus()	#显示集群的sharding信息
			
> sh.isBalancerRinning()		#手工均衡sharding
> sh.getBalancerState()		#查看集群sharding均衡状态
```

  [1]: https://www.cnblogs.com/zhuminghui/p/8330429.html