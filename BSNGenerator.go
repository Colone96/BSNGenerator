package main

import (
	"fmt"
	"math/rand"
	"os"
	"time"
)

const numbers = "0123456789"

var path = "BSNs.txt"

func generateBSN() string {
	// creates a empty slice of 9 bytes
	b := make([]byte, 9)
	// adds a random number until it matches the length of the slice
	for i := range b {
		b[i] = numbers[rand.Intn(len(numbers))]
	}
	return string(b)
}

func createFile() {
	// check if file exists
	var _, err = os.Stat(path)

	// create file if not exists
	if os.IsNotExist(err) {
		var file, err = os.Create(path)
		if isError(err) {
			return
		}
		defer file.Close()
	}
	fmt.Println("Created new file:", path)
}

func writeFile(bsn string) {
	//opens file and sets readwrite and append rights to file
	file, err := os.OpenFile(path, os.O_RDWR|os.O_APPEND, 0644)
	if isError(err) {
		return
	}
	defer file.Close()

	// write bsn with a white line to file
	_, err = file.WriteString(bsn)
	_, err = file.WriteString("\n")
	if isError(err) {
		return
	}

	// save file
	err = file.Sync()
	if isError(err) {
		return
	}
}

// checks if the error is nil
func isError(err error) bool {
	if err != nil {
		fmt.Println(err.Error())
	}
	return (err != nil)
}

func main() {
	createFile()
	// creates random seed
	rand.Seed(time.Now().Unix())
	//creates 10 BSN numbers and write them to a file
	for i := 1; i <= 10; i++ {
		writeFile(generateBSN())
	}
}
