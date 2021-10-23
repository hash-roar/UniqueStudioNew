# bash script

## script _ exmple

 #!/bin/bash
 echo "Hello World !"

## variable

===> var=value
echo var
echo {var}
echo "hello {var}"

===> var = `commond`
var1=(commond)

===> readonly var  //make var  readonly

===> unset var

====> 

## special var

# --> var nums

  * $@ //$ * 和 $@ 都表示传递给函数或脚本的所有参数，不被双引号(" ")包含时，都以"$1" "$2" … "$n" 的形式输出所有参数。

 但是当它们被双引号(" ")包含时，" * " 会将所有的参数作为一个整体，以"$1 $2 … $n"的形式输出所有参数；"$@" 会将各个参数分开，以"$1" "$2" … "$n" 的形式输出所有参数。 

===> tr1="hello"
str2="world"  
str3=$str1"  "$str2

## array

==> arrayName=(value0 value1 value2 value3)

==> arrayname[0]=value0

==> valuen=${arrayName[2]} 

==> length=${#arrayName}

==> 
ary=()
read -p "arg>" -a ary
echo ${ary[@]}
for item in   ${ary[@]}    
do
	echo $item
done
 
==> 

