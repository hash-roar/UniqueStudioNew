# tmuted

## replace

1. 每一行前面加东西:
s/^/something/g

2. 末尾加东西
s/$/something/g

3. 删除tag<>

4. 语法

region option /s/t/ regionLine

3,6s/my/your/g

s/s/S/2

s/s/S/3g

### muti

'1,3s/my/your/g; 3,$s/This/That/g'


