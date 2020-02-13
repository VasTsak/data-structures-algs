
#include <cassert>
#include <stdexcept>
#include <iostream>

// TODO: Define class Pyramid
class Pyramid{
    // public class members
    public:
    
        Pyramid(double length, double width,double height) {
            // This is a constructors
            SetLength(length);
            SetWidth(width);
            SetHeight(height);      
        }
        double Length() const;
        double Width() const;
        double Height() const;
        void SetLength(double length);
        void SetWidth(double width);
        void SetHeight(double height);
        double Volume();
    // private class members
    private:
        double length_;
        double width_;
        double height_;
};

// define accessors

double Pyramid::Length() const { 
    return Pyramid::length_; 
};
double Pyramid::Width() const {
    return Pyramid::width_; 
};
double Pyramid::Height()const  { 
    return Pyramid::height_; 
};

// define mutators

void Pyramid::SetLength(double length) {
    if (length >= 0){
        Pyramid::length_ = length;
    } else{
        throw std::invalid_argument("value not valid");
    } 
void Pyramid::SetWidth(double width) {
    if (width >= 0){
        Pyramid::width_ = width;
    } else {
        throw std::invalid_argument("value not valid");
    } 
}
void Pyramid::SetHeight(double height) {
    if (height >= 0) {
        Pyramid::height_ = height;
    } else {
        throw std::invalid_argument("value not valid");
    }
}       


double Pyramid::Volume(){
    double l;
    double w;
    double h;
    l = Pyramid::Length();
    w = Pyramid::Width();
    h = Pyramid::Height();
    double volume = l * w * h / 3;
    std::cout<< volume;
    return volume;
    
};

// Test
int main() {
  Pyramid pyramid(4, 5, 6);
  assert(pyramid.Length() == 4);
  assert(pyramid.Width() == 5);
  assert(pyramid.Height() == 6);
  assert(pyramid.Volume() == 40);

  bool caught{false};
  try {
    Pyramid invalid(-1, 2, 3);
  } catch (...) {
    caught = true;
  }
  assert(caught);
}
