build:
	go get ./...
	go build

test: build
	./risgo examples/test.yml
	g++ examples/test.cpp examples/resource.cpp -Iexamples -I. -std=c++11 -o test
	./test
