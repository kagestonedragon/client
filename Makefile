PROJECT=client
BINARY=bin/${PROJECT}

clean:
	rm -f ${BINARY}

build: clean
	cd ./cmd/app && \
	go build -o ../../${BINARY} .
