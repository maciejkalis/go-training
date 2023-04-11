// Implementation of the simple program that reads the content of the file,
// written on the hard drive, in terminal.
// The file name is passed to the program as an argument
// added when running the program in the terminal.

package main

import (
	"fmt"
	"io"
	"os"
)

func main() {
	f, err := os.Open(os.Args[1])
	if err != nil {
		fmt.Println("Error:", err)
		os.Exit(1)
	}

	io.Copy(os.Stdout, f)
}
