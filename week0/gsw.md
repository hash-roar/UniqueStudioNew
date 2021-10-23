# awk


## if -else

if(condition){statement;…}[else statement]  //双分支

if(condition1){statement1}else if(condition2){statement2}else{statement3}  // 多分支


## while

while(condition){statement;…} // 循环

## 数组
==> NR!=1{a[$6]++;} END {for (i in a) print i ", " a[i];}' netstat.txt  // 统计状态

## 变量

==> awk -v val=$x '{print $1, $2, $3, $4+val, $5+ENVIRON["y"]}' OFS="\t" score.txt //变量 及环境变量

# grep

## option
-a search binary

## 字符串匹配

$6 ~ /FIN/ || NR==1 {print NR,$4,$5,$6}

## 拆分文件
==> NR!=1{print > $6} // 按第六列拆分文件

==> NR!=1{if($6 ~ /TIME|ESTABLISHED/) print > "1.txt";
else if($6 ~ /LISTEN/) print > "2.txt";
else print > "3.txt" }  //自定义拆分文件

# find

-type --> 

-ok  --> excute commond  and query

-excute --> no query
==> find . -type f -name " * .txt" -exec cat {} \;> /all.txt //
查找当前目录下所有.txt文件并把他们拼接起来写入到all.txt文件中

-name -->  search by name 

-iname --> search by name ignore case

-regex --> 

! --> not 
==>  find /home ! -name " * .txt"  // 找出/home下不是以.txt结尾的文件

-maxdepth -->  max depth

-size  --> b k M G c
==> find . -type f -size -10k  // 搜索小于10KB的文件

-delete  --> 

-user  --> 

! -path --> ignore path
 -path -prune --> ignore path

-empty -->

===> find . -name " * .java"|xargs cat|grep -v ^$|wc -l # 代码行数统计, 排除空行

