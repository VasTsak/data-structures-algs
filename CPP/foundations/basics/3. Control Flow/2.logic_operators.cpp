#include<iostream>
#include<string>

int main()
{
    int A = 1;
    int B = 2;
    int C = 2;
    int D = 3;

    std::string TorF[] = {"False", "True"}; 

    //A true and false = false
    std::cout << "\n\n(A==C) && (A == D) is " << TorF[(A == C) && (A == D)];

    //A true and true = true
	std::cout<<"\n(B == C) && (B < D) is "<<TorF[(B == C) && (B < D)];

    //The || operator
	//A false || true = true
	std::cout<<"\n\n(A == C) || (B == C) is "<<TorF[(A == C) || (B == C)];  
	//A false || false = false
	std::cout<<"\n(A == C) || (B > D) is "<<TorF[(A == C) || (B > D)];  

    //The 'Not' operator
	// true = true
	// !(true) = false
	std::cout<<"\n\nA < B is "<<TorF[A<B];
	std::cout<<"\n!(A < B) is "<<TorF[!(A<B)];

    // false= false
	// !(false) = true
	std::cout<<"\n\nA == C is "<<TorF[A==C];
	std::cout<<"\n!(A == C) is "<<TorF[!(A==C)] << "\n";  
}