package main

import (
	"fmt"
	"log"
	"os"
)

func main() {
	lockFile := ".b4.lock"

	file, err := os.Open(lockFile)
	if os.IsNotExist(err) {
		file.Close()
		os.WriteFile(lockFile, []byte(""), 0644)
	} else {
		file.Close()
		log.Fatal(fmt.Errorf("File %s exists, run ./b4.sh to perform pre-check and delete\n", lockFile))
	}
}
