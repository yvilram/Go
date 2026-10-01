package main

import (
	"fmt"
	"io"
	"os"
)


func main(){

	file, err := os.Open(os.Args[1])

	if err != nil{
		fmt.Println(err)
		return
	}

	/* defer file.Close()

	deck := make([]byte, 100)

	read, err := file.Read(deck)

	if err != nil {
		fmt.Println(err)
		return
	}

	fmt.Println(string(deck[:read])) */

	io.Copy(os.Stdout,file)
}
