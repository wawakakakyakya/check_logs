# /bin/sh

mod := github.com/wawakakakyakya/check_logs_by_mail
fileName := check_logs_by_mail
version := $(shell git describe --tags --abbrev=0)
revision := $(shell git rev-parse --short HEAD)
arch := nil

mod_tidy:
	go mod tidy
set_amd64:
	$(eval arch := amd64)
set_arm:
	$(eval arch := arm)

build:
	GOOS=linux GOARCH=$(arch) go build -ldflags "-X $(mod)/version.Version=$(version) -X $(mod)/version.Revision=$(revision)" -o ${fileName}_for_$(arch) main.go
build_linux: clean set_amd64 build
build_arm: clean set_arm build

clean:
	rm -f ${fileName}_for_*
test:
	./${fileName}_for_$(arch) -dest ../ -src ../
