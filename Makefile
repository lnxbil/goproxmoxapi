.PHONY: love

all: clean build install test-all

install:
	$(info Installing...)
	go install

build:
	$(info Building...)
	go clean

clean:
	$(info Cleaning...)
	go clean

test-all:
	$(info testing...)
	go test -test.v -parallel 20

test:
	$(info testing $run)
	go test -test.v -run $(run)

love:
	@echo not war
