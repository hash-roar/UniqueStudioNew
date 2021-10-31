#include "./hash.h"
#include "./infixToSuffix.h"
#include "./LRU.h"
#include "./avlTree.h"
#include "./skipList.h"
// #include "./Tree.h"
#include <iostream>
#include <stdio.h>
#include <string>

using namespace std;
using str = string;
// const int HASH_SIZE = 97;
int main()
{
    //test fir hash table

    // hashTable hashtable;
    // hashtable.add("key", "value");
    // string stri = "key";
    // if (hashtable.get(stri))
    // {
    //     cout<<hashtable.get(stri)->value<<endl;
    // }

    //test for  infixToSuffix

    // infixToSuffix infixFormula("a+b*c+(a+b)/2");
    // cout<<infixFormula.getSuffix()<<endl;

    //LRU

    // LRU lrulist(4);
    // lrulist.put("5","hello");
    // lrulist.put("2","hello");
    // lrulist.put("3","hello");
    // lrulist.get("5");
    // lrulist.put("1","hello");
    // lrulist.put("3","hello");
    // lrulist.put("5","hello");
    // lrulist.get("2");
    // lrulist.put("7","hello");

    // avlTree

    avlTree tree(5);
    tree.root = tree.Insert(tree.root, 4);
    tree.root = tree.Insert(tree.root, 2);
    tree.root = tree.Insert(tree.root, 5);
    tree.root = tree.Insert(tree.root, 7);
    tree.root = tree.Insert(tree.root, 9);
    tree.root = tree.Insert(tree.root, 1);
    tree.root = tree.Delete(tree.root,7);
    cout << tree.getMax(tree.root)->value << endl;

    // skipList

    // skipList skiplist;
    // skiplist.Insert(1, 1);
    // skiplist.Insert(2,5);
    // auto result = skiplist.Find(2);
    // if (result)
    // {
    //     cout << result->data << endl;
    // }

    //tree

    // Tree tree("key", "value");
    // tree.root->left = new struct treeNode("1", "1");
    // tree.root->right = new struct treeNode("2", "2");
    // tree.root->left->right = new struct treeNode("3", "3");
    // tree.root->right->right = new struct treeNode("4", "4");
    // tree.postOrder(tree.root);
    // cout << endl;
    // tree.noRecurInOrder(tree.root);
    // cout << endl;
    // tree.inOrder(tree.root);
    // cout << endl;
    // tree.levelOrder(tree.root);
    // auto result = tree.dfs(tree.root, "2");
    // cout << result->value << endl;
    

    return 0;
}