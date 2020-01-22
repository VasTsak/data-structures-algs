#include <iostream>

int integer = 456;

int main()
{
    std::cout<<"The variable value is (no escape sequence): " <<integer<<"\n";
    // The most comon is to use \n or \t
    std::cout<<"The variable value is (with new line): \n"<<integer<<"\n";
    std::cout<<"The variable value is (with tab): \t"<<integer<<"\n";
    return 0;
}