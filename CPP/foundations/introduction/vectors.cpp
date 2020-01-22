#include <iostream>
#include <vector>
using std::vector;
using std::cout;

int main() {
    // Declare and initialize a vector v here.
    vector<int> v{6,7,8};
    for(int i : v)
        cout << i << " ";
    cout << "\n";
}   