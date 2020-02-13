
#include <string>
#include <cstring>
#include <iostream>

class Car {
    // TODO: Declare private attributes
    private:
        int horsepower_;
        int weight_;
        char *brand;
        
    // TODO: Declare getter and setter for brand
    public:
        int GetHorsePower() const;
        int GetWeight() const;
        std::string GetBrand() const;
        void SetHorsePower(int horsepower);
        void SetWeight (int weight);
        void SetBrand(std::string brand_name);
};


// Define setters
void Car::SetHorsePower(int horsepower){
    Car::horsepower_ = horsepower;
}

void Car::SetWeight(int weight){
    Car::weight_ = weight;
}

void Car::SetBrand(std::string brand_name) {
    // Initialization of char array
    Car::brand = new char[brand_name.length() + 1];
    // copying every character from string to char array;
    strcpy(Car::brand, brand_name.c_str());
}
// Define getters

int Car::GetHorsePower() const {
    return Car::horsepower_;
};

int Car::GetWeight() const {
    return Car::weight_;
};

std::string Car::GetBrand() const {
    std::string result = "Brand name: ";
    // Specifying string for output of brand name
    result += Car::brand;
    return result;
};

// Test in main()
int main() 
{
    Car car;
    car.SetBrand("Peugeot");
    std::cout << car.GetBrand() << "\n";   
}
