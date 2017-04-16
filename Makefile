
SRC := $(wildcard *.go)

export CGO_CFLAGS = -I$(shell pg_config --includedir-server)

build: $(SRC)
	go build -o ptgo.so -buildmode=c-shared $(SRC)

