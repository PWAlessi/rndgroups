package main

import (
	"bytes"
	"github.com/docopt/docopt-go"
	"io/ioutil"
	"log"
	"math/rand"
	"time"
)

func main() {
	// Set up docs
	usage := `Usage: 
				rndgroups <filename> <groupsize>

`
	arguments, err := docopt.ParseDoc(usage)

	if err != nil {
		println(err)
	}
	// Read in file
	fileName, _ := arguments.String("<filename>")
	stringList := readFile(fileName)

	// shuffle
	shuffle(stringList, time.Now().UnixNano())

	// output in groups
	groupSize, _ := arguments.Int("<groupsize>")
	printGroups(stringList, groupSize)

}

func shuffle(data []string, seed int64) {
	// Get random groups of groupSize from data
	rand.Seed(seed)

	for i := 0; i < 100000; i++ {
		i, j := rand.Intn(len(data)), rand.Intn(len(data))

		// swap
		temp := data[i]
		data[i] = data[j]
		data[j] = temp
	}
}

func readFile(path string) []string {
	content, err := ioutil.ReadFile(path)
	if err != nil {
		log.Fatal(err)
	}

	return convertByteSliceToList(content)
}

func convertByteSliceToList(byteSlice []byte) []string {
	byteList := bytes.Split(byteSlice, []byte("\n"))

	retVal := make([]string, len(byteList))

	for i, val := range byteList {
		retVal[i] = string(val)
	}

	return retVal
}

func printGroups(data []string, groupSize int) {
	for i := 0; i < len(data); {
		for j := 0; j < groupSize; j++ {
			if i < len(data) {
				print(data[i])

				// Omit trailing dash
				if j < groupSize-1 {
					print(" - ")
				}
			}
			i++
		}
		print("\n")
	}
}
