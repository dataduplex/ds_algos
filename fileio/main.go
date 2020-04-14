package main

import (
	"fmt"
	"io"
	"io/ioutil"
	"log"
	"os"
)

func main() {
	read("main.go")
	readDirFiles(".")
	//archiveDir(".", "my_tar.tar")
	archiveDir2(".", "my_tar.tar")
}

func read(fileName string) {

	// Read entire file into a byte slice
	content, err := ioutil.ReadFile(fileName)
	if err != nil {
		log.Fatal(err)
	}

	fmt.Printf("File contents: %s", content)

	// Read file contents into a byte slice of specified length
	file, err := os.Open(fileName) // For read access.
	if err != nil {
		log.Fatal(err)
	}

	for {
		data := make([]byte, 50)
		count, err := file.Read(data)
		if err == io.EOF {
			break
		}
		if err != nil {
			log.Fatal(err)
		}
		fmt.Printf("read %d bytes: %q\n", count, data[:count])
	}

}

func readDirFiles(dir string) {

	files, err := ioutil.ReadDir(dir)
	if err != nil {
		log.Fatal(err)
	}

	for _, file := range files {
		fmt.Println(file.Name())
	}

}
