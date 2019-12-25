标签: tomcat

# Java技术体系

---

[toc]

## 编程语言

> 系统级: C,C++,go,erlang

> 应用级: C#,Java,Python,Perl,Ruby,PHP

## 虚拟机

> 提供程序自身的运行环境,如: jvm(java),pvm(python)

> 动态站点: asp/.net,java/jsp,python/django等开发框架;
		
## 动态网站

> 客户端动态

> 服务端动态
> * CGI:Common Gateway Interface
> * webapp server: `jsp->tomcat,jboss,jetty`和`php->php-fpm`
			
## java编程语言

> Java发展历史
```
SUN,James Gosling,Green Project(Oak --> Java),SunWorld大会
1995: Java 1.0,Write Once,Run Aynwhere
1996: JDK(Java Development Kit),包含一个JVM(Sun Classic VM)
    JDK 1.0: JVM,Applet(小应用程序),AWT
1997: JDK 1.1,JAR文档格式,JDBC,JavaBeans
1998:JDK 1.2
    Sun把Java技术拆分为三个方向:Java 2
        J2SE:Standard Edition/桌面级应用
        J2EE:Enterprise Edition/企业级应用
        J2ME:Mobile Edition/移动端应用
    代表性技术: EJB(Enterprise JavaBeans),Java Plug-in(java插件),Swing,
    JIT编译器: Just In Time,即时编译器
1999: 引入HostSpot虚拟机(HostSpot由Sun公司收购而来)
2000: JDK 1.3
2002: JDK 1.4
2006: Sun开源了Java技术;遵循GPL规范,并建立了OpenJDK组织管理此些代码;
2009: Oracle收购Sun公司
虚拟机: JRockit,HostSpot;
```
> 编程语言的类别:
> * 面向过程:以指令为中心,围绕指令组织数据;
> * 面向对象:以数据为中心,围绕数据组织指令;		

> Java体系结构:
> * Java编程语言
> * Java Class文件格式
> * Java API: Java语言提供的类,供编程人员调用
> * Java VM

> JVM的核心组成部分:
> * Class Loader: 装载需要加载的各种class
> * 执行引擎: C语言研发

> Java编程语言的特性:
> * 面向对象,多线程,结构化错误处理
> * 垃圾收集,动态链接,动态扩展

> 三个技术流派: 重命名
> * J2EE ==> Java 2 EE
> * J2SE ==> Java 2 SE
> * J2ME ==> Java 2 ME

> JDK: Java + Development Technology + APIs + JRE

> JRE: Java Runtime Environment(运行Java代码的环境,不提供程序开发功能)

> JVM运行时区域: 运行为多个线程
> * 方法区: 线程共享,用于存储被虚拟机加载的类信息,常量,静态变量等,永久代;
> * 堆: 线程共享,Java堆是jvm所管理的内存中最大的一部分,也是GC管理的主要区域,主流的算法都基于分代收集方式进行: 新生代和老年代;
> * Java栈: 线程私有,存放线程自己的局部变量等信息;
> * PC寄存器(Program Counter Register),线程独占的内存空间;
> * 本地方法栈:平台独有(platform: windows,linux,unix)

> Java 2 EE:
> * Java 2 SE
> * Servlet(服务端应用程序),JSP,EJB,JMS,JMX,JavaMail
> * Servlet Containler:
```
`println("<h1>")`
html标签要硬编码在应用程序中;
```

> JSP: JavaServer Pages, Servlet的前端, 使用jasper(引擎)翻译
			
> Web Container: JDK,Servlet,JSP

> 商业实现:
> * WebSphere(IBM)
> * WebLogic(BEA --> Oracle)
> * Oc4j
> * Glassfish
> * Geronimo
> * JOnAS
> * JBoss

> 开源实现:
> * Tomcat
> * Jetty
> * Resin
		
## Tomcat,Jetty,Resin区别
	
> * Apache Tomcat
```
Stand alone web server supporting the J2EE web technologies (e.g Servlets, JSP).
Free to use for any purpose.
Support remote deployment and administration.
Compile time of all JSPs,it took 8 seconds.
```
> * Jetty
```
Opensource Java HTTP Server and Servlet Container.
Small, Efficient, Production proven and Embeddable.
Designed be embedded in other java code, means jetty as a set of JAR files.
Jetty integrations with Geronimo, JBoss, Jonas, etc.
Compile time of all JSPs, it took 13 seconds.
```
> * Resin
```
High-performance J2EE Application server (eg. Servlets, JSP, EJB).
Free only for open source/student purpose use, not for commercial purpose use.
Supports load balancing for increased reliability, separation of content from style with its fast XML and XSL support.
Ability to run PHP applications under the JVM through Quercus.
Servlets and their related classes compile/recompile is done automatically.
Support servlet cron job.
Configuration and deployment is much simpler.
Does not support remote deployment, done through another system(eg. SSH).
Compile time of all JSPs, it took 7 seconds.
```
	  
## Java性能故障排除工具

> jconsole是随着JDK 1.5而推出的,这是一个Java监测和管理控制台-JMX兼容的图形工具来监测Java虚拟机,它能够同时监测本地和远程的JVMs;

> VisualVM集成了几个现有的JDK软件工具,轻量级内存和CPU信息概要能力,这个工具被设计为同时在生产和开发时使用,进一步提高监视的能力和Java SE平台的性能分析能力;

> HeapAnalyzer: 能够通过它采用启发式搜索引擎和分析Java堆栈存储信息发现可能的Java堆栈泄漏区域,它通过解析Java堆栈存储信息,创建定向图表,变换他们成定向树和执行启发式搜索引擎;    

> PerfAnal是在Java 2平台上为了分析应用表现的一个基于GUI的工具,您能使用PerfAnal的辩认性来查找出您需要调整的代码;

> JAMon是一个免费,简单,高性能,安全,允许开发者容易地监测生产应用程序的Java API;

> Eclipse Memory Analyzer: 帮助您发现内存泄漏和减少记忆消耗量的一台快速和功能丰富的Java堆分析仪; 

> GCViewer: 一个免费开源工具,使用Java VM属性-verbose:gc 和－Xloggc生成可视化数据,它也计算垃圾收集相关的性能指标(生产量、积累停留、最长的停留等等);

> 在/usr/local/tomcat/bin目录下的catalina.sh
```
添加: JAVA_OPTS=''-Xms512m -Xmx1024m''
要加"m"说明是MB，否则就是KB了，在启动tomcat时会报内存不足。
-Xms: 初始值
-Xmx: 最大值
-Xmn: 最小值
Windows
```
		
## MVC架构

> Controller,Model和View各自独立,一个流行的开源实现是Apache Structs框架;目今,设计优良的Web应用程序通常用相就的技术实现相应的功能,比如:
> * 1、Servlet用于实现应用逻辑;
> * 2、JSP用于内容展示;
> * 3、标签库和JSP扩展语言用于替换在JSP内部嵌入Java代码,进而降低了HTML维护的复杂度;
> * 4、MVC框架用于实现展示和应用逻辑的分离;