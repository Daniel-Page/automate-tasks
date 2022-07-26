package main

import (
	"fmt"
	"io/ioutil"
	"log"
	"os"
	"time"
)

func deleteFiles() {
	fmt.Printf("%s\n", getTime())
	//fmt.Printf("Testing files...")
	deleteFile("./data.txt")
}

func deleteFile(fileName string) {
	e := os.Remove(fileName)
	if e != nil {
		log.Fatal(e)
	}
}

func addStartUp() {

}

func generateTestFiles() {
	f, err := os.Create("data.txt")

	if err != nil {
		log.Fatal(err)
	}

	defer f.Close()

	_, err2 := f.WriteString("Testing\n")

	if err2 != nil {
		log.Fatal(err2)
	}

	fmt.Println("done")
}

func findFiles() {
	files, err := ioutil.ReadDir("./")

	if err != nil {
		log.Fatal(err)
		fmt.Print("test")
	}

	for _, file := range files {
		fmt.Println(file.Name(), file.IsDir())
	}
}

func getTime() string {
	return time.Now().String()
}
