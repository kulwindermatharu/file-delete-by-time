package main

import (
	"fmt"
	"io/ioutil"
	"log"
	"os"
	"time"
)

//cutoff is How Old File Delete
var cutoff = 4 * time.Hour

func main() {
	fileInfo, err := ioutil.ReadDir("File Path")
	if err != nil {
		log.Fatal(err.Error())
	}
	now := time.Now()
	for _, info := range fileInfo {
		if diff := now.Sub(info.ModTime()); diff > cutoff {
			os.Remove("File Path" + info.Name())
			fmt.Printf("Deleting %s which is %s old\n", info.Name(), diff)
		}
	}
}
