all: compile

depend:
	go get -u "github.com/vit1251/materialize"

compile:
	go build 01_materialize_grid.go

check:

.PHONY: all compile check

