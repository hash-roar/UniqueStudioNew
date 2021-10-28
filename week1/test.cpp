#include "hash.h"
#include <iostream>
#include <stdio.h>
#include <string>

using namespace std;
using str = string;
// const int HASH_SIZE = 97;
int hashFunc(str &key)
{
    auto p = key.c_str();
    unsigned int h = *p;  
    if(h)  
    {  
        for(p += 1; *p != '\0'; ++p)  
            h = (h << 5) - h + *p;  
    }  
    return h%HASH_SIZE;  
}

int main()
{
    str test = "hello3";
    cout << hashFunc(test) << endl;
    return 0;
}