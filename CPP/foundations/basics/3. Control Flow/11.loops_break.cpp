#include <iostream>

int main()
{
    int a = 0;
    while (a < 5)
    {
        std::cout << "a = "<< a << "\n";
        a++;
        if (a == 3)
        {
            break;
        }
    }
    std::cout << "The first statement after the first while loop \n\n";

    return 0;
}