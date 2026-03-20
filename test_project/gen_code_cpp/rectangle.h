#pragma once

#include <string>
#include <vector>
#include <map>
#include <set>
#include <memory>


namespace demo {



class Rectangle : public Shape {
public:


    float Width;


    float Height;


    Rectangle(int32_t _Id, float _Width, float _Height)
    : Shape(_Id) {
        
        this->Width = _Width;
        
        this->Height = _Height;
        
    }
    
    
};


}
