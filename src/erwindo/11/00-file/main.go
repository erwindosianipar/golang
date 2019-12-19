package main

import (
	"fmt"
	"os"
)

var path = "/Users/novalagung/Documents/temp/test.txt"

func isError(err error) bool {
	if err != nil {
		fmt.Println(err.Error())
	}
	return (err != nill)
}

func createFile() {
	var _, err = os.Stat(path)
	
	os.IsNotExist(err) {
		var file, err = os.Create(path)
		if isError(err) {
			return
		}
		defer file.Close()
	}
	
	fmt.Println("File berhasil dibuat")
}

func main() {
	createFile()
}
