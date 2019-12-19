package main

import (
	"fmt"
	"log"
	"os"
)

type student struct {
	name string
	age  int
}

var students = []student{
	{"Devi", 19},
	{"Debi", 20},
}

var path = "D:/golang/src/enigma/10/p3/folder/file.txt"

func main() {
	//createFile()
	writeFile()
	readFile()
}

func createFile() {
	var file, _ = os.Create(path)
	file.Close()
}

func writeFile() {
	var file, err = os.OpenFile(path, os.O_APPEND|os.O_CREATE, 0644)
	if err != nil {
		log.Fatal(err)
	}

	for _, student := range students {
		strStudent := fmt.Sprintf("%v, %v\n", student.name, student.age)
		file.WriteString(strStudent)
	}

	file.Sync()
}

func readFile() {
	var file, err = os.OpenFile(path, os.O_RDONLY, 0644)
	if err != nil {
		log.Fatal(err)
	}

	var text = make([]byte, 1)
	for {
		n, _ := file.Read(text)
		if n == 0 {
			break
		}
		fmt.Print(string(text))
	}
}
