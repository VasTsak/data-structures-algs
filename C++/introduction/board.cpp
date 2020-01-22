#include <iostream>
#include <vector>
using std::vector;
using std::cout;

void PrintBoard(vector<vector<int>> b) {
  for (int i = 0; i < b.size(); i ++) {
    for (int j : b[i]) {
      cout << j << " ";
    }    
    cout << "\n";
  }
}

int main() {
  // TODO: Declare a "board" variable here, and store
  // the data provided above.
  vector<vector<int>> board{{0, 1, 0, 0, 0, 0},
                            {0, 1, 0, 0, 0, 0},
                            {0, 1, 0, 0, 0, 0},
                            {0, 1, 0, 0, 0, 0},
                            {0, 0, 0, 0, 1, 0}};

    // display the board
    PrintBoard(board);
}