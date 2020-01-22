// Print out even numbers
#include <iostream>
using std::cout;

int main() 
{
    auto i = 1;
    // Write your code here.
    while (i < 11) {
        if (i % 2 == 0) {
            cout << i << "\n";
        }
        i++;
    }

}