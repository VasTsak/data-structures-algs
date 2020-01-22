#include <iostream>

int main()
{
    int menuItem = 1;

    std::cout << "What is your favourite winter sport? \n";
    std::cout << "1. Skiing \n 2. Sledding \n 3. Sitting by the fire. \n";
    std::cout << "4. Drinking hot chocolate. \n";
    std::cout << "\n\n";

    switch(menuItem)
    {
        case(1): std::cout << "Skiing?! That sounds dangerous! \n";
        break;
        case(2): std::cout << "Sledding?! That sounds less dangerous \n";
        break;
        case(3): std::cout << "Sitting by the fire? Isn't that boring? \n";
        break;
        case(4): std::cout << "Hot chocolate? More like my type of guy! \n";
        break;
    }

    char begin;
    std::cout<<"\n\nWhere do you want to begin?\n";
    std::cout<<"B. At the beginning?\nM. At the middle?";
    std::cout<<"\nE. At the end?\n\n";
    begin = 'M';  

    switch(begin)
    {
        case('B'): std::cout<<"Once upon a time there was a wolf.\n";
        case('M'): std::cout<<"The wolf hurt his leg.\n";
        case('E'): std::cout<<"The wolf lived happily ever after\n";
    }
    return 0;
}