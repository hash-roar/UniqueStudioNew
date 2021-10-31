#include "Tree.h"

Tree::Tree(std::string k, std::string v)
{
    root = new treeNode(k, v);
}

Tree::~Tree()
{
}

void Tree::visit(treeNode *node)
{
    printf("key:%s,value:%s\n", node->key.c_str(), node->value.c_str());
}
void Tree::preOrder(treeNode *node)
{
    if (node)
    {
        visit(node);
        preOrder(node->left);
        preOrder(node->right);
    }
}

void Tree::noRecurPreOrder(treeNode *root)
{
    if (!root)
    {
        return;
    }

    std::vector<treeNode *> stack;
    treeNode *node = root;
    while (node || !stack.empty())
    {
        while (node)
        {
            visit(node);
            stack.push_back(node);
            node = node->left;
        }
        if (!stack.empty())
        {
            node = stack.back();
            stack.pop_back();
            node = node->right;
        }
    }
}

void Tree::inOrder(treeNode *node)
{
    if (node)
    {
        preOrder(node->left);
        visit(node);
        preOrder(node->right);
    }
}

void Tree::noRecurInOrder(treeNode *root)
{
    if (!root)
    {
        return;
    }
    std::vector<treeNode *> stackQ;
    treeNode *node = root;
    while (node || !stackQ.empty())
    {
        while (node)
        {
            stackQ.push_back(node);
            node = node->left;
        }
        if (!stackQ.empty())
        {
            node = stackQ.back();
            stackQ.pop_back();
            visit(node);
            node = node->right;
        }
    }
}

void Tree::postOrder(treeNode *node)
{
    if (node)
    {
        preOrder(node->left);
        preOrder(node->right);
        visit(node);
    }
}

void Tree::noRecurPostOrder(treeNode *root)
{

    if (!root)
    {
        return;
    }
    std::vector<treeNode *> stackQ;
    treeNode *node = root;
    treeNode *flag = root; //judge the right tree
    while (node || !stackQ.empty())
    {
        while (node)
        {
            stackQ.push_back(node);
            node = node->left;
        }
        node = stackQ.back();
        if (!node->right || flag == node->right)
        {
            visit(node);
            flag = node;
            stackQ.pop_back();
            node = nullptr; //important
        }
        else
        {
            node = node->right;
        }
    }
}

void Tree::levelOrder(treeNode *root)
{
    if (!root)
    {
        return;
    }
    std::queue<treeNode *> nodeQ;
    nodeQ.push(root);
    while (!nodeQ.empty())
    {
        treeNode *temp = nodeQ.front();
        nodeQ.pop();
        visit(temp);
        if (temp->left)
        {
            nodeQ.push(temp->left);
        }
        if (temp->right)
        {
            nodeQ.push(temp->right);
        }
    }
    printf("\n");
}

struct treeNode *Tree::dfs(treeNode *node, std::string key)
{
    treeNode *result = nullptr;
    if (!node)
    {
        return nullptr;
    }
    if (node->key == key)
    {
        return node;
    }
    if ((result = dfs(node->left, key)))
    {
        return result;
    }
    if ((result = dfs(node->right, key)))
    {
        return result;
    }
    return result;
}

struct treeNode *Tree::bfs(treeNode *root, std::string key)
{
    if (!root)
    {
        return nullptr;
    }
    std::queue<treeNode *> nodeQ;
    while (!nodeQ.empty())
    {
        if (nodeQ.front()->key == key)
        {
            return nodeQ.front();
        }
        if (nodeQ.front()->left)
        {
            nodeQ.push(nodeQ.front()->left);
        }
        if (nodeQ.front()->right)
        {
            nodeQ.push(nodeQ.front()->right);
        }
        nodeQ.pop();
    }
    return nullptr;
}