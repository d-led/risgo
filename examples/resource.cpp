/* This file has been generated using risgo, do not modify! */
#include <unordered_map>
#include "examples/resource.h"

namespace test {

    std::string res::bla() {
        static char const literal[] = {
104, 101, 108, 
108, 111, 44, 
32, 119, 111, 
114, 108, 100, 

        ;
        return std::string(literal, sizeof(literal)/sizeof(char));
    }

std::string res::Get(std::string const& key) {
    static std::unordered_map<std::string,ResourceGetter> getters = {
    };
    auto getter = getters.find(key);
    if (getter == getters.end())
         return OnNoKey(key);
    return getter->second();
}

}
