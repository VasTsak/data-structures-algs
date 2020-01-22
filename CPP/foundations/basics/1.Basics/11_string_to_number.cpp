// Step 1: Include the Stringstream library

//#include <sstream>

//Step 2: Use getline to get the string from the user

//std::getline(std::cin, stringVariable);

//Step 3: Convert the string variable to float variable

//std::stringstream(stringVariable) >> numericVariable;

#include<iostream>
#include<string>
#include<sstream>

int main()
{
    std::string stringLength;
    float inches = 0;
    float yardage = 0;

    std::cout << "Enter number of inches: \n";
    std::getline(std::cin, stringLength);
    std::stringstream(stringLength) >> inches;
    std::cout << "You entered: " << inches << "\n";
    yardage = inches / 36;
    std::cout << "Yardage is: " << yardage <<"\n";
    return 0;
}