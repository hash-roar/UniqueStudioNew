#ifndef __SKIPLIST_H_
#define __SKIPLIST_H_

#include <cstdlib>
#include <vector>
#define MAX_INDEX_LEVEL 4
#define SKIP_P 0.5

struct listNode
{
    int key=0;
    int data=0;
    struct listNode *next = nullptr, *down = nullptr;
};
typedef struct listNode listNode;

class skipList
{
private:
    listNode *headNode;
    int level_now;

private:
    double getrand()
    {
        return rand() / double(RAND_MAX);
    }
    int randomLevel()
    {
        int level = 1;
        while (getrand() < SKIP_P && level < MAX_INDEX_LEVEL)
        {
            level += 1;
        }
        return level;
    }

public:
    skipList(/* args */);
    ~skipList();
    listNode *Find(int key);
    listNode *Find(int key1,int key2);
    void Delete(int key);
    void Insert(const struct listNode &node);
    void Insert(int key, int value);
};

#endif