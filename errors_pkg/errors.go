package main

import (
	"log"

	"github.com/pkg/errors"
)

func mainx() {

	baseErr := errors.New("Example Error")
	err := errors.WithStack(baseErr)
	log.Printf("%+v", err)

	err = errors.Wrap(baseErr, "GOat Poznań #2")
	log.Printf("%v", err)

}
