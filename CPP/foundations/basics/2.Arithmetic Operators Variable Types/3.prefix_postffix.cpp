/*
n C++, as in many languages, there are postfix and prefix 
operators.
The form for each is:
Incrementing
    prefix: ++a
    postfix: a++ 
Decrementing
    prefix: --a
    postfix: a-- 
The difference between prefix and postfix is subtle, 
but crucial.
Prefix operators increment the value of the variable, 
then return the reference to the variable.
Postfix operators create a copy of the variable and 
increments the value of the variable. Then it returns 
a copy from BEFORE the increment.
*/

#include<iostream>

using namespace std;

int main()
{
    int a, b = 1;
    int post, pre = 0;
    cout<<"Inital values: \t\t\tpost = "<<post<<" pre= "<<pre<<"\n";
    post = a++;
    pre = ++b;
    cout<<"After one postfix and prefix: \tpost = "<<post<<" pre= "<<pre<<"\n";
    post = a++;
    pre = ++b;
    cout<<"After two postfix and prefix: \tpost = "<<post<<" pre= "<<pre<<"\n";  
    return 0;
}