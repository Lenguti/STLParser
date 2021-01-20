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

This parser will parse the contents of an STL file and output how many triangles, the surface area, and the bounding box of your object.
If you would like to parse an stl file, place the file inside the `files` directory. you can run the parser with go or with docker as such.
```bash
file=files/sample.stl make run
or
file=files/sample.stl make docker-run
```

## Design/Improvements

For the design of the parser I decided to create Token identifiers of what is pertinent to the contents of an STL file. The Lexer reads the file per byte and determines the tokenzation. The Parser consumes the Tokens and determines if we have a valid sequence of tokens for an STL file and is in charge of building our object from the data values of the tokens. Once we have built our object from the contents I created helper methods to calculate how many triangles, surface area, and bounding box. As the current design is loading the whole file in memory, we would need about 2MB for a million of triangles. I am doing deffered calculations once the whole file has been parsed. Improvements that can be made is do calculations onces each triangle has been parsed. Also, instead of loading the file into memory we can stream the contents of the file and parse/calculate chunk by chunk. I think those two improvements could give a potentially unlimited threshhold of triangles to compute.
