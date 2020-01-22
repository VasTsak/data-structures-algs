/*
Arithmetic operations in C++ are similar to other languages.
We’ll have a quick discussion about them here and then you can practice programming.
When doing math operations you may need to include the cmath library, it contains a number of useful functions.
*/


/*Goal: practice arithmetic operations in C++
**Write a program that calculates the volumes of:
**a cube, sphere, cone.
**Cube Volume = side^3
**Sphere Volume = (4/3) * pi * radius^3
**Cone Volume = pi * radius^2 * (height/3)
**Write the values to the console.
*/

#include <iostream>
#include <cmath>

int main()
{
    // Dimension of the cube
    float cubeSide = 5.4;
    // Dimensions of sphere
    float sphereRadius = 2.33;
    float pi = 3.14159265358979323846;
    // Dimensions of cone
    float coneRadius = 7.65;
    float coneHeight = 14;

    float volCube, volSphere, volCone = 0;
    volCube = std::pow(cubeSide, 3);
    volSphere = (4/3) * pi * std::pow(sphereRadius, 3);
    volCone =  pi * std::pow(coneRadius, 2) * coneHeight/3;

    std::cout << "The volume of cube is: "<< volCube << "\n";
    std::cout << "The volume of sphere is: "<< volSphere << "\n";
    std::cout << "The volume of cone is: "<< volCone << "\n";

    return 0;
}