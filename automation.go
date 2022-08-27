package main

import (
	"fmt"
	"io/ioutil"
	"log"
	"os"
	"os/user"
	"time"
)

//  To do: Check if program added to startup menu

type recycler struct {
	dir string
}

// Periodically moves files in a given folder to recycle bin
func Recycler(dir string) *recycler {
	return &recycler{dir}
}

func (r *recycler) run() {
	findFiles(r.dir)
}

func (r *recycler) print() {
	fmt.Print(r.dir)
}

func getDownloadsDirectory() string {
	currentUser, err := user.Current()
	if err != nil {
		log.Fatalf(err.Error())
	}
	homeDir := currentUser.HomeDir // e.g. C:\Users\Admin
	downloadsDir := homeDir + "\\Downloads"
	return downloadsDir
}

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

func findFiles(directory string) {
	files, err := ioutil.ReadDir(directory)

	if err != nil {
		log.Fatal(err)
		fmt.Print("test")
	}

	for _, file := range files {
		fmt.Println(file.Name(), file.IsDir())
	}
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

func getTime() string {
	return time.Now().String()
}
