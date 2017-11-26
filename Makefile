test:
	go run main.go examples/test.yml
	g++ examples/test.cpp examples/resource.cpp -Iexamples -I. -std=c++11 -o test
	./test
