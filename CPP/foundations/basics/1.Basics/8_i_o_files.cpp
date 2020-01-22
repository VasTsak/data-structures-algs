#include<iostream>
#include<fstream>
#include<string>

using namespace std;

int main()
{
    string line;
    //Create an output stream to write to the file 
    //append the new lines to the end of the file 

    ofstream myfileI("input.txt", ios::app);
    if(myfileI.is_open())
    {
        myfileI << "\n I am adding a line \n";
        myfileI << "I am adding another line";
        myfileI.close();
    }
    else cout << "Unable to open the file for writing";

    //Create an input stream to read the file 
    ifstream myfile0("input.txt");
    //During the creation of ifstream, the file is opened.
    //So we do not have explicitly open the file. 
    if (myfile0.is_open())
    {
        while (getline(myfile0, line))
        {
            cout << line << '\n';
        }
        myfile0.close();
    }
    else cout<<"Unable to open the file for reading";

    return 0;
}
