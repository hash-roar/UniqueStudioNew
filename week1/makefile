# src=$(wildcard *.cpp)

# obj=$(patsubst %.cpp, %.o, $(src))
# target=server
# CXXFLAGS=-g

# $(target):$(obj)
#         g++ -g $(obj)  -lyaml-cpp -lpthread -L/usr/local/lib -I/usr/local/include  -o $(target)

# %.o:%.c
#         g++ -g $< -c 

# .PHONY:clean
# clean:
#         rm $(target) $(obj)
app:test.cpp
	g++ -g test.cpp avlTree.cpp LRU.cpp hash.cpp infixToSuffix.cpp skipList.cpp Tree.cpp -o app

clean:
	rm app