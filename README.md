# Error Code Handling Using Stringer

## Prerequisites

- Go >= 1.16
- Install `stringer` by running `go get golang.org/x/tools/cmd/stringer`

## How To Use

- Add new error code into `err/err.go`.
- Re-generate error code `String()` in `errorcode_string.go`.
- Run `go run main.go` to see how the error code can be printed.
