package main

import (
	"encoding/json"
	"fmt"
	"log"

	"demo/err"
)

type AppError struct {
	Code err.ErrorCode `json:"code,omitempty"`
	Name string        `json:"name,omitempty"`
}

func main() {
	res, e := json.Marshal(appError(err.ERR_INVALID_PARAM))
	if e != nil {
		log.Fatal(e)
	}

	fmt.Printf("%+v\n", string(res))

	res, e = json.Marshal(appError(err.ERR_INVALID_EMAIL, true))
	if e != nil {
		log.Fatal(e)
	}

	fmt.Printf("%+v\n", string(res))
}

func appError(e err.ErrorCode, noName ...bool) AppError {
	name := e.String()

	if len(noName) > 0 && noName[0] {
		name = ""
	}

	return AppError{
		Code: e,
		Name: name,
	}
}
