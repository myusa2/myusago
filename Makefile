main: main.go
	go build main.go

test: main
	./test.sh

clean:
	rm -f main *.o tmp*

.PHONY: test clean