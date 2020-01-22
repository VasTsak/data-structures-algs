#include<iostream>
#include<string>

int main()
{
	std::string TorF[] = {"False", "True"};

	int a = 1;
	int b = 2;
	int c = 2;
	    
	//Print out the string values of each relational operation
	std::cout<<"a < b is "<<TorF[a<b];
	std::cout<<"\na > b is "<<TorF[a>b];
	std::cout<<"\na != b is "<<TorF[a!=b];
	std::cout<<"\nc >= b is "<<TorF[c>=b];
	std::cout<<"\nc <= b is "<<TorF[c<=b]<<"\n"; 
}