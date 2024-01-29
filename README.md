# Evolucion

## Developping

1. Install dependencies `go mod tidy`.
2. Run the backend `go run .`.
3. Visit [http://localhost:1323](http://localhost:1323).

If you want to use [air](https://github.com/cosmtrek/air) for live reloading:
1. Install air `go install github.com/cosmtrek/air@latest`.
2. Be sure that `$GOPATH/bin` is in your `$PATH`: `export PATH=$PATH:$GOROOT/bin:$GOPATH/bin` or `export PATH=$PATH:$(go env GOPATH)/bin`
3. Run `air` instead of `go run .`.
