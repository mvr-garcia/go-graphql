gqlgen-install:
	go get github.com/99designs/gqlgen@v0.17.24

generate: gqlgen-install
	go run github.com/99designs/gqlgen generate
