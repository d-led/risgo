#pragma once
/* This file has been generated using risgo, do not modify! */
#include <string>

namespace test {

class res /*final*/ {
public:
    static std::string string_test();
    static std::string binary_file_test();
    static std::string itself();
    static std::string custom_member_name();
    static std::string multiline_example();
public:
    typedef std::string(*ResourceGetter)();
public: // key/value api
template <typename TInserter>
static void GetKeys(TInserter inserter) {
    static const char* keys[] = {
        "string_test",
        "binary_file_test",
        "itself",
        "custom.member/name",
        "multiline_example",
    };
    for (auto key : keys) {
        inserter(key);
    }
}
public: // key/value api
    static std::string Get(std::string const& key);

public:
    static std::string OnNoKey(std::string const& key="") {
        return "";
    }
};

}
