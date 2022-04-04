PROJECT=client
BINARY=bin/${PROJECT}

clean:
	rm -f ${BINARY}

build: clean
	cd ./cmd/app && \
	CGO_ENABLED=0 go build -o ../../${BINARY} .
