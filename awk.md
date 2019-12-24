标签:三剑客

# awk

[toc]

## AWK介绍

> * Awk是一个维护和处理文本数据文件的强大语言, 在文本数据有一定的格式, 即每行数据包含多个以分界符分隔的字段时, 显得尤其有用; 即便是输入文件没有一定的格式, 你仍然可以使用awk进行基本的处理, Awk当然也可以没有输入文件, 那不是必须的; 简而言之, AWK是一种能处理从琐碎的小事到日常例行公事的强大语言; 
> * 学习AWK的难度要比学习其他任意语言的难度都小, 如果你已经掌握了C语言, 那么你会发现学习AWK将会是如此简单和容易; 
> * AWK最开始由三个人开发 --> A.Aho、B.W.Kernighan 和P.Weinberger, 所以AWK的名字来自他们名字的第一个字母;
> * AWK的几个变种:
 - AWK是最原始的AWK 
 - NAWK是new AWK 
 - GAWK是GNU AWK, 所有linux发行版都默认使用GAWK, 它和AWK以及NAWK完全兼容
> * 本内容包含了原始AWK的所有基础功能, 以及GAWK特有的一些高级功能; 在安装了NAWK或GAWK的操作系统上, 你仍然可以直接使用awk命令, 它会根据情况调用nawk或gawk;
> * 以linux系统为例, 你会看到awk是一个指向gawk的符号链接, 所以在linux上执行awk或gawk将会调用gawk:
```
$ ls -l /bin/awk /bin/gawk 
lrwxrwxrwx. 1 root root      4 Apr 27 15:06 /bin/awk -> gawk
-rwxr-xr-x  1 root root 382456 Aug  7  2012 /bin/gawk 
```

## 测试文件创建

> * 本次实验环境将用到下面三个文件, 请先建立它们, 然后用它们来运行所有示例
> * example.txt文件 
 - example.txt文件以逗号作为字段分界符,包含5个雇员的记录,其格式如下：number, name, title 
```
$ cat example.txt 
101,John Doe,CEO 
102,Jason Smith,IT Manager 
103,Raj Reddy,Sysadmin 
104,Anand Ram,Developer 
105,Jane Miller,Sales Manager
```
> * items.txt 文件 
 - items.txt是一个以逗号作为字段分界符的文本文件,包含5条记录,其格式如下: items-number,item-description,item-category,cost,quantity-available 
```
$ cat items.txt 
101,HD Camcorder,Video,210,10 
102,Refrigerator,Appliance,850,2 
103,MP3 Player,Audio,270,15 
104,Tennis Racket,Sports,190,20 
105,Laser Printer,Office,475,5
```
> * items-sold.txt是一个以空格作为字段分界符的文本文件,包含5条记录, 每条记录都是特定商品的编号以及当月的销售量(6个月), 因此每条记录有7个字段; 第一个字段是商品编号, 第二个字段到第七个字段是6个月内每月的销售量;
 - 其格式如下: item-number qty-sold-month1 qty-sold-month2 qty-sold-month3 qty-sold-month4 qty-sold-month5 qty-sold-month6
```
$cat items-sold.txt
101 2 10 5 8 10 12 
102 0 1 4 3 0 2 
103 10 6 11 20 5 13 
104 2 3 4 0 6 5 
105 10 2 5 7 12 6 
```

## Awk命令语法

> * Awk基础语法: 
```
Awk -Fs '/pattern/ {action}' input-file  
Awk -Fs '{action}' input-file 

-F 为字段分界符, 如果不指定,默认会使用空格作为分界符; 
/pattern/和{action}需要用单引号引起来; 
/pattern/是可选的, 如果不指定, awk将处理输入文件中的所有记录, 如果指定一个模式, awk则只处理匹配指定的模式的记录; 
{action} 为awk命令, 可以是单个命令, 也可以多个命令; 整个action(包括里面的所有命令)都必须放在{ }之内; 
Input-file 即为要处理的文件 
```
> * 下面是一个演示awk语法的非常简单的例子: 
```
$awk -F ":" '/mail/{print $1}' /etc/passwd
mail
 
-F 指定字段分界符为冒号, 即各个字段以冒号分隔; 请注意, 你也可以把分界符用双引号引住, 即-F":" 也是正确的; 
/mail/ 指定模式, awk只会处理包含关键字mail的记录; 
{print $1} 动作部分, 该动作只包含一个awk命令, 它打印匹配mail的每条记录的第1个字段; 
/etc/passwd即是输入文件
```
 > * 把awk命令放入单独的文件中(awk脚本) 
 - 当需要执行很多awk命令时, 可以把/pattern/{action}这一部分放到单独的文件中, 然后调用它: 
   awk -Fs -f myscript.awk input-file 
 - myscript.awk可以使用任意扩展名(或者不用扩展名), 但是加上扩展名.awk便于维护, 也可以在这个文件中设置字段分界符调用: `awk -f myscript.awk input-file` 

## Awk程序结构区域

> * 典型的awk程序包含下面三个区域:
 - BEGIN
 - body
 - END
> * BEGIN 区域 
 - Begin区域的语法: BEGIN { awk-commands } 
 - BEGIN 区域的命令只最开始、在awk执行body区域命令之前执行一次。 
 - BEGIN区域很适合用来打印报文头部信息,以及用来初始化变量。 
 - BEGIN区域可以有一个或多个awk命令 
 - 关键字BEGIN必须要用大写 
 - BEGIN区域是可选的 
> * body区域 
 - body区域的语法: /pattern/ {action}  
 - body区域的命令每次从输入文件读取一行就会执行一次 
 - 如果输入文件有10行,那body区域的命令就会执行10次(每行执行一次) 
 - Body区域没有用任何关键字表示,只有用正则模式和命令。
> * END区域
 - END区域的语法: END { awk-commands } 
 - END区域在awk执行完所有操作后执行,并且只执行一次。 
 - END区域很适合打印报文结尾信息,以及做一些清理动作 
 - END 区域可以有一个或多个awk命令 
 - 关键字END必须要用大写 
 - END区域是可选的 
> * 下面的例子包含上上述的三个区域： 
```
$ awk 'BEGIN{FS=":";print "----header----"} \
/mail/ {print $1} \
END{print "----footer----"}' \
/etc/passwd 
----header----
mail
----footer----

#提示: 如果命令很长, 即可以放到单行执行, 也可以用\折成多行执行; 上面的例子用\把命令折成了3行
BEGIN { FS=”:”;print “----header----“ } 为BEGIN区域, 它设置了字段分界符变量FS的值, 然后打印报文头部信息; 这个区域仅在body区域循环之前执行一次;
/mail/{print $1} body区域, 包含一个正则模式和一个动作, 即在输入文件中搜索包含关键字mail的行, 并打印第一个字段;
END {print “----footer----“ } END区域,打印报文尾部信息;
/etc/passwd 是输入文件, 每行记录都会执行一次body区域里的动作; 
```
> * 上面的例子中, 除了可以在命令行上执行外, 还可以通过脚本执行 
 - 首先建立下面的文件myscript.awk, 它包含了begin,body和end, 然后如下所示, 在/etc/passwd上执行:
```
$ cat mysctipt.awk 
BEGIN { 
FS=":" 
print "---header---" 
} 
/mail/ { 
print $1 
} 
END { 
print "---footer---" 
} 

$ awk -f mysctipt.awk /etc/passwd
---header---
mail
---footer---
```
> * awk脚本中, 注释以#开头, 如果要编写复杂的awk脚本, 在*awk文件中写上足够多的注释, 这样以后再次使用该脚本时, 更易于读懂;
> * 下面是随机列出的一些简单的例子,用例演示awk各个区域的不同组合方式：
``` 
# 只有body区域: 
$ awk -F : '{ print $1 }' /etc/passwd

# 同时具有begin,body和end区域: 
$awk -F: 'BEGIN{printf "username\n-------\n"} \
{ print $1 } \
END {print "----------" }' \
/etc/passwd

# 只有begin和body区域: 
$ awk -F ":" 'BEGIN{print "UID"}{print $3}' /etc/passwd  
```
> * 关于使用BEGIN区域的提示 
 - 只使用BEGIN区域在awk中是符合语法的, 在没有使用body区域时, 不需要指定输入文件, 因为body区域只在输入文件上执行; 所以在执行和输入文件无关的工作时, 可以只使用BEGIN区域;
 - 下面的不少例子中, 只包含BEGIN区域,用来说明awk的不同部分是如何执行的; 
 - 只包含BEGIN的简单示例: 
```
$ awk 'BEGIN{ print "Hello,World!"}'    
Hello,World!
```
> * 多个输入文件: 
 - 可以为awk指定多个输入文件, 如果指定了两个文件, 那么body区域会首先在第一个文件的所有行上执行, 然后在第二个文件的所有行上执行; 
 - 多个输入文件示例: 
```
awk 'BEGIN{FS=":";print "---header---"} \
/mail/{print $1} \
END{ print "---footer---"}'  \
/etc/passwd /etc/group
# 注意,即是指定了多个文件,BEGIN和END区域,仍然只会执行一次。
```

## 打印命令

> * 默认情况下, awk的打印命令print(不带任何参数)会打印整行数据, 下面的例子等价于"cat example.txt"命令
```
$ awk '{print}' example.txt   
101,John Doe,CEO 
102,Jason Smith,IT Manager 
103,Raj Reddy,Sysadmin 
104,Anand Ram,Developer 
105,Jane Miller,Sales Manager 
```

> * 可以通过传递变量"$字段序号"作为print的参数来指定要打印的字段, 例如只打印雇员名称(第2个字段);
``` 
$ awk '{print $2}' example.txt 
Doe,CEO
Smith,IT
Reddy,Sysadmin
Ram,Developer
Miller,Sales

# 输出和预期不符, 它打印了从姓氏开始直到记录结尾的所有内容, 这是因为awk默认的字段分隔符是空格, awk准确地执行了我们要求的动作, 它以空格作为分隔符, 打印第2个字段; 当使用默认的空格作为字段分隔符时, 101,Johne变成了第一条记录的第一个字段, Doe,CEO变成了第二个字段; 因此上面例子中, awk把Doe,CEO作为第二个字段打印出来了;
# 要解决这个文件, 应该使用-F选项为awk指定一个逗号",", 最为字段分隔符;
$ awk -F ',' '{print $2}' example.txt 
John Doe
Jason Smith
Raj Reddy
Anand Ram
Jane Miller
```

> * 当字段分隔符是单个字符时, 下面的所有写法都是正确的, 即可以把它放在单引号或双引号中,或者不使用引号 
``` 
$ awk -F ',' '{print $2}' example.txt 
$ awk -F "," '{print $2}' example.txt  
$ awk -F , '{print $2}' example.txt   
```

> * 输出雇员姓名, 职位, 同时附带header和footer信息:
``` 
$ awk 'BEGIN{FS=",";print "---------\nName  Title\n------------\n"} \
{print $2,"\t",$3} \
END{print "-------------------"}'  \
example.txt  

---------
Name  Title
------------
John Doe         CEO 
Jason Smith      IT Manager 
Raj Reddy        Sysadmin 
Anand Ram        Developer 
Jane Miller      Sales Manager 
-------------------

# 这个例子中,输出结果各字段并没有很好地对齐,后面章节将会介绍如何处理这个问题。这个例子还展示了如何使用BEGIN来打印header以及如何使用END来打印footer. 
# 请注意,$0代表整条记录。下面两个命令是等价的,都打印example.txt的所有行: 
$ awk '{print}' example.txt 
$ awk '{print $0}' example.txt
```

## 模式匹配

> * 可以只在匹配特殊模式的行数执行awk命令; 
> * 下面的例子只打印管理者的姓名和职位;
``` 
$awk -F',' '/Manager/{print $2,$3}' example.txt       
Jason Smith IT Manager 
Jane Miller Sales Manager 
```

> * 下面的例子只打印雇员id为102的雇员的信息;
``` 
$awk -F',' '/^102/{print "Emp id 102 is",$2}' example.txt   
Emp id 102 is Jason Smith 
```

## awk内置变量 

### FS输入字段分隔符

> * awk默认的字段分隔符是空格, 如果你的输入文件中不是一个空格作为字段分隔符, 通过-F选项指定其他分隔符;
``` 
$ cat example.txt
101,John Doe,CEO 
102,Jason Smith,IT Manager 
103,Raj Reddy,Sysadmin 
104,Anand Ram,Developer 
105,Jane Miller,Sales Manager

$ awk -F ',' '{print $2,$3}' example.txt  
John Doe CEO 
Jason Smith IT Manager 
Raj Reddy Sysadmin 
Anand Ram Developer 
Jane Miller Sales Manager
```

> * 也可以使用awk内置变量FS来完成, FS只能在BEGIN区域中使用;
``` 
$awk 'BEGIN {FS=","}{print $2,$3}' example.txt     
John Doe CEO 
Jason Smith IT Manager 
Raj Reddy Sysadmin 
Anand Ram Developer 
Jane Miller Sales Manager
```

> * BEGIN区域可以包含多个命令, 如下, BEGIN区域包含一个FS和一个print命令, BEGIN区域的多个命令之间要用分号分隔; 
 - 注意: 默认的字段分隔符不仅仅是单个空格字符, 它实际上是一个或多个空白字符; 
```
$ awk 'BEGIN{ FS=",";print "-------------------\nName\tTitle\n-------------------"} \
{print $2,"\t",$3;} \
END{print "------------------------------"}'  \
example.txt
---------------------------
Name    Title
------------------------
John Doe         CEO 
Jason Smith      IT Manager 
Raj Reddy        Sysadmin 
Anand Ram        Developer 
Jane Miller      Sales Manager 
-----------------------------------------
```

> * 下面的examples.txt文件, 每行记录都包含3个不同的字段分隔符: 
 -  ,雇员id后面的分隔符是逗号 
 -  :雇员姓名后面的分隔符是分号 
 - %雇员职位后面的分隔符是百分号
```
# 创建文件 
$ cat examples.txt
101,John Doe:CEO%10000 
102,Jason Smith:IT Manager%5000 
103,Raj Reddy:Sysadmin%4500 
104,Anand Ram:Developer%4500 
105,Jane Miller:Sales Manager%3000  
```

> * 当遇到一个包含多个字段分隔符的文件时, FS可以使用正则表达式来指定多个字段分隔符, 如FS = "[,:%]"指定字段分隔符可以是逗号,或者分号:或者百分号% 
 - 下面的例子将打印examples.txt文件中雇员名称和职位: 
```
$ awk 'BEGIN {FS="[,:%]"}{print $2,$3}' examples.txt              
John Doe CEO
Jason Smith IT Manager
Raj Reddy Sysadmin
Anand Ram Developer
Jane Miller Sales Manager
```

### OFS输出字段分隔符

> * FS是输入字段分隔符, OFS是输出字段分隔符;
> * OFS会被打印在输出行的连续的字段之间, 默认情况下, awk在输出字段中间以空格分开; 
  - 请注意,我们没有指定IFS作为输入字段分隔符,我们从简地使用FS。 
> * 打印雇员姓名和薪水,并以空格分开, 使用单个print语句打印多个以逗号分开:
``` 
$ awk -F ',' '{print $2,$3}' example.txt 
John Doe CEO 
Jason Smith IT Manager 
Raj Reddy Sysadmin 
Anand Ram Developer 
Jane Miller Sales Manager 
```

> * 如果你尝试人为地在输出字段之间加上冒号, 会有如下输出; 请注意在冒号前后均有一个多余的空格, 这是因为awk仍然以空格作为输出字段分隔符。;
 - 下面的print语句实际上会打印3个值(以逗号分割)$2,:和$3, 当使用单个print语句打印多个变量时, 输出内容会包含多余的空格;
``` 
$awk -F',' '{print $2,":",$3}' example.txt 
John Doe : CEO 
Jason Smith : IT Manager 
Raj Reddy : Sysadmin 
Anand Ram : Developer 
Jane Miller : Sales Manager

# 正确的方法是使用awk内置变量OFS(输出字段分隔符), 如下面了示例, 请注意这个例子中分号前后没有多余的空格,因为OFS使用冒号取代了awk默认的分隔符; 
# 下面的print语句打印两个变量($2和$3), 输出结果却是却以分号分隔(而不是空格), 因为OFS被设置分号
$ awk -F ',' 'BEGIN {OFS=":"} {print $2,$3}' example.txt 
John Doe:CEO 
Jason Smith:IT Manager 
Raj Reddy:Sysadmin 
Anand Ram:Developer 
Jane Miller:Sales Manager 

# 同时请注意在print语句中使用和不使用逗号的细微差别(打印多个变量时), 当在print语句中指定了逗号, awk会使用OFS, 不指定逗号,awk不会使用输出分隔符。;
$ awk 'BEGIN { print "test1","test2"}'
test1 test2
$ awk 'BEGIN { print "test1""test2"}' 
test1test2 
```

### RS记录分隔符

> * 假定有下面一个文件, 雇员的id和名称都在单一的一行内;
```
$ cat example-one-line.txt
101,John Doe;102,Jason Smith;103,Raj Reddy;104,Anand Ram;105,Jane, Miller  
```

> * 这个文件中, 每条记录包含两个字段并且以逗号分隔; 
 - Awk默认的记录分隔符是换行符(回车), 如果要尝试只打印雇员姓名,下面的例子无法完成:
``` 
$ awk -F, '{print $2}' example-one-line.txt 
John Doe:102

# 这个例子把example-one-line.txt的内容作为单独一行, 把逗号作为字段分隔符, 所以它打印"John Doe:102"作为第二个字段;
```
> * 如果要把文件内容作为5行记录来处理(而不是单独的一行), 并且打印每条记录中雇员的姓名, 就必须把记录分隔符指定为分号: 
```
$ awk -F, 'BEGIN { RS=":" }{print $2}' example-one-line.txt   
John Doe
Jason Smith
Raj Reddy
Anand Ram
Jane 
```

> * 假设有下面的文件, 记录之间用-分隔, 独占一行, 所有的字段都占单独的一行;
``` 
$ cat example-enter.txt
101 
John Doe 
CEO 
-
102 
Jason Smith 
IT Manager 
- 
103 
Raj Reddy 
Sysadmin 
- 
104 
Anand Ram 
Developer 
- 
105 
Jane Miller 
Sales Manager  
```

> * 上面例子中, 字段分隔符FS是换行符, 记录分隔符RS是"-"和换行符, 所以如果要打印雇员名称和职位,需要:
``` 
$ awk 'BEGIN{RS="-";FS=" ";OFS=":"}{print $2" "$3,$4" "$5}' example-enter.txt 
John Doe:CEO 
Jason Smith:IT Manager
Raj Reddy:Sysadmin 
Anand Ram:Developer 
Jane Miller:Sales Manager 
```

### ORS输出记录分隔符

> * RS是输入字段分隔符, ORS是输出字段分隔符;  
> * 下面的例子在每个输出行后面追加"---------", awk默认使用换行符"\n"作为ORS, 我们使用"\n---\"作为ORS;
``` 
$ awk 'BEGIN {FS=",";ORS="\n---\n"}{print $2,$3}' example.txt 
John Doe CEO 
---
Jason Smith IT Manager 
---
Raj Reddy Sysadmin 
---
Anand Ram Developer 
---
Jane Miller Sales Manager 
--- 
```

> * 下面的例子从example.txt获取输入, 把每个字段打印成单独一行, 每条记录用"---"分隔;
``` 
$ awk 'BEGIN { FS=",";OFS="\n";ORS="\n---\n"}{print $1,$2,$3}' example.txt  
101
John Doe
CEO 
---
102
Jason Smith
IT Manager 
---
103
Raj Reddy
Sysadmin 
---
104
Anand Ram
Developer 
---
105
Jane Miller
Sales Manager 
--- 
```

### NR记录序号

> * NR非常有用, 在循环内部标识记录序号, 用于END区域时, 代表输入文件的总记录数; 
> * 尽管你会认为NR代表"记录的数量(Number of Records)", 但它跟确切的叫法是"记录的序号(Number of the Record)", 也就是当前记录在所有记录中的行号; 
> * 下面的例子演示了NR在block和END区域是怎么运行的:
``` 
$ awk 'BEGIN {FS=","} \
{print "Emp Id of record number",NR,"is",$1;} \
END {print "Total number of records:",NR}'  \
example.txt
  
Emp Id of record number 1 is 101
Emp Id of record number 2 is 102
Emp Id of record number 3 is 103
Emp Id of record number 4 is 104
Emp Id of record number 5 is 105
Total number of records: 5
```

### FILENAME当前处理的文件名

> * 当使用awk处理多个输入文件时, FILENAME就显得很有用, 它代表awk当前正在处理的文件;
``` 
$ awk '{print FILENAME}' example.txt examples.txt 
example.txt
example.txt
example.txt
example.txt
example.txt
examples.txt
examples.txt
examples.txt
examples.txt
examples.txt
examples.txt
```
 
> * 如果awk从标准输入获取内容, FILENAME的值将会是"-", 下面的例中, 我们不提供任何输入文件, 所以你应该手动输入内容以代替标准输入; 
  - 例如: 我们只输入一个人名"John Doe"作为第一条记录, 然后awk打印出该人的姓氏; 这种情况下, 必须按Ctrl-C才能停止标准输入;
``` 
$ awk '{print "Last name:",$2;print "Filename:",FILENAME}'
John Deo
Last name: Deo
Filename: - 
```
	
> * 上面这个例子在使用管道向awk传递数据时, 同样适用; 如下所示,打印出来的FILENAME仍然是"-"
``` 
$ echo "John Doe"|awk '{print "Last name:",$2;print "Filename:",FILENAME}'
Last name: Doe
Filename: -
# 注意: 在BEGIN区域内,FILENAME的值是空,因为BEGIN区域只针对awk本身,而不处理任何文件。 
```

### FNR文件中的NR 

> * 我们已经知道NR是"记录条数"(或者叫"记录的序号"), 代表awk当前处理的记录的行号; 
> * 在给awk传递了两个输入文件时NR会是什么, NR会在多个文件中持续增加, 当处理到第二个文件时, NR不会被重置为1,而是在前一个文件的NR基础上继续增加; 
> * 下面的例子中, 第一个文件有5条记录, 第二个文件也有5条记录; 如下所示, 当body区域的循环处理到第二个文件时, NR从6开始递增(而不是1), 最后在END区域, NR返回两个文件的总记录条数;
``` 
$ awk 'BEGIN {FS=","} \
{print FILENAME ": record number",NR,"is",$1;} \
END {print "Total number of records:",NR}'  \
example.txt examples.txt
example.txt: record number 1 is 101
example.txt: record number 2 is 102
example.txt: record number 3 is 103
example.txt: record number 4 is 104
example.txt: record number 5 is 105
examples.txt: record number 6 is 101
examples.txt: record number 7 is 102
examples.txt: record number 8 is 103
examples.txt: record number 9 is 104
examples.txt: record number 10 is 105
examples.txt: record number 11 is 
Total number of records: 11 
```

## awk变量的操作符 

### 变量

> * Awk变量以字母开头, 后续字符可以是数字、字母、或下划线; 关键字不能用作awk变量, 和其他编程语言不同的是awk变量可以直接使用而不需事先声明; 如果要初始化变量, 最好在BEGIN区域内作,它只会执行一次; 
> * Awk中没有数据类型的概念, 一个awk变量是number还是string取决于该变量所处的上下文; 
> * 创建下面文件: 
```
$ cat example-sal.txt
101,John Doe,CEO,10000 
102,Jason Smith,IT Manager,5000 
103,Raj Reddy,Sysadmin,4500 
104,Anand Ram,Developer,4500 
105,Jane Miller,Sales Manager,3000
```

> * 下面的例子演示awk中创建和使用自己的变量, "total"便是用户建立的用来存储公司所有雇员工资总和的变量;
``` 
# awk脚本
$ cat total.awk                    
BEGIN{ 
    FS=","; 
    total=0; 
} 
{ 
    print $2 "'s salary is: " $4; 
    total=total+$4 
} 
END{ 
    print "---\nTotal company salary =$"total; 
} 

# 执行脚本
$ awk -f total.awk example-sal.txt 
John Doe's slary is: 10000 
Jason Smith's slary is: 5000 
Raj Reddy's slary is: 4500 
Anand Ram's slary is: 4500 
Jane Miller's slary is: 3000 
---
Total company salary =$27000
```

### 一元操作符

> * 只接受单个操作数的操作符叫做一元操作符
|操作符(Operator)|描述(Description)|
|:------|:------|
|+|取正(返回数字本身)|
|-|取反|
|++|自增|
|--|自减|

> * 下面的例子使用取反操作 
```
$ awk -F, '{print -$4}' example-sal.txt 
-10000
-5000
-4500
-4500
-3000
``` 
	
> * 下面的例子演示取正、取反操作符对文件中存放的复数的作用： 
```
$ cat negative.txt 
-1
-2
-3
$ awk '{print +$1}' negative.txt                               
-1
-2
-3
$ awk '{print -$1}' negative.txt                               
1
2
3 
```

> * 自增和自减操作 
 - 自增和自减改变变量的值, 它可以在使用变量"之前"或"之后"改变变量的值; 在表达式中, 使用的可能是改变前的值(post)或改变后的值(pre);
 - 使用改变后的变量值(pre)即是在变量前面加上++(或--), 首先把变量的值加1(或减1), 然后把改变后的值应用到表达式的其它操作中。;
 - 使用改变前的变量值(post)即是在变量后面加上++(或--), 首先把变量值应用到表达式中进行计算, 然后把变量的值加1(或减1);
```	
# Pre自增示例: 
$ awk -F, '{print ++$4}' example-sal.txt
10001
5001
4501
4501
3001 

# Pre自减示例: 
$ awk -F, '{print --$4}' example-sal.txt 
9999
4999
4499
4499
2999 

# Post自增示例: (因为++ 在print语句中,所以变量的原始值被打印) 
$ awk -F, '{print $4++}' example-sal.txt  
10000
5000
4500
4500
3000

# Post自减示例: (因为++是单独语句,所以自增后的值被打印) 
$ awk -F, '{$4++;print $4}' example-sal.txt 
10001
5001
4501
4501
3001 

# Post自减示例: (因为—在print语句中,所以变量原始值被打印) 
$awk -F, '{print $4--}' example-sal.txt 
10000
5000
4500
4500
3000 
	
# Post自减示例: (因为—在单独语句中,所以自减后的值被打印) 
$ awk -F, '{$4--;print $4}' example-sal.txt    
9999
4999
4499
4499
2999 
```

> * 下面例子, 显示所有登录到shell的用户数, 即哪些用户可以登录shell并获得命令提示符;
 - 使用算后自增运算符(尽管变量值只到END区域才打印出来,算前自增仍会产生同样的结果) 
 - 脚本的body区域包含一个模式匹配, 因此只有最后一个字段匹配模式/bin/bash时, body的代码才会执行 
 - 提示: 正则表达式应该包含在//之间, 但如果正则表达式中的/必须转移, 以避免被解析为正则表达式的结尾 
 - 当有匹配到模式的行时,变量n的值增加1,最终的值在END区域打印出来 
> * 打印所有可登陆shell的用户总数:
``` 
$ awk -F':' '$NF ~ /\/bin\/bash/{n++}END{print n}' /etc/passwd 
2 
```

### 算术操作符

> * 需要两个操作数的操作符, 成为二元操作符;
> * Awk中有多种基本二元操作符(如算术操作符、字符串操作符、赋值操作符,等等);

|操作符(Operator)|描述(Description)|
|:---|:---|
|+|加|
|-|减|
|*|乘|
|/|除|
|%|取模(取余)|
> * 下面的例子展示+,-,*,/的用法. 
 - 下面例子完成两个事情: 
    1. 将每件单独的商品价格减少20% 
    2. 将每件单独的商品的数量减少1
 - 创建并运行awk算术运算脚本:
```
$ cat compute.awk
BEGIN { 
        FS=","; 
        OFS=","; 
        item_discount=0; 
} 
 
{ 
        item_discount=$4*20/100; 
        print $1,$2,$3,$4-item_discount,$5-1; 
}

$ cat items.txt
101,HD Camcorder,Video,210,10 
102,Refrigerator,Appliance,850,2 
103,MP3 Player,Audio,270,15 
104,Tennis Racket,Sports,190,20 
105,Laser Printer,Office,475,5

$ awk -f compute.awk items.txt 
101,HD Camcorder,Video,168,9
102,Refrigerator,Appliance,680,1
103,MP3 Player,Audio,216,14
104,Tennis Racket,Sports,152,19
105,Laser Printer,Office,380,4 
```

> * 下面的例子只打印偶数行, 打印前会检查行号是否能被2整除, 如果整除, 则执行默认的操作(打印整行)
 - 取模运算演示:
```
$ awk 'NR % 2 == 0' items.txt 
102,Refrigerator,Appliance,850,2 
104,Tennis Racket,Sports,190,20
 
$ awk 'NR % 2 != 1' items.txt  
102,Refrigerator,Appliance,850,2 
104,Tennis Racket,Sports,190,20
```

### 字符串操作符

> * (空格)是连接字符串的操作符
> * 下面例子中, 有三处使用了字符串连接; 在语句"string3=string1 string2"中, string3包含了string1和string2连接后的内容, 每个print语句都把一个静态字符串和awk变量做了连接;
```
$ cat items.txt
101,HD Camcorder,Video,210,10 
102,Refrigerator,Appliance,850,2 
103,MP3 Player,Audio,270,15 
104,Tennis Racket,Sports,190,20 
105,Laser Printer,Office,475,5

$ cat string.awk 
BEGIN { 
        FS=","; 
        OFS=","; 
        string1="Audio"; 
        string2="Video"; 
        numberstring="100"; 
        string3=string1 string2; 
        print "Concatenate string is:" string3; 
        numberstring=numberstring+1; 
        print "String to number:" numberstring; 
}

$ awk -f string.awk items.txt 
Concatenate string is:AudioVideo
String to number:101 
```

### 赋值操作符

> * 同其他大部分编程语言一样, awk使用=作为赋值操作符; 和C语言一样, awk支持赋值的缩写方式
|操作符(Operator)|描述(Description)|
|:------|:------|
|=|赋值|
|+=|加法赋值的缩写|
|-=|减法赋值的缩写|
|*=|乘法赋值的缩写|
|/=|除法赋值的缩写|
|%=|取模赋值的缩写|

> * 下面的例子演示如何使用赋值: 
```
$ cat assignment.awk
BEGIN { 
        FS=","; 
        OFS=","; 
        total1 = total2 = total3 = total4 = total5 = 10; 
        total1 += 5; print total1; 
        total2 -= 5; print total2; 
        total3 *= 5; print total3; 
        total4 /= 5; print total4; 
        total5 %= 5; print total5; 
} 

$ awk -f assignment.awk 
15
5
50
2
0 
```

> * 下面的例子使用加法赋值的缩写形式 
 - 显示所有商品的清单:
``` 
$ awk -F',' 'BEGIN{total=0}{total+=$5}END{print "Total Quantity: "total}' items.txt             
Total Quantity: 52 
```

> * 下面的例子统计输入文件中所有的字段数, Awk读取每一行, 并把字段数量增加到变量total中, 然后在END区域打印该变量; 
```
$ awk -F',' 'BEGIN {total=0}{total+=NF}END{print total}' items.txt
25 
```

### 比较操作符

> * Awk支持下面标准比较操作符

|操作符(Operator)|描述(Description)|
|:------|:------|
|>|大于|
|\>=|大于等于|
|<|小于|
|<=|小于等于|
|==|等于|
|!=|不等于|
|&&|且(and)|
|\|\||或(or)|

> * 提示: 下面的例子,如果不指定操作,awk会打印符合条件的整条记录。
> * 打印数量小于等于临界值5的商品信息:
``` 
$ awk -F',' '$5 <= 5' items.txt  
102,Refrigerator,Appliance,850,2 
105,Laser Printer,Office,475,5 
```

> * 打印编号为103的商品信息: 
```
$ awk -F "," '$1 == 103' items.txt 
103,MP3 Player,Audio,270,15
提示:不要把==(等于)和=(赋值)搞混了。
```

> * 打印编号为103的商品的描述信息: 
```
$ awk -F "," '$1 == 103 { print $2}' items.txt 
MP3 Player 
```

> * 打印除Video以外的所有商品:
``` 
$ awk -F "," '$3 != "Video"' items.txt
102,Refrigerator,Appliance,850,2 
103,MP3 Player,Audio,270,15 
104,Tennis Racket,Sports,190,20 
105,Laser Printer,Office,475,5 
```
	
> * 和上面相同, 但只打印商品描述信息:
``` 
$ awk -F "," '$3 != "Video"{ print $2}' items.txt  
Refrigerator
MP3 Player
Tennis Racket
Laser Printer 
```

> * 使用&&比较两个条件, 打印价钱低于900并且数量小于等于临界值5的商品信息:
``` 
$ awk -F "," '$4 < 900 && $5 <= 5' items.txt 
102,Refrigerator,Appliance,850,2 
105,Laser Printer,Office,475,5 
```

> * 和上面相同, 但只打印商品描述信息
``` 
$ awk -F "," '$4 < 900 && $5 <= 5 {print $2}' items.txt  
Refrigerator
Laser Printer 
```
	
> * 使用||比较两个条件, 打印价钱低于900或者数量小于等于临界值5的商品信息:
``` 
$ awk -F "," '$4 < 900 || $5 <= 5' items.txt 
101,HD Camcorder,Video,210,10 
102,Refrigerator,Appliance,850,2 
103,MP3 Player,Audio,270,15 
104,Tennis Racket,Sports,190,20 
105,Laser Printer,Office,475,5 
```
	
> * 和上面相同, 但只打印商品描述信息:
```
$ awk -F "," '$4 < 900 || $5 <= 5 {print $2}' items.txt
HD Camcorder
Refrigerator
MP3 Player
Tennis Racket
Laser Printer  
```
	
> * 下面例子使用>条件, 打印/etc/password中最大的UID(以及其所在的整行); Awk把最大的UID(第3个字段)放在变量maxuid中, 并且把包含最大UID的行复制到变量maxline中; 循环执行完后,打印最大的UID和其所在的行:
``` 
$ awk -F':' '$3>maxuid{maxuid=$3;maxline=$0}END{print maxuid,maxline}' /etc/passwd
501 rsync:x:501:501::/home/rsync:/sbin/nologin 
```

> * 打印/etc/passwd中UID和GROUP ID相同的用户信息 
```
$ awk -F':' '$3 == $4' /etc/passwd 
```

> * 打印/etc/passwd中UID >= 100 并且用户的shell是/bin/sh的用户: 
```
$awk -F':' '$3 >= 100 && $NF~/\/bin\/bash/' /etc/passwd 
cluster:x:500:500::/home/cluster:/bin/bash 
```

> * 打印/etc/passwd中没有注释信息(第5个字段)的用户: 
```
$ awk -F':' '$5 == ""' /etc/passwd 
abrt:x:173:173::/etc/abrt:/sbin/nologin
ntp:x:38:38::/etc/ntp:/sbin/nologin
postfix:x:89:89::/var/spool/postfix:/sbin/nologin
tcpdump:x:72:72::/:/sbin/nologin
cluster:x:500:500::/home/cluster:/bin/bash
rsync:x:501:501::/home/rsync:/sbin/nologin 
```

### 正则表达式操作符 

|操作符(Operator)|描述(Description)|
|:------|:------|
|~ |匹配|
|!~ |不匹配 |
> * 使用==时, awk检查精确匹配; 下面的例子不会打印任何信息, 因为items.txt中, 没有任何一条记录的第二个字段精确匹配关键字"Tennis","Tennis Racket"不是精确匹配:
> * 打印第二个字段为"Tennis"的记录: 
```
$ awk -F "," '$2 == "Tennis"' items.txt 
```

> * 当使用~时, awk执行模糊匹配,检索"包含"关键字的记录; 
 - 打印第二个字段包含”Tennis”的行: 
```
$ awk -F "," '$2~ "Tennis"' items.txt     
104,Tennis Racket,Sports,190,20  
```
> * !~ 对应 ~, 即不匹配;
 - 打印第二个字段不包含"Tennis"的行:
``` 
$ awk -F "," '$2 !~ "Tennis"' items.txt 
101,HD Camcorder,Video,210,10 
102,Refrigerator,Appliance,850,2 
103,MP3 Player,Audio,270,15 
105,Laser Printer,Office,475,5
```
	
> * 下面的例子打印shell为/bin/bash的用户的总数, 如果最后一个字段包含"/bin/bash",则变量n增加1;
``` 
$ awk -F':' '$NF~/\/bin\/bash/{ a++ }END{print a}' /etc/passwd
2 
```

## awk分支和循环

> * Awk支持条件判断,控制程序流程。Awk的大部分条件判断语句很像C语言的语法。 
> * Awk支持下面三种if语句. 
 - 单个if语句 
 - If-else语句 
 - 多级If-elseif语句 

### if结构 

> * 单个if语句检测条件, 如果条件为真, 执行相关的语句;
 - 单条语句, 语法:
```  	 
if(conditional-expression ) 
action 

if是关键字 
conditional-expression是要检测的条件表达式 
action是要执行的语句 
```
> * 多条语句 
 - 如果要执行多条语句,需要把他们放在{ } 中,每个语句之间必须用分号或换行符分开, 语法如下;
   if (conditional-expression) 
    { 
     action1; 
     action2; 
    } 
 - 如果条件为真, { }中的语句会依次执行; 当所有语句执行完后, awk会继续执行后面的语句;

> * 打印数量小于等于5的所有商品;
``` 
$ cat items.txt
101,HD Camcorder,Video,210,10 
102,Refrigerator,Appliance,850,2 
103,MP3 Player,Audio,270,15 
104,Tennis Racket,Sports,190,20 
105,Laser Printer,Office,475,5

$ awk -F',' '{if($5<= 5) print "Only",$5,"qty of",$2 "is available"}' items.txt      
Only 2  qty of Refrigeratoris available
Only 5  qty of Laser Printeris available
```

> * 使用多个条件, 可以打印价钱在500至100, 并且总数不超过5的商品
```
$ awk -F',' '{if(($4 >= 500 && $4<=1000)&&($5<= 5))print "Only",$5,"qty of",$2,"is available"}' items.txt         
Only 2  qty of Refrigerator is available 
```

### if else 结构

> * 在if else结构中, 还可以指定判断条件为false时要执行的语句; 下面的语法中, 如果条件为true, 那么执行action1, 如果条件为false,则执行action2;
``` 
# 语法: 
if (conditional-expression) 
 action1 
else 
 action2 
``` 
> * awk还有个条件操作符( ? : )和C语言的三元操作符等价。 
 - 和if-else结构相同,如果codintional-expresion是true,执行action1,否则执行action2 
 - 三元操作符: 
   codintional-expression ? action1 : action2; 

> * 如果商品数量不大于5, 打印"Buy More", 否则打印商品数量; 
```
$ cat if-else.awk 
BEGIN { 
        FS=","; 
} 
{ 
        if( $5 <= 5) 
                print "Buy More: Order",$2,"immediately!" 
        else 
                print "Shell More: Give discount on",$2,"immediately!" 
} 

$ cat items.txt
101,HD Camcorder,Video,210,10 
102,Refrigerator,Appliance,850,2 
103,MP3 Player,Audio,270,15 
104,Tennis Racket,Sports,190,20 
105,Laser Printer,Office,475,5

$ awk -f if-else.awk items.txt 
Shell More: Give discount on HD Camcorder immediately!
Buy More: Order Refrigerator immediately!
Shell More: Give discount on MP3 Player immediately!
Shell More: Give discount on Tennis Racket immediately!
Buy More: Order Laser Printer immediately! 
```

> * 下面的例子, 使用三元操作符, 把items.txt文件中的每两行都以逗号分隔合并起来;
``` 
$ awk 'ORS=NR%2?",":"\n"' items.txt 
101,HD Camcorder,Video,210,10 ,102,Refrigerator,Appliance,850,2 
103,MP3 Player,Audio,270,15 ,104,Tennis Racket,Sports,190,20 
105,Laser Printer,Office,475,5 ,[root@kvm ~]#

#复杂的就是: awk '{if(NR%2==0) ORS="\n"; else ORS=",";print}' items.txt 
```

### while循环

> * Awk循环用例执行一系列需要重复执行的动作, 只要循环条件为true, 就一直保持循环; awk支持多种循环结构; 
> * 首先是while循环结构,语法如下:
```
while (codition) 
 Actions 

while是awk的关键字 
condition是条件表达式 
actions是循环体,如果有多条语句,必须放在{ }中 
	
while首先检查condtion,如果是true,执行actions,执行完后,再次检查condition,如果是true,再次执行actions,直到condition为false时,退出循环。 
 注意,如果第一次执行前condition返回false,那么所有actions都不会被执行。 
```

> * 下面的例子中, BEGIN区域中的语句会先于其他awk语句执行; While循环将50个'x'追加到变量string中; 每次循环都检查变量count, 如果其小于50, 则执行追加操作; 因此循环体会执行50次, 之后, 变量string的值被打印出来;
``` 
$ awk 'BEGIN{while(count++<50) string=string"x";print string}'  
xxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxx
```

> * 下面的例子计算items-sold.txt文件中, 每件商品出售的总数; 
 - 对于每条记录, 需要把从第2至第7个字段累加起来(第一个字段是编号,因此不用累加); 所以当循环条件从第2个字段(因为while之前设置了i=2)开始, 检查是否到达最后一个字段(i<=NF); NF代表每条记录的字段数;
``` 
$ cat while.awk                   
{ 
        i=2; total=0; 
        while (i <= NF ){ 
                total = total + $i; 
                i++; 
        } 
        print "Item",$1,":",total,"quantities sold"; 
} 

$awk -f while.awk items-sold.txt      
Item 101 : 47 quantities sold
Item 102 : 10 quantities sold
Item 103 : 65 quantities sold
Item 104 : 20 quantities sold
Item 105 : 42 quantities sold 
```

### do-while循环 

> * While循环是一种进入控制判断的循环结构, 因为在进入循环体前执行判断; 而do-while循环是一种退出控制循环, 在退出循环前执行判断; do-while循环至少会执行一次, 如果条件为true, 它将一直执行下去;
``` 
# 语法: 
do 
action 
while(condition) 
```
> * 下面的例子中, print语句仅会执行一次, 因为我们已经确认判断条件为false; 如果这是一个while结构, 在相同的初始化条件下, print一次都不会执行;
``` 
$ awk 'BEGIN{ 
count=1; 
do 
print "This gets printed at least once"; 
while(count!=1) 
}'
This gets printed at least once
```

> * 下面打印items-sold.txt文件中, 每种商品的总销售量,输出结果和之前的while.awk脚本相同, 不同的是, 这次试用do-while结构来实现;
``` 
$ cat dowhile.awk 
{ 
  i=2;   
  total=0; 
  do {   
    total = total + $i;  
    i++;         
    }            
  while(i<=NF) 
print "Item",$1,":",total,"quantities sold"; 
} 

$ awk -f dowhile.awk items-sold.txt 
Item 101 : 47 quantities sold
Item 102 : 10 quantities sold
Item 103 : 65 quantities sold
Item 104 : 20 quantities sold
Item 105 : 42 quantities sold 
```

### for循环

> * Awk的for循环和while循环一样实用, 但语法更简洁易用;
```	
# 语法： 
for(initialization;condition;increment/decrement) 
    for循环一开始就执行initialization, 然后检查condition, 如果condition为true, 执行actions, 然后执行increment 或 decrement; 如果condition为true,就会一直重复执行actions和increment/decrement。 
```

> * 下面例子打印文件总字段数总和; i的初始值为1, 如果i小于等于字段数, 则当前字段会被追加到总数中, 每次循环i的值会增加1;
```
$ echo "1 2 3 4" |awk '{for(i=1;i<=NF;i++)total=total+$i }END{print total}'             
10
```

> * 下面的例子, 使用for循环把文件中的字段反序打印出来; 注意这次在for中使用decrement而不是increment; 
 - 提示: 每读入一行数据后,awk会把NF的值设置为当前记录的总字段数。 
 - 该例用相反的顺序,从最后一个自动开始到第一个字段,逐个输出每个字段,然后输出换行。 
 - 反转示例: 
```
$ cat forreverse.awk
BEGIN { 
  ORS=""; 
} 
{ 
  for (i=NF;i>0;i--) 
  print $i," " 
  print "\n";  
} 

$ awk -f forreverse.awk items-sold.txt 
12  10  8  5  10  2  101  
2  0  3  4  1  0  102  
13  5  20  11  6  10  103  
5  6  0  4  3  2  104  
6  12  7  5  2  10  105   
```

> * 使用for循环输出了items-sold.txt中每种商品的销售量;
``` 
$ cat for.awk 
{ 
  total=0; 
  for(i=2;i<=NF;i++) 
    total = total + $i 
  print "Item ",$1," : ",total," quantities sold" 
} 

$ awk -f for.awk items-sold.txt 
Item  101  :  47  quantities sold
Item  102  :  10  quantities sold
Item  103  :  65  quantities sold
Item  104  :  20  quantities sold
Item  105  :  42  quantities sold 
```

### break语句

> * Break语句用来跳出它所在的最内层的循环(while,do-while,或for循环); 请注意,break语句只有在循环中才能使用; 
> * 打印某个月销售量为0的任何商品,即从第2至第7个字段中出现0的记录;
```
$ cat break.awk 
{ 
  i=2;total=0; 
  while(i++<=NF) 
  { 
    if($i == 0) 
    { 
      print "Item",$0,"had a month without item sold" 
      break 
    } 
  } 
} 

$ awk -f break.awk items-sold.txt 
Item 102 0 1 4 3 0 2  had a month without item sold
Item 104 2 3 4 0 6 5  had a month without item sold
```

> *  如果执行下面的命令, 要按Ctrl+c才能停止;
```
$ awk 'BEGIN{while(1) print "forever"}' 
#这条语句一直打印字符串”forever”,因为条件永远不为false。尽管死循环会用于操作系统和进程控制,但通常不是什么好事
``` 
 
> * 下面我们修改这个死循环,让它执行10次后退出;
``` 
$ awk 'BEGIN{ 
x=1; 
while(1) 
{ 
print "Iteration" 
if( x==10) 
break; 
x++ 
}}'
其输出结果如下: 
Iteration 
Iteration 
Iteration 
Iteration 
Iteration 
Iteration 
Iteration 
Iteration 
Iteration 
Iteration 
```

### continue语句

> * Continue语句跳过后面剩余的循环部分, 立即进入下次循环;  
> * 下面打印items-sold.txt文件中所有商品的总销售量; 其输出结果和while.awk、dowhile.awk以及for.awk一样, 但是这里的while循环中使用contine, 使循环从1而不是从2开始;
``` 
$ cat continue.awk 
{ 
  i=1;total=0; 
  while(i++<=NF) 
  { 
    if(i==1) 
    continue 
    total = total + $i 
  } 
  print "Item",$1,":",total,"quantities sold" 
} 
 
$ awk -f continue.awk items-sold.txt 
Item 101 : 47 quantities sold 
Item 102 : 10 quantities sold 
Item 103 : 65 quantities sold 
Item 104 : 20 quantities sold 
Item 105 : 42 quantities sold 
```

> * 下面的脚本在每次循环时都打印x的值, 除了第5次循环, 因为continue导致跳打印语句; 
```
$ awk 'BEGIN{ x=1; while(x<=10) { if(x==5) { x++; continue } print "Value of x:",x;x++; } }'
输出结果如下: 
Value of x: 1 
Value of x: 2 
Value of x: 3 
Value of x: 4 
Value of x: 6 
Value of x: 7 
Value of x: 8 
Value of x: 9 
Value of x: 10 
```

### exit语句

> * exit命令立即停止脚本的运行 ,并忽略脚本中其余的命令;
> * exit命令接受一个数字参数最为awk的退出状态码, 如果不提供参数, 默认的状态码是0;
> * 下面的脚本执行到第5次循环时退出, 因为print命令位于exit之后, 所以输出的值只到4为止, 到第5次循环时就退出了;
``` 
$ awk 'BEGIN{ x=1; while(x<=10) { if(x==5) { x++ ;exit } print "Value of x:",x;x++; } }'
其输出结果如下: 
Value of x: 1 
Value of x: 2 
Value of x: 3 
Value of x: 4 
```

> * 下面例子打印第一次出现的有个月没有卖出一件的商品的信息, 和break.awk脚本很相似; 
 - 区别在于, 遇到某月为出售的商品时, 退出脚本, 而不是继续执行;
``` 
$ cat exit.awk 
{ 
  i=2;total=0; 
  while(i++<=NF) 
  if($i==0) { 
    print "Item",$1,"had a month with no item sold" 
    exit 
  } 
} 
 
$ awk -f exit.awk items-sold.txt 
Item 102 had a month with no item sold   
提示：104号商品有的月份也没有卖出一件,但是并没有被打印,因为我们在循环中使用了exit命令。
```

## awk关联数组

> * 相比较与其他编程语言中的传统数组, awk的数组更为强大;
> * Awk的数组都是关联数组, 即一个数组包含多个"索引/值"的元素; 索引没必要是一系列连续的数字, 实际上它可以是字符串或者数字, 并且不需要指定数组长度；
> * 语法如下:
```
arrayname[string]=value 
   arrayname是数组名称 
   string是数组索引 
   value是为数组元素赋的值 
访问awk数组的元素 
   如果要访问数组中的某个特定元素,使用arrayname[index] 即可返回该索引中的值;
```

> * 一个简单的数组赋值示例:
```
$ cat array-assign.awk 
BEGIN { 
  item[101]="HD Camcorder"; 
  item[102]="Refrigerator"; 
  item[103]="MP3 Player"; 
  item[104]="Tennis Racket"; 
  item[105]="Laser Printer"; 
  item[1001]="Tennis Ball"; 
  item[55]="Laptop"; 
  item["na"]="Not Available"; 
  print item["101"]; 
  print item[102]; 
  print item["103"]; 
  print item[104]; 
  print item["105"]; 
  print item[1001]; 
  print item["na"]; 
} 
 
$ awk -f array-assign.awk 
HD Camcorder 
Refrigerator 
MP3 Player 
Tennis Racket 
Laser Printer 
Tennis Ball 
Not Available 
 
#请注意: 
  数组索引没有顺序,甚至没有从0或1开始,而是直接从101….105开始,然后直接跳到1001,又降到55,还有一个字符串索引”na”. 
  数组索引可以是字符串,数组的最后一个元素就是字符串索引,即”na” 
  Awk中在使用数组前,不需要初始化甚至定义数组,也不需要指定数组的长度。 
  Awk数组的命名规范和awk变量命名规范相同。 

# 从awk的角度来说, 数组的索引通常是字符串, 即是你使用数组作为索引, awk也会当做字符串来处理; 下面的写法是等价的: 
   Item[101]="HD Camcorder" 
   Item[“101”]= "HD Camcorder" 
```

### 引用数组元素

> * 可以使用print命令直接打印数组的元素, 也可以把元素的值赋给其他变量以便后续使用; 
 - print item[101] 
 - x=item[105] 
> * 如果试图访问一个不存在的数组元素, awk会自动以访问时指定的索引建立该元素, 并赋予null值; 为了避免这种情况,在使用前应该检测元素是否存在; 
 - 使用if语句可以检测元素是否存在, 如果返回true, 说明改元素存在于数组中; 
   if ( index in array-name ) 
> * 一个简单的引用数组元素的例子: 
```
$ cat array-refer.awk 
BEGIN { 
  x = item[55]; 
  if ( 55 in item ) 
    print "Array index 55 contains",item[55]; 
  item[101]="HD Camcorder"; 
  if ( 101 in item ) 
    print "Array index 101 contains",item[101]; 
  if ( 1010 in item ) 
    print "Array index 1010 contains",item[1010]; 
} 
 
$ awk -f array-refer.awk 
Array index 55 contains 
Array index 101 contains HD Camcorder 
# 该例中： 
   Item[55] 在引用前没有赋任何值, 所以在引用是awk自动创建该元素并赋null值 
   Item[101]是一个已赋值的元素, 所以在检查索引值时返回true, 打印该元素 
   Item[1010]不存在, 因此检查索引值时,返回false,不会被打印 
```

### 使用循环遍历awk数组

> * 如果要访问数组中的所有元素, 可以使用for的一个特殊用法来遍历数组的所有索引; 
> * 语法: 
```
for ( var in arrayname ) 
actions 
 
var 是变量名称 
in是关键字 
arrayname是数组名 
actions是一系列要执行的awk语句, 如果有多条语句, 必须包含在{ }中; 通过把索引值赋给变量var, 循环体可以把所有语句应用到数组中所有的元素上; 
```
> * 在示例"for (x in item)"中, x是变量名, 用来存放数组索引; 
> * 我们并没有指定循环执行的条件, 实际上我们不必关心数组中有多少个元素, 因为awk会自动判断, 在循环结束前遍历所有元素; 
 > * 下面的例子遍历数组中所有元素并打印出来;
``` 
$ cat array-for-loop.awk 
BEGIN { 
  item[101]="HD Camcorder"; 
  item[102]="Refrigerator"; 
  item[103]="MP3 Player"; 
  item[104]="Tennis Racket"; 
  item[105]="Laser Printer"; 
  item[1001]="Tennis Ball"; 
  item[55]="Laptop"; 
  item["no"]="Not Available"; 
 
  for(x in item) 
  print item[x] 
} 
 
$ awk -f array-for-loop.awk  
Not Available 
Laptop 
HD Camcorder 
Refrigerator 
MP3 Player 
Tennis Racket 
Laser Printer 
Tennis Ball 
```

### 删除数组元素

> * 如果要删除特定的数组元素,使用delete语句。一旦删除了某个元素,就再也获取不到它的值了。 
> * 语法:
``` 
delete arrayname[index]; 
```
> * 删除数组内所有元素
``` 
for (var in array) 
 delete array[var] 
```
> * 在GAWK中, 可以使用单个delete命令来删除数组的所有元素:
``` 
 Delete array 
```
> * 下面例子中, item[103]=””并没有删除整个元素, 仅仅是给它赋了null值;
```
$ cat array-delete.awk 
BEGIN { 
  item[101]="HD Camcorder"; 
  item[102]="Refrigerator"; 
  item[103]="MP3 Player"; 
  item[104]="Tennis Racket"; 
  item[105]="Laser Printer"; 
  item[1001]="Tennis Ball"; 
  item[55]="Laptop"; 
  item["no"]="Not Available"; 
 
  delete item[102] 
  item[103]="" 
  delete item[104] 
  delete item[1001] 
  delete item["na"] 
 
  for(x in item)  
  print "Index",x,"contains",item[x] 
} 
 
$ awk -f array-delete.awk  
Index no contains Not Available 
Index 55 contains Laptop 
Index 101 contains HD Camcorder 
Index 103 contains  
Index 105 contains Laser Printer 
```

### 多维数组

> * 虽然awk只支持一维数组, 但其奇妙之处在于, 可以使用一维数组来模拟多维数组; 
> * 假定要创建下面的2X2维数组: 
 - 10 20 
 - 30 40 
   其中位于"1,1"的元素是10, 位于"1,2"的元素是20, 等等…,下面把10赋值给"1,1"的元素: 
 item[“1,1”]=10 
  即使使用了”1,1"作为索引值, 它也不是两个索引, 仍然是单个字符串索引, 值为"1,1"; 所以上面的写法中, 实际上是把10赋给一维数组中索引"1,1"代表的值; 
```
$ cat array-multi.awk 
BEGIN { 
  item["1,1"]=10; 
  item["1,2"]=20; 
  item["2,1"]=30; 
  item["2,2"]=40 
 
  for (x in item)  
  print item[x] 
} 
 
$ awk -f array-multi.awk  
30 
20 
40 
10 
```
> * 现在把索引外面的引号去掉,会发生什么情况？`即item[1,1](而不是item[“1,1”])`:
``` 
$ cat array-multi2.awk  
BEGIN { 
  item[1,1]=10; 
  item[1,2]=20; 
  item[2,1]=30; 
  item[2,2]=40 
 
  for (x in item) 
  print item[x] 
} 
$ awk -f array-multi2.awk  
10 
30 
20 
40
# 上面的例子仍然可以运行, 但是结果有所不同; 在多维数组中, 如果没有把下标用引号引住, awk会使用"\034"作为下标分隔符;  
```
> * 当指定元素item[1,2]时,它会被转换为item[“1\0342”]。Awk用把两个下标用”\034”连接起来并转换为字符串;
> * 当使用["1,2"]时, 则不会使用"\034", 它会被当做一维数组; 
```
# 如下示例 
$ cat array-multi3.awk 
BEGIN { 
  item["1,1"]=10; 
  item["1,2"]=20; 
  item[2,1]=30; 
  item[2,2]=40; 
 
  for(x in item)  
  print "Index",x,"contains",item[x]; 
} 
 
$ awk -f array-multi3.awk  
Index 1,2 contains 20 
Index 21 contains 30 
Index 22 contains 40 
Index 1,1 contains 10 

索引"1,1"和"1,2"放在了引号中, 所以被当做一维数组索引, awk没有使用下标分隔符, 因此索引值被原封不动地输出; 
索引2,1和2,2没有放在引号中, 所以被当做多维数组索引, awk使用下标分隔符来处理, 因此索引变成"2\0341"和"2\0342", 于是在两个下标直接输出了非打印字符"\034" 
```

###  SUBSEP下标分隔符

> * 通过变量SUBSEP可以把默认的下标分隔符改成任意字符, 下面例子中, SUBSEP被改成了分号;
``` 
$ cat array-multi4.awk  
BEGIN { 
  SUBSEP=":"; 
  item["1,1"]=10; 
  item["1,2"]=20; 
 
  item[2,1]=30; 
  item[2,2]=40; 
 
  for(x in item) 
  print "Index",x,"contains",item[x]; 
} 

$ awk -f array-multi4.awk  
Index 1,2 contains 20 
Index 2:1 contains 30 
Index 2:2 contains 40 
Index 1,1 contains 10 
# 这个例子中,索引"1,1"和"1,2"由于放在了引号中而没有使用SUBSEP变量; 
 
# 所以, 使用多维数组时, 最好不要给索引值加引号,如: 
$ cat array-multi5.awk  
BEGIN { 
  SUBSEP=":"; 
 
  item[1,1]=10; 
  item[1,2]=20; 
 
  item[2,1]=30; 
  item[2,2]=40; 
 
  for(x in item) 
  print "Index",x,"contains",item[x]; 
} 
 
$ awk -f array-multi5.awk  
Index 1:1 contains 10 
Index 2:1 contains 30 
Index 1:2 contains 20 
Index 2:2 contains 40 
```

### 用asort为数组排序

> * asort函数重新为元素值排序, 并且把索引重置为从1到n的值, 此处n代表数组元素个数; 
> * 假定一个数组有两个元素: item["something"]="B - I'm big b"和item["notsure"]="A – I'm big a" 
 - 调用asort函数手, 数组会以元素值排序, 变成:item[1]="A – I'm big a" 和 item[2]="B – I'm big b" 
> * 下面例子中, 数组索引是非连续的数字和字符串, 调用asort后, 元素值被排序, 并且索引值变成1,2,3,4…. 请注意, asort函数会返回数组元素的个数;
``` 
$ cat asort.awk 
BEGIN { 
  item[101]="HD Camcorder"; 
  item[102]="Refrigerator"; 
  item[103]="MP3 Player"; 
  item[104]="Tennis Racket"; 
  item[105]="Laser Printer"; 
  item[1001]="Tennis Ball"; 
  item[55]="Laptop"; 
  item["na"]="Not Available"; 
 
  print "---------- Before asort -------------" 
  for(x in item) 
  print "Index",x,"contains",item[x] 
  total = asort(item); 
 
  print "---------- After asort -------------" 
  for(x in item) 
  print "Index",x,"contains",item[x] 
  print "Return value from asort:",total; 
} 
 
$ awk -f asort.awk  
---------- Before asort ------------- 
Index 55 contains Laptop 
Index 101 contains HD Camcorder 
Index 102 contains Refrigerator 
Index 103 contains MP3 Player 
Index 104 contains Tennis Racket 
Index 105 contains Laser Printer 
Index na contains Not Available 
Index 1001 contains Tennis Ball 
---------- After asort ------------- 
Index 4 contains MP3 Player 
Index 5 contains Not Available 
Index 6 contains Refrigerator 
Index 7 contains Tennis Ball 
Index 8 contains Tennis Racket 
Index 1 contains HD Camcorder 
Index 2 contains Laptop 
Index 3 contains Laser Printer 
Return value from asort: 8 
```
> * 上面例子中, asort之后, 数组打印顺序不是按索引值从1到8, 而是随机的; 可以用下面的方法, 按索引值顺序打印: 
```
$ cat asort1.awk 
BEGIN { 
  item[101]="HD Camcorder"; 
  item[102]="Refrigerator"; 
  item[103]="MP3 Player"; 
  item[104]="Tennis Racket"; 
  item[105]="Laser Printer"; 
  item[1001]="Tennis Ball"; 
  item[55]="Laptop"; 
  item["na"]="Not Available"; 
  total = asort(item); 
 
  for(i=1;i<=total;i++) 
  print "Index",i,"contains",item[i] 
} 
 
$ awk -f asort1.awk  
Index 1 contains HD Camcorder 
Index 2 contains Laptop 
Index 3 contains Laser Printer 
Index 4 contains MP3 Player 
Index 5 contains Not Available 
Index 6 contains Refrigerator 
Index 7 contains Tennis Ball 
Index 8 contains Tennis Racket 

# 一旦调用asort函数,数组原始的索引值就不复存在了; 因此你可能想在不改变原有数组索引的情况下,使用新的索引值创建一个新的数组; 
```
> * 下面的例子中, 原始数组item不会被修改, 相反使用排序后的新索引值创建新数组itemnew;
 - 即itemnew[1],itemnew[2],itemnew[3],等等。 
 - total = asort(item,itemnew); 
 - **asort函数按元素值排序,但排序后使用从1开始的新索引值,原先的索引被覆盖掉了;** 

### 用asorti为索引排序

> * 和以元素值排序相似, 也可以取出所有索引值, 排序, 然后把他们保存在新数组中; 
> * 下面的例子展示了asort和asorti的不同, 请牢记下面两点: 
 - asorti函数为索引值(不是元素值)排序,并且把排序后的元素值当做元素值保存。 
 - 如果使用asorti(state)将会丢失原始元素值,即索引值变成了元素值;
   通常方式是给asorti传递两个参数, 即asorti(state,statebbr), 这样一来, 原始数组state就不会被覆盖了; 
```
$ cat asorti.awk 
BEGIN { 
  state["TX"]="Texas"; 
  state["PA"]="Pennsylvania"; 
  state["NV"]="Nevada"; 
  state["CA"]="California"; 
  state["AL"]="Alabama"; 
 
  print "-------------- Function: asort -----------------" 
  total = asort(state,statedesc); 
  for(i=1;i<=total;i++) 
  print "Index",i,"contains",statedesc[i]; 
 
  print "-------------- Function: asorti -----------------" 
  total = asorti(state,stateabbr); 
  for(i=1;i<=total;i++) 
  print "Index",i,"contains",stateabbr[i]; 
} 
 
$ awk -f asorti.awk  
-------------- Function: asort ----------------- 
Index 1 contains Alabama 
Index 2 contains California 
Index 3 contains Nevada 
Index 4 contains Pennsylvania 
Index 5 contains Texas 
-------------- Function: asorti ----------------- 
Index 1 contains AL 
Index 2 contains CA 
Index 3 contains NV 
Index 4 contains PA 
Index 5 contains TX 
```

## awk命令 

### 使用printf格式化输出

> * printf可以非常灵活、简单地以你期望的格式输出结果; 
> * 语法: 
 - printf "print format", variable1,variable2,etc. 
> * printf中可以使用下面的特殊字符
 
|特殊字符|描述|
|:------|:------|
|\n|换行| 
|\t|制表符| 
|\v|垂直制表符| 
|\b|退格| 
|\r|回车符| 
|\f|换页| 
> * 使用换行符把Line1和Line2打印在单独的行里: 
```
$ awk'BEGIN {printf "Line 1\nLine 2\n"}'
Line 1 
Line 2 
``` 
> * 以制表符分隔字段,Field 1后面有两个制表符: 
```
$ awk'BEGIN { printf "Field 1\t\tField 2\tField 3\tField 4\n"}'
Field 1         Field 2 Field 3 Field 4 
``` 
> * 每个字段后面使用垂直制表符: 
```
$ awk'BEGIN { printf "Field 1\vField 2\vField 3\vField 4\n" }'
Field 1 
       Field 2 
              Field 3 
                     Field 4 
```
> * 下面的例子中, 除了第4个字段外, 每个字段后面使用退格符, 这会擦除前三个字段最后的数字;如”Field 1”会被显示为"Field", 因为最后一个字符被退格符擦除了; 然而"Field 4"会照旧输出, 因为它后面没有使用\b;
```
$ awk'BEGIN { printf "Field 1\bField 2\bField 3\bField 4\n" }'
Field Field Field Field 4 
```
> * 下面的例子,打印每个字段后, 执行一个"回车", 在当前打印的字段的基础上, 打印下一个字段; 这就意味着, 最后只能看到"Field 4",因为其他的字段都被覆盖掉了;
``` 
$ awk'BEGIN { printf "Field 1\rField 2\rField 3\rField 4\n" }'
Field 4 
```
> * 使用OFS,ORS 
 - 当使用print(不是printf)打印多个以逗号分隔的字段时,awk默认会使用内置变量OFS和ORS处理输出;
 - 下面例子展示OFS和ORS对单个print的影响: 
```
$ cat print.awk  
BEGIN { 
  FS=","; 
  OFS=":"; 

  ORS="\n--\n"; 
} 
{ 
  print $2,$3 
} 
 
$ awk -f print.awk items.txt  
HD Camcorder:Video 
-- 
Refrigerator:Appliance 
-- 
MP3 Player:Audio 
-- 
Tennis Racket:Sports 
-- 
Laser Printer:Office 
-- 
```
> * Printf不受OFS,ORS影响 
 - printf不会使用OFS和ORS, 它只根据"format"里面的格式打印数据, 如下所示: 
```
$ cat printf1.awk  
BEGIN { 
  FS=","; 
  OFS=":"; 
  ORS="\n--\n"; 
} 
{ 
   printf "%s^^%s\n",$2,$3 
} 
 
$ awk -f printf1.awk  items.txt  
HD Camcorder^^Video 
Refrigerator^^Appliance 
MP3 Player^^Audio 
Tennis Racket^^Sports 
Laser Printer^^Office 
```
> * printf格式化字符 

|格式化字符|描述|
|:------|:------|
|s|字符|
|c|单个字符| 
|d|数值| 
|e|指数| 
|f|浮点数| 
|g|根据值决定使用e或f中较短的输出| 
|o|八进制| 
|x|十六进制| 
|%|百分号| 

> * 下面展示各个格式化字符的基本用法:
``` 
$ cat printf-format.awk 
BEGIN { 
  printf "s--> %s\n", "String" 
  printf "c--> %c\n", "String" 
  printf "s--> %s\n", 101.23 
  printf "d--> %d\n", 101,23 

  printf "e--> %e\n", 101,23 
  printf "f--> %f\n", 101,23 
  printf "g--> %g\n", 101,23 
  printf "o--> %o\n", 0x8 
  printf "x--> %x\n", 16 
  printf "percentage--> %%\n", 17 
} 
 
$ awk -f printf-format.awk  
s--> String 
c--> S 
s--> 101.23 
d--> 101 
e--> 1.010000e+02 
f--> 101.000000 
g--> 101 
o--> 10 
x--> 10 
percentage--> % 
```
> * 指定打印列的宽度 
 - 要指定打印列的宽度, 必须在%和格式化字符之间设置一个数字; 该数字代表输出列的最小宽度, 如果字符串的宽度比该数字小,会在输出列左侧加上空格以凑足该宽度; 
 - 下面例子演示如何指定输出列宽度:
```
$ cat printf-width.awk 
BEGIN { 
  FS="," 
  printf "%3s\t%10s\t%10s\t%5s\t%3s\n", "Num","Description","Type","Price","Qty" 
  printf "------------------------------------------------------------------\n" 
} 
{ 
  printf "%3d\t%10s\t%10s\t%g\t%d\n", $1,$2,$3,$4,$5 
} 

$ awk -f printf-width.awk items.txt 
Num     Description           Type      Price   Qty 
------------------------------------------------------------------ 
101     HD Camcorder         Video      210     10 
102     Refrigerator     Appliance      850     2 
103     MP3 Player           Audio      270     15 
104     Tennis Racket       Sports      190     20 
105     Laser Printer       Office      475     5 
``` 
> * 即使我们指定了输出列的宽度, 输出结果仍然没有对齐; 因为我们指定的是最小宽度, 而不是绝对宽度; 如果字符串长度超过了指定的宽度, 整个字符仍然会打印出来; 所以要在到底打印多少宽度的字符上下点功夫; 
> * 如果想在字符串超出指定宽度时, 仍然以指定的宽度把字符串打印出来,可以使用substr函数(或者)在指定宽度的数字前面加一个小数点; 
> * 在左边补空格, 把字符串"Good"打印成6个字符:
``` 
$ awk'BEGIN { printf "%6s\n","Good" }'
  Good 
```  
> * 指定宽度为6, 但仍然输出所有字符: 
```
$ awk'BEGIN { printf "%6s\n", "Good Boy!" }'
Good Boy! 
``` 
> * 打印指定宽度(左对齐)
 - 当字符串长度小于指定宽度时,如果要让它靠左对齐(右边补空格), 那么要在%和格式化字符之间加上一个减号(-);
```
#"%6s"是右对齐: 
$ awk'BEGIN { printf "|%6s|\n", "Good" }'
|  Good| 

"%-6s"是左对齐: 
$ awk'BEGIN { printf "|%-6s|\n", "Good" }'
|Good  | 
```
 
> * 打印美元标识 
 - 如果要在价钱之前加上美元符号, 只需在格式化字符串之前(%之前)加上$即可;
```
$ cat printf-width2.awk 
BEGIN { 
  FS="," 
  printf "%3s\t%10s\t%10s\t%5s\t%3s\n", "Num","Description","Type","Price","Qty" 
  printf "------------------------------------------------------------------\n" 
} 
{ 
  printf "%3d\t%10s\t%10s\t$%-.2f\t%d\n", $1,$2,$3,$4,$5 
} 
 
$ awk -f printf-width2.awk  items.txt 
Num     Description           Type      Price   Qty 
------------------------------------------------------------------ 
101     HD Camcorder         Video      $210.00 10 
102     Refrigerator     Appliance      $850.00 2 
103     MP3 Player           Audio      $270.00 15 
104     Tennis Racket       Sports      $190.00 20 
105     Laser Printer       Office      $475.00 5  
``` 
> * 字符串长度不足时补0 
```
# 默认情况下,右对齐时左边会补空格 
$ awk'BEGIN { printf "|%5s|\n", "100" }'
|  100| 

# 为了在右对齐是,左边补0 (而不是空格),在指定宽度的数字前面加一个0,即使用”%05s”代替”%5s” 
$ awk'BEGIN { printf "|%05s|\n", "100" }'
|00100| 
```
> * 下面的例子中,在打印的商品数量之前补0:
``` 
$ vim printf-width3.awk 
BEGIN { 
  FS="," 
  printf "%-3s\t%-10s\t%-10s\t%-5s\t%-3s\n", "Num","Description","Type","Price","Qty" 
  printf "---------------------------------------------------------------------\n" 
} 
 
{ 
  printf "%-3d\t%-10s\t%-10s\t$%-.2f\t%03d\n", $1,$2,$3,$4,$5 
} 
 
$ awk -f printf-width3.awk  items.txt 
Num     Description     Type            Price   Qty 
--------------------------------------------------------------------- 
101     HD Camcorder    Video           $210.00 010 
102     Refrigerator    Appliance       $850.00 002 
103     MP3 Player      Audio           $270.00 015 
104     Tennis Racket   Sports          $190.00 020 
105     Laser Printer   Office          $475.00 005 
```
> * 以绝对宽度打印字符串 
 - 通过之前的例子可以知道, 如果字符串长度超过指定的宽度, 字符串仍然会整个被打印出来;
``` 
$ awk'BEGIN { printf "%6s\n", "Good Boy!" }'
Good Boy!  
```

> * 如果要最多打印6个字符, 要在指定宽度的数字前面加一个小数点, 即使用"%.6s"代替"%6s", 这样即使字符串比指定宽度长,也只打印字符串中的前6个字符;
``` 
$ awk'BEGIN { printf "%.6s\n", "Good Boy!" }'
Good B 

#这个例子并非适用于所有版本的awk,在GAWK 3.1.5上可以,但在GAWK 3.1.7上则不行。 
``` 
> * 以绝对宽度打印字符串, 最可取的方法是使用substr函数
```
$ awk'BEGIN { printf "%6s\n", substr("Good Boy!",1,6) }'
Good B 
```
> * 控制精度 
 - 数字前面的点,用来指定其数值精度 
 - 下面的例子说明如何控制精度,展示了使用 .1 和 .4时数值"101.23"的精度(使用的格式化字符有d,e,f和g)
```
$ cat dot.awk 
BEGIN { 
  print "---------- Using .1 -----------" 
  printf ".1d--> %.1d\n", 101.23 
  printf ".1e--> %.1e\n", 101.23 
  printf ".1f--> %.1f\n", 101.23 
  printf ".1g--> %.1g\n", 101.23 
  print "---------- Using .4 -----------" 
  printf ".4d--> %.4d\n", 101.23 
  printf ".4e--> %.4e\n", 101.23 
  printf ".4f--> %.4f\n", 101.23 
  printf ".4g--> %.4g\n", 101.23 
} 
$ awk -f dot.awk 
---------- Using .1 ----------- 
.1d--> 101 
.1e--> 1.0e+02 
.1f--> 101.2 
.1g--> 1e+02 
---------- Using .4 ----------- 
.4d--> 0101 
.4e--> 1.0123e+02 
.4f--> 101.2300 
.4g--> 101.2 
```
> * 把结果重定向到文件 
 - Awk中可以把print语句打印的内容重定向到指定的文件中; 
 - 下面的例子中, 第一个print语句使用">report.txt"创建report.txt文件并把内容保存到该文件中; 随后的所有print语句都使用">>report.txt", 把内容追加到已存在的report.txt文件中;
```
$ cat printf-width4.awk 
BEGIN { 
  FS="," 
  printf "%-3s\t%-10s\t%-10s\t%-5s\t%-3s\n", "Num","Description","Type","Price","Qty" > "report.txt" 
  printf "---------------------------------------------------------------------\n" >> "report.txt" 
} 
 
{ 
  if($5 > 10) 
  printf "%-3d\t%-10s\t%-10s\t$%-.2f\t%03d\n", $1,$2,$3,$4,$5 >> "report.txt" 
} 
 
$ awk -f printf-width4.awk items.txt 
$ cat report.txt 
Num     Description     Type            Price   Qty 
--------------------------------------------------------------------- 
103     MP3 Player      Audio           $270.00 015 
104     Tennis Racket   Sports          $190.00 020 
```
 > * 另一中方法是不在print语句中使用">"或">>", 而是在执行awk脚本时使用重定向
```
$ vim printf-width5.awk 
BEGIN { 
  FS="," 
  printf "%-3s\t%-10s\t%-10s\t%-5s\t%-3s\n", "Num","Description","Type","Price","Qty" 
  printf "---------------------------------------------------------------------\n" 
} 
 
{ 
  if($5 > 10) 
  printf "%-3d\t%-10s\t%-10s\t$%-.2f\t%03d\n", $1,$2,$3,$4,$5 
} 
 
$ awk -f printf-width5.awk items.txt >report.txt 
$ cat report.txt 
Num     Description     Type            Price   Qty 
--------------------------------------------------------------------- 
103     MP3 Player      Audio           $270.00 015 
104     Tennis Racket   Sports          $190.00 020 
```

### awk内置数值函数

> * Awk有很多内置的数值、字符串、输入输出函数,下面介绍其中的一部分;
> * int(n)函数 
 - int()函数返回给定参数的整数部分值; 
 - n可以是整数或浮点数,;
 - 如果使用整数做参数, 返回值即是它本身,;
 - 如果指定浮点数,小数部分会被截断; 
```
# int函数示例: 
$ awk 'BEGIN { 
print int(3.534); 
print int(4); 
print int(-5.223); 
print int(-5); 
}'

输出结果为: 
3 
4 
-5 
-5 
```
> * log(n)函数:
 - log(n)函数返回给定参数的自然对数, 参数n必须是正数, 否则会抛出错误; 
```
# log函数示例: 
$ awk 'BEGIN { 
print log(12); 
print log(0); 
print log(1); 
print log(-1); 
}'
2.48491 
-inf 
0 
awk: cmd. line:4: warning: log: received negative argument -1 
nan 

# 可以看到,该例子中log(0)的值是无穷大,显示为-inf, log(-1)抛出了错误(非数字)  
# 注意: 你可以同时会看到log(-1)抛出如下错误信息awk: cmd. line:4: warning: log: received negative argument -1 
```
> * exp(n)函数: 返回e的n次幂 
``` 
# exp函数示例: 
$ awk 'BEGIN { 
print exp(123434346); 
print exp(0); 
print exp(-12); 
}'

awk: cmd. line:1: warning: exp: argument 1.23434e+08 is out of range 
inf 
1 
6.14421e-06 
# 这个例子中,exp(1234346)返回的值是inf,因为这个值已经超出范围(溢出)了; 
```
> * sin(n)函数 
 - sin(n)返回n的正弦值,n是弧度值 
```
# sin函数示例: 
$ awk 'BEGIN { 
print sin(90); 
print sin(45); 
}'

0.893997 
0.850904
```
> * cos(n)函数 
 - cos(n)返回n的余弦值,n是弧度值 
```
# cos函数示例: 
$ awk'BEGIN { 
print cos(90); 
print cos(45); 
}'
-0.448074 
0.525322 
```
>* atan2(m,n)函数 
 - 该函数返回m/n的反正切值,m和n是弧度值。 
``` 
atan2函数示例: 
$ awk'BEGIN { print atan2(30,45) }'
0.588003 
``` 

### 随机数生成器 

#### rand()函数

> * rand()函数用于产生0~1之间的随机数, 它只返回0~1之间的数, 绝不会返回0或1; 这些数在awk运行时是随机的, 但是在多次运行中,又是可预知的; 
 - awk使用一套算法产生随机数, 因为这个算法是固定的, 所以产生的数也有重复的; 
 - 下面的例子产生1000个0到100之间的随机数,并且演示了每个数是怎么产生的; 
```
# 产生1000个随机数(0到100之间): 
$ cat rand.awk 
BEGIN { 
while(i<1000) 
  { 
    n = int(rand()*100); 
    rnd[n]++; 
    i++; 
  } 
 
for(i=0;i<=100;i++) 
{ 
  print i,"Occured",rnd[i],"times"; 
} 
}  
 
$ awk -f rand.awk 
0 Occured 11 times 
1 Occured 8 times 
2 Occured 9 times 
3 Occured 15 times 
4 Occured 16 times 
5 Occured 5 times 
6 Occured 8 times 
7 Occured 9 times 
8 Occured 7 times 
9 Occured 7 times 
10 Occured 11 times 
11 Occured 7 times 
12 Occured 10 times 
13 Occured 9 times 
14 Occured 6 times 
15 Occured 18 times 
16 Occured 10 times 
17 Occured 10 times 
18 Occured 9 times 
19 Occured 8 times 
20 Occured 11 times 
21 Occured 13 times 
22 Occured 10 times 
23 Occured 9 times 
24 Occured 15 times 
25 Occured 8 times 
26 Occured 3 times 
27 Occured 17 times 
28 Occured 9 times 
29 Occured 13 times 
30 Occured 11 times 
31 Occured 9 times 
32 Occured 12 times 
33 Occured 12 times 
34 Occured 9 times 
35 Occured 6 times 
36 Occured 13 times 
37 Occured 15 times 
38 Occured 6 times 
39 Occured 9 times 
40 Occured 7 times 
41 Occured 8 times 
42 Occured 6 times 
43 Occured 8 times 
44 Occured 10 times 
45 Occured 7 times 
46 Occured 10 times 
47 Occured 8 times 
48 Occured 16 times 
49 Occured 12 times 
50 Occured 6 times 
51 Occured 15 times 
52 Occured 6 times 
53 Occured 12 times 
54 Occured 8 times 
55 Occured 13 times 
56 Occured 6 times 
57 Occured 16 times 
58 Occured 5 times 
59 Occured 7 times 
60 Occured 11 times 
61 Occured 12 times 
62 Occured 14 times 
63 Occured 11 times 
64 Occured 9 times 
65 Occured 6 times 
66 Occured 7 times 
67 Occured 10 times 
68 Occured 8 times 
69 Occured 12 times 
70 Occured 13 times 
71 Occured 9 times 
72 Occured 10 times 
73 Occured 11 times 
74 Occured 7 times 
75 Occured 13 times 
76 Occured 13 times 
77 Occured 10 times 
78 Occured 5 times 
79 Occured 12 times 
80 Occured 17 times 
81 Occured 8 times 
82 Occured 7 times 
83 Occured 10 times 
84 Occured 12 times 
85 Occured 12 times 
86 Occured 11 times 
87 Occured 14 times 
88 Occured 4 times 
89 Occured 8 times 
90 Occured 15 times 
91 Occured 10 times 
92 Occured 15 times 
93 Occured 8 times 
94 Occured 11 times 
95 Occured 5 times 
96 Occured 12 times 
97 Occured 11 times 
98 Occured 7 times 
99 Occured 11 times 
100 Occured  times 
通过这个例子可以看出,rand()函数产生的随机数有很高的重复率。 
```

#### srand(n)函数

> * srand(n)函数使用给定的参数n作为种子来初始化随机数的产生过程; 不论何时启动,awk只会从n开始产生随机数,如果不指定参数n,awk默认使用当天的时间作为产生随机数的种子; 
> * 产生5个从5到50的随机数:
```
$ cat srand.awk 
BEGIN { 
  #Initialize the sed d with 5. 
  srand(5); 
 
  #Totally I want to generate 5 numbers 
  total = 5; 
 
  #maximun number is 50 
  max = 50; 
  count = 0; 
  while(count < total) 
  { 
    rnd = int(rand()*max); 
    if( array[rnd] == 0 ) 
    { 
      count++; 
      array[rnd]++; 
    } 
  } 
 
  for ( i=5;i<=max;i++) 
  { 
    if (array[i]) 
      print i; 
  } 
} 
 
$ awk -f srand.awk 
14 
16 
23 
33 
35 
 
首先使用rand()函数产生随机数,然后乘以期望的最大值,获得一个小于50的数 
检测产生的数是否存在于数组中,如果不存在,增加数组的索引和循环数。本例中产生5个数 
最后在for循环中,从最小到最大一次打印每个索引对应的元素值。 
```

### 常用字符串函数

> * 下面是一些可以在所有风格的awk上运行的常用字符串函数。 
 
#### index函数

> * index函数用来获取给定字符串在输入字符串中的索引(位置)。 
- 下面的例子中,字符串”Cali”在字符串”CA is California”中的位置是7. 
  - 可以用index来检测指定的字符串(或者字符)是否存在于输入字符串中。如果指定的字符串没有出现,返回0,就说明指定的字符串不存在,如下所示。 
```
$ cat index.awk 
BEGIN { 
  state="CA is California" 
  print "String CA starts at location",index(state,"CA"); 
  print "String Cali starts at location",index(state,"Cali"); 
  if(index(state,"NY")==0) 
  print "String NY is not found in:",state 
} 
 
$ awk -f index.awk 
String CA starts at location 1 
String Cali starts at location 7 
String NY is not found in: CA is California 
``` 

#### length函数
 
> * length函数返回字符串的长度,下面例子将打印items.txt文件中每行字符串的总数; 
``` 
$ awk '{print length($0)}' items.txt 
30 
33 
28 
32 
30 
```

#### split函数

> * 语法: 
 - split(input-string,output-array,separator) 
 - split函数把字符串分割成单个数组元素,其接受如下参数： 
input-string:这个是需要被分割的字符串 
output-array:这个是分割后的字符串存放的数组 
separator:分割字符串的字段分隔符  
> * 统计某件特定商品的销售量,我们需要取出第2个字段(以逗号分隔的销售量数字列表),然后使用逗号作为分隔符,将其分隔并保存在一个数组中,然后使用循环遍历数组统计总和;
``` 
$ cat items-sold1.txt 
101:2,10,5,8,10,12 
102:0,1,4,3,0,2 
103:10,6,11,20,5,13 
104:2,3,4,0,6,5 
105:10,2,5,7,12,6 
 
$ cat split.awk 
BEGIN { 
        FS=":" 
} 
{ 
        split($2,quantity,","); 
        total=0; 
        for(x in quantity) 
        total=total+quantity[x]; 
        print "Item",$1,":",total,"quantities sold"; 
} 
 
$ awk -f split.awk items-sold1.txt 
Item 101 : 47 quantities sold 
Item 102 : 10 quantities sold 
Item 103 : 65 quantities sold 
Item 104 : 20 quantities sold 
Item 105 : 42 quantities sold 
```

#### substr函数
 
> * 语法： 
 - substr(input-string,location,length) 
 - substr函数从字符串中提取指定的部分(子串),上面语法中： 
   input-string:包含子穿的字符串 
   location:子串的开始位置 
   length:从location开始起,出去的字符串的总长度, 这个选项是可选的,如果不指定长度,那么从location开始一直取到字符串的结尾 
> * 下面的例子从字符串的第5个字符开始,取到字符串结尾并打印出来。开始的3个字符是商品编号,第4个字符时逗号。所以下面的例子会跳过商品编号,打印剩余的内容;
``` 
$ awk '{ print substr($0,5) }' items.txt 
HD Camcorder,Video,210,10 
Refrigerator,Appliance,850,2 
MP3 Player,Audio,270,15 
Tennis Racket,Sports,190,20 
Laser Printer,Office,475,5 
 
从第2个字段的第1个字符起,打印5个字符: 
$ awk -F"," '{ print substr($2,1,5) }' items.txt 
HD Ca 
Refri 
MP3 P 
Tenni 
Laser 
```

### GAWK/NAWK的字符串函数

> * 下面这些函数只能在GAWK和NAWK中使用;

#### sub函数

> * 语法： 
 - sub(original-string,replacement-string,string-variable) 
   sub代表substitution(替换)的意思 
   original-string:将被替换掉的字符串。也可以是一个正则表达式。 
   replacement-string:用来替换的字符串 
   string-variable:既是输入字符串,也是输出字符串。必须小心,一旦替换操作成功执行,你将丢失该字符串原来的值。 
> * 下面的例子中:
 - original-string:这里是一个正则表达式C[Aa],匹配CA或者Ca 
 - replacement-string:如果匹配到original-string,则用”KA”替换之 
 - string-variable:执行替换操作之前,该变量保存的是输入字符串(替换前的字符串),一旦执行替换操作,该变量保存的是输出字符串(替换后的字符串) 
 - 需要注意的是,sub函数只替换第一次出现的original-string
``` 
$ cat sub.awk 
BEGIN { 
   state="CA is California" 
   sub("C[Aa]","KA",state); 
   print state; 
} 
$ awk -f sub.awk 
KA is California 
``` 
> * 第3个参数string-variable是可选的, 如果没有指定, awk会使用$0(当前记录)做为第3个参数;
 - 如下这个例子把头两个字符从"10"替换为"20", 所有商品编号101变成了201,102变成了202,依此类推; 
```
$ awk '{ sub("10","20"); print $0 }' items.txt 
201,HD Camcorder,Video,210,10 
202,Refrigerator,Appliance,850,2 
203,MP3 Player,Audio,270,15 
204,Tennis Racket,Sports,190,20 
205,Laser Printer,Office,475,5 
 
# 如果替换操作执行成功,sub函数返回1,否则返回0. 
 
# 仅打印替换成功的记录 ： 
$ awk '{ if(sub("HD","High-Def")) print $0; }' items.txt 
101,High-Def Camcorder,Video,210,10 
``` 

#### gsub函数
 
> * gsub代表全局替换, 它把所有出现的original-string都替换为replacement-string; 
> * 下面例子中, "CA"和"Ca"都将被替换成"KA": 
``` 
$ cat gsub.awk 
BEGIN { 
   state="CA is California" 
   gsub("C[Aa]", "KA", state); 
   print state; 
} 
$ awk -f gsub.awk 
KA is KAlifornia 

#和sub函数相同, 第3个参数也是可选的, 如果没有指定, 则使用$0作为第3个参数; 
``` 
> * 如下把所有出现的"10"都替换为"20", 它不仅会替换商品编号, 如果其他字段包含了10, 也会进行替换;
```
$ awk'{ gsub("10", "20"); print $0 }'items.txt 
201, HD Camcorder, Video, 220, 20 
202, Refrigerator, Appliance, 850, 2 
203, MP3 Player, Audio, 270, 15 
204, Tennis Racket, Sports, 190, 20 
205, Laser Printer, Office, 475, 5 
```

#### match函数和RSTART, RLENGTH变量

> * match函数从输入字符串中检索给定的字符串(或正则表达式), 当检索到字符串时, 返回一个正数值; 
> * 语法： 
 - match(input-string, search-string) 
 - input-string: 这是需要被检索的字符串 
search-string: 要检索的字符串, 需要包含在input-string中, 它可以是一个正则表达式 

> * 下面例子在state字符串中检索"Cali", 如果Cali出现, 则打印一条检索成功的消息; 
```
$ cat match.awk 
BEGIN { 
  state="CA is California" 
  f(match(state, "Cali")) 
  { 
    print substr(state, RSTART, RLENGTH), "is present in: ", state; 
  } 
} 
$ awk -f match.awk 
Cali is present in:  CA is California 
# match函数设置了两个特殊变量, 这个例子在调用substr函数时使用了它们用来打印检索成功的消息; 
RSTART – search-string的开始位置 
RLENGTH – search-string的长度 
```

#### tolower和toupper

> * tolower和toupper函数仅在GAWK中可以使用;
> * 正如函数名一样, 这两个函数把给定的字符串转换成小写或大写形式, 如下所示：
``` 
$ awk '{ print tolower($0) }' items.txt 
101, hd camcorder, video, 210, 10 
102, refrigerator, appliance, 850, 2 
103, mp3 player, audio, 270, 15 
104, tennis racket, sports, 190, 20 
105, laser printer, office, 475, 5 
 
$ awk'{ print toupper($0) }'items.txt 
101, HD CAMCORDER, VIDEO, 210, 10 
102, REFRIGERATOR, APPLIANCE, 850, 2 
103, MP3 PLAYER, AUDIO, 270, 15 
104, TENNIS RACKET, SPORTS, 190, 20 
105, LASER PRINTER, OFFICE, 475, 5 
```

### 处理参数(ARGC, ARGV, ARGIND)
 
> * awk内置变量如: FS, NFS, RS, NR, FILENAME, OFS和ORS, 这些变量在所有的awk版本(包括nawk和gawk)上都可以使用; 
 - 本节中提到的环境变量仅仅适用于nawk和gawk 
 - 可以使用ARGC和ARGV从命令行传递一些参数给awk脚本 
 - ARGC保存着传递给awk脚本的所有参数的个数 
 - ARGV是一个数组, 保存着传递给awk脚本的所有参数, 其索引范围从0到ARGC 
 - 当传递5个参数是, ARGC的值为6 
 - ARGV[0]的值永远是awk 
> * 下面的例子arguments.awk演示ARGC和ARGV的作用: 
``` 
$ cat arguments.awk 
BEGIN { 
  print "ARGC=", ARGC 
  for(i=0;i<ARGC;i++) 
    print ARGV[i] 
}
 
$ awk -f arguments.awk arg1 arg2 arg3 arg4 arg5 
ARGC= 6 
awk 
arg1 
arg2 
arg3 
arg4 
arg5 
``` 

> * 在下面的例子中:  
 - 我们以"--参数名 参数值"的格式给awk脚本传递一些参数 
 - awk脚本获取传递元素的内容和数量作为参数 
 - 如果把"--item 104 --qty 25"作为参数传递给awk脚本, awk会把商品104的数量设置为25 
 - 如果把"--item 105 --qty 3"作为参数传递给awk脚本, awk会把商品105的数量设置为3 
```
$ cat argc-argv.awk 
BEGIN { 
  FS=", "; 
  OFS=", "; 
  for(i=0;i<ARGC;i++) 
  { 
    if(ARGV[i] == "--item") 
    { 
      itemnumber=ARGV[i+1]; 
      delete ARGV[i] 
      i++; 
      delete ARGV[i] 
    } 
    else if (ARGV[i]=="--qty") 
    { 
      quantity=ARGV[i+1] 
      delete ARGV[i] 
      i++; 
      delete ARGV[i] 
    } 
  } 
} 
 
{ 
  if ($1==itemnumber) 
  print $1, $2, $3, $4, quantity 
  else 
  print $0 
} 
 
$ awk -f argc-argv.awk --item 104 --qty 25 items.txt 
101, HD Camcorder, Video, 210, 10 
102, Refrigerator, Appliance, 850, 2 
103, MP3 Player, Audio, 270, 15 
104, Tennis Racket, Sports, 190, 25 
105, Laser Printer, Office, 475, 5 
``` 
> * 在gawk中, 当前处理的文件被存放在数组ARGV中, 该数组在body区域被访问;ARGIND是ARGV的一个索引, 其对应的值是当前正在处理的文件名; 
> * 当awk脚本仅处理一个文件时, ARGIND的值是1, ARGV[ARGIND]会返回当前正在处理的文件名; 
- 下面的例子只有body区域, 打印ARGIND的值以及ARGV[ARGIND] 
```
$ cat argind.awk 
{ 
  print "ARGIND: ", ARGIND 
  print "Current file: ", ARGV[ARGIND] 
} 

# 调用这个脚本时, 可以传递两个文件给它, 没处理一行记录就会打印两条数据;这个例子意在让你搞清楚ARGIND和ARGV[ARGIND]的值是怎么存储的; 
$ awk -f argind.awk items.txt items-sold1.txt 
ARGIND:  1 
Current file:  items.txt 
ARGIND:  1 
Current file:  items.txt 
ARGIND:  1 
Current file:  items.txt 
ARGIND:  1 
Current file:  items.txt 
ARGIND:  1 
Current file:  items.txt 
ARGIND:  2 
Current file:  items-sold1.txt 
ARGIND:  2 
Current file:  items-sold1.txt 
ARGIND:  2 
Current file:  items-sold1.txt 
ARGIND:  2 
Current file:  items-sold1.txt 
ARGIND:  2 
Current file:  items-sold1.txt 
```

### OFMT 
> * 内置变量OFMT仅适用于NAWK和GAWK; 
> * 当一个数值被转换成字符串并打印时, awk适用OFMT格式来决定如何打印这些值; OFMT默认值是"%.6g", 包括小数点两边的数字, 它打印一共6个长度的字符; 
- 在使用g时, 必须数清楚小数点两边的数字位数, 如"%.4g"表示包括小数点两侧, 一共打印4个字符; 
- 在使用f时, 只需要数清楚小数点后面的数字位数, 如"%.4f"表示小数点后面会打印4个字符;在此不必关系小数点左边有多少个字符; 
> * 下面的脚本ofmt.awk演示在使用不同的OFMT值(g和f)时, 如何打印字符串
``` 
$ cat ofmt.awk 
BEGIN { 
  total=143.123456789; 
  print "--- using g ---" 
  print "Default OFMT: ", total; 
  OFMT="%.3g" 
  print "%.3g OFMT: ", total; 
  OFMT="%.4g" 
  print "%.4g OFMT: ", total; 
  OFMT="%.5g" 
  print "%.5g OFMT: ", total; 
  OFMT="%.6g" 
  print "%.6g OFMT: ", total; 
  print "--- using f ---" 
  OFMT="%.0f"; 
  print "%.0f OFMT: ", total; 
  OFMT="%.1f"; 
  print "%.1f OFMT: ", total; 
  OFMT="%.2f"; 
  print "%.2f OFMT: ", total; 
  OFMT="%.3f"; 
  print "%.3f OFMT: ", total; 
} 
$ awk -f ofmt.awk 
--- using g --- 
Default OFMT:  143.123 
%.3g OFMT:  143 
%.4g OFMT:  143.1 
%.5g OFMT:  143.12 
%.6g OFMT:  143.123 
--- using f --- 
%.0f OFMT:  143 
%.1f OFMT:  143.1 
%.2f OFMT:  143.12 
%.3f OFMT:  143.123 
```

### GAWK内置的环境变量

> * 本节讨论的内置变量仅适用于GAWK; 

#### ENVIRON
 
> * 如果能在awk脚本中访问shell环境变量会十分有用;ENVIRON是一个包含所有shell环境变量的数组, 其索引就是环境变量的名称; 
> * 如元素ENVIRON["PATH"]的值是环境变量PATH的值; 
> * 下面的例子打印所有环境变量名称和值; 
```
$ cat environ.awk 
BEGIN { 
  OFS="=" 
  for(x in ENVIRON) 
  print x, ENVIRON[x] 
} 
$ awk -f environ.awk 
MAIL=/var/mail/root 
CPU=x86_64 
XDG_CONFIG_DIRS=/etc/xdg 
LC_CTYPE=en_US.UTF-8 
INPUTRC=/etc/inputrc 
HOST=mkey 
PWD=/root 
FROM_HEADER= 
………..
``` 
 
#### IGNORECASE
 
> * 默认情况下, IGNORECASE的值是0, 所有awk区分大小写; 
> * 当把IGNORECASE的值设置为1时, awk则不区分大小写, 这在使用正则表达式和比较字符串时很有效率; 
> * 下面的例子不会打印任何内容, 因为它想用匹配小写的"video", 但items.txt中包含的却是大写的"Video". 
- awk'/video/ {print}'items.txt 
> * 然而, 当把IGNORECASE设置为1时, 就能匹配并打印包含"Video"的行, 因为现在awk不区分大小写; 
```
$ awk'BEGIN{IGNORECASE=1} /video/{print}'items.txt 
101, HD Camcorder, Video, 210, 10 
``` 
> * 下面的例子同时支持字符串比较和正则表达式; 
```
$ cat ignorecase.awk 
BEGIN { 
 FS=", "; 
 IGNORECASE=1; 
} 
{ 
 if ($3 == "video") print $0; 
 if ($2 ~ "TENNIS") print $0; 
} 
$ awk -f ignorecase.awk items.txt 
101, HD Camcorder, Video, 210, 10 
104, Tennis Racket, Sports, 190, 20 
```

#### ERRNO 

> * 当执行I/O操作(比如getline)出错时, 变量ERRNO会保存错误信息; 
> * 下面的例子试图用getline读取一个不存在的文件, 此时ERRNO的内容将会是"No such file or directory"; 
```
$ cat errno.awk 
{ 
  print $0; 
  x = getline < "dummy-file.txt" 
  if ( x == -1 ) 
    print ERRNO 
  else 
    print $0 
} 
 
$ awk -f errno.awk items.txt 
101, HD Camcorder, Video, 210, 10 
No such file or directory 
102, Refrigerator, Appliance, 850, 2 
No such file or directory 
103, MP3 Player, Audio, 270, 15 
No such file or directory 
104, Tennis Racket, Sports, 190, 20 
No such file or directory 
105, Laser Printer, Office, 475, 5 
No such file or directory 
``` 

### pgawk - awk运行分析器

> * pgawk程序用来生成awk执行的结果报告;用pgawk可以看到akw每次执行了多少条语句(以及用户自定义的函数);  
> * 首先建立下面的awk脚本作为样本, 以供pgawk执行, 然后分析其结果; 
```
$ cat profiler.awk 
BEGIN { 
  FS=", "; 
  print "Report Generate On: ", strftime("%a %b %d %H: %M: %S %Z %Y", systime()); 
} 
 
{ 
  if ( $5 <= 5 ) 
    print "Buy More:  Order", $2, "immediately!" 
  else 
    print "Sell More:  Give discount on", $2, "immediately!" 
} 
 
END { 
  print "----------" 
} 
```
> * 接下来使用pgawk(不是直接调用awk)来执行该样本脚本; 
```
$ pgawk -f profier.awk items.txt

Report Generate On:  Tue Apr 09 15: 56: 26 CST 2013 
Sell More:  Give discount on HD Camcorder immediately! 
Buy More:  Order Refrigerator immediately! 
Sell More:  Give discount on MP3 Player immediately! 
Sell More:  Give discount on Tennis Racket immediately! 
Buy More:  Order Laser Printer immediately! 
---------- 

# pgawk默认会创建输出文件profier.out(或者awkprof.out), 使用—profier选项可以指定输出文件, 如下所示; 
$ pgawk --profile=myprofiler.out -f profier.awk items.txt 
``` 
> * 查看默认的输出文件awkprof.out来弄清楚每条单独的awk语句的执行次数; 
```
$ cat awkprof.out 
  # gawk profile,  created Tue Apr  9 15: 56: 26 2013 
 
  # BEGIN block(s) 
 
  BEGIN { 
     1    FS = ", " 
     1    print "Report Generate On: ",  strftime("%a %b %d %H: %M: %S %Z %Y",  systime()) 
  } 
 
  # Rule(s) 
 
     5  { 
     5    if ($5 <= 5) { # 2 
     2      print "Buy More:  Order",  $2,  "immediately!" 
     3    } else { 
     3      print "Sell More:  Give discount on",  $2,  "immediately!" 
    } 
  } 
 
  # END block(s) 
 
  END { 
     1    print "----------" 
  } 

# 查看awkprof.out文件时, 务必牢记:  
  左侧一列有一个数字, 标识着该awk语句执行的次数;如BEGIN区域里的print语句仅执行了一次(duh!);而while循环执行了5次; 
  对于任意一个条件判断语句, 左边有一个数字, 括号右边也有一个数字;左边数字代表该判断语句执行了多少次, 右边数字代表判断语句为true的次数;上面的例子中, 由(# 2)可以判定, if语句执行了5次, 但只有2次为true; 
```

### 位操作
 
> * 和C语言类似, awk也可以进行位操作;在日常工作中用不到, 但可以说明你可以利用位操作来做什么; 
> * 下面表格列出了十进制数字及其二进制形式\

|十进制|二进制|
|:------|:------|
|2|10| 
|3|11| 
|4|100|
|5|101|
|6|110| 
|7|111|
|8|1000|
|9|1001|

> * AND (按位与)
``` 
# 要使AND结果为1, 两个操作数都必须为1. 
    0 and 0 = 0 
    0 and 1 = 0 
    1 and 0 = 0 
    1 and 1 =1 
# 例如, 在十进制数15和25上执行AND操作, 结果是二进制的01001, 也就是十进制的9. 
    15 = 01111 
    25 = 11001 
    15 and 25 = 01001
```
> * OR(按位或) 
```
# 要使OR结果为1, 任意一个操作数为1即可 
    0 or 0 = 0 
    0 or 1 = 1 
    1 or 0 = 1 
    1 or 1 = 1 
# 例如, 在十进制数15和25上执行OR操作, 结果是二进制的11111, 也就是十进制的31. 
    15 = 01111 
    25 = 11001 
    15 and 25 = 11111 
```

> * XOR(按位异或)
```
# 要使XOR结果为1, 必须只有一个操作数为1  
    0 xor 0 = 0 
    0 xor 1 = 1 
    1 xor 0 = 1 
    1 xor 1 = 0 
# 例如, 在十进制数15和25上执行XOR操作, 结果是二进制的10110, 也就是十进制的22. 
    15 = 01111 
    25 = 11001 
    15 and 25 = 10110
```
> * complement(取反码)
``` 
# 反码把0变成1, 把1变成0  
# 如, 给15取反码; 
   15 = 01111 
   15 compl= 10000
``` 
> * Left Shift(左移) 
```
 - 该函数把操作数向左位移, 可以指定位移多少次, 位移后右边补0; 
 - 例如把十进制数15向左位移(移两次), 结果将是二进制的111100, 即十进制的60. 
   15 = 1111 
   lshift twice = 111100
``` 
> * Right Shift(右移) 
```
 - 该函数把操作数向右位移, 可以指定位移多少次, 位移后左边补0; 
 - 例如把十进制数15向右位移(移两次), 结果将是二进制的0011, 即十进制的3. 
   15 = 1111 
   lshift twice = 0011
```
> * awk位移函数示例 
```
$ cat bits.awk 
BEGIN { 
  number1=15 
  number2=25 
  print "AND:  " and(number1, number2); 
  print "OR:  " or(number1, number2); 
  print "XOR:  " xor(number1, number2); 
  print "LSHIFT:  " lshift(number1, 2); 
  print "RSHIFT:  " rshift(number1, 2); 
} 
 
$ awk -f bits.awk 
AND:  9 
OR:  31 
XOR:  22 
LSHIFT:  60 
RSHIFT:  3 
```

### 用户自定义函数
 
> * awk运行用户自定义函数, 这在编写大量代码同时又要多次重复执行其中某些片段时特别有用, 这些片段就适合定义成函数;  
> * 语法： 
```
function fn-name(parameters) 
{ 
 function-body 
} 
 
fn-name:  函数名, 和awk变量名一样, 用户定义的函数名应该以字母开头, 后续字符可以数字、字母或下划线, 关键字不能用做函数名 
parameters: 多个参数要使用逗号分开, 也可以定义一个没有参数的函数 
function-body: 一条或多条awk语句 
```
> * 如果你在awk中已经使用了某个名字作为变量名, 那么它就不能再用来作函数名;  
> * 下面的例子创建了一个简单的用户自定义函数——discount, 它返回商品打折后的价钱, 如discount(10)返回打九折后的价钱; 
  - 对于任意一种商品, 如果数量不大于10, 则打九折, 否则打5折; 
```
$ cat function.awk 
BEGIN { 
  FS=", " 
  OFS=", " 
} 
{ 
  if ($5 <= 10) 
  print $1, $2, $3, discount(10), $5 
  else 
  print $1, $2, $3, discount(50), $5 
} 
 
function discount(percentage) 
{ 
  return $4 - ($4*percentage/100); 
} 
 
$ awk -f function.awk items.txt 
101, HD Camcorder, Video, 189, 10 
102, Refrigerator, Appliance, 765, 2 
103, MP3 Player, Audio, 135, 15 
104, Tennis Racket, Sports, 95, 20 
105, Laser Printer, Office, 427.5, 5 
```
> * 自定义函数另外一个作用是打印debug信息; 
- 下面是一个简单的mydebug函数:  
```
$cat function-debug.awk 
{ 
  i=2; total=0; 
  while (i <= NF ) { 
  mydebug("quantity is "$i); 
  total = total + $i; 
  i++; 
  } 
  print "Item", $1, ": ", total, "quantities sold"; 
} 
 
function mydebug( message ) 
{ 
  printf("DEBUG[%d]>%s\n", NR, message); 
} 
 
$ awk -f function-debug.awk items-sold.txt 
DEBUG[1]>quantity is 2 
DEBUG[1]>quantity is 10 
DEBUG[1]>quantity is 5 
DEBUG[1]>quantity is 8 
DEBUG[1]>quantity is 10 
DEBUG[1]>quantity is 12 
Item 101 :  47 quantities sold 
DEBUG[2]>quantity is 0 
DEBUG[2]>quantity is 1 
DEBUG[2]>quantity is 4 
DEBUG[2]>quantity is 3 
DEBUG[2]>quantity is 0 
DEBUG[2]>quantity is 2 
Item 102 :  10 quantities sold 
DEBUG[3]>quantity is 10 
DEBUG[3]>quantity is 6 
DEBUG[3]>quantity is 11 
DEBUG[3]>quantity is 20 
DEBUG[3]>quantity is 5 
DEBUG[3]>quantity is 13 
Item 103 :  65 quantities sold 
DEBUG[4]>quantity is 2 
DEBUG[4]>quantity is 3 
DEBUG[4]>quantity is 4 
DEBUG[4]>quantity is 0 
DEBUG[4]>quantity is 6 
DEBUG[4]>quantity is 5 
Item 104 :  20 quantities sold 
DEBUG[5]>quantity is 10 
DEBUG[5]>quantity is 2 
DEBUG[5]>quantity is 5 
DEBUG[5]>quantity is 7 
DEBUG[5]>quantity is 12 
DEBUG[5]>quantity is 6 
Item 105 :  42 quantities sold 
``` 

### 使输出摆脱语言依赖-国际化

> * 在使用awk脚本打时, 可能需要用print打印指定的header和footer信息;你或许会用英文把header和footer写成固定的内容;
> * 但在其他语言中, 你希望它会打印什么信息？最终你可能会把脚本复制过去, 然后修改要打印的固定信息, 来适应当前的语言环境; 
> * 有个更容易的方法来实现这个目的--国际化, 这样使用同一个脚本, 仅需要在运行脚本时修改那些固定的输出信息即可; 
> * 在运行大型程序, 而出于某些原因你要频繁地修改那些固定的输出信息时；或者希望用户能自行修改要输出的内容时, 这个方法也很有用; 
> * 下面的例子演示了在awk中实现国际化的关键4步; 
```
###### 步骤1 – 建立文本域 
# 创建一个文本域文件, 并把它和awk要搜寻的目录绑定;下面以当前目录为例; 
$ cat iteminfo.awk 
BEGIN { 
  FS=", " 
  TEXTDOMAIN = "item" 
  bindtextdomain(".") 
  print _"START_TIME: " strftime("%a %b %d %H: %M: %S %Z %Y", systime()); 
  printf "%-3s\t", _"Num"; 
  printf "%-10s\t", _"Description" 
  printf "%-10s\t", _"Type" 
  printf "%-5s\t", _"Price" 
  printf "%-3s\n", _"Qty" 
  printf _"---------------------------------------------------\n" 
} 
 
{ 
  printf "%-3d\t%-10s\t%-10s\t%-.2f\t%03d\n", $1, $2, $3, $4, $5 
} 
# 注意：这个例子中, 前面带"_"的字符串均可以自行定义;字符串前面的_(下划线)不会影响字符串内容的打印, 即它和下面的输出完全相同; 
$ awk -f iteminfo.awk items.txt 
START_TIME: Thu Apr 11 11: 37: 40 CST 2013 
Num     Description     Type            Price   Qty 
--------------------------------------------------- 
101     HD Camcorder    Video           210.00  010 
102     Refrigerator    Appliance       850.00  002 
103     MP3 Player      Audio           270.00  015 
104     Tennis Racket   Sports          190.00  020 
105     Laser Printer   Office          475.00  005 
 
########### 步骤2：生成.po文件 
# 建立如下可移植对象文件(扩展名.po), 注意, 除了使用—gen-po之外, 也可以使用"-W gen-po" 
$ gawk --gen-po -f iteminfo.awk > iteminfo.po 
$ cat iteminfo.po 
#:  iteminfo.awk: 5 
msgid "START_TIME: " 
msgstr "" 
 
#:  iteminfo.awk: 6 
msgid "Num" 
msgstr "" 
 
#:  iteminfo.awk: 7 
msgid "Description" 
msgstr "" 
 
#:  iteminfo.awk: 8 
msgid "Type" 
msgstr "" 
 
#:  iteminfo.awk: 9 
msgid "Price" 
msgstr "" 
 
#:  iteminfo.awk: 10 
msgid "Qty" 
msgstr "" 
 
#:  iteminfo.awk: 11 
msgid "---------------------------------------------------\n" 
"" 
msgstr "" 
# 现在修改该文件中相应的内容;例如要显示"Report Generated on: "(替换原来的"START_TIME"), 那么修改iteminfo.po文件, 把START_TIME右下方的msgstr修改为"Reprot Genterated On: " 
$ cat iteminfo.po 
#:  iteminfo.awk: 5 
msgid "START_TIME: " 
msgstr "Report Generated On: " 
提示：这个例子中, 其余的msgstr字符串均被置空; 
 
####### 步骤3：生成消息对象 
# 使用msgfmt命令(从可移植对象文件生成)生成消息对象文件; 
# 如果iteminfo.po文件中所有的msgstr都是空, 那么将不会生成任何消息对象文件, 如： 
$ msgfmt –v iteminfo.po 
0 translated messages,  7 untranslated messages. 
 
# 我们已经创建了一条消息, 所以会生成messages.mo文件; 
$ msgfmt -v iteminfo.po 
1 translated message,  6 untranslated messages. 
$ ls -l messages.mo 
-rw-r--r-- 1 root root 89 Apr 11 12: 11 messages.mo 
 
# 把message.mo文件复制到消息目录下, 消息目录应该在当前目录下创建 
$ mkdir -p en_US/LC_MESSAGES  
$ mv messages.mo en_US/LC_MESSAGES/item.mo 
# 注意: 目标文件的名字要和初始awk文件中的TEXTDOAMIN后面的值相同;之前的awk文件中TEXTDOMAIN="item"; 
 
######### 步骤4: 核证消息 
# 现在可以看到, awk将不再显示"START_TIME", 而是转换为"Report Generated On: "然后打印出来： 
$ gawk -f iteminfo.awk items.txt 
Reprot Generated On: Thu Apr 11 12: 16: 19 CST 2013 
Num     Description     Type            Price   Qty 
--------------------------------------------------- 
101     HD Camcorder    Video           210.00  010 
102     Refrigerator    Appliance       850.00  002 
103     MP3 Player      Audio           270.00  015 
104     Tennis Racket   Sports          190.00  020 
105     Laser Printer   Office          475.00  005 
```

### 双向管道
 
> * awk可以使用"|&"和外部进程通信, 这个过程是双向的; 
> * 下面的例子把关键字"Awk"替换为"Sed  and Awk"; 
```
$ echo "Awk is great" | sed 's/Awk/Sed  and Awk/'
Sed  and Awk is great 
```
> * 下面的例子使用"|&"了模拟实现上面的例子, 以说明awk是如何双向管道的; 
```
$ cat two-way.awk 
BEGIN { 
  command = "sed 's/Awk/Sed  and Awk/'" 
  print "Awk is Great!" |& command 
  close(command, "to"); 
  command |& getline tmp 
  print tmp; 
  close(command); 
} 
$ awk -f two-way.awk 
Sed  and Awk is Great! 

command = "sed 's/Awk/Sed  and Awk/'" –这是要和awk双向管道对接的命令;它是一个简单的sed 替换命令, 把"Awk"替换为"Sed  and Awk"; 
print "Awk is Great!" |& command – command的输入, 即sed 替换命令的输入是"Awk is Great!";"|&"表示这里是双向管道;"|&"右边命令的输入来自左边命令的输出; 
close(command, "to") – 一旦命令执行完成, 应该关闭"to"进程; 
command |& getline tmp –既然命令已经执行完成, 就要用getline获取其输出;前面命令的输出会被存在变量"tmp"中; 
print tmp –打印输出 
close(command) –最后, 关闭命令; 

# 双向管道迟早会派上用场, 尤其是当awk对外部程序的输出依赖比较大时; 
```

### 系统函数
 
> * 可以使用操作系统内置的函数来执行操作系统命令, 但请注意, 调用系统命令和使用双向管道是不同的; 
> * 使用"|&"时, 可以把任意awk命令的输出作为外部命令的输入；也可以接收外部命令的输出作为awk的输入(要不怎么叫双向管道呢); 
> * 执行系统命令时, 可以传递任意的字符串作为命令的参数, 它会被当做操作系统命令准确第执行, 并返回结果(这和双向管道有所不同); 
> * 下面的例子在awk中调用pwd和date命令： 
```
$ awk 'BEGIN { system("pwd") }'
/root 
$ awk 'BEGIN { system("date") }'
Wed Apr 10 17: 07: 08 CST 2013 
```
> * 在执行比较大的awk脚本时, 你或许想在脚本开始和结束时发送一封电子邮件;下面的例子展示如何在BEGIN和END区域中调用系统命令, 在脚本开始和结束时发送邮件; 
```
$ cat system.awk 
BEGIN { 
  system("echo'Started'| mail -s'Program system.awk started ..'ramesh@thegeekstuff.com"); 
} 
{ 
  split($2, quantity, ", "); 
  total=0; 
  for (x in quantity) 
  total=total+quantity[x]; 
  print "Item", $1, ": ", total, "quantities sold"; 
} 
 
END { 
  system("echo'Completed'| mail -s'Program system.awk completed..'ramesh@thegeekstuff.com"); 
} 
 
$ awk -f system.awk items-sold.txt 
Item 101 :  2 quantities sold 
Item 102 :  0 quantities sold 
Item 103 :  10 quantities sold 
Item 104 :  2 quantities sold 
Item 105 :  10 quantities sold 
``` 

### 时间函数
 
> * 这些命令仅适用于GAWK; 
> * 看下面的例子, systime()函数返回系统的POSIX时间, 即自1970年1月1日起至今经过的秒数;
``` 
$ awk 'BEGIN { print systime() }'
1365585325  
``` 
> * 如果使用strftime函数把POSIX时间转换为可读的格式, systime函数就变动很有用了; 
> * 下面例子使用systime和strftime以可读的格式打印当前时间; 
```
$ awk'BEGIN { print strftime("%c", systime()) }'
Wed Apr 10 17: 17: 35 2013  
``` 
> * 下面的例子展示了多种可用的时间格式; 
``` 
$ cat strftime.awk 
BEGIN { 
  print "--- basic formats ---" 
  print strftime("Format 1:  %m/%d/%Y %H: %M: %S", systime()) 
  print strftime("Format 2:  %m/%d/%y %I: %M: %S %p", systime()) 
  print strftime("Format 3:  %m-%b-%Y %H: %M: %S", systime()) 
  print strftime("Format 4:  %m-%b-%Y %H: %M: %S %Z", systime()) 
  print strftime("Format 5:  %a %b %d %H: %M: %S %Z %Y", systime()) 
  print strftime("Format 6:  %A %B %d %H: %M: %S %Z %Y", systime()) 
  print "--- quick formats ---" 
  print strftime("Format 7:  %c", systime()) 
  print strftime("Foramt 8:  %D", systime()) 
  print strftime("Format 9:  %F", systime()) 
  print strftime("Format 10:  %x", systime()) 
  print strftime("Format 11:  %X", systime()) 
  print "--- single line format with %t ---" 
  print strftime("%Y %t%B %t%d", systime()) 
  print "--- multi line format with %n ---" 
  print strftime("%Y%n%B%n%d", systime()) 
} 
$ awk -f strftime.awk 
--- basic formats --- 
Format 1:  04/10/2013 18: 04: 06 
Format 2:  04/10/13 06: 04: 06 PM 
Format 3:  04-Apr-2013 18: 04: 06 
Format 4:  04-Apr-2013 18: 04: 06 CST 
Format 5:  Wed Apr 10 18: 04: 06 CST 2013 
Format 6:  Wednesday April 10 18: 04: 06 CST 2013 
--- quick formats --- 
Format 7:  Wed Apr 10 18: 04: 06 2013 
Foramt 8:  04/10/13 
Format 9:  2013-04-10 
Format 10:  04/10/13 
Format 11:  18: 04: 06 
--- single line format with %t --- 
2013    April   10 
--- multi line format with %n --- 
2013 
April 
10 
```
> * 下面是strftime函数可用的时间格式标识符, 需要注意的是, 下面所有的标识符依赖于本地系统的设置, 下面的示例是打印英文时间;

|格式标识符|描述|
|:------|:------|
|%m 	|两位数字月份, 一月显示为01| 
|%b 	|月份缩写, 一月显示为Jan |
|%B 	|月份完整单词, 一月显示为January |
|%d 	|两位数字日期, 4号显示为04 |
|%Y 	|年份的完整格式, 如2011 |
|%y 	|两位数字的年份, 如2011显示为11| 
|%H 	|24小时格式,  1 p.m显示为13 |
|%l 	|12小时格式,  1 p.m显示为01 |
|%p 	|显示AM或PM, 和%l搭配使用 |
|%M 	|两位数字分钟, 9分显示为09 |
|%S 	|两位数字描述, 5秒显示为05 |
|%a 	|三位字符星期, 周一显示为Mon| 
|%A 	|完整的日期, 周一显示为Monday| 
|%Z 	|时区, 太平洋地区时区显示为PST |
|%c 	|显示本地时间的完整格式, 如Fri 11 Feb 2011 02: 45: 03 AM PST |
|%D 	|简单日期格式, 和%m/%d/%y相同 |
|%F 	|简单日期格式, 和%Y-%m-%d相同 |
|%x 	|基于本地设置的时间格式 |
|%X| 	基于本地设置的时间格式 |

### getline命令

> * 使用geline命令可以控制awk从输入文件(或其他文件)读取数据;注意, 一旦getline执行完成, awk脚本会重置NF, NR, FNR和$0等内置变量; 
> * getline示例:
```
$ awk -F", " '{getline;print $0;}' items.txt 
102, Refrigerator, Appliance, 850, 2 
104, Tennis Racket, Sports, 190, 20 
105, Laser Printer, Office, 475, 5 
 
# 当在body区域使用了getline时, 会直接读取下一行数据; 这个例子中, 第一条语句便是getline, 所以即使awk已经读取了第一行数据, getline也会继续读取下一行, 因为我们强制它读取下一行, 因此, getline后面的print $0会打印输入文件的第二行; 

开始执行body区域时, 执行任何命令之前, awk从items.txt文件中读取第一行数据, 保存在变量$0中 
getline   我们用getline命令强制awk读取下一行数据, 保存在变量$0中(之前的内容被覆盖掉了) 
print $0  既然现在$0中保存的是第二行数据, print $0会打印文件第二行(而不是第一行) 
body区域继续执行, 只打印偶数行的数据;(注意到最后一行105也打印了么？) 
```

> * 把getline的内容保存的变量中, 除了把getline的内容放到$0中, 还可以把它保存在变量中; 
```
# 只打印奇数行内容 
$ awk -F", "'{getline tmp; print $0;}'items.txt 
101, HD Camcorder, Video, 210, 10 
103, MP3 Player, Audio, 270, 15 
105, Laser Printer, Office, 475, 5 

# 开始执行body区域时, 执行任何命令之前, awk从items.txt文件中读取第一行数据, 保存在变量$0中 
# getline tmp – 强制awk读取下一行, 并保存在变量tmp中 
# print $0 – 此时$0仍然是第一行数据, 因为getline tmp没有覆盖$0, 因此会打印第一行数据(而不是第二行) 
# body区域继续执行, 只打印奇数行的数据; 
```

> * 下面的例子同时打印\$0和tmp, 可以看到, $0是奇数行, 而tmp是偶数行
``` 
$ awk -F", " '{getline tmp; print "$0->", $0;print "tmp->", tmp;}' items.txt 
$0-> 101, HD Camcorder, Video, 210, 10 
tmp-> 102, Refrigerator, Appliance, 850, 2 
$0-> 103, MP3 Player, Audio, 270, 15 
tmp-> 104, Tennis Racket, Sports, 190, 20 
$0-> 105, Laser Printer, Office, 475, 5 
tmp-> 104, Tennis Racket, Sports, 190, 20 
``` 
 
> * getline也可以从其他文件(非当前输入文件)读取内容, 如下所示; 
 - 在两个文件中循环切换, 打印所有内容; 
```
$ awk -F", " '{print $0;getline <"items-sold.txt"; print $0;}' items.txt 
101, HD Camcorder, Video, 210, 10 
101 2 10 5 8 10 12 
102, Refrigerator, Appliance, 850, 2 
102 0 1 4 3 0 2 
103, MP3 Player, Audio, 270, 15 
103 10 6 11 20 5 13 
104, Tennis Racket, Sports, 190, 20 
104 2 3 4 0 6 5 
105, Laser Printer, Office, 475, 5 
105 10 2 5 7 12 6 

# 开始执行body区域时, 执行任何命令之前, awk从items.txt文件中读取第一行数据, 保存在变量$0中 
# print $0 – 打印items.txt文件的第一行 
# getline <"items-sold.txt" – 读取items-sold.txt中第一行并保存在$0中 
# print $0 – 打印items-sold.txt文件中的第一行 
# body区域继续执行, 轮番打印items.txt和items-sold.txt的剩余内容 
```

> * 除了把两个文件的内容都读到$0中之外, 也可以使用"getline var"把读取的内容保存到变量中; 
 - 继续使用上面那两个文件, 打印所有内容(使用tmp变量) 
```
$ awk -F", " '{print $0; getline tmp < "items-sold.txt";print tmp;}' items.txt 
101, HD Camcorder, Video, 210, 10 
101 2 10 5 8 10 12 

102, Refrigerator, Appliance, 850, 2 
102 0 1 4 3 0 2 
103, MP3 Player, Audio, 270, 15 
103 10 6 11 20 5 13 
104, Tennis Racket, Sports, 190, 20 
104 2 3 4 0 6 5 
105, Laser Printer, Office, 475, 5 
105 10 2 5 7 12 6 
```

> * getline执行外部命令 
 - getline也可以执行UNIX命令并获取其输出; 
> * 下面的例子使用getline获取date命令的输出并打印出来;请注意这里也要使用close刚执行的命令, date命令的输出保存在变量$0中, 如下所示; 
 - 使用这个方法可以在输出报文的header和footer中显示时间戳; 
```
$ cat getline1.awk 
BEGIN { 
  FS=", "; 
  "date" | getline 
  close("date") 
  print "Timestamp: " $0 
} 
{ 
  if ( $5 <= 5) 
  print "Buy More: Order", $2, "immediately!" 
  else 
  print "Sell More: Give discount on", $2, "immediatelty!" 
} 
 
$ awk -f getline1.awk items.txt 
Timestamp: Thu Apr 11 11: 07: 47 CST 2013 
Sell More: Give discount on HD Camcorder immediatelty! 
Buy More: Order Refrigerator immediately! 
Sell More: Give discount on MP3 Player immediatelty! 
Sell More: Give discount on Tennis Racket immediatelty! 
Buy More: Order Laser Printer immediately! 
``` 

> * 除了把命令输出保存在$0中之外, 也可以把它保存在任意的awk变量中(如timestamp), 如下所示; 
```
$ cat getline2.awk 
BEGIN { 
  FS=", "; 
  "date" | getline timestamp 
  close("date") 
  print "Timestamp: " timestamp 
} 
{ 
  if ( $5 <= 5) 
  print "Buy More:  Order", $2, "immediately!" 
  else 
  print "Sell More:  Give discount on", $2, "immediately!" 
} 
$ awk -f getline2.awk items.txt 
Timestamp: Thu Apr 11 11: 20: 05 CST 2013 
Sell More:  Give discount on HD Camcorder immediately! 
Buy More:  Order Refrigerator immediately! 
Sell More:  Give discount on MP3 Player immediately! 
Sell More:  Give discount on Tennis Racket immediately! 
Buy More:  Order Laser Printer immediately! 
```

## awk分析网络状态

```
$ netstat -an | awk '/^tcp\>/ {state[$NF]++} END {for(key in state) printf "state: %-15scount: %d\n",key,state[key]}'
```

## awk统计日志流量

> * 访问日志原型
```
127.0.0.1 - - [28/Feb/2019:14:45:21 +0800] "GET / HTTP/1.1" 200 612 "-" "curl/7.54.0"
127.0.0.1 - - [28/Feb/2019:14:45:21 +0800] "GET / HTTP/1.1" 200 612 "-" "curl/7.54.0"
127.0.0.2 - - [28/Feb/2019:14:45:21 +0800] "GET / HTTP/1.1" 200 612 "-" "curl/7.54.0"
127.0.0.3 - - [28/Feb/2019:14:45:25 +0800] "GET / HTTP/1.1" 200 612 "-" "curl/7.54.0"
127.0.0.2 - - [28/Feb/2019:14:45:21 +0800] "GET / HTTP/1.1" 200 612 "-" "curl/7.54.0"
127.0.0.3 - - [28/Feb/2019:14:45:25 +0800] "GET / HTTP/1.1" 200 612 "-" "curl/7.54.0"
127.0.0.4 - - [28/Feb/2019:14:45:25 +0800] "GET / HTTP/1.1" 200 612 "-" "curl/7.54.0"
127.0.0.5 - - [28/Feb/2019:14:45:26 +0800] "GET / HTTP/1.1" 200 612 "-" "curl/7.54.0"
127.0.0.6 - - [28/Feb/2019:14:46:01 +0800] "GET / HTTP/1.1" 200 612 "-" "Mozilla/5.0 (Macintosh; Intel Mac OS X 10_14_3) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/72.0.3626.119 Safari/537.36"
127.0.0.4 - - [28/Feb/2019:14:45:25 +0800] "GET / HTTP/1.1" 200 612 "-" "curl/7.54.0"
127.0.0.4 - - [28/Feb/2019:14:45:25 +0800] "GET / HTTP/1.1" 200 612 "-" "curl/7.54.0"
127.0.0.5 - - [28/Feb/2019:14:45:26 +0800] "GET / HTTP/1.1" 200 612 "-" "curl/7.54.0"
127.0.0.6 - - [28/Feb/2019:14:46:01 +0800] "GET / HTTP/1.1" 200 612 "-" "Mozilla/5.0 (Macintosh; Intel Mac OS X 10_14_3) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/72.0.3626.119 Safari/537.36"
127.0.0.5 - - [28/Feb/2019:14:45:26 +0800] "GET / HTTP/1.1" 200 612 "-" "curl/7.54.0"
127.0.0.6 - - [28/Feb/2019:14:46:01 +0800] "GET / HTTP/1.1" 200 612 "-" "Mozilla/5.0 (Macintosh; Intel Mac OS X 10_14_3) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/72.0.3626.119 Safari/537.36"
127.0.0.1 - - [28/Feb/2019:14:45:09 +0800] "\xFF\xF4\xFF\xFD\x06" 400 157 "-" "-"
127.0.0.2 - - [28/Feb/2019:14:45:21 +0800] "GET / HTTP/1.1" 200 612 "-" "curl/7.54.0"
127.0.0.3 - - [28/Feb/2019:14:45:25 +0800] "GET / HTTP/1.1" 200 612 "-" "curl/7.54.0"
```
> * awk分析实现
```
$ awk '{count[$1]++;sum[$1]+=$10}END{for(key in sum) print "IP:",key,"total:",sum[key],"count:",count[key]}' access.log
IP: 127.0.0.1 total: 1224 count: 3
IP: 127.0.0.2 total: 1836 count: 3
IP: 127.0.0.3 total: 1836 count: 3
IP: 127.0.0.4 total: 1836 count: 3
IP: 127.0.0.5 total: 1836 count: 3
IP: 127.0.0.6 total: 1836 count: 3
```