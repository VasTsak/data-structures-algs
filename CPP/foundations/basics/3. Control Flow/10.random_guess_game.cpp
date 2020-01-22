#include <iostream>
#include<sstream>
#include <time.h> //added for the random number generator seed
#include <cstdlib>//added to use the rand function

int main()
{
    int target;
    std::string userString;
    int guess = -1;

    srand(time(NULL)); //set the seed for the random number generator
    target = rand() %10000 + 1; //generate the 'random' number

    while (guess != target)
    {
        std::cout << "Guess a number between 0 and 10000: ";
        std::getline(std::cin, userString);
        // convert string to an integer
        std::stringstream(userString) >> guess;
        std::cout <<userString<<"\n";
        if (guess > target)
        {
            std::cout << "Your guess is too high. \n";
        }
        else if (guess < target)
        {
            std::cout << "Your guess is too low. \n";
        }
        else
        {
            std::cout << "You have guessed correctly :)\n";
        }
        //Not I had to use double quotes around q, because it is string
        if(userString == "q")
        {
            std::cout << "good bye! \n";
            std::cout << "The number was " << target << "\n";
            break;
        }
    }
    return 0;
}
