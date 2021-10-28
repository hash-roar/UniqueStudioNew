#ifndef __HASH_H_
#define __HASH_H_

#include <string>
#include <vector>

using str = std::string;
const int HASH_SIZE = 97;

//shoud be a templete
// class hashNode
// {
// private:
//     str key;
//     str value;
//     hashNode *next;

// public:
//     hashNode(/* args */);
//     ~hashNode();
// };

 struct hashNode
{
    str key;
    str value;
    hashNode *next;
    hashNode(str& k,str &v):key(k),value(v)
    {
    }
};

class hashTable
{
private:
    /* data */
    std::vector<struct hashNode *> tbItem;

public:
    hashTable(/* args */);
    ~hashTable();
    int hashFunc(str &key); //hash function
    const struct hashNode* get(str &key);
    int add( struct hashNode &item);
    int add(str key,str value);
    int set(str &key,str &value);
    int remove(str &key);
};

#endif