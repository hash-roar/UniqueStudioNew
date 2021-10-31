#ifndef __TREE_H_
#define __TREE_H_

#include <vector>
#include <string>
#include <stdio.h>
#include <queue>

struct treeNode
{
    std::string key;
    std::string value;
    struct treeNode *left = nullptr, *right = nullptr;
    // std::vector<struct treeNode*> children;
    treeNode(std::string k = "", std::string v = "") : key(k), value(v){};
};
typedef struct treeNode treeNode;

class Tree
{
private:
    /* data */

public:
    treeNode *root;
    Tree(std::string k = "", std::string v = "");
    ~Tree();
    void visit(treeNode *node);
    void preOrder(treeNode *node);
    void noRecurPreOrder(treeNode *node);
    void noRecurInOrder(treeNode *node);
    void noRecurPostOrder(treeNode *node);
    void inOrder(treeNode *node);
    void postOrder(treeNode *node);
    void levelOrder(treeNode *node);
    struct treeNode *bfs(treeNode *node, std::string key);
    struct treeNode *dfs(treeNode *node, std::string key);
};

#endif