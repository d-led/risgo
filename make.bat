@echo off

go get ./...
if %ERRORLEVEL% neq 0 exit /b %ERRORLEVEL%

go build
if %ERRORLEVEL% neq 0 exit /b %ERRORLEVEL%

risgo.exe examples/test.yml
rem g++ examples/test.cpp examples/resource.cpp -Iexamples -I. -std=c++11 -o test
rem ./test

@rem testing backslashes
risgo.exe examples\test.yml --template examples\cpp03.yml --header examples\resource03.h --source examples\resource03.cpp
rem g++ examples/test.cpp examples/resource03.cpp -DRES_INCLUDE=\"resource03.h\" -Iexamples -I. -std=c++03 -o test03
rem ./test03
