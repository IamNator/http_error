package main

import (
	"errors"
	"log"
	"os"

	"github.com/IamNator/httperror"
)

func main() {
	testEr := httperror.New(nil)
	print("\n error object: ")
	n, err := testEr.WriteToWriter(os.Stdout)
	if err != nil {
		log.Println(err.Error())
	}

	log.Println(returnError("\nsome error"))

	println("\n number of bytes written: ", n)

}

func returnError(er string) httperror.Error {
	return httperror.New(errors.New(er))
}
