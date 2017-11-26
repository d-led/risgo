build:
	go get ./...
	go build

test: build
	./risgo examples/test.yml
	g++ examples/test.cpp examples/resource.cpp -Iexamples -I. -std=c++11 -o test
	./test

	./risgo examples/test.yml --template examples/cpp03.yml --header examples/resource03.h --source examples/resource03.cpp
	g++ examples/test.cpp examples/resource03.cpp -DRES_INCLUDE=\"resource03.h\" -Iexamples -I. -std=c++03 -o test03
	./test03
