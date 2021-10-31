#include "hash.h"
// hashItem::hashItem(/* args */)
// {
// }

hashTable::hashTable(/* args */)
{
    tbItem.resize(HASH_SIZE);
    for (auto &item : tbItem)
    {
        item = new struct hashNode;
    }
}

hashTable::~hashTable()
{
    for (auto &itemP : tbItem)
    {
        if (itemP != nullptr)
        {
            auto head = itemP;
            while (itemP != nullptr)
            {
                head = itemP->next;
                delete itemP;
                itemP = head;
            }
        }
    }
}

int hashTable::hashFunc(str &key)
{
    auto p = key.c_str();
    unsigned int h = *p;
    if (h)
    {
        for (p += 1; *p != '\0'; ++p)
            h = (h << 5) - h + *p;
    }
    return h % HASH_SIZE;
}

int hashTable::add(str key, str value)
{
    int hashIndex = hashFunc(key);
    auto itemP = tbItem[hashIndex];
    if (key == "")
    {
        return -1;
    }

    while (itemP->next)
    {
        itemP = itemP->next;
    }
    auto newP = new struct hashNode;
    newP->key = key;
    newP->value = value;
    if (!newP)
    {
        printf("new error");
        return -1;
    }
    itemP->next = newP;
    itemP->next->next = nullptr;
    return hashIndex;
}

// int hashTable::add(str key, str value)
// {
//     struct hashNode node;
//     return add(node);
// }

const struct hashNode *hashTable::get(str key)
{
    if (key == "")
    {
        return nullptr;
    }
    int hashIndex = hashFunc(key);
    auto itemP = tbItem[hashIndex];
    while (itemP != nullptr)
    {
        if (itemP->key == key)
        {
            return itemP;
        }
        else
        {
            itemP = itemP->next;
        }
    }
    return nullptr;
}

int hashTable::set(str &key, str &value)
{

    if (key == "")
    {
        return -1;
    }
    int hashIndex = hashFunc(key);
    auto itemP = tbItem[hashIndex];
    while (itemP != nullptr)
    {
        if (itemP->key == key)
        {
            itemP->value = value;
            return hashIndex;
        }
        else
        {
            itemP = itemP->next;
        }
    }
    return -1;
}

int hashTable::remove(str &key)
{
    if (key == "")
    {
        return -1;
    }
    int hashIndex = hashFunc(key);
    auto itemP = tbItem[hashIndex];
    // auto itemPBack = itemPFront;
    if (itemP->key == key && itemP->next == nullptr)
    {

        delete itemP;
        itemP = nullptr;
    }
    else if (itemP->key == key && itemP->next != nullptr)
    {
        auto temp = itemP;
        temp = itemP->next;
        delete itemP;
    }
    while (itemP->next != nullptr)
    {
        if (itemP->next->key == key)
        {
            auto temp = itemP->next;
            itemP->next = temp->next;
            delete temp;
            return hashIndex;
        }
        else
        {
            itemP = itemP->next;
        }
        return -1;
    }
    return hashIndex;
}
