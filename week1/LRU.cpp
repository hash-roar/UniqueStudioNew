#include "LRU.h"

std::string LRU::get(std::string key)
{
    for (auto &item : _nodeList)
    {
        if (item.key == key)
        {
            return item.value;
        }
    }
    return "";
}

void LRU::put(std::string key, std::string value)
{
    for (auto &item : _nodeList)
    {
        if (item.key == key)
        {
            item.value = value;
            return;
        }
    }
    _nodeList.emplace_front(new lruNode(key,value));
    if (_nodeList.size()<_capacity)
    {
        _nodeList.pop_back();
    }
}