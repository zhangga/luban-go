#pragma once

#include <string>
#include <vector>
#include <map>
#include <set>
#include <memory>


namespace demo {



class Item {
public:


    int32_t Id;


    std::string Name;


    std::string Desc;


    Item(int32_t _Id, std::string _Name, std::string _Desc)
     {
        
        this->Id = _Id;
        
        this->Name = _Name;
        
        this->Desc = _Desc;
        
    }
    
    
};


}
