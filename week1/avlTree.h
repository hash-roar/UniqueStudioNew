#ifndef __AVLTREE_H_
#define __AVLTREE_H_

#include <string>
#include <stdio.h>
#include <math.h>

struct treeNode
{
    int value;
    int height = 1;
    struct treeNode *left = nullptr, *right = nullptr;
};
typedef struct treeNode treeNode;

class avlTree
{
private:

public:
    treeNode *root;
    avlTree(int value);
    ~avlTree();

public:
    //balance
    treeNode *reBalance(treeNode *node);
    int getBalanceFactor(treeNode *node);
    int getHeight(treeNode *node);
    treeNode *rightRotation(treeNode *node);
    treeNode *leftRotation(treeNode *node);

public:
    //curd
    treeNode* Insert(treeNode *node, int value);
    treeNode *Delete(treeNode *node, int value);
    treeNode *Find(treeNode *node, int value);
    void change(treeNode *node);
    treeNode *getMax(treeNode *node)
    {
        while (node->right)
        {
            node = node->right;
        }
        return node;
    }
    treeNode *getMin(treeNode *node)
    {
        while (node->left)
        {
            node = node->left;
        }
        return node;
    }
};

#endif
