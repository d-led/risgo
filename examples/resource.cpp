/* This file has been generated using risgo, do not modify! */
#include <unordered_map>
#include "examples/resource.h"

namespace test {

    std::string res::string_test() {
        static char const literal[] = {
112, 108, 97, 105, 110, 32, 116, 101, 120, 116 }
        ;
        return std::string(literal, sizeof(literal)/sizeof(char));
    }
    std::string res::binary_file_test() {
        static char const literal[] = {
49, 50, 51, 10, 49, 50, 51, 10 }
        ;
        return std::string(literal, sizeof(literal)/sizeof(char));
    }
    std::string res::itself() {
        static char const literal[] = {
123, 10, 32, 32, 32, 34, 104, 101, 97, 100, 101, 114, 34, 58, 32, 34, 97, 99, 99, 101, 112, 116, 97, 110, 99, 101, 95, 116, 101, 115, 116, 47, 114, 101, 115, 111, 117, 114, 99, 101, 46, 104, 34, 44, 10, 32, 32, 32, 34, 110, 97, 109, 101, 115, 112, 97, 99, 101, 34, 58, 32, 34, 116, 101, 115, 116, 34, 44, 10, 32, 32, 32, 34, 99, 108, 97, 115, 115, 34, 58
, 32, 34, 114, 101, 115, 34, 44, 10, 32, 32, 32, 34, 114, 101, 115, 111, 117, 114, 99, 101, 115, 34, 58, 32, 91, 10, 32, 32, 32, 32, 32, 32, 123, 10, 32, 32, 32, 32, 32, 32, 32, 32, 32, 34, 110, 97, 109, 101, 34, 58, 32, 34, 115, 116, 114, 105, 110, 103, 95, 116, 101, 115, 116, 34, 44, 10, 32, 32, 32, 32, 32, 32, 32, 32, 32, 34, 115, 111, 117, 114
, 99, 101, 34, 58, 32, 34, 112, 108, 97, 105, 110, 32, 116, 101, 120, 116, 34, 44, 10, 32, 32, 32, 32, 32, 32, 32, 32, 32, 34, 115, 111, 117, 114, 99, 101, 95, 116, 121, 112, 101, 34, 58, 32, 34, 115, 116, 114, 105, 110, 103, 34, 10, 32, 32, 32, 32, 32, 32, 125, 44, 10, 32, 32, 32, 32, 32, 32, 123, 10, 32, 32, 32, 32, 32, 32, 32, 32, 32, 34, 110
, 97, 109, 101, 34, 58, 32, 34, 98, 105, 110, 97, 114, 121, 95, 102, 105, 108, 101, 95, 116, 101, 115, 116, 34, 44, 10, 32, 32, 32, 32, 32, 32, 32, 32, 32, 34, 115, 111, 117, 114, 99, 101, 34, 58, 32, 34, 116, 101, 115, 116, 46, 98, 105, 110, 34, 44, 10, 32, 32, 32, 32, 32, 32, 32, 32, 32, 34, 115, 111, 117, 114, 99, 101, 95, 116, 121, 112, 101, 34, 58
, 32, 34, 102, 105, 108, 101, 34, 10, 32, 32, 32, 32, 32, 32, 125, 44, 10, 32, 32, 32, 32, 32, 32, 123, 10, 32, 32, 32, 32, 32, 32, 32, 32, 32, 34, 110, 97, 109, 101, 34, 58, 32, 34, 105, 116, 115, 101, 108, 102, 34, 44, 10, 32, 32, 32, 32, 32, 32, 32, 32, 32, 34, 115, 111, 117, 114, 99, 101, 34, 58, 32, 34, 116, 101, 115, 116, 46, 106, 115, 111
, 110, 34, 44, 10, 32, 32, 32, 32, 32, 32, 32, 32, 32, 34, 115, 111, 117, 114, 99, 101, 95, 116, 121, 112, 101, 34, 58, 32, 34, 102, 105, 108, 101, 34, 10, 32, 32, 32, 32, 32, 32, 125, 44, 10, 32, 32, 32, 32, 32, 32, 123, 10, 32, 32, 32, 32, 32, 32, 32, 32, 32, 34, 99, 111, 109, 112, 114, 101, 115, 115, 105, 111, 110, 34, 58, 32, 34, 76, 90, 52
, 70, 34, 44, 10, 32, 32, 32, 32, 32, 32, 32, 32, 32, 34, 110, 97, 109, 101, 34, 58, 32, 34, 105, 116, 115, 101, 108, 102, 95, 112, 97, 99, 107, 101, 100, 95, 108, 122, 52, 34, 44, 10, 32, 32, 32, 32, 32, 32, 32, 32, 32, 34, 115, 111, 117, 114, 99, 101, 34, 58, 32, 34, 116, 101, 115, 116, 46, 106, 115, 111, 110, 34, 44, 10, 32, 32, 32, 32, 32, 32
, 32, 32, 32, 34, 115, 111, 117, 114, 99, 101, 95, 116, 121, 112, 101, 34, 58, 32, 34, 102, 105, 108, 101, 34, 10, 32, 32, 32, 32, 32, 32, 125, 44, 10, 32, 32, 32, 32, 32, 32, 123, 10, 32, 32, 32, 32, 32, 32, 32, 32, 32, 34, 109, 101, 109, 98, 101, 114, 95, 110, 97, 109, 101, 34, 58, 32, 34, 99, 117, 115, 116, 111, 109, 95, 109, 101, 109, 98, 101, 114
, 95, 110, 97, 109, 101, 34, 44, 10, 32, 32, 32, 32, 32, 32, 32, 32, 32, 34, 110, 97, 109, 101, 34, 58, 32, 34, 99, 117, 115, 116, 111, 109, 46, 109, 101, 109, 98, 101, 114, 47, 110, 97, 109, 101, 34, 44, 10, 32, 32, 32, 32, 32, 32, 32, 32, 32, 34, 115, 111, 117, 114, 99, 101, 34, 58, 32, 34, 116, 101, 115, 116, 34, 44, 10, 32, 32, 32, 32, 32, 32
, 32, 32, 32, 34, 115, 111, 117, 114, 99, 101, 95, 116, 121, 112, 101, 34, 58, 32, 34, 115, 116, 114, 105, 110, 103, 34, 10, 32, 32, 32, 32, 32, 32, 125, 44, 10, 32, 32, 32, 32, 32, 32, 123, 10, 32, 32, 32, 32, 32, 32, 32, 32, 32, 34, 110, 97, 109, 101, 34, 58, 32, 34, 109, 117, 108, 116, 105, 108, 105, 110, 101, 95, 101, 120, 97, 109, 112, 108, 101, 34
, 44, 10, 32, 32, 32, 32, 32, 32, 32, 32, 32, 34, 115, 111, 117, 114, 99, 101, 34, 58, 32, 34, 116, 104, 101, 115, 101, 32, 97, 114, 101, 92, 110, 92, 110, 116, 104, 114, 101, 101, 32, 108, 105, 110, 101, 115, 34, 44, 10, 32, 32, 32, 32, 32, 32, 32, 32, 32, 34, 115, 111, 117, 114, 99, 101, 95, 116, 121, 112, 101, 34, 58, 32, 34, 115, 116, 114, 105, 110, 103
, 34, 10, 32, 32, 32, 32, 32, 32, 125, 10, 32, 32, 32, 93, 44, 10, 32, 32, 32, 34, 115, 111, 117, 114, 99, 101, 34, 58, 32, 34, 97, 99, 99, 101, 112, 116, 97, 110, 99, 101, 95, 116, 101, 115, 116, 47, 114, 101, 115, 111, 117, 114, 99, 101, 46, 99, 112, 112, 34, 10, 125 }
        ;
        return std::string(literal, sizeof(literal)/sizeof(char));
    }
    std::string res::custom_member_name() {
        static char const literal[] = {
116, 101, 115, 116 }
        ;
        return std::string(literal, sizeof(literal)/sizeof(char));
    }
    std::string res::multiline_example() {
        static char const literal[] = {
116, 104, 101, 115, 101, 32, 97, 114, 101, 10, 10, 116, 104, 114, 101, 101, 32, 108, 105, 110, 101, 115 }
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
