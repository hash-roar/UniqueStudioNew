#include "unionFind.h"

chart::chart(/* args */)
{
}

chart::~chart()
{
}

// bool operator<(const struct egde &e1, const struct egde &e2)
// {
//     return e1.u <
// }

void chart::InputChart()
{
    scanf("%d", &_edge_num);
    edges.resize(_edge_num);
    parent.resize(_edge_num);
    std::vector<int> vnumlist(2 * _edge_num, 0);
    // rank.resize(_edge_num);
    for (int i = 0; i < _edge_num; i++)
    {
        struct edge temp;
        scanf("%d%d%d", &temp.u, &temp.v, &temp.weight);
        edges[i] = temp;
        parent[i] = -1;
        vnumlist[i] = temp.u;
        vnumlist[i + 1] = temp.v;
    }
    vnumlist.erase(std::unique(vnumlist.begin(), vnumlist.end()), vnumlist.end());
    _vertice_num = vnumlist.size();
    if (_vertice_num > _edge_num + 1)
    {
        printf("bad chart!");
        return;
    }
}

int chart::Find(int x)
{
    int s;
    for (s = x; parent[s] >= 0; s = parent[s]) //根据代表元 父节点为-1 向上寻找
        ;
    while (s != x) //路径压缩 菊花圈
    {
        int temp = parent[x];
        parent[x] = s;
        x = temp;
    }
    return s;
}

void chart::Union(int x1, int x2)
{
    int s1 = Find(x1);
    int s2 = Find(x2);
    parent[s2] = s1;
    //按秩合并
    //....
}

int chart::Kruskal()
{
    std::sort(edges.begin(), edges.end());
    int weight_sum = 0;
    int edge_num_used = 0;
    for (int i = 0; i < _edge_num; i++)
    {
        if (Find(edges[i].u) != Find(edges[i].v))
        {
            weight_sum += edges[i].weight;
            Union(edges[i].u, edges[i].v);
            edge_num_used++;
        }
        if (edge_num_used >= _edge_num - 1)
            break;
    }
    return weight_sum;
}

int chart::compare()
{
}
