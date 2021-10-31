#include "miniTree.h"

miniTree::miniTree(/* args */)
{
}

miniTree::~miniTree()
{
}

void miniTree::getGraph()
{
    scanf("%d", &_vnum);
    for (int i = 0; i < _vnum; i++)
    {
        for (int j = 0; j < _vnum; j++)
        {
            G[i][j] = -1;
        }
    }
    for (int i = 0; i < _vnum; i++)
    {
        for (int j = 0; j < _vnum; j++)
        {
            int temp;
            scanf("%d", temp);
            G[i][j] = temp;
        }
    }
}

int miniTree::getMiniV(int lowcost[], int length, bool visit[])
{
    int min = INT32_MAX, min_v;
    for (int i = 0; i < length; i++)
    {
        if (!visit[i] && lowcost[i] < min)
        {
            min = lowcost[i];
            min_v = i;
        }
    }
    return min_v;
}

int miniTree::prim()
{
    bool S[_vnum];
    int lowcost[_vnum];
    int sumWeight = 0;
    for (int i = 0; i < _vnum; i++)
    {
        lowcost[i] = INT32_MAX;
        S[i] = false;
    }
    lowcost[0] = 0;
    S[0] = true;

    for (int i = 0; i < _vnum - 1; i++)
    {
        int u = getMiniV(lowcost, _vnum, S);
        S[u] = true;
        sumWeight += lowcost[u];
        for (int j = 0; j < _vnum; j++)
        {
            if (G[u][j] != -1 && S[j] == false && G[u][j] < lowcost[j])
            {
                lowcost[j] = G[u][j];
            }
        }
    }
    return sumWeight;
}