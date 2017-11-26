#pragma once
/* This file has been generated using risgo, do not modify! */
#include <string>

namespace test {

class res /*final*/ {
public:
    static std::string bla;
public:
    typedef std::string(*ResourceGetter)();
public: // key/value api
template <typename TInserter>
static void GetKeys(TInserter inserter) {
    static const char* keys[] = {
        "bla",
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
