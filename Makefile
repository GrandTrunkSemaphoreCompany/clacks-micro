GOFILES = $(shell find . -name '*.go')
#FEATHER_M4_DEV_PATH=/dev/ttyACM0

default: build-feather-m4

build-feather-m4: $(GOFILES)
	tinygo flash -target=feather-m4 main.go

test:
	go test ./...

flash:
	tinygo flash -target=feather-m4 main.go

flash-with-debug:
	tinygo flash -target=feather-m4 -port ${FEATHER_M4_DEV_PATH} main.go