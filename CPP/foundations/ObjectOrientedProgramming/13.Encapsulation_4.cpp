
#include <cassert>
#include <stdexcept>

// TODO: Define "Student" class
class Student{
 public:
  // constructor
    Student(std::string name, int grade, float gpa){
        SetName(name);
        SetGrade(grade);
        SetGPA(gpa);
    }
  // accessors
    std::string Name() const;
    int Grade() const;
    float GPA() const;
    
  // mutators
    void SetName(std::string name);
    void SetGrade(int grade);
    void SetGPA(float gpa);
    
 private:
    std::string name_;
    int grade_;
    float gpa_;
};

std::string Student::Name() const {return Student::name_ ;}
int Student::Grade() const {return Student::grade_ ;}
float Student::GPA() const {return Student::gpa_ ;}

void Student::SetName(std::string name){
    Student::name_ = name;
}
void Student::SetGrade(int grade) {
    if (grade >= 0 && grade <= 12){
        Student::grade_ = grade;
    } else {
        throw std::invalid_argument("value not valid");
    }
}

void Student::SetGPA(float gpa) {
    if (gpa >= 0 && gpa <= 4){
        Student::gpa_ = gpa;
    } else {
        throw std::invalid_argument("value not valid");
    }
}

// TODO: Test
int main() {
    Student vasilis("vasilis", 12, 4);
    assert(vasilis.Name() == "vasilis");
    assert(vasilis.Grade() == 12);
    assert(vasilis.GPA() == 4);
    
    bool caught{false};
    try {
    Student invalid("invalid", 13, -2);
  } catch (...) {
    caught = true;
  }
  assert(caught);
}
