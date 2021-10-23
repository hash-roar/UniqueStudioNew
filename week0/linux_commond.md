# linux 

## ls

-a --> list hidden files

-A --> not list . ..

-color --> color distinction

-d --> lis dir like files 

-h --> size human readbale

-l --> detail informatio

-R --> recursive

-sort 
extension -X status -c
none -U time -t
size -S atime -utime -t access -u
version -v use -u
--> sort by

-1 -->  one file one line

-F --> dir+ "/ "

//蓝色：目录

绿色：可执行文件

白色：一般性文件，如文本文件，配置文件等

红色：压缩文件或归档文件

浅蓝色：链接文件

红色闪烁：链接文件存在问题

黄色：设备文件

青黄色：管道文件

## cd

cd ~ --> cd user home

cd - --> cd last dir

cd !$ --> use last commond arg as cd arg

## mkdir

-m --> mode 
==> mkdir -m 777 test //mkdir mode 777

-p --> recursive 

-v --> show info 

===> mkdir -vp  project/{lib/,bin/,doc/{info,product},logs/{info,product},service/deploy/{info,product}} //mdir

## touch

-a --> only change time

-d --> use date

-m --> change modify time

-c --> do not create

-r --> set time reference 
==> touch -r log2.log log2.txt

## cp

-a --archive --> make backup

-R -r --> recursive

-l --> link not copy

-i --> interactive

-p --> reserve origin info

## rm

-f --> force

-i 

-r

-v --> show info

===> self define rm
myrm(){ D=/tmp/$(date +%Y%m%d%H%M%S); mkdir -p $D; mv "$@" $D && echo "moved to $D ok"; } 
alias rm='myrm' 

## mv

-b --> backup	

-f --> force

-i

# linux authority

## rwx

r 4
w 2
x 1

## chgrp

===> chgrp group1 install.log 

## chown
-c --> -v

-R 

-deference --> change use file


===> chown [-R] 所有者:所属组 文件或目录

## chmod 

-c -v --> info

-R 

--reference=Rfile --> set mod referenced to

===> chmod u=rwx, g=rw, o=r ./test.log


# process

## start process 


===> find / -name install.log &  //start backstatge process

## get process info

### ps
-a --> all in the same tty

-e -A --> all process

e --> show env-var

-p --> get process name


### top
-d --> delimition (second)

-n --> times

-u --> user

### top interactive
h --> help

P --> sort by cpu

M --> memory

T --> tmulated time

k --> kill


### kill

-s --> signal

-n --> sigNum

-l --> list signal

-p --> print pid not send signal

### kill all

-i --> interactive

-e --> exact

-u --> user

### pkill

-g --> process group

# port 

## lsof
-d --> file number

+d --> file opend 
+D --> recursive

-i --> 


-p --> file opened by process

-u --> user

-c --> program

===>
lsof -i udp
lsof -i :8080
lsof -i udp:5500
lsof -a -u test -i

### netstat

-a --> all connenct

-t --> tcp

-u --> udp

-l --> listen

-p --> program name

# linux service


servicserviceservice








# io redirect

## 
命令 < 文件 将文件作为命令的标准输入
命令 << 分界符 从标准输入中读入，直到遇见分界符才停止
命令 < 文件 1 > 文件 2 将文件 1 作为命令的标准输入并将标准输到到文件 2

命令 > 文件 将标准输出重定向到一个文件中（清空原有文件的数据）
命令 2> 文件 将错误输出重定向到一个文件中（清空原有文件的数据）
命令 >> 文件 将标准输出重定向到一个文件中（追加到原有内容的后面）
命令 2>> 文件 将错误输出重定向到一个文件中（追加到原有内容的后面）
命令 >> 文件 2>&1
或
命令 &>> 文件
将标准输出与错误输出共同写入到文件中（追加到原有内容的
后面）

## 
