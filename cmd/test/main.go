package main

import (
	"crypto/rand"
	"encoding/binary"
	"fmt"
	"os"
)

func random() (result int64, err error) {
	const desired = 16
	b := make([]byte, desired)
	size, err := rand.Read(b)
	if err != nil {
		return 0, fmt.Errorf("Failed to read system entropy: %w", err)
	} else if size < desired {
		return 0, fmt.Errorf("Failed to read desired size from system entropy")
	}
	return int64(binary.BigEndian.Uint64(b)), nil
}

func main() {
	n, err := random()
	if err != nil {
		fmt.Println(err)
		os.Exit(2)
	}
	fmt.Printf("%v\n", n)
}
