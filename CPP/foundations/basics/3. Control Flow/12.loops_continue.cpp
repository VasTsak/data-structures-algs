#include <iostream>

int main()
{
    int a = 0;
    while (a < 15)
    {
        a++;
        if (a == 10)
        {
            std::cout << "\t If a is 10, then go back to the top of the loop\n";
            std::cout << "\t This, a = 10 is skipped \n";
            continue;
        }
        std::cout << "After continue a = " << a << "\n"; 
    }
    return 0;
}