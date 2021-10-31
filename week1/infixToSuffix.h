#ifndef __INFIXTOSUFFIX_
#define __INFIXTOSUFFIX_

#include <vector>
#include <stdio.h>
#include <string>

class infixToSuffix
{
private:
    /* data */
    std::string _formula;
    std::vector<char> stack; //符号栈
    std::vector<char> list;  //保存后缀表达式

public:
    infixToSuffix(std::string formula);
    ~infixToSuffix();
    std::string getSuffix();
    int getPrio(char c);
    int getType(char c); //返回1表示表达式,返回2表示符号,返回0表示未识别

private:
};

#endif