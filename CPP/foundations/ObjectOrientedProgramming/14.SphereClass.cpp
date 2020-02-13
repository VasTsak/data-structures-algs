
#include <cassert>
#include <cmath>
#include <stdexcept>
#include <iostream>

// TODO: Define class Sphere
class Sphere {
 public:
  // Constructor
    Sphere(float radius){
        SetRadius(radius);
    }

  // Accessors
    float Radius() const;
    float Volume() const;
  // Mutators
    void SetRadius(float radius);

 private:
  // Private members
    float radius_;
};

float Sphere::Radius() const {return Sphere::radius_ ;}

void Sphere::SetRadius(float radius){
    if (radius >= 0){
        Sphere::radius_ = radius;
    } else{
        throw std::invalid_argument("value is not valid");
    }
}

float Sphere::Volume() const {
    float radius;
    float volume;
    radius = Sphere::Radius();
    volume =(pow(radius, 3) * M_PI) * (4 / 3.0) ;
    return volume ;
}


// Test
int main(void) {
  Sphere sphere(5);
  assert(sphere.Radius() == 5);
  assert(abs(sphere.Volume() - 523.6) < 1);
}
