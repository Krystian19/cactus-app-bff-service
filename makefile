run : build
	./app

gen :
	(cd gql/ && go run github.com/99designs/gqlgen)

build : **/*.go
	go build -o app -i bin/*.go