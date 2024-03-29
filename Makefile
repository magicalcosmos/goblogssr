all: v8build build

test: v8build gotest

v8build:
	cd v8worker/depsc++;chmod 755 build.sh;./build.sh

build:
	go build -o shareus

gql:
	go generate ./...

gqlgen:
	go run github.com/99designs/gqlgen generate.

gotest:
	go test -v ./v8worker

clean:
	rm -f ./v8worker/depsc++/libdepsc++.a
	rm -f ./v8worker/depsc++/v8binding.o
	go clean
