# stringsvc

A dockerised string service built in Go.

## Running stringsvc

To run with docker just enter

```bash
  docker run --rm -it -p 8080:8080 camjw/stringsvc
```

and go to `localhost:8080`.

If you want to build the Docker image yourself then first build the binary by running
```bash
docker run --rm -v "$PWD":/go/src/github.com/camjw/stringsvc -w /go/src/github.com/camjw/stringsvc iron/go:dev go build -o myapp
docker build -t camjw/stringsvc:latest .
```
and then run the app on port 8080:

```bash
docker run --rm -p 8080:8080 camjw/stringsvc
```

Otherwise, clone this repo, build the app with `go build main.go` and execute the build file `./main`.

## Endpoints

At the moment there are three endpoints: `/uppercase`, `/count` and `/metrics`.

The first two take a JSON of the form `{"s":"hello, world"}'` by a POST request and either return the uppercased string or the number of characters in the string.

```bash
curl -X POST -d'{"s":"hello, world"}' localhost:8080/uppercase
# {"v":"HELLO, WORLD"}
```

The `/metrics` endpoint returns a load of statistics about the usage of the app since it was instantiated.

## Licence

MIT

## Acknowledgements

This was built by following the `go-kit` example.
