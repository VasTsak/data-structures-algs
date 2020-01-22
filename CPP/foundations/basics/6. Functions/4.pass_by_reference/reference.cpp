/*
Now write a program that uses passing variables by reference.

Write a program that uses two functions:

*    calculate(input1, input2, operation, result);
*    printEquation(input1, input2, operation, result); 

Input, input2, and result are floats.
Operation is a char. The choices are '+', '-', '*', '/'
The result is modified by the function called calculate.

*/

#include "main.hpp"

int main()
{
    char operation = '/';
    float input1 = 9.8;
    float input2 = 2.3;
    float result;

    calculate(input1, input2, operation, result);
    printEquation(input1, input2, operation, result);
    return 0;
}
