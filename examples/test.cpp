#ifndef RES_INCLUDE
#include "resource.h"
#else
#include RES_INCLUDE
#endif

#include <iostream>

int main() {
    std::cout<<test::res::multiline_example()<<std::endl;
}