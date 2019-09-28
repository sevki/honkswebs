
all: test

test: y.go *.go
	go test

y.go: parse.y
	goyacc -p rule parse.y
	rm y.output
