all:
	$(CC) -g -fPIC -c -o lib/lib.o lib/lib.c
	$(CC) -g -fPIC -shared -o liblib.so lib/lib.o
	cp liblib.so wrapper/
	export GOPATH=$GOPATH:$PWD/../../../../
	gofmt -e -s -w .
	# go vet .
	LD_LIBRARY_PATH=. go run main.go
	LD_LIBRARY_PATH=. go run main2.go
