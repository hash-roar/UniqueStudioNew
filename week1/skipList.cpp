#include "skipList.h"

skipList::skipList(/* args */)
{
    headNode = new listNode;
    level_now = 1;
}

skipList::~skipList()
{
}

listNode *skipList::Find(int key)
{
    auto node = headNode;
    while (node)
    {
        if (node->key == key)
        {
            return node;
        }
        else if (!node->next)
        {
            node = node->down;
        }
        else if (node->next->key > key)
        {
            node = node->down;
        }
        else
        {
            node = node->next;
        }
    }
    return nullptr;
}

void skipList::Delete(int key)
{
    auto node = headNode;
    while (node)
    {
        if (!node->next)
        {
            node = node->next;
        }
        else if (node->next->key = key)
        {
            delete node->next;
            node->next = node->next->next;
            node = node->down;
        }
        else if (node->next->key > key)
        {
            node = node->down;
        }
        else
        {
            node = node->next;
        }
    }
}

void skipList::Insert(int key, int value)
{
    listNode temp;
    temp.key = key;
    temp.data = value;
    Insert(temp);
}

void skipList::Insert(const struct listNode &node)
{
    int key = node.key;
    // auto tempNode = Find(key);
    // if (!tempNode)
    // {
    //     tempNode->data=node.data;
    //     return;
    // }
    std::vector<listNode *> stack;
    auto indexNode = headNode;
    while (indexNode)
    {
        if (!indexNode->next)
        {
            stack.push_back(indexNode);
            indexNode = indexNode->down;
        }
        else if (indexNode->next->key > key)
        {
            stack.push_back(indexNode);
            indexNode = indexNode->down;
        }
        else if (indexNode->key == key)
        {
            while (indexNode)
            {
                indexNode->data = node.data;
                indexNode = indexNode->down;
            }
            return;
        }
        else
        {
            indexNode = indexNode->next;
        }
    }
    int index_level = randomLevel();
    listNode *nodeDown = nullptr;
    int level = 1;
    for (int i = 0; i < index_level; i++)
    {
        if (!stack.empty())
        {
            auto nodeUp = stack.back();
            stack.pop_back();
            auto temp = nodeUp->next;
            nodeUp->next = new listNode(node);
            nodeUp->next->key = node.key;
            nodeUp->next->data = node.data;
            nodeUp->next->down = nodeDown;
            nodeUp->next->next = temp;
            nodeDown = nodeUp->next;
        }
        else
        {
            level_now += 1;
            auto headNodeUp = new listNode;
            headNodeUp->down = headNode;
            headNode = headNodeUp;
            headNodeUp->next = new listNode(node);
            headNodeUp->next->key = node.key;
            headNodeUp->next->data = node.data;
            headNodeUp->next->down = nodeDown;
            nodeDown = headNodeUp->next;
        }
    }
}
