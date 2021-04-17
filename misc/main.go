package main

import (
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

	println("\n number of bytes written: ", n)
}
