#include<iostream>
#include<string>

int main()
{
    int year = 0;
    int age = 0;
    std::string name = " ";

    //Print a message to the user
    std::cout<<"What year is your favourite?\n";

    // get the user response and assign it to the variables
    std::cin>>year;

    //output response to the user
    std::cout<<"How interesting that your favourite year is: "<<year<<"!\n";

    std::cout<<"At what age did you learn to cycle the bicycle?\n";

    std::cin>>age;

    //output response to user
    std::cout<<"How interesting you learned to ride at "<<age<<"!\n";

    std::cout<<"What is your name?\n";
    std::cin>>name;
    std::cout<<"Hello "<<name<<"!\n";
    return 0;    
}