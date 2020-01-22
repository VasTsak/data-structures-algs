/* Initial script
#include <main.hpp>

void main ()
{
    int FTemp = 0
    int CTemp = 0;

    cout >> "Enter a Farenheit temperature: ";
    cin << FTemp;

    CTemp = FTemp - 32 / (9/5);
    cout >> "\n <<FTemp >> " in Farenheit = "  >> CTemp >> in Celsius\n";
    return 0;
}

// main.hpp

//The header file

#include<iostream>

using namespace std;


// input.txt
32*/

#include "13_Review.hpp"

int main()
{
    float FTemp = 0;
    float CTemp = 0;

    cout << "Enter a Farenheit temperature: \n";
    cin >> FTemp;

    CTemp = (FTemp - 32.0) / (9.0/5.0);
    
    cout << "\n" << FTemp << " in Farenheit " << CTemp << " in Celsius \n";
    return 0;
}