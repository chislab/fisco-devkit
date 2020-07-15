.PHONY: init deps up down clean

ROOT = $(PWD)

SOLC = solc:0.6.0

deps:
	rm -rf build
	$(call compile,hello)

init:
	rm -rf bin nodes build vendor && mkdir -p bin build/hello && cp hello.go build/hello/hello.go
	go mod tidy && go mod vendor
	git submodule update --init --recursive
	rm -rf vendor/github.com/chislab/go-fiscobcos/crypto/secp256k1 && tar -xvf secp256k1.tar.gz
	mv -f secp256k1 vendor/github.com/chislab/go-fiscobcos/crypto/
	bash build_chain.sh -l "127.0.0.1:4" -i -p 30300,20200,8545

down:
	bash nodes/127.0.0.1/stop_all.sh

up:
	rm -rf keys_certs
	bash nodes/127.0.0.1/start_all.sh

clean:
	rm -rf $(ROOT)/nodes/127.0.0.1/node0/data
	rm -rf $(ROOT)/nodes/127.0.0.1/node1/data
	rm -rf $(ROOT)/nodes/127.0.0.1/node2/data
	rm -rf $(ROOT)/nodes/127.0.0.1/node3/data

define compile
	mkdir -p ./build/$(1)
	docker run --rm -v $(ROOT)/contracts:/sources -v $(ROOT)/build/:/output ethereum/$(SOLC) --overwrite --abi --bin -o /output /sources/$(1).sol
	$(ROOT)/bin/abigen --bin=./build/$(1).bin --abi=./build/$(1).abi --pkg=$(1) --out=./build/$(1)/$(1).go
endef
