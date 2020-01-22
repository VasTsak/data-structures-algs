#include<iostream> 

float add(float m1, float m2)
{
	return m1  + m2;
}

float sub(float m1, float m2)
{
	return m1 - m2;
}

float div(float m1, float m2)
{
	return m1 / m2;
}

float mul(float m1, float m2)
{
	return m1 * m2;
}

int main()
{
	std::cout<<"4 + 5 = "<<add(4.0, 5.0)<<"\n";
	std::cout<<"4 - 5 = "<<sub(4.0, 5.0)<<"\n";
	std::cout<<"4 / 5 = "<<div(4.0, 5.0)<<"\n";
	std::cout<<"4 * 5 = "<<mul(4.0, 5.0)<<"\n";
}