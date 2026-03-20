#pragma once

#include <string>
#include <vector>
#include <map>
#include <set>
#include <memory>


namespace demo {



class Circle : public Shape {
public:


    float Radius;


    Circle(int32_t _Id, float _Radius)
    : Shape(_Id) {
        
        this->Radius = _Radius;
        
    }
    
    
};


}
