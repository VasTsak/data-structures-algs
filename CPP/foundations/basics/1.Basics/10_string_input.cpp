#include<iostream>
#include<string>

int main()
{
    std::string name;
    std::string address;
    std::string phone_number;
    std::cout<<"What is your name? \n";
    std::getline(std::cin, name);
    std::cout<<"What is your address? \n";
    std::getline(std::cin, address);
    std::cout<<"what is your phone number? \n";
    std::getline(std::cin, phone_number);
    std::cout<<name<<'\n';
    std::cout<<address<<"\n";
    std::cout<<phone_number<<"\n";
    return 0;
}