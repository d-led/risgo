#ifndef RES_INCLUDE
#define RES_INCLUDE "resource.h"
#endif

#include RES_INCLUDE

#define DOCTEST_CONFIG_IMPLEMENT_WITH_MAIN
#include <doctest.h>

#include <iostream>
#include <fstream>
#include <string>

namespace {
std::string readAllBytes(std::string const& filename)
{
    std::ifstream file(filename, std::ios::binary);
    if (!file)
        throw std::runtime_error(std::string("cannot open ") + filename);

    return std::string(std::istreambuf_iterator<char>(file),
                       std::istreambuf_iterator<char>());
}
}

TEST_CASE("files are embedded exactly")
{
    std::cout<<"Used header: " RES_INCLUDE<<std::endl;
    CHECK(readAllBytes("examples/test.yml") == test::res::itself());
}

TEST_CASE("getting strings via custom member names and getters")
{
    CHECK(test::res::Get("custom.member/name") == test::res::custom_member_name());
    CHECK(test::res::custom_member_name().size() > 0);
}

TEST_CASE("files are not null-terminated (test.bin should be <<0,1,2,3>>)") {
    std::string binary(test::res::binary_file_test());
    CHECK(readAllBytes("examples/test.bin") == binary);
    CHECK(binary.size()>1);
    // by design of the test file
    CHECK(binary[0]==0);
}
