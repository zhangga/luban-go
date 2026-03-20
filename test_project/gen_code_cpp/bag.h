#pragma once

#include <string>
#include <vector>
#include <map>
#include <set>
#include <memory>


namespace demo {



class Bag {
public:


    int32_t Id;


    std::string Name;


    std::vector<int32_t> PriceList;


    std::map<std::string, int32_t> PropMap;


    Bag(int32_t _Id, std::string _Name, std::vector<int32_t> _PriceList, std::map<std::string, int32_t> _PropMap)
     {
        
        this->Id = _Id;
        
        this->Name = _Name;
        
        this->PriceList = _PriceList;
        
        this->PropMap = _PropMap;
        
    }
    
    
};


}
