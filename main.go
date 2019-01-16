package main

import (
	"fmt"
	"io"
	"os"

	"golang.org/x/text/encoding/charmap"
)

func main() {
	if len(os.Args) != 3 {
		fmt.Println(usage)
		return
	}
	fmt.Printf("Converting %q to %q \n", os.Args[1], os.Args[2])
	f, err := os.Open(os.Args[1])
	if err != nil {
		panic(err)
	}
	defer func() {
		if err := f.Close(); err != nil {
			panic(err)
		}
	}()

	fout, err := os.Create(os.Args[2])
	if err != nil {
		panic(err)
	}
	defer func() {
		if err := fout.Close(); err != nil {
			panic(err)
		}
	}()

	r := charmap.Windows1256.NewDecoder().Reader(f)
	_, err = io.Copy(fout, r)
	if err != nil {
		panic(err)
	}
}

var usage = `Converting ASCII (Windows1256 encoded) text file to UTF-8 encoded text file.
Usage:
ascii2utf8 input.txt output.txt`
