@echo off

IF "%CI%"=="" (set DOCTEST_OPT=) ELSE (set DOCTEST_OPT=-nc)

go get ./...
if %ERRORLEVEL% neq 0 exit /b %ERRORLEVEL%

go build
if %ERRORLEVEL% neq 0 exit /b %ERRORLEVEL%

risgo.exe examples/test.yml
if %ERRORLEVEL% neq 0 exit /b %ERRORLEVEL%

cl examples\test.cpp examples\resource.cpp /EHsc /Iexamples /I. /Fetest
if %ERRORLEVEL% neq 0 exit /b %ERRORLEVEL%

test.exe %DOCTEST_OPT%
if %ERRORLEVEL% neq 0 exit /b %ERRORLEVEL%


@rem testing backslashes
risgo.exe examples\test.yml --template examples\cpp03.yml --header examples\resource03.h --source examples\resource03.cpp
if %ERRORLEVEL% neq 0 exit /b %ERRORLEVEL%

cl examples\test.cpp examples\resource03.cpp /EHsc /DRES_INCLUDE=\"resource03.h\" /Iexamples /I. /Fetest03
if %ERRORLEVEL% neq 0 exit /b %ERRORLEVEL%

test03.exe %DOCTEST_OPT%
if %ERRORLEVEL% neq 0 exit /b %ERRORLEVEL%
