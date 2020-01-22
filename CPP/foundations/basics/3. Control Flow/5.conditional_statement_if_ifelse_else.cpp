#include <iostream>

int main()
{
    int a_if_1 = 1;
    int a_if_2 = 2;

    if (a_if_1 == 0)
    {
        std::cout << "a_if_1 is equal to 0. \n";
    }
    else if (a_if_2 == 1)
    {
        std::cout << "a_if_2 is 1. \n";
    }
    else 
    {
        std::cout << "a_if_1 is not 0 and a_if_2 is not 1. \n";
    }
    return 0;
}