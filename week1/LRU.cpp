#include "LRU.h"
LRU::~LRU()
{
}

std::string LRU::get(std::string key)
{
    auto it = _nodeList.begin();
    while (it != _nodeList.end())
    {
        if ((*it)->key == key)
        {
            auto temp = *it;
            _nodeList.erase(it);
            _nodeList.emplace_front(temp);
            return temp->key;
        }
        it++;
    }
    return "";
}

void LRU::put(std::string key, std::string value)
{
    for (auto &item : _nodeList)
    {
        if (item->key == key)
        {
            item->value = value;
            return;
        }
    }
    _nodeList.emplace_front(new lruNode(key, value));
    if (_nodeList.size() > _capacity)
    {
        printf("淘汰: %s\n", _nodeList.back()->key.c_str());
        _nodeList.pop_back();
    }
}
