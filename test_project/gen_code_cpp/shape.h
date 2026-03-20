#pragma once

#include <string>
#include <vector>
#include <map>
#include <set>
#include <memory>


namespace demo {



class Shape {
public:


    int32_t Id;


    Shape(int32_t _Id)
     {
        
        this->Id = _Id;
        
    }
    
    virtual ~Shape() = default;
};


}
