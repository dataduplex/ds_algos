package main

import (
	"bufio"
	"fmt"
	"io/ioutil"
	"log"
	"os"
)

func main () {
	//readFile()
	//readFileLineByLine()
	stringTests()
}


/*

Read from a file
 - line by line
 - bytes

 */

func readFile() {
	content, err := ioutil.ReadFile("/Users/jelumala/.ssh/config")
	if err != nil {
		log.Fatal(err.Error())
	}
	fmt.Printf("content: %s\n", content)
}

func readFileLineByLine() {

	file, err := os.Open("/Users/jelumala/.ssh/config")
	if err != nil {
		log.Fatal(err)
	}
	defer file.Close()

	//bufio.NewReader()
	//r := strings.NewReader("Hello, Reader!")
	//file.Read()
	//file.Write()
	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		fmt.Println(scanner.Text())
	}
}

func stringTests() {
	str := "jayanthan"
	b := []byte{}
	for i:=0; i<len(str); i++ {
		fmt.Printf("%c\n",str[i])
		fmt.Printf("%b\n",str[i])
		b = append(b, str[i])
	}

	fmt.Printf("%+v\n",b)

	b2 := []byte{'a','b','c'}
	fmt.Printf("byte slice #2 contents: %c\n",b2)

}