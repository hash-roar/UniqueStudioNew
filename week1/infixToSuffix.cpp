#include "infixToSuffix.h"

infixToSuffix::infixToSuffix(std::string formula)
{
    _formula = formula;
}

infixToSuffix::~infixToSuffix()
{
}

int infixToSuffix::getType(char c)
{
    if (c >= '0' && c <= '9' || c >= 'a' && c <= 'z' || c >= 'A' && c <= 'Z')
    {
        return 1;
    }
    if (c == '+' || c == '-' || c == '/' || c == '*')
    {
        return 2;
    }
    return 0;
}

int infixToSuffix::getPrio(char c)
{
    int priority = -1;
    if (c == '*' || c == '/')
    {
        priority = 2;
    }
    if (c == '+' || c == '-')
    {
        priority = 1;
    }
    if (c == '(')
    {
        priority = 0;
    }
    return priority;
}

std::string infixToSuffix::getSuffix()
{
    for (int i = 0; i < _formula.size(); i++)
    {
        if (getType(_formula[i]) == 0)
        {
            perror("bad formula");
            exit(1);
        }
        if (getType(_formula[i]) == 1)
        {
            list.push_back(_formula[i]);
        }
        else
        {
            if (stack.empty())
            {
                stack.push_back(_formula[i]);
            }
            else if (_formula[i] == '(')
            {
                stack.push_back(_formula[i]);
            }
            else if (_formula[i] == ')')
            {
                while (stack.back() != '(')
                {
                    list.push_back(stack.back());
                    stack.pop_back();
                }
                stack.pop_back();
            }
            else
            {
                while (getPrio(_formula[i]) <= getPrio(stack.back()))
                {
                    list.push_back(stack.back());
                    stack.pop_back();
                    if (stack.empty())
                    {
                        break;
                    }
                }
                stack.push_back(_formula[i]);
            }
        }
    }
    while (!stack.empty())
    {
        list.push_back(stack.back());
        stack.pop_back();
    }
    std::string temp;
    // temp.resize(list.size());
    // for (int i = 0; i < list.size(); i++)
    // {
    //     temp[i] = list[i];
    // }
    for (auto c : list)
    {
        temp += c;
    }
    return temp;
}
