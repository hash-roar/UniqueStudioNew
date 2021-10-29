#ifndef __UNION_FIND_
#define __UNION_FIND_

#include <string>
#include <math.h>
#include <stdio.h>
#include <vector>
#include <algorithm>

#define EDGE_NUM 12

struct edge
{
    bool operator<(const struct edge e)
    {
        return u < e.v;
    }
    int u, v;
    int weight;
    edge(int U = -1, int V = -1, int Weight = -1) : u(U), v(V), weight(Weight) {}
};

// struct edge edges[EDGE_NUM];
// int parent[EDGE_NUM];

class chart
{
private:
    std::vector<struct edge> edges;
    std::vector<int> parent;
    // std::vector<int> rank; 按秩合并
    int _edge_num;
    int _vertice_num;

public:
    int compare();
    // friend bool operator<(const struct egde &e1, const struct egde &e2);
    chart(/* args */);
    ~chart();
    void InputChart();
    int Find(int x);
    void Union(int x1, int x2);
    int Kruskal();
};

#endif