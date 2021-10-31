#include "avlTree.h"

avlTree::avlTree(int value)
{
    root = new treeNode;
    root->value = value;
}

avlTree::~avlTree()
{
    //delete
}

int avlTree::getHeight(treeNode *node)
{
    if (node!=nullptr)
    {
        return std::max(getHeight(node->left), getHeight(node->right)) + 1;
    }
    return 0;
}

int avlTree::getBalanceFactor(treeNode *node)
{
    if (node)
    {
        return getHeight(node->left) - getHeight(node->right);
    }
    return 0;
}

treeNode *avlTree::rightRotation(treeNode *node)
{
    treeNode *nodeL = node->left;
    treeNode *nodeLR = nodeL->right;
    nodeL->right = node;
    node->left = nodeLR;
    node->height = std::max(getHeight(node->left), getHeight(node->right)) + 1;
    nodeL->height = std::max(getHeight(nodeL->left), getHeight(nodeL->right)) + 1;
    return nodeL;
}

treeNode *avlTree::leftRotation(treeNode *node)
{
    treeNode *nodeR = node->right;
    treeNode *nodeRL = nodeR->left;
    nodeR->left = node;
    node->left = nodeRL;
    node->height = std::max(getHeight(node->left), getHeight(node->right)) + 1;
    nodeR->height = std::max(getHeight(nodeR->left), getHeight(nodeR->right)) + 1;
    return nodeR;
}

treeNode *avlTree::reBalance(treeNode *node)
{
    int balance_factor = getBalanceFactor(node);
    if (balance_factor > 1 && getBalanceFactor(node->left) > 0) //LL
    {
        return rightRotation(node);
    }
    else if (balance_factor > 1 && getBalanceFactor(node->left) <= 0) //LR
    {
        node->left = leftRotation(node->left);
        return rightRotation(node);
    }
    else if (balance_factor < -1 && getBalanceFactor(node->right) < 0) //RR
    {
        return leftRotation(node);
    }
    else if (balance_factor < -1 && getBalanceFactor(node->right) >= 0)
    {
        node->right = rightRotation(node->right);
        return leftRotation(node);
    }
    else
    {
        return node;
    }
}

treeNode *avlTree::Insert(treeNode *node, int value)
{
    if (!node)
    {
        node = new treeNode;
        node->value = value;
        return node;
    }
    else if (node->value == value)
    {
        return node;
    }
    else if (node->value > value)
    {
        node->left = Insert(node->left, value);
    }
    else
    {
        node->right = Insert(node->right, value);
    }
    node->height = std::max(getHeight(node->left), getHeight(node->right)) + 1;
    node = reBalance(node);
    return node;
}

treeNode *avlTree::Delete(treeNode *node, int value)
{
    if (!node)
    {
        return nullptr;
    }
    if (value < node->value)
    {
        node = Delete(node->left, value);
        return node;
    }
    else if (value > node->value)
    {
        node = Delete(node->right, value);
        return node;
    }
    if (!node->left || !node->right)
    {
        if (!node->left && !node->right)
        {
            delete node;
            node = nullptr;
            return node;
        }
        else if (node->left && !node->right)
        {
            node = node->left;
            return node;
        }
        else
        {
            node = node->right;
            return node;
        }
    }
    else
    {
        treeNode *minSon = getMin(node->right);
        node->value = minSon->value;
        node->right = Delete(node->right, minSon->value);
    }
    node = reBalance(node);
    node->height = std::max(getHeight(node->left), getHeight(node->right)) + 1;
    return node;
}

treeNode *avlTree::Find(treeNode *node, int value)
{
    if (!node)
    {
        return nullptr;
    }
    if (node->value = value)
    {
        return node;
    }
    else if (node->value > value)
    {
        return Find(node->left, value);
    }
    else
    {
        return Find(node->right, value);
    }
}