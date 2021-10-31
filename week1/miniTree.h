#ifndef __MINITREE_H_
#define __MINITREE_H_

#include <string>
#include <stdio.h>
#include <vector>
class miniTree
{
private:
    std::vector<std::vector<int>> G;
    int _vnum;

public:
    miniTree(/* args */);
    ~miniTree();
    void getGraph();
    int getMiniV(int lowcost[], int length,bool visit[]);
    int prim();
};



#endif