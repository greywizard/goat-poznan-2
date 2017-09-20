package main

import (
	"errors"
	"fmt"
	"log"
)

func main() {
	err := test()
	if err != nil {
		log.Println(err)
	} 

	test2()								// Errors unhandled.,LOW,HIGH (gas)
										// error return value not checked (test2()) (errcheck)
}

func test() (err error) {
	if true {
		err := errors.New("Test")		// declaration of "err" shadows declaration at example.go:21 (vetshadow)
		_ = err
	}
	return err
}

func test2() (err error) {
	err = errors.New("Error1")			// ineffectual assignment to err (ineffassign)
										// this value of err is never used (SA4006) (staticcheck)
	err = errors.New("Error2")
	return err
}

