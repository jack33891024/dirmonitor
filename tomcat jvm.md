标签: tomcat

# tomcat jvm

---

[toc]

## JVM HEAP内存空间

> 新生代:
> * 新生区(Eden): 初创对象
> * 存活区(Survivor): 步入成熟期的初创对象
```
ss1: Survivor Space 1
ss2: Survivor Space 2
```

> 老年代: mark --> compact

> 持久代: 不会被删除
		
## 垃圾回收器

> 新生代回收器: Minor GC, 频率高

> 老年代回收器: Major GC(FULL GC),遍历内存空间,标记打包,然后清除
			
## 堆内存空间的调整参数
		
> `-XmX`: 新生代和老年代总共可用的最大空间

> `-Xms`: 新生代和老年代初始空间之和

> `-XX:NewSize`: 新生代初始空间

> `-XX:MaxNewSize`: 新生代的最大空间

> * `-xx:MaxPermSize`: 持久代最大空间

> * `-xx:PermSize`: 持久代初始空间

![image_1d4sqs5ujmdm1l3l19r11ske10bt9.png-135.8kB][1]
		
## tomcat堆内存参数

> catalina.sh中有两个环境变量:
> * `CATALINA_OPTS`：仅对启动运行tomcat实例的java虚拟机有效;
> * `JAVA_OPTS`：对本机上的多有java虚拟机有效;
	
## 性能监控工具

> 存在问题:
> * OutOfMemoryError: 内存不足;
> * 内存泄露
> * 线程死锁
> * 锁竞争(Lock Contention)
> * Java消耗过多的CPU
			
> jps(java virtual machine process status tool): 监控jvm进程状态的信息
```
jps [options] [hostid]
  -m：输出传入main方法的参数
  -l：显示main类或jar的完全限定名称
  -v：显示jvm虚拟机指定的参数			
```
		
> jstack: 查看某个java进程内的线程堆栈信息
```
jstack [options] pid
  选项:
    -l long listings: 输出完成的锁信息
    -m	混合模式,即会输出java堆栈及C/C++堆栈信息		
```
		
> jmap: jvm memory map	
```			
jmap [options] pid
  -heap	详细输出堆内存空间内存使用状态；
  -histo:live	查看堆内存中的对象数目、大小统计结果;
```

> jstat: jvm统计监测工具
```
jstat -<options> [-t] [-h<lines>] <vmid> [<interval>] [<count>]
  其中<options>为必须提供的选项,所有可用选项可使用jstat -options列出:
    -class
    -compiler
    -gc
    -gccapacity
    -gccause
    -gcnew
    -gcnewcapacity
    -gcold
    -gcoldcapacity
    -gcpermacpacity
    -gcutil
    -printcompilation		
  字段意义:
    S0C,S1C,S0U,S1U：C表示容量,U表示已用量
    EC,EU：eden区域的容量和已用量
    OC,OU
    PC,PU
    YGC,YGT: 新生代的GC次数和耗时
    TGC,FGCT: FULL GC的次数和耗时
    GCT: GC总耗时
```
					
> JVM GUI TOOLS:
 - jconsole
 - jvisualvm


  [1]: http://static.zybuluo.com/yujianfeng/a9snn5r6zehzjl60qbt06ujn/image_1d4sqs5ujmdm1l3l19r11ske10bt9.png