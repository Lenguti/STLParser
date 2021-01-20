# STLParser

## Installation

To run this project you will need either golang or docker installed.
You can install the latest go version [here](https://golang.org/doc/install). If you would like to use docker instead you can install docker [here](https://docs.docker.com/get-docker).

Once inside the root of the project you can build the binary from go or docker as such.
```bash
make build
or
make docker-build
```

## Usage
If you would like to parse an STL file, place the file inside the `files` directory. You can run the parser with go or with docker as such.
```bash
FILE=files/sample.stl make run
or
FILE=files/sample.stl make docker-run
```
