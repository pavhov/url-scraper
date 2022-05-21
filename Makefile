# Golang global paths

export GOPATH     := $(shell pwd)/
export GOVERSION  := go1.18
export GORELACE   := rc
export PATH       := ${GOPATH}bin:${PATH}

p				  := ""
s				  := ""
n				  := ""

init:
	mkdir -p bin pkg src/service/src

mod-init:
	cd ./src/service && \
	go mod init service

install:
	cd ./src/service/src && \
	go get -d -v -t

add:
	cd ./src/service/src && \
	go get -d -v -t ${p}

install-bin:
	cd ./src/service/src && \
	go install

download:
	cd ./src/service/src && \
	go mod download

update: clean
	cd ./src/service/src && \
	go get -u -d -v

clean:
	cd ./src/service/src && \
	go clean --modcache

tidy:
	cd ./src/service/src && \
	go mod tidy

build:
	cd ./src/service/src && \
	go build -ldflags="-w -s" -o ./pkg/app ./src && \
	chmod +x ./pkg/app

mock:
	mockgen -source=src\service\src\${s} -self_package=src\service\src\${s} -destination=src\service\src\mock\${s} -package=${p} -mock_names=${n} --build_flags=--mod=mod

run:
	cd ./src/service && \
	./pkg/app

enable-git-hooks:
	chmod u+x ./hooks/*
	git config core.hooksPath ${GOPATH}hooks
	git config advice.ignoredHook false
