#include<iostream>

using namespace std;

int main()
{
    //Define months as having 12 possible values
    enum MONTHS {Jan, Feb, Mar, Apr, May, Jun, Jul, Aug, Sep, Oct, Nov, Dec}

    //Assign bestMonth one of the values of MONTHS
    bestMonth = Jan;

    //Now we check the values of bestMonths just like any other variable
    if(bestMonth == Jan)
    {
        cout<<"I'm not sure that January is the best month";
    }
    return 0;
}