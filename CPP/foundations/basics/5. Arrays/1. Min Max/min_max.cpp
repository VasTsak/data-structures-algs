/*Find the min and max and average of 15 numbers that a user will input.
**The numbers range from 0 to 100. 
**We will do it now for practice and again later when we learn arrays and 
**functions. So you do not have to keep all fifteen numbers stored in memory.
**HINT: you will need to use scanf for input text.
*/
#include<iostream>
 int main()
 {
 	int userInput = 0;
 	int maxNumber = 0;
 	int minNumber = 100;
 	int sumTotal = 0;
 	float average = 0;

 	for (int i = 0; i < 15; i++)
 	{
 		std::cout << "Enter the number: ";
 		std::cin>>userInput;
 		std::cout << userInput << "\n";
 		if (userInput > maxNumber)
 		{
 			maxNumber = userInput;
 		}
 		if (userInput < minNumber)
 		{
 			minNumber = userInput;
 		}
 		sumTotal =+ userInput;
 	}
 	std::cout << "Maximum number = "<<maxNumber << "\n";
 	std::cout << "Minimum number = "<<minNumber << "\n";
 	average = sumTotal / 15.0;
 	std::cout << "Average = " << average << "\n";
 	return 0;
 }