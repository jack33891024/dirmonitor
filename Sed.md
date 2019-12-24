标签: 三剑客

# Sed

------
[toc]

## Sed简介

> * SED: Stream EDitor, 流编辑器;
> * sed是操作、过滤和转换文本内容的强大工具, sed可以从文本和管道中读取输入;
> * 在bash启动文件中,就可能有不少用来设置各种环境变量的sed命令;

## 创建文件

> * 所有的示例都要用到下面的example.txt文件,请先行创建改文件
```
$ cat example.txt 
 101,John Doe,CEO 
 102,Jason Smith,IT Manager 
 103,Raj Reddy,Sysadmin 
 104,Anand Ram,Developer 
 105,Jane Miller,Sales Manager

# 上面的雇员数据，每行记录都包含下面几列:
  雇员ID 
  雇员姓名 
  雇员职位
```

## Sed语法

> * Sed基本语法:`sed [options] {sed-commands} {input-file}`
```
# Sed执行过程说明:
  1.sed每次从input-file中读取一行记录,并在该记录上执行sed-commands; 
  2.sed首先从input-file中读取第一行,然后执行所有的sed-commands; 再读取第二行,执行所有sed-commands,重复这个过程,直到input-file结束;

# 通过指定[options]还可以给sed传递一些可选的选项;
```
> * 下面例子演示了sed的基本语法,它打印出/etc/passwd文件中所有行;
 - sed -n 'p' /etc/passwd
 该例子的重点在于,{sed-commands}既可以是单个命令,也可以是多个命令;
 可以把多个sed命令合并到一个文件中,这个文件被称为sed脚本,然后使用-f选项调用它,
 > * 使用sed脚本的基本语法: 
 - sed [options] -f {sed-commands-in-a-file} {input-file}
 > * 下面例子演示了使用sed脚本的用法, 打印/etc/passwd中以root和nobody开头的行:
```
# 脚本
$ cat >test-script.sed
/^root/p
/^nobody/p

# sed执行脚本
$ sed -n -f test-script.sed /etc/passwd
root:x:0:0:root:/root:/bin/bash
nobody:x:99:99:Nobody:/:/sbin/nologin
```
> * 使用-e选项,执行多个sed命令, -e的使用方法如下所示:
 - sed [options] -e {sed-command-1} -e {sed-command2} {input-file}
 - 下面例子演示-e的使用方法,打印/etc/passwd中以root和nobody开头的行
```
$ sed -n -e '/^root/p' -e'/^nobody/p' /etc/passwd
root:x:0:0:root:/root:/bin/bash
nobody:x:99:99:Nobody:/:/sbin/nologin

# 使用-e执行多个命令,可以使用\换行
$ sed -n \
-e '/^root/p' \
-e '/^nobody/p' \
/etc/passwd
root:x:0:0:root:/root:/bin/bash
nobody:x:99:99:Nobody:/:/sbin/nologin
```
> * 可以使用{}将多个命令分组执行
```
# 语法
sed [options] '{ 
sed-command-1 
sed-command-2 
}' input-file

# 演示{}的使用方法,打印/etc/password中以root和nobody开头的行
$ sed -n '{
/^root/p 
/^nobody/p 
}' /etc/passwd 
root:x:0:0:root:/root:/bin/bash
nobody:x:99:99:Nobody:/:/sbin/nologin
```
> * sed不会修改原始文件input-file,它只是将结果内容输出到标准输出设备,如果要保持变更,应该使用重定向`>filename.txt`

## Sed执行流程

> * Sed脚本执行遵从下面简单易记的顺序: Read, Execute, Print, Repeat(读取, 执行, 打印, 重复), 简称REPR;
> * 分析脚本执行顺序:
 - 读取一行到模式空间(sed内部的一个临时缓存,用于存放读取到的内容)
 - 在模式空间中执行命令, 如果使用了{}或-e指定了多个命令, sed将依次执行每个命令
 - 打印模式空间的内容, 然后清空模式空间
 - 重复上述过程, 直到文件结束
> * sed执行流程图
![image_1cijgc16b8631938p10eb9v5a19.png-66.5kB][1]

## 打印模式空间

> * 使用命令p, 可以打印当前模式空间的内容
> * sed在执行完命令后会默认打印模式空间的内容, 使用命令p时,每行记录会输出两次;
> * 命令p可以控制只输出你指定的内容, 通常使用p时,还需要使用-n选项来屏蔽sed的默认输出;
> * 下面例子打印example.txt文件,每行会输出两次
```
$ sed 'p' example.txt 
 101,John Doe,CEO 
 101,John Doe,CEO 
 102,Jason Smith,IT Manager 
 102,Jason Smith,IT Manager 
 103,Raj Reddy,Sysadmin 
 103,Raj Reddy,Sysadmin 
 104,Anand Ram,Developer 
 104,Anand Ram,Developer 
 105,Jane Miller,Sales Manager 
 105,Jane Miller,Sales Manager
```
> * 输出example.txt的内容,只打印一行
```
$ sed -n 'p' example.txt 
 101,John Doe,CEO 
 102,Jason Smith,IT Manager 
 103,Raj Reddy,Sysadmin 
 104,Anand Ram,Developer 
 105,Jane Miller,Sales Manager
```
> * 指定地址范围,如果在命令前面不指定地址范围, 默认会匹配所有行
```
#只打印第2行:
$ sed -n '2p' example.txt 
 102,Jason Smith,IT Manager

#打印第1行至第4行:
$ sed -n '1,4p' example.txt  
 101,John Doe,CEO 
 102,Jason Smith,IT Manager 
 103,Raj Reddy,Sysadmin 
 104,Anand Ram,Developer

#打印第2行至最后一行: $代表最后一行
$ sed -n '2,$p' example.txt      
 102,Jason Smith,IT Manager 
 103,Raj Reddy,Sysadmin 
 104,Anand Ram,Developer 
 105,Jane Miller,Sales Manager
```
> * 修改地址范围
```
# 可以使用逗号、加号、波浪号来修改地址范围
# 使用逗号参与地址范围的指定, 例如: n,m代表第n行至第m行
# 加号配合逗号使用, 可以指定相邻的若干行，而不是绝对的几行, 例如: n,+m代表从第n行开始后的m行
# 波浪号~可以指定地址范围, 它指定每次要跳过的行数, 如: n~m代表从第n行开始, 每次跳过m行
    1~2: 匹配1,3,5,7, ... 
    2~2: 匹配2,4,6,8, ...
    1~3: 匹配1,4,7,10, ...
    2~3: 匹配2,5,8,11, ...
```
> * 只打印奇数行
```
$ sed -n '1~2p' example.txt 
 101,John Doe,CEO 
 103,Raj Reddy,Sysadmin 
 105,Jane Miller,Sales Manager
```

> * 模式匹配
 - 可以使用数字指定地址(或地址范围),也可以使用一个模式(或模式范围)来匹配,如下:
 - 打印匹配模式"Jane"的行:`sed -n '/Jane/p' exmaple.txt`
 - 打印第一次匹配Jason的行至第四行的内容:`sed -n '/Jason/,4p' example.txt`
   如果开始的4行中，没有匹配到Jason,那么sed会打印第4行以后匹配到Jason的内容
 - 打印从第一次匹配Raj的行到最后所有的行:`
   sed -n '/Raj/,$p' example.txt`
 - 打印自匹配Raj的行开始到匹配Jane的行之间的所有内容:`sed -n '/Raj/,/Jane/p' example.txt` 
 - 打印匹配Jason的行和其后面的两行:`sed -n '/Raj/,+2p' example.txt`         

## 删除行

> * 命令d用来删除行,需要注意的是它只删除模式空间的内容,和其他sed命令一样,命令d不会修改原始文件的内容;
> * 如果不提供地址范围,sed默认匹配所有行,下面例子什么都不会输出,因为它匹配了所有行并删除了它们
```
sed 'd' example.txt
```
> * 指定要删除的地址范围更有用,以下是几个例子
```
# 只删除第2行:
   sed '2d' exmaple.txt
# 删除第1行至第4行:
   sed '1,4d' exmaple.txt
# 删除第2行至最后一行:
   sed '2,\$d' example.txt   
# 只删除奇数行:
   sed '1~2d' example.txt 
# 删除匹配Manager的行:
   sed '/Manager/d' example.txt
# 删除从第一次匹配Jason的行至第4行:
   sed '/Jason/,4d' example.txt 
   如果开头的4行中，没有匹配Jason的行，那么上述命令将删除第4行以后匹配Manager的行
# 删除从第一次匹配Raj的行至最后一行:
   sed '/Raj/,$d' example.txt     
# 删除第一次匹配Jason的行和紧跟着它后面的两行:
   sed '/Jason/,+2d' example.txt 
# 删除所有空行
   sed '/^\$/d' example.txt
# 删除所有注释行(假定注释行以#开头) 
   sed '/^#/d' example.txt
```
> * 注意: 如果有多个命令,sed遇到命令d时,会删除匹配到的整行数据,其余的命令将无法操作被删除的行

## 保存模式空间内容至文件中

> * 命令w可以把当前模式空间的内容保存到文件中;
> * 默认情况下模式空间的内容每次都会打印到标准输出, 如果要把输出保存到文件同时不显示到屏幕上,还需要使用-n选项;
> * 把example.txt的内容保存到文件output.txt,同时显示在屏幕上
```
$ sed 'w output.txt' example.txt 
 101,John Doe,CEO 
 102,Jason Smith,IT Manager 
 103,Raj Reddy,Sysadmin 
 104,Anand Ram,Developer 
 105,Jane Miller,Sales Manager 
$ cat output.txt 
 101,John Doe,CEO 
 102,Jason Smith,IT Manager 
 103,Raj Reddy,Sysadmin 
 104,Anand Ram,Developer 
 105,Jane Miller,Sales Manager
```
> * 把example.txt的内容保存到文件output.txt, 但不在屏幕上显示
```
$ sed -n 'w output.txt' example.txt 
$ cat output.txt 
 101,John Doe,CEO 
 102,Jason Smith,IT Manager 
 103,Raj Reddy,Sysadmin 
 104,Anand Ram,Developer 
 105,Jane Miller,Sales Manager
```
> * 只保存第2行:
```
$ sed -n '2 w output.txt' example.txt 
$ cat output.txt 
 102,Jason Smith,IT Manager
```
> * 保存第1至第4行
```
$ sed -n '1,4 w output.txt' example.txt 
$ cat output.txt 
 101,John Doe,CEO 
 102,Jason Smith,IT Manager 
 103,Raj Reddy,Sysadmin 
 104,Anand Ram,Developer
```
> * 保存第2行起至最后一行:
```
$ sed -n '2,$ w output.txt' example.txt 
$ cat output.txt 
 102,Jason Smith,IT Manager 
 103,Raj Reddy,Sysadmin 
 104,Anand Ram,Developer 
 105,Jane Miller,Sales Manager
```
> * 只保存奇数行
```
$ sed -n '1~2 w output.txt' example.txt 
$ cat output.txt                       
 101,John Doe,CEO 
 103,Raj Reddy,Sysadmin 
 105,Jane Miller,Sales Manager
```
> * 保存匹配Jane的行:
```
$ sed -n '/Jane/w output.txt' example.txt
$ cat output.txt                          
 105,Jane Miller,Sales Manager
```
> * 保存第一次匹配Jason的行至第4行:
```
$ sed -n '/Jason/,4w output.txt' example.txt
$ cat output.txt 
 102,Jason Smith,IT Manager 
 103,Raj Reddy,Sysadmin 
 104,Anand Ram,Developer
注意: 如果开始的4行里没有匹配到Jason,那么该命令只保存第4行以后匹配到Jason行 
```
> * 保存第一次匹配Raj的行至最后一行:
```
$ sed -n '/Raj/,$w output.txt' example.txt
$ cat output.txt                           
 103,Raj Reddy,Sysadmin 
 104,Anand Ram,Developer 
 105,Jane Miller,Sales Manager
```
> * 保存匹配Raj的行至匹配Jane的行
```
$ sed -n '/Raj/,/Jane/w output.txt' example.txt 
$ cat output.txt 
 103,Raj Reddy,Sysadmin 
 104,Anand Ram,Developer 
 105,Jane Miller,Sales Manager
```
> * 保存匹配Jason的行以及紧跟在其后面的两行:
```
$ sed -n '/Jason/,+2w output.txt' example.txt
$ cat output.txt                              
 102,Jason Smith,IT Manager 
 103,Raj Reddy,Sysadmin 
 104,Anand Ram,Developer
```

## Sed替换

> * 流编辑器中最强大的功能就是替换(substitute)
> * sed替换命令语法:
```
sed '[address-range|pattern-range] s/original-string/replacement-string/[substitute-flags]' input-file

# address-range或pattern-range: 即地址范围和模式范围是可选的,若无,那么sed将在所有行上进行替换;
# s: 即执行替换命令substitute
# original-string: 是被sed搜索然后被替换的字符串,它可以是一个正则表达式;
# replacement-string: 替换后的字符串
# substitute-flags: 可选的, 下面具体解释
# 谨记: 原始输入文件不会被修改, sed只在模式空间中执行替换命令, 然后输出模式空间内容;
```
> * 用Director替换所有行中的Manager:
```
$ cat example.txt 
 101,John Doe,CEO 
 102,Jason Smith,IT Manager 
 103,Raj Reddy,Sysadmin 
 104,Anand Ram,Developer 
 105,Jane Miller,Sales Manager 

$ sed 's#Manager#Director#' example.txt 
 101,John Doe,CEO 
 102,Jason Smith,IT Director 
 103,Raj Reddy,Sysadmin 
 104,Anand Ram,Developer 
 105,Jane Miller,Sales Director
```
> * 只把包含Sales的行中的Manager替换为Director:
```
$ cat example.txt 
 101,John Doe,CEO 
 102,Jason Smith,IT Manager 
 103,Raj Reddy,Sysadmin 
 104,Anand Ram,Developer 
 105,Jane Miller,Sales Manager
$ sed '/Sales/ s/Manager/Director/' example.txt   
 101,John Doe,CEO 
 102,Jason Smith,IT Manager 
 103,Raj Reddy,Sysadmin 
 104,Anand Ram,Developer 
 105,Jane Miller,Sales Director
注意：本例由于使用了地址范围限制，所以只有一个Manager被替换了
```

### 全局标志g

> * g代表全局(global) 默认情况下, sed只会替换每行中第一次出现的original-string;
> * 如果需要替换每行中出现的所有original-string,就需要使用g;
> * 用大写A替换第一次出现的小写字母a:
```
$ cat example.txt 
 101,John Doe,CEO 
 102,Jason Smith,IT Manager 
 103,Raj Reddy,Sysadmin 
 104,Anand Ram,Developer 
 105,Jane Miller,Sales Manager 

# 替换每行第一次出现的a为A
$ sed 's#a#A#' example.txt 
 101,John Doe,CEO 
 102,JAson Smith,IT Manager 
 103,RAj Reddy,Sysadmin 
 104,AnAnd Ram,Developer 
 105,JAne Miller,Sales Manager
注意：上述例子会在所有行上替换，因为没有指定地址范围

# 替换每行中所有的a为A
$ sed 's#a#A#g' example.txt
 101,John Doe,CEO 
 102,JAson Smith,IT MAnAger 
 103,RAj Reddy,SysAdmin 
 104,AnAnd RAm,Developer 
 105,JAne Miller,SAles MAnAger
```

###  数字标志(1,2,3...)

> * 使用数字可以指定original-string出现的次序, 只有第n次出现的original-string才会触发替换,;
> * 每行的数字从1开始, 最大为512;
> * 比如/11 会替换每行中第11次出现的original-string
> * 把第二次出现的小写字母a替换为大写字母A:
```
$ cat example.txt 
 101,John Doe,CEO 
 102,Jason Smith,IT Manager 
 103,Raj Reddy,Sysadmin 
 104,Anand Ram,Developer 
 105,Jane Miller,Sales Manager 
$ sed 's#a#A#2' example.txt 
 101,John Doe,CEO 
 102,Jason Smith,IT MAnager 
 103,Raj Reddy,SysAdmin 
 104,Anand RAm,Developer 
 105,Jane Miller,SAles Manager
```
> * 新建文件, 把每行中第二次出现的locate替换为find:
```
$ cat >sub.txt
locate command is used to locate files 
locate command uses database to locate files 
locate command can also use regex for searching 
$ sed 's#locate#find#2' sub.txt  
locate command is used to find files 
locate command uses database to find files 
locate command can also use regex for searching
注意：第3行中locate只出现了一次，所以没有替换任何内容
```

### 打印标志p(print)

> * 命令p代表打印,当替换操作完成后,打印替换的行;
> * 与其他打印命令类似,sed中比较有用的方法是和`-n`一起使用,用来抑制默认的打印操作;
> * 只打印替换后的行:
```
$ cat example.txt 
 101,John Doe,CEO 
 102,Jason Smith,IT Manager 
 103,Raj Reddy,Sysadmin 
 104,Anand Ram,Developer 
 105,Jane Miller,Sales Manager
$ sed -n 's#John#Johnny#p' example.txt 
 101,Johnny Doe,CEO
```
> * 把每行中第二次出现的locate替换为find并打印出来:
```
$ cat sub.txt
locate command is used to locate files 
locate command uses database to locate files 
locate command can also use regex for searching 

$ sed -n 's#locate#find#2p' sub.txt 
locate command is used to find files 
locate command uses database to find files
```

### 写标志w

> * 标志w代表write, 当替换操作执行成功后, 它把替换后的结果保存的文件中;
> * 多数人更倾向于使用p打印内容, 然后重定向到文件中;
> * 只把替换后的内容写到output.txt中:
```
$ cat example.txt 
 101,John Doe,CEO 
 102,Jason Smith,IT Manager 
 103,Raj Reddy,Sysadmin 
 104,Anand Ram,Developer 
 105,Jane Miller,Sales Manager 
$ sed -n 's#John#Johnny#w output.txt' example.txt 
$ cat output.txt 
 101,Johnny Doe,CEO
```
> * 把每行第二次出现的locate替换为find, 把替换的结果保存到文件中, 同时显示输入文件所有内容
```
$ cat sub.txt 
locate command is used to locate files 
locate command uses database to locate files 
locate command can also use regex for searching

$ sed -n 's#locate#find#2w output.txt' sub.txt 
$ cat output.txt 
locate command is used to find files 
locate command uses database to find files
```

### 忽略大小写标志i (ignore)

> * 替换标志i代表忽略大小写, 可以使用i来以小写字符的模式匹配original-string;
> * 该标志只有GNU Sed中才可使用
> * 下面的例子不会把John替换为Johnny,因为original-string字符串是小写形式:
```
## 不加i
$ cat example.txt 
 101,John Doe,CEO 
 102,Jason Smith,IT Manager 
 103,Raj Reddy,Sysadmin 
 104,Anand Ram,Developer 
 105,Jane Miller,Sales Manager

$ sed 's#John#Johnny#' example.txt    
 101,Johnny Doe,CEO 
 102,Jason Smith,IT Manager 
 103,Raj Reddy,Sysadmin 
 104,Anand Ram,Developer 
 105,Jane Miller,Sales Manager

## 加i
$ sed 's#john#Johnny#i' example.txt  
 101,Johnny Doe,CEO 
 102,Jason Smith,IT Manager 
 103,Raj Reddy,Sysadmin 
 104,Anand Ram,Developer 
 105,Jane Miller,Sales Manager
```

### 执行命令标志e (excuate)

> * 替换标志e代表执行(execute),该标志可以将模式空间中的任何内容当做shell命令执行,并把命令执行的结果返回到模式空间;
> * 该标志只有GNU Sed中才可使用;
> * 创建测试文件:
```
$ cat file.txt
/etc/passwd
/etc/group
```
> * 在files.txt文件中的每行前面添加ls –l 并打结果:
```
sed 's#^#ls -l #' file.txt 
ls -l /etc/passwd
ls -l /etc/group
```
> * 使用标志e把结果作为命令执行:
```
$ sed 's#^#ls -l #e' file.txt 
-rw-r--r-- 1 root root 1402 Jul 16 10:00 /etc/passwd
-rw-r--r-- 1 root root 652 Jul 16 10:00 /etc/group
```

### 使用替换标志组合

> * 根据需要可以把一个或多个替换标志组合起来使用
> * 把每行中出现的所有Manager或manager替换为Director, 然后把替换后的内容打印到屏幕上, 同时把这些内容保存到output.txt文件中: **使用g,I,p和w的组合**
```
$ cat example.txt 
 101,John Doe,CEO 
 102,Jason Smith,IT Manager 
 103,Raj Reddy,Sysadmin 
 104,Anand Ram,Developer 
 105,Jane Miller,Sales Manager

$ sed -n 's#manager#Director#igpw output.txt' example.txt 
 102,Jason Smith,IT Director 
 105,Jane Miller,Sales Director 

$ cat output.txt 
 102,Jason Smith,IT Director 
 105,Jane Miller,Sales Director	
```

### 替换命令分界符

> * 以上都是用了sed默认的分界符/, 即`s/original-string/replacement-string/g`
> * 如果在`original-string`或`replacement-string`中有`/`,那么需要使用反斜杠\来转义;
> * 创建测试文件:
```
$ cat path.txt 
reading /usr/local/bin directory
```
> * 限制使用sed把/usr/local/bin替换为/usr/bin, 以下sed默认的分界符/都被\转义了:
```
$ sed 's/\/usr\/local\/bin/\/usr\/bin/' path.txt 
reading /usr/bin directory
```
> * 可以使用任何一个字符作为sed替换命令的分界符, 如|或 ^ 或@或者 !
 - sed 's|/usr/local/bin|/usr/bin|' path.txt 
 - sed 's^/usr/local/bin^/usr/bin^' path.txt 
 - sed 's@/usr/local/bin@/usr/bin@' path.txt 
 - sed 's!/usr/local/bin!/usr/bin!' path.txt

### 单行内容上执行多个命令

> * sed执行的过程是读取内容、执行命令、打印结果、重复循环;
> * 其中执行命令部分,可以由多个命令执行,sed将一个一个地依次执行它们;
> * 假如: 你有两个命令,sed将在模式空间中执行第一个命令,然后执行第二个命令;如果第一个命令改变了模式空间的内容,第二个命令会在改变后的模式空间上执行(此时模式空间的内容已经不是最开始读取进来的内容了)
> * 以下例子演示了在模式空间内执行两个替换命令的过程:
```
$ cat example.txt 
 101,John Doe,CEO 
 102,Jason Smith,IT Manager 
 103,Raj Reddy,Sysadmin 
 104,Anand Ram,Developer 
 105,Jane Miller,Sales Manager 
$ sed '{
s/Developer/IT Manager/ 
s/Manager/Director/ 
}' example.txt
 101,John Doe,CEO 
 102,Jason Smith,IT Director 
 103,Raj Reddy,Sysadmin 
 104,Anand Ram,IT Director 
 105,Jane Miller,Sales Director
# 分析下第4行的执行过程： 
1.读取数据: 在这一步,sed读取内容到模式空间,此时模式空间的内容为: 104,Anand Ram,Developer 
2.执行命令：第一个命令,s/Developer/IT Manager/执行后,模式空间的内容为: 104,Anand Ram,IT Manager 
 现在在模式空间上执行第二个命令s/Manager/Director/,执行后模式空间内容为: 104,Anand Ram,IT Director 
 谨记：sed在第一个命令执行的结果上,执行第二个命令。 
3.打印内容,打印当前模式空间的内容: 104,Anand Ram,IT Director 
4.重复循环:移动的输入文件的下一行,然后重复执行第一步,即读取数据 
```

### 获取匹配到的模式(&的作用)

> * 当在replacement-string中使用&时, 它会被替换成匹配到的original-string或正则表达式, 这是个很有用的东西;
> * 案例: 给雇员ID(即第一列的3个数字)加上[],如101改成[101]
```
$ cat example.txt 
 101,John Doe,CEO 
 102,Jason Smith,IT Manager 
 103,Raj Reddy,Sysadmin 
 104,Anand Ram,Developer 
 105,Jane Miller,Sales Manager

$ sed 's#[0-9][0-9][0-9]#[&]#g' example.txt    
 [101],John Doe,CEO 
 [102],Jason Smith,IT Manager 
 [103],Raj Reddy,Sysadmin 
 [104],Anand Ram,Developer 
 [105],Jane Miller,Sales Manager
```
> * 把每一行放进<>中:
```
$ cat example.txt 
 101,John Doe,CEO 
 102,Jason Smith,IT Manager 
 103,Raj Reddy,Sysadmin 
 104,Anand Ram,Developer 
 105,Jane Miller,Sales Manager 
$ sed 's#^.*#<&>#' example.txt 
< 101,John Doe,CEO >
< 102,Jason Smith,IT Manager >
< 103,Raj Reddy,Sysadmin >
< 104,Anand Ram,Developer >
< 105,Jane Miller,Sales Manager >
```

### 分组替换(单个分组)

> * 跟在正则表达式中一样,sed中也可以使用分组,分组以\(开始,以\)结束,分组可以用在回溯引用中;
> * 回溯引用即重新使用分组所选择的部分正则表达式,在sed替换命令的replacement-string
中和正则表达式中,都可以使用回溯引用;
> * 单个分组:
```
$ cat example.txt 
 101,John Doe,CEO 
 102,Jason Smith,IT Manager 
 103,Raj Reddy,Sysadmin 
 104,Anand Ram,Developer 
 105,Jane Miller,Sales Manager
$ sed 's#\([^,]*\).*#\1#g' example.txt  
 101
 102
 103
 104
 105
# 正则表达式\([^,]*\)匹配字符串从开头到第一个逗号之间的所有字符(并将其放入第一个分组中) 
# replacement-string中的\1 将替代匹配到的分组 
# g即是全局标志
```
> * 以下例子,如果单词第一个字符为大写,那么会给这个大写字符加上()
```
$  echo "The Geek Stuff"|sed 's#\([A-Z]\)#\(\1\)#g'    
(T)he (G)eek (S)tuff
```
> * 创建测试文件
```
$ cat num.txt 
1 
12 
123 
1234 
12345 
123456
```
> * 格式化数字, 增加其可读性:
```
$ sed 's#\(^\|[^0-9.]\)\([0-9]\+\)\([0-9]\{3\}\)#\1\2,\3#g' num.txt    
1 
12 
123 
1,234 
12,345 
123,456
```

### 分组替换(多个分组)

> * 使用多个\(和\)划分多个分组,使用多个分组时,需要在replacement-string中使用\n来指定第n个分组;
> * 只打印第一列(雇员ID)和第三列(雇员职位):
```
$ cat example.txt 
 101,John Doe,CEO 
 102,Jason Smith,IT Manager 
 103,Raj Reddy,Sysadmin 
 104,Anand Ram,Developer 
 105,Jane Miller,Sales Manager
$ sed 's/^\([^,]*\),\([^,]*\),\([^,]*\)/\1,\3/' example.txt 
 101,CEO 
 102,IT Manager 
 103,Sysadmin 
 104,Developer 
 105,Sales Manager
在这个例子中，可以看到，original-string中，划分了3个分组，以逗号分隔。 
 ([^,]*\) 第一个分组，匹配雇员ID 
，为字段分隔符 
 ([^,]*\) 第二个分组，匹配雇员姓名 
，为字段分隔符 
 ([^,]*\) 第三个分组，匹配雇员职位 
，为字段分隔符，上面的例子演示了如何使用分组 
\1 代表第一个分组(雇员ID) 
，出现在第一个分组之后的逗号 
\3 代表第二个分组(雇员职位) 
```
> * 交换第一列(雇员ID)和第二列(雇员姓名):
```
$ sed 's/^\([^,]*\),\([^,]*\),\([^,]*\)/\2,\1,\3/' example.txt 
John Doe, 101,CEO 
Jason Smith, 102,IT Manager 
Raj Reddy, 103,Sysadmin 
Anand Ram, 104,Developer 
Jane Miller, 105,Sales Manager
注意：sed最多能处理9个分组，分别用\1至\9表示。
```

### GNU Sed专有的替换标志

> * 下面的标志,只有GNU版的sed才能使用,它们可以用在替换命令中的replacement-string里面;
> * `\l`标志
 - 当在replacement-string中使用\l标志时, 它会把紧跟在其后面的字符当做小写字符来处理;
 - 下面的例子将把John换成JOHNNY:
```
$ cat example.txt 
 101,John Doe,CEO 
 102,Jason Smith,IT Manager 
 103,Raj Reddy,Sysadmin 
 104,Anand Ram,Developer 
 105,Jane Miller,Sales Manager 
$ sed 's#John#JOHNNY#' example.txt 
 101,JOHNNY Doe,CEO 
 102,Jason Smith,IT Manager 
 103,Raj Reddy,Sysadmin 
 104,Anand Ram,Developer 
 105,Jane Miller,Sales Manage
$ sed -n 's#John#JO\lHNNY#p' example.txt  
101,JOhNNY Doe,CEO
```
> * \L标志
 - 当在replacement-string中使用\l标志时, 它会把后面所有的字符都当做小写字符来处理;
 - 下面的例子, 在replacement-string中的H前面放置了\L标志,它会把H和它后面的所有字符都换成小写;
```
$ sed -n 's#John#JO\LHNNY#p' example.txt        
 101,JOhnny Doe,CEO
```
> * \u标志
 - 和\l类似, 只不过是把字符换成大写; 
 - 当在replacement-string中使用\u标志时，它会把紧跟在其后面的字符当做大写字符来处理; 
 - 下面的例子中, replacement-string里面的h前面有 \u标志, 所以h将被换成大写的H:
```
$ sed -n 's#John#jo\uhnny#p' example.txt       
 101,joHnny Doe,CEO
```
> * \U标志
 - 当在replacement-string中使用\U标志时, 它会把后面所有的字符都当做大写字符来处理; 
 - 下面的例子中, replacement-string里面的h前面有\U标志, 所以h及其以后的所有字符, 都将被换成大写:
```
$ sed -n 's#John#jo\Uhnny Boy#p' example.txt   
 101,joHNNY BOY Doe,CEO
```
> * \E标志
 - \E标志需要和\U或\L一起使用, 它将关闭\U或\L的功能; 
 - 下面的例子将把字符串"Johnny Boy"的每个字符都以大写的形式打印出来, 因为在replacement-string前面使用了\U标志:
 - 下面将把John换成JOHNNY Boy:
```
$ sed -n 's#John#\Ujohnny\E Boy#p' example.txt    
 101,JOHNNY Boy Doe,CEO
这个例子只把Johnny显示为大写，因为在Johnny后面使用了\E标志(关闭了\U的功能)
```
> * 替换标志的用法
 - 上面的例子仅仅展示了这些标志的用法和功能, 如果你使用的是具体的字符串, 那么这些选项未必有什么作用, 因为你可以在需要的地方写出精确的字符串, 而不需要使用这些标志进行转换;
 - 和分组配合使用时, 这些选项就显得很有用了, 前面例子中我们已经学会了如何使用分组调换第一列和第三列的位置, 使用上述标志, 可以把整个分组转换为小写或大写;
 - 下面的例子, 雇员ID都显示为大写, 职位都显示为小写:
```
$ cat example.txt 
 101,John Doe,CEO 
 102,Jason Smith,IT Manager 
 103,Raj Reddy,Sysadmin 
 104,Anand Ram,Developer 
 105,Jane Miller,Sales Manager
$ sed 's#\([^,]*\),\([^,]*\),\([^,]*\)#\U\2\E,\1,\L\3#' example.txt 
JOHN DOE, 101,ceo 
JASON SMITH, 102,it manager 
RAJ REDDY, 103,sysadmin 
ANAND RAM, 104,developer 
JANE MILLER, 105,sales manager
这个例子中： 
\U\2\E 把第二个分组转换为大写，然后用\E关闭转换 
\L\3 把第三个分组转换为小写
```
> * 使用sed把DOS格式的文件转换为Unix格式: 
 - sed 's#.$##' filename

## 执行sed

### 单行内执行多个sed命令

> * 使用多命令选项 -e
 - 多命令选项-e使用方法如下: sed -e 'command1' -e 'command2' -e 'command3'
 - 在/etc/passwd文件中, 搜索root、nobody或mail:
```
$ sed -n -e '/^root/ p' -e '/^nobody/ p' -e '/^mail/ p' /etc/passwd 
root:x:0:0:root:/root:/bin/bash
mail:x:8:12:mail:/var/spool/mail:/sbin/nologin
nobody:x:99:99:Nobody:/:/sbin/nologin
``` 
> * 使用\折行执行多个命令
 - 在执行很长的命令,比如使用-e选项执行多个sed命令时,可以使用\来把命令折到多行 
> * 使用{}把多个命令组合
 - 如果要执行很多sed命令,可以使用{}把他们组合起来执行: 
```
$ sed -n '{       
/^root/p
/^nobody/p
/^mail/p
}' /etc/passwd
root:x:0:0:root:/root:/bin/bash
mail:x:8:12:mail:/var/spool/mail:/sbin/nologin
nobody:x:99:99:Nobody:/:/sbin/nologin
``` 

###  sed脚本文件

> * 如果用重复使用一组sed命令,那么可以建立sed脚本文件,里面包含所有要执行的sed命令,然后用-f选项来使用;
> * 首先建立下面文件,里面包含了所有要执行的sed命令;
```
$ cat example.txt 
 101,John Doe,CEO 
 102,Jason Smith,IT Manager 
 103,Raj Reddy,Sysadmin 
 104,Anand Ram,Developer 
 105,Jane Miller,Sales Manager

$ cat cmd.sed 
s/\([^,]*\),\([^,]*\),\(.*\).*/\2,\1, \3/g 
s/^.*/<&>/ 
s/Developer/IT Manager/ 
s/Manager/Director/

$ sed -f cmd.sed example.txt 
<John Doe, 101, CEO >
<Jason Smith, 102, IT Director >
<Raj Reddy, 103, Sysadmin >
<Anand Ram, 104, IT Director >
<Jane Miller, 105, Sales Director >
```

###  sed注释

> * sed注释以#开头,因为sed是比较晦涩难懂的语言,因此建议把写脚本时的初衷作为注释,写到脚本里面:
> * 注释信息如下:
```
$ cat cmd.sed
#交换第一列和第二列
s/\([^,]*\),\([^,]*\),\(.*\).*/\2,\1, \3/g
#把整行内容放入<>中
s/^.*/<&>/
#把Developer替换为IT Manager
s/Developer/IT Manager/
#把Manager替换为Director
s/Manager/Director/
注意：如果sed脚本第一行开始的两个字符是#n的话，sed会自动使用-n选项(即不自动打印模式空间的内容)
```

### 把sed当做命令解释器使用

> * 我们可以把命令放进一个shell脚本中,然后调用脚本名称来执行它们一样;
> * 你也可以把sed用作命令解释器,要实现这个功能,需要在sed脚本最开始加入`#!/bin/sed –f`,如下所示:
```
$ cat mycmd.sed 
#!/bin/sed -f
#交换第一列和第二列 
s/\([^,]*\),\([^,]*\),\(.*\).*/\2,\1, \3/g 
#把整行内容放入<>中 
s/^.*/<&>/ 
#把Developer替换为IT Manager 
s/Developer/IT Manager/ 
#把Manager替换为Director 
s/Manager/Director/
```
> * 给这个脚本加上可执行权限,然后直接在命令行调用它:
```
$ chmod u+x mycmd.sed 
$ ./mycmd.sed example.txt 
<John Doe, 101, CEO >
<Jason Smith, 102, IT Director >
<Raj Reddy, 103, Sysadmin >
<Anand Ram, 104, IT Director >
<Jane Miller, 105, Sales Director > 
``` 
> * 指定-n选项来屏蔽默认输出:
```
$ vim scripts.sed
#!/bin/sed -nf
/^root/p
/nobody/p
/mail/p 
$ chmod u+x scripts.sed 
$ ./scripts.sed /etc/passwd 
root:x:0:0:root:/root:/bin/bash
mail:x:8:12:mail:/var/spool/mail:/sbin/nologin
nobody:x:99:99:Nobody:/:/sbin/nologin
``` 

### 直接修改输入文件

> * sed默认不会修改输入文件,它只会把输出打印到标准输出上,当想保存结果时,把输出重定向到文件中(或使用w命令);
> * 执行下面的例子之前,先备份一下example.txt文件:
```
$ cp example.txt{,.ori}
```
> * 为了修改输入文件,通常方法是把输出重定向到一个临时文件,然后重命名该临时文件:
```
$ sed 's/John/Johnny/' example.txt >example.txt.tmp
$ mv example.txt.tmp example.txt
```
> * 可以在sed命令中使用-i选项, 使sed可以直接修改输入文件:
 - 在原始文件example.txt中, 把John替换为Johnny:
```
$ cat example.txt                     
 101,John Doe,CEO 
 102,Jason Smith,IT Manager 
 103,Raj Reddy,Sysadmin 
 104,Anand Ram,Developer 
 105,Jane Miller,Sales Manager
 
$ sed -i 's/John/Johnny/' example.txt
$ cat example.txt
 101,Johnny Doe,CEO 
 102,Jason Smith,IT Manager 
 103,Raj Reddy,Sysadmin 
 104,Anand Ram,Developer 
 105,Jane Miller,Sales Manager
# -i会修改输入文件,一个保护性的措施是在-i后面加上备份扩展,这样sed就会在修改原始文件之前,备份一份;
```
> * 在原始文件example.txt中,把John替换为Johnny,但在替换前备份example.txt:
```
$ sed -i.bak 's#John#Johnny#g' example.txt

#备份的文件如下:
$ cat example.txt.bak 
 101,John Doe,CEO 
 102,Jason Smith,IT Manager 
 103,Raj Reddy,Sysadmin 
 104,Anand Ram,Developer 
 105,Jane Miller,Sales Manager
	
#修改后的原始文件为:
$ cat example.txt
 101,Johnny Doe,CEO 
 102,Jason Smith,IT Manager 
 103,Raj Reddy,Sysadmin 
 104,Anand Ram,Developer 
 105,Jane Miller,Sales Manager

# 除了使用-i,也可以使用完整样式-in-place,下面两个命令是等价的:
sed -i.bak 's/John/Johnny/' example.txt 
sed -in-place=bak 's/John/Johnny/' example.txt
```

## sed附加命令

### 追加命令a

> * 使用命令a可以在指定位置的后面插入新行;
> * 语法:`sed '[address] a the-line-to-append' input-file`
> * 在第2行后面追加一行
```
$ cat example.txt
 101,Johnny Doe,CEO 
 102,Jason Smith,IT Manager 
 103,Raj Reddy,Sysadmin 
 104,Anand Ram,Developer 
 105,Jane Miller,Sales Manager

$ sed '2a 203,Jack Johnson,Engineer' example.txt  
 101,Johnny Doe,CEO 
 102,Jason Smith,IT Manager 
 203,Jack Johnson,Engineer
 103,Raj Reddy,Sysadmin 
 104,Anand Ram,Developer 
 105,Jane Miller,Sales Manager
```
> * 在example.txt文件结尾追加一行:
```
$ cat example.txt
 101,Johnny Doe,CEO 
 102,Jason Smith,IT Manager 
 103,Raj Reddy,Sysadmin 
 104,Anand Ram,Developer 
 105,Jane Miller,Sales Manager

$ sed '$a 106,Jack Johnson,Engineer' example.txt
 101,Johnny Doe,CEO 
 102,Jason Smith,IT Manager 
 103,Raj Reddy,Sysadmin 
 104,Anand Ram,Developer 
 105,Jane Miller,Sales Manager 
 106,Jack Johnson,Engineer
```
> * sed也可以追加多行:
```
$ sed '/Jason/a \
203,Jack Johnson,Engineer \
204,Mark Smith,Sales Engineer' \
example.txt
101,Johnny Doe,CEO 
102,Jason Smith,IT Manager 
203,Jack Johnson,Engineer 
204,Mark Smith,Sales Engineer
103,Raj Reddy,Sysadmin 
104,Anand Ram,Developer 
105,Jane Miller,Sales Manager

# 追加多行之间可以用\n来换行
sed '/Jason/a 203,Jack Johnson,Engineer\n204,Mark Smith,Sales Engineer' example.txt
```

### 插入命令i

> * 插入命令insert命令和追加命令类似,只不过是在指定位置之前插入行;
> * 语法: `sed '[address] i the-line-to-insert' input-file`
> * 在example.txt的第2行之前插入一行:
```
$ sed '2 i 203,Jack Johnson,Engineer' example.txt 
101,Johnny Doe,CEO 
203,Jack Johnson,Engineer
102,Jason Smith,IT Manager 
103,Raj Reddy,Sysadmin 
104,Anand Ram,Developer 
105,Jane Miller,Sales Manager
```
> * 在example.txt最后一行之前, 插入一行:
```
$ sed '$i 108,Jack Johnson,Engineer' example.txt 
101,Johnny Doe,CEO 
102,Jason Smith,IT Manager 
103,Raj Reddy,Sysadmin 
104,Anand Ram,Developer 
108,Jack Johnson,Engineer
105,Jane Miller,Sales Manager
```
> * sed也可以插入多行:
```
$ sed '/Jason/i \
203,Jack Johnson,Engineer \
204,Mark Smith,Sales Engineer' \
example.txt

101,Johnny Doe,CEO 
203,Jack Johnson,Engineer 
204,Mark Smith,Sales Engineer
102,Jason Smith,IT Manager 
103,Raj Reddy,Sysadmin 
104,Anand Ram,Developer 
105,Jane Miller,Sales Manager
```

### 修改命令c

> * 修改命令change可以用新行取代旧行;
> * 语法: sed '[address] c the-line-to-insert’ input-file
> * 用新数据取代第2行:
```
$ sed '2 c 202,Jack,Johnson,Engineer' example.txt   
101,Johnny Doe,CEO 
202,Jack,Johnson,Engineer
103,Raj Reddy,Sysadmin 
104,Anand Ram,Developer 
105,Jane Miller,Sales Manager

#命令c等价于替换下面
$ sed '2s/.*/202,Jack,Johnson,Engineer/' example.txt 
101,Johnny Doe,CEO 
202,Jack,Johnson,Engineer
103,Raj Reddy,Sysadmin 
104,Anand Ram,Developer 
105,Jane Miller,Sales Manager
```

### 命令a、i和c组合使用

> * 命令a、i和c可以组合使用,下面的例子将完成三个操作:
 - a 在”Jason”后面追加”Jack Johnson”
 - i 在”Jason”前面插入”Mark Smith” 
 - c 用”Joe Mason”替代”Jason” 
```
$ sed '/Jason/{
a\
204,Jack Johnson,Engineer
i\
202,Mark Smith,Sales Engineer
c\
203,Joe Mason,Sysadmin
}' example.txt
101,Johnny Doe,CEO 
202,Mark Smith,Sales Engineer
203,Joe Mason,Sysadmin
204,Jack Johnson,Engineer
103,Raj Reddy,Sysadmin 
104,Anand Ram,Developer 
105,Jane Miller,Sales Manager
```

### 打印不可见字符(命令l)

> * 命令l可以打印不可见的字符,比如制表符\t,行尾标志$等;
> * 创建测试文件,请确保字段之间使用制表符(Tab键)分开:
```
$ cat tabfile.txt 
fname   First Name
lname   Last Name
mname   Middle Name
```
> * 使用命令l, 把制表符显示为\t,行尾标志显示为EOL:
```
$ sed -n 'l' tabfile.txt 
fname\tFirst Name$
lname\tLast Name$
mname\tMiddle Name$
```
> * 如果在l后面指定了数字,那么会在第n个字符处使用一个不可见自动折行,效果如下:
```
$ sed -n 'l20' example.txt  
101,John Doe,CEO $
102,Jason Smith,IT \
Manager $
103,Raj Reddy,Sysad\
min $
104,Anand Ram,Devel\
oper $
105,Jane Miller,Sal\
es Manager$
这个功能只有GNU sed才有。
```

### 打印行号(命令=)

> * 命令=会在每一行前面显示该行的行号:
```
$ sed '=' example.txt       
1
101,John Doe,CEO 
2
102,Jason Smith,IT Manager 
3
103,Raj Reddy,Sysadmin 
4
104,Anand Ram,Developer 
5
105,Jane Miller,Sales Manager
```
> * 只打印1,2,3行的行号:
```
$ sed '1,3 =' example.txt 
1
101,John Doe,CEO 
2
102,Jason Smith,IT Manager 
3
103,Raj Reddy,Sysadmin 
104,Anand Ram,Developer 
105,Jane Miller,Sales Manager
```
> * 打印包含关键字”Jane”的行的行号，同时打印输入文件中的内容：
```
$ sed '/Jane/ =' example.txt 
101,John Doe,CEO 
102,Jason Smith,IT Manager 
103,Raj Reddy,Sysadmin 
104,Anand Ram,Developer 
5
105,Jane Miller,Sales Manager
```
> * 如果你想只显示行号但不显示行的内容，那么使用-n选项来配合命令=:
```
$ sed  -n '/Raj/ =' example.txt 
3
$ sed -n '=' /etc/services|tail -1 #统计services文件行数
10774
```

### 转换字符(命令y)

> * 命令y根据对应位置转换字符, 好处之一便是把大写字符转换为小写, 反之亦然
> * 下面例子中,将把a换为A,b换为B,c换为C,以此类推:
```
$ cat example.txt
101,John Doe,CEO 
102,Jason Smith,IT Manager 
103,Raj Reddy,Sysadmin 
104,Anand Ram,Developer 
105,Jane Miller,Sales Manager

$ sed 'y/abcde/ABCDE/' example.txt 
101,John DoE,CEO 
102,JAson Smith,IT MAnAgEr 
103,RAj REDDy,SysADmin 
104,AnAnD RAm,DEvElopEr 
105,JAnE MillEr,SAlEs MAnAgEr
```
> * 把所有小写字符转换为大写字符:
```
$ sed 'y/abcdefghijklmnopqrstuvwxyz/ABCDEFGHIJKLMNOPQRSTUVWXYZ/' example.txt 
101,JOHN DOE,CEO 
102,JASON SMITH,IT MANAGER 
103,RAJ REDDY,SYSADMIN 
104,ANAND RAM,DEVELOPER 
105,JANE MILLER,SALES MANAGER
```

### 操作多个文件

> * 之前都只操作了单个文件,sed 也可以同时处理多个文件;
> * 在/etc/passwd中搜索root并打印出来： 
```
$ sed -n '/root/ p' /etc/passwd  
root:x:0:0:root:/root:/bin/bash
operator:x:11:0:operator:/root:/sbin/nologin 
```
> * 在/etc/group中搜索root并打印出来： 
```
$ sed -n '/root/ p' /etc/group
root:x:0:
```
> * 同时在/etc/passwd和/etc/group中搜索root: 
```
$ sed -n '/^root/p' /etc/passwd /etc/group   
root:x:0:0:root:/root:/bin/bash
root:x:0:
```

### 退出sed (命令q)

> * 命令q终止正在执行的命令并退出sed;
> * 正常的sed执行流程是: 读取数据、执行命令、打印结果、重复循环; 
> * 当sed遇到q命令,便立刻退出,当前循环中的后续命令不会被执行,也不会继续循环;
> * 打印第1行后退出:
```
$ sed 'q' example.txt 
101,John Doe,CEO
```
> * 打印第3行后退出，即只打印前3行：
```
$ sed '3q' example.txt  
101,John Doe,CEO 
102,Jason Smith,IT Manager 
103,Raj Reddy,Sysadmin
```
> * 打印所有行，直到遇到包含关键字Manager的行:
```
$ sed '/Manager/ q' example.txt 
101,John Doe,CEO 
102,Jason Smith,IT Manager
注意：q命令不能指定地址范围(或模式范围),只能用于单个地址(或单个模式)。
```

### 从文件读取数据(命令r)

> * sed在处理输入文件时,命令r会从另外一个文件读取内容,并在指定的位置打印出来;
> * 下面的例子将读取log.txt的内容,并在打印example.txt最后一行之后,把读取的内容打印出来,事实上它把example.txt和log.txt合并然后打印出来;
``` 
$ cat log.txt 
oldboy
oldgirl

$ sed '$ r log.txt' example.txt 
101,John Doe,CEO 
102,Jason Smith,IT Manager 
103,Raj Reddy,Sysadmin 
104,Anand Ram,Developer 
105,Jane Miller,Sales Manager
oldboy
oldgirl
```
> * 给命令r指定一个模式,下面的例子将读取log.txt的内容,并且在匹配'Raj'的行后面打印出来:
```
$ cat log.txt                        
oldboy
oldgirl

$ cat example.txt                    
101,John Doe,CEO 
102,Jason Smith,IT Manager 
103,Raj Reddy,Sysadmin 
104,Anand Ram,Developer 
105,Jane Miller,Sales Manager

$ sed '/Raj/r log.txt' example.txt     
101,John Doe,CEO 
102,Jason Smith,IT Manager 
103,Raj Reddy,Sysadmin 
oldboy
oldgirl
104,Anand Ram,Developer 
105,Jane Miller,Sales Manager
```
### 打印模式空间(命令n)

> * 命令n打印当前模式空间的内容, 然后从输入文件中读取下一行; 如果在命令执行过程中遇到n,那么它会改变正常的执行流程;
```
$ sed  n example.txt    
101,Johnnyny Doe,CEO 
102,Jason Smith,IT Manager 
103,Raj Reddy,Sysadmin 
104,Anand Ram,Developer 
105,Jane Miller,Sales Manager
```
> * 如果使用了-n选项,将没有任何输出：
```
$ sed  -n n example.txt
不要把选项-n和命令n弄混了
```
> * sed正常流程是读取数据、执行命令、打印输出、重复循环; 
> * 命令n可以改变这个流程,它打印当前模式空间的内容,然后清除模式空间,读取下一行进来,然后继续执行后面的命令; 
> * 假定命令n前后各有两个其他命令, 如下: 
 ```
 sed -command-1 
 sed -command-2 
 n 
 sed -command-3 
 sed -command-4 

# 这种情况下,sed -command-1和sed -command-2会在当前模式空间中执行,然后遇到n,它打印当前模式空间的内容,并清空模式空间,读取下一行,然后把sed -command-3和sed -command-4应用于新的模式空间的内容; 
```
 
### 用sed 模拟Unix命令(cat,grep,read)

> * 之前的例子的完成的功能都很像标准的Unix命令,使用sed可以模拟很多Unix命令;
> * 模拟cat命令
```
$ cat example.txt 
101,John Doe,CEO 
102,Jason Smith,IT Manager 
103,Raj Reddy,Sysadmin 
104,Anand Ram,Developer 
105,Jane Miller,Sales Manager 
```
> * 下面的每个sed命令的输出都和上面的cat命令的输出一样： 
```
$ sed 's/JUNK/&/p' example.txt   
101,John Doe,CEO 
102,Jason Smith,IT Manager 
103,Raj Reddy,Sysadmin 
104,Anand Ram,Developer 
105,Jane Miller,Sales Manager

$ sed -n 'p' example.txt           
101,John Doe,CEO 
102,Jason Smith,IT Manager 
103,Raj Reddy,Sysadmin 
104,Anand Ram,Developer 
105,Jane Miller,Sales Manager

$ sed 'n' example.txt     
101,John Doe,CEO 
102,Jason Smith,IT Manager 
103,Raj Reddy,Sysadmin 
104,Anand Ram,Developer 
105,Jane Miller,Sales Manager

$ sed 'N' example.txt  
101,John Doe,CEO 
102,Jason Smith,IT Manager 
103,Raj Reddy,Sysadmin 
104,Anand Ram,Developer 
105,Jane Miller,Sales Manager 
```
> * 模拟grep命令 
```
$ grep Jane example.txt 
105,Jane Miller,Sales Manager
```
> * 下面的每个sed 命令的输出都和上面的grep命令的输出一样： 
```
$ sed -n 's/Jane/&/ p' example.txt
105,Jane Miller,Sales Manager
$ sed -n '/Jane/ p' example.txt   
105,Jane Miller,Sales Manager

# grep –v (打印不匹配的行): 
$ sed -n '/Jane/ !p' example.txt      
101,John Doe,CEO 
102,Jason Smith,IT Manager 
103,Raj Reddy,Sysadmin 
104,Anand Ram,Developer
注意：这里不能使用sed -n 's/Jane/&/ !p' 了。
```
> * 模拟head命令
```
$ head -10 /etc/passwd 
root:x:0:0:root:/root:/bin/bash
bin:x:1:1:bin:/bin:/sbin/nologin
daemon:x:2:2:daemon:/sbin:/sbin/nologin
adm:x:3:4:adm:/var/adm:/sbin/nologin
lp:x:4:7:lp:/var/spool/lpd:/sbin/nologin
sync:x:5:0:sync:/sbin:/bin/sync
shutdown:x:6:0:shutdown:/sbin:/sbin/shutdown
halt:x:7:0:halt:/sbin:/sbin/halt
mail:x:8:12:mail:/var/spool/mail:/sbin/nologin
uucp:x:10:14:uucp:/var/spool/uucp:/sbin/nologin
```
> * 下面的每个sed 命令的输出都和上面的head命令的输出一样：
```
$ sed '11,$d' /etc/passwd  
$ sed -n '1,10 p' /etc/passwd
$ sed '10q' /etc/passwd
```

## sed命令选项

### -n选项

> * 该选项屏蔽sed的默认输出,也可以使--quiet,或--silent来代替-n,它们的作用是相同的;
```
$ sed  -n 'p' example.txt 
$ sed  --quiet 'p' example.txt 
$ sed  --silent 'p' example.txt
```

### -f选项

> * 可以把多个sed命令保存在sed脚本文件中,然后使用-f选项来调用;
```
$ sed  -n -f test-script.sed  /etc/passwd 
$ sed  -n --file=test-script.sed  /etc/passwd
```

### -e选项

> * 该选项执行来自命令行的一个sed命令,可以使用多个-e来执行多个命令,也可以使用—expression来代替;
```
# 下面所有命令都是等价的: 
$ sed -n -e '/root/p' /etc/passwd
$ sed -n --expression '/root/p' /etc/passwd
```

### -i 选项

> * sed 不会修改输入文件,只会把内容打印到标准输出,或则使用w命令把内容写到不同的文件中;
> * 可以使用-i选项来直接修改输入文件;
> * 下面所有命令都是等价的:
```
# 在原始文件example.txt中, 用Johnny替换John
$ sed  -i 's/John/Johnny/' example.txt
# 执行和上面相同的命令，但在修改前备份原始文件:
$ sed  -i.bak 's/John/Johnny/' example.txt
# 也可以使用--in-place来代替-i: 以下两个是等价的
$ sed  -i.bak 's/John/Johnny/' example.txt 
$ sed  --in-place=bak 's/John/Johnny/' example.txt
```

## 保持空间与模式空间命令

> * sed有两个内置空间;
> * 模式空间:
 - 模式空间用来sed执行的正常流程中,该空间sed内置的一个缓冲区,用来存放、修改从输入文件读取的内容
> * 保持空间:
 - 保持空间是另外一个缓冲区,用来存放临时数据;sed可以在保持空间和模式空间交换数据,但是不能在保持空间上执行普通的sed命令,每次循环读取数据过程中,模式空间的内容都会被清空,然而保持空间的内容则保持不变,不会再循环中被删除;
> * 创建测试文件,用于保持空间的示例:
```
$ cat space.txt 
John Doe 
CEO 
Jason Smith 
IT Manager 
Raj Reddy 
Sysadmin 
Anand Ram 
Developer 
Jane Miller 
Sales Manager
```

### 用保持空间替换模式空间(命令x)

> * 命令x(Exchange)交换模式空间和保持空间的内容;
> * 假定目前模式空间内容为"line 1",保持空间内容为"line 2",那么执行命令x后,模式空间的内容变为"line2",保持空间的内容变为"line1";
> * 案例如下打印管理者的名称, 它搜索关键字'Manager'并打印之前的那一行:
```
$ sed -n -e '{x;n}' -e '/Manager/{x;p}' space.txt 
Jason Smith 
Jane Miller

$ sed  -n 'x;n;/Manager/{x;p}' space.txt   
Jason Smith 
Jane Miller

# 如果你的space.txt文件,雇员名称和职位不是连续的,那么得不到上面的结果;
# {x;n}: x交换模式空间和保持空间的内容,n读取下一行到模式空间, 在示例文件中保持空间保存的是雇员名称,模式空间保存的是职位。 
# /Manager/{x;}: 如果当前模式空间的内容包含关键'Manager', 那么就交换保持空间和模式空间的内容, 然后打印模式空间的内容, 这就意味着, 如果雇员职位中包含Manager,那么该雇员的名称将被打印出来;
```
> * 把上述命令保存在sed 脚本中, 然后执行:
```
$ vim x.sed 
#!/bin/sed  -nf 
x;n
/Manager/{x;p}
$ chmod u+x x.sed 
$ ./x.sed space.txt 
Jason Smith 
Jane Miller
```

### 把模式空间的内容复制到保持空间(命令h)

> * 命令h(hold)把模式空间的内容复制到保持空间, 和命令x不同, 命令h不会修改当前模式空间的内容; 执行命令h时, 当前保持空间的内容会被模式空间的内容覆盖;  
> * 假定目前模式空间内容为"line 1", 保持空间内容为"line 2", 那么执行命令h后, 模式空间的内容仍然为"line 1", 保持空间的内容则变为"line 1";
> * 打印管理者的名称:
```
$ sed  -n -e '/Manager/!h' -e'/Manager/{x;p}' space.txt           
Jason Smith 
Jane Miller 

/Manager/!h: 如果模式空间内容不包含关键字'Manager'(模式后面的!表示不匹配该模式), 那么复制模式空间内容到保持空间,(这样一来, 保持空间的内容可能会是雇员名称或职位, 而不是'Manager'), 这个例子中没有使用命令n来获取下一行, 而是通过正常的流程来读取后续内容。 

/Manager/{x;p}:  如果模式空间内容包含关键字'Manager'，那么交换保持空间和模式空间的内容，并打印模式空间的内容;
```

### 把模式空间内容追加到保持空间(命令H) 

> * 大写H命令表示把模式空间的内容追加到保持空间, 追加之前保持空间的内容不会被覆盖; 相反, 它在当前保持空间内容后面加上换行符\n, 然后把模式空间内容追加进来;  
> * 假定目前模式空间内容为"line 1", 保持空间内容为"line 2", 那么执行命令H后, 模式空间的内容没有改变, 仍然为"line 1", 保持空间的内容则变为"line2\nline 1" 
 > * 打印管理者的名称和职位(在不同的行上): 
```
$ sed -n -e '/Manager/!h' -e '/Manager/{H;x;p}' space.txt     
Jason Smith 
IT Manager 
Jane Miller 
Sales Manager

/Manager/!h: 如果模式空间内容不包含关键字'Manager'(模式后面的!表示不匹配该模式), 那么复制模式空间内容到保持空间;(这样一来, 保持空间的内容可能会是雇员名称或职位, 而不是'Manager'), 这和之前使用命令h的方法是一样的; 

/Manager/{H;x;p}: 如果模式空间内容包含关键字'Manager', 那么命令H把模式空间的内容(也就是管理者的职位)作为新行追加到保持空间, 所以保持空间内容会变为"雇员名称\n职位"(职位包含关键字Manager), 然后命令x交换模式空间和保持空间的内容, 随后命令p打印模式空间的内容;
```
> * 把雇员名称和职位显示在同一行, 以分号分开, 如下:
```
$ sed  -n -e '/Manager/!h' -e '/Manager/{H;x;;s/\n/:/p}' space.txt 
Jason Smith :IT Manager 
Jane Miller :Sales Manager

# 这个例子除了在第二个-e后面的命令中加入了替换命令之外, H、x和p都完成和之前相同的操作, 在交换模式空间和保持空间之后, 命令s把换行符\n替换为分号, 然后打印出来:
```

### 把保持空间内容复制到模式空间(命令g)

> * 命令g(get)把保持空间的内容复制到模式空间; 
  - 这样理解: 命令h保持(hold)住保持空间(hold space), 命令g从保持空间获取(get)内容; 
 > * 假定当前模式空间内容为"line 1", 保持空间内容为"line 2"; 执行命令g之后, 模式空间内容变为"line 2", 保持空间内容仍然为"line 2"; 
> * 打印管理者的名称: 
```
$ sed  -n -e '/Manager/!h' -e '/Manager/{g;p}' space.txt  
Jason Smith 
Jane Miller
 
/Manager/!h: 如果模式空间内容不包含关键字'Manager',那么就把他复制到保持空间。 
/Manager/{g;p}: 把保持空间的内容丢到模式空间中,然后打印出来
```

### 把保持空间追加到模式空间(命令G)

> * 大写G命令把当前保持空间的内容作为新行追加到模式空间中, 模式空间的内容不会被覆盖, 该命令在模式空间后面加上换行符\n, 然后把保持空间内容追加进去; 
 > * G和g的用法类似于H和h; 小写命令替换原来的内容, 大写命令追加原来的内容; 
 > * 假定当前模式空间内容为"line 1", 保持空间内容为"line 2"; 命令G执行后, 模式空间内容变为"line 1\nline 2", 同时保持空间内容不变, 仍然为"line 2"; 
> * 以分号分隔, 打印管理者的名称和职位: 
```
$ sed -n -e '/Manager/!h' -e '/Manager/{x;G;s/\n/:/;p}' space.txt
Jason Smith :IT Manager 
Jane Miller :Sales Manager
 
/Manager/!h: 如果模式空间内容不包含关键字'Manager',那么就把他复制到保持空间。 
/Manager/{x;G;s/\n/:/;p}: 如果模式空间包含'Manager', x: 交换模式空间和保持空间的内容, G: 把保持空间的内容追加到模式空间; s/\n/;/ 在模式空间中, 把换行符替换为分号:, p打印模式空间内容; 
# 注意: 如果舍去命令x,即使用/Manager/{G;s/\n/:/;p}, 那么结果会由"雇员职位: 雇员名称" 变成"雇员名称; 雇员职位"
```

## sed 多行模式及循环

> * Sed默认每次只处理一行数据, 除非使用H,G或者N等命令创建多行模式, 每行之间用换行符分开;
 - 提示: 在处理多行模式时, 请务必牢记^只匹配该模式的开头, 即最开始一行的开头, 且$只匹配该模式的结尾, 即最后一行的结尾。	

### 读取下一行数据并附加到模式空间(命令N)

> * 就像大写的命令H和G一样, 只会追加内容而不是替换内容, 命令N从输入文件中读取下一行并追加到模式空间, 而不是替换模式空间;  
> * 前面提到过, 小写命令n打印当前模式空间的内容, 并清空模式空间, 从输入文件中读取下一行到模式空间, 然后继续执行后面的命令;  
> * 大写命令N, 不会打印模式空间内容, 也不会清除模式空间内容, 而是在当前模式空间内容后加上换行符\n, 并且从输入文件中读取下一行数据, 追加到模式空间中, 然后继续执行后面的命令;
> * 以分号分隔,打印雇员名称和职位:
``` 
$ sed -e '{N;s/\n/:/}' space.txt 
John Doe :CEO 
Jason Smith :IT Manager 
Raj Reddy :Sysadmin 
Anand Ram :Developer 
Jane Miller :Sales Manager
 
N追加换行符\n到当前模式空间(雇员名称)的最后, 然后从输入文件读取下一行数据, 追加进来; 因此, 当前模式空间内容变为"雇员名称\n雇员职位";
 
s/\n/:/ 把换行符\n替换为分号, 把分号作为雇员名称和雇员职位的分隔符; 
```
> * 打印example.txt文件内容的同时, 以文本方式显示每行的行号:
``` 
$ sed -e '=' example.txt|sed '{N;s/\n/ /}'         
1 101,Johnnyny Doe,CEO 
2 102,Jason Smith,IT Manager 
3 103,Raj Reddy,Sysadmin 
4 104,Anand Ram,Developer 
5 105,Jane Miller,Sales Manager

命令=先打印行号, 然后打印原始的行的内容; 
命令N在当前模式空间后面加上\n(当前模式空间内容为行号), 然后读取下一行, 并追加到模式空间中, 因此模式空间内容变为"行号\n原始内容", 然后用s/\n/ /把换行符\n替换成空格。
```

### 打印多行模式中的第一行(命令P)

> * 三个大写的命令(H,N,G), 每个命令都是追加内容而不是替换内容; 大写的D和P, 虽然他们的功能和小写的d和p非常相似, 但他们在多行模式中有特殊的功能;  
> * 小写的命令p打印模式空间的内容, 大写的P也打印模式空间内容, 直到它遇到换行符\n;
> * 打印所有管理者的名称: 
```
$ sed -n -e 'N' -e'/Manager/P' space.txt  
Jason Smith 
Jane Miller
```

### 删除多行模式中的第一行(命令D)

> * 小写命令d会删除模式空间内容, 然后读取下一条记录到模式空间, 并忽略后面的命令, 从头开始下一次循环; 
> * 大写命令D, 既不会读取下一条记录, 也不会完全清空模式空间(除非模式空间内只有一行), 它只会： 
 - 删除模式空间的部分内容, 直到遇到换行符\n 
 - 忽略后续命令, 在当前模式空间中从头开始执行命令
> * 下面文件, 每个雇员的职位都用@包含起来作为注释, 需要注意的是, 有些注释是跨行的; 如@Information Technology officer@就跨了两行, 创建测试文件:
``` 
$ cat modules.txt 
John Doe 
CEO @Chief Executive Officer@ 
Jason Smith 
IT Manager @Infromation Technology 
Officer@ 
Raj Reddy 
Sysadmin @System Administrator@ 
Anand Ram 
Developer @Senior 
Programmer@ 
Jane Miller 
Sales Manager @Sales 
Manager@
```

> * 现在我们的目标是,去掉文件里的注释: 
```
$ sed -e '/@/{N;/@.*@/{s/@.*@//;P;D}}' modules.txt 
John Doe 
CEO  
Jason Smith 
IT Manager  
Raj Reddy 
Sysadmin  
Anand Ram 
Developer  
Jane Miller 
Sales Manager

/@/{ 这是外传循环, Sed 搜索包含@符号的任意行, 如果找到, 就执行后面的命令; 如果没有找到, 则读取下一行; 为了便于说明, 以第4行, 即"@Information Technology"(这条注释跨了两行)为例, 它包含一个@符合, 所以后面的命令会被执行; 

N 从输入文件读取下一行, 并追加到模式空间, 以上面提到的那行数据为例, 这里N会读取第5行, 即"Officer@"并追加到模式空间,因此模式空间内容变为"@Informatioin Technology\nOfficer@" 

/@.*@/ 在模式空间中搜索匹配/@.*@/的模式, 即以@开头和结尾的任何内容, 当前模式空间的内容匹配这个模式, 因此将继续执行后面的命令;
 
s/@.*@//;P;D 这个替换命令把整个"@Information Technology\nOfficer@"替换为空(相当于删除), P打印模式空间中的第一行, 然后D删除模式空间中的第一行, 然后从头开始执行命令(即不读取下一条记录,又返回到/@/处执行命令)
```

### 循环和分支(命令b和 :label标签)

> * 使用标签和分支命令b,可以改变sed 的执行流程: 
 - :label 定义一个标签 ;
 - b lable 执行该标签后面的命令, Sed 会跳转到该标签,然后执行后面的命令。;
 - 注意: 命令b后面可以不跟任何标签, 这种情况下, 它会直接跳到sed 脚本的结尾 ;
> * 把space.txt文件中的雇员名称和职位合并到一行内, 字段之间以分号;分隔, 并且在管理者的名称前面加上一个*
``` 
$ cat label.sed 
#!/bin/sed  -nf 
h;n;H;x 
s/\n/:/ 
/Manager/!b end 
s/^/*/ 
:end 
p
 
/Manager/!b end  如果行内不包含关键字"Manager", 则跳转到'end'标签, 你可以任意设置你想要的标签名称,; 因此只有匹配Manager的雇员名称签名, 才会执行s/^/*/(在行首加上星号*) 

:end 即是标签 

# 给这个脚本加上可执行权限,然后执行： 
$ ./label.sed space.txt 
John Doe :CEO 
*Jason Smith :IT Manager 
Raj Reddy :Sysadmin 
Anand Ram :Developer 
*Jane Miller :Sales Manager

$  sed  'N;s/\n/:/;/Manager/s/^/\*/' space.txt  
John Doe :CEO 
*Jason Smith :IT Manager 
Raj Reddy :Sysadmin 
Anand Ram :Developer 
*Jane Miller :Sales Manager
```

### 使用命令t进行循环

> * 命令t的作用是: 如果前面的命令执行成功,那么就跳转到t指定的标签处,继续往下执行后续命令;否则仍然继续正常的执行流程; 
> * 把space.txt文件中的雇员名称和职位合并到一行内,字段之间以分号;分隔,并且在管理者的名称前面加上三个*。 
  - 提示: 我们只需把前面例子中的替换命令改为s/^/***/即可达到该目的,下面这个例子仅仅是为了解释命令t是如何运行的;
``` 
$ cat label-t.sed 
#!/bin/sed  -nf 
h;n;H;x 
s/\n/:/ 
: repeat 
/Manager/s/^/*/ 
/\*\*\*/! t repeat  
p

$ chmod u+x label-t.sed 
$ ./label-t.sed space.txt 
John Doe :CEO 
***Jason Smith :IT Manager 
Raj Reddy :Sysadmin 
Anand Ram :Developer 
***Jane Miller :Sales Manager

 
# 下面的代码执行循环 
:repeat 
/Manager/s/^/*/ 
/\*\*\*/! t repeat 

/Manager/s/^/*/ 如果匹配到Manager, 在行首加上* 

/\*\*\*/!t repeat 如果没有匹配到三个连续的*(用/\*\*\*/!来表示),并且前面一行的替换命令成功执行了, 则跳转到名为repeat的标签处(即 t repeat)
 
:repeat 标签 
```

## sed小技巧

```
sed -n 'n;p' FILE:显示偶数行
$ sen '1!G;h;$!d' FILE:逆向显示文件内容
$ sed '$!N;$!D' FILE:取出文件后两行
$ sed '$!d' FILE:取出文件最后一行
$ sed 'G' FILE:
$ sed '/^$/d;G' FILE:合并空白行
$ sed 'n;d' FILE:显示奇数行
```
  [1]: http://static.zybuluo.com/yujianfeng/cezyjufe9pt7m0hab0p8ofbe/image_1cijgc16b8631938p10eb9v5a19.png