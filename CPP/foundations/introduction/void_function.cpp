/*Sometimes a function doesn't need to return anything. For example, 
a function might simply modify an object that is passed into it, or 
it might just print to the terminal. If a function doesn't need to 
return a value, the void type can be used for the return type. Using 
the function syntax provided above, write a function PrintStrings that 
takes two strings as arguments and prints both of them. If you are 
having trouble, click the solution button for help.*/

#include <iostream>
#include <string>
using std::cout;
using std::string;

// Write the PrintStrings function here.

void PrintStrings(string s1, string s2){
    cout << s1 << s2 << "\n";
}
int main() 
{
    string s1 = "C++ is ";
    string s2 = "fast.";
    
    // Uncomment the following line to call your function:
    PrintStrings(s1, s2);
}