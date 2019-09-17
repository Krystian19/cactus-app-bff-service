run : build
	./app

# gen :
# 	(cd gqlgen/ && go run github.com/99designs/gqlgen)

build : **/*.go
	go build -o app -i bin/*.go