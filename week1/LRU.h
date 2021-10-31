#ifndef __LRU_H_
#define __LRU_H_

#include <list>
#include <string>

// const int LRU_CAPCITY = 9;

struct lruNode
{
    std::string key;
    std::string value;
    lruNode(std::string k, std::string v) : key(k), value(v)
    {
    }
};

class LRU
{
    using str = std::string;
    std::list<lruNode*> _nodeList;
    int _capacity;

private:
public:
    ~LRU();
    str get(str key);
    void put(str key, str value);
    LRU(int capacity) : _capacity(capacity)
    {
    }
};


#endif