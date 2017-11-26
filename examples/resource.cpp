/* This file has been generated using risgo, do not modify! */
#include <unordered_map>
#include "examples/resource.h"

namespace test {

    std::string res::string_test() {
        static char const literal[] = {
112, 108, 97, 105, 110, 32, 116, 101, 120, 116, 
        ;
        return std::string(literal, sizeof(literal)/sizeof(char));
    }
    std::string res::binary_file_test() {
        static char const literal[] = {
46, 46, 46, 116, 111, 100, 111, 
        ;
        return std::string(literal, sizeof(literal)/sizeof(char));
    }
    std::string res::itself() {
        static char const literal[] = {
46, 46, 46, 116, 111, 100, 111, 
        ;
        return std::string(literal, sizeof(literal)/sizeof(char));
    }
    std::string res::itself_packed_lz4() {
        static char const literal[] = {
46, 46, 46, 116, 111, 100, 111, 
        ;
        return std::string(literal, sizeof(literal)/sizeof(char));
    }
    std::string res::custom_member_name() {
        static char const literal[] = {
116, 101, 115, 116, 
        ;
        return std::string(literal, sizeof(literal)/sizeof(char));
    }
    std::string res::multiline_example() {
        static char const literal[] = {
116, 104, 101, 115, 101, 32, 97, 114, 101, 10, 10, 116, 104, 114, 101, 101, 32, 108, 105, 110, 101, 115, 
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
