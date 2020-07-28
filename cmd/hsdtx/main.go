package main

import (
	"bufio"
	"fmt"
	"os"
	"encoding/hex"

	"goshake/primitives"
)

func main() {
	var data []byte
	var err error
    args := os.Args
	if len(args) != 2 {
		scanner := bufio.NewScanner(os.Stdin)
		fmt.Print("Enter transaction data in hex: ")
		scanner.Scan()
		data, err = hex.DecodeString(scanner.Text())
	} else {
		data, err = hex.DecodeString(args[1])
	}
	if err != nil {
		panic(err)
	}

	var offset uint64
	fmt.Printf("\n")

	tx := new(primitives.TX)
	tx.Read(data[:], &offset)
	tx.Print()
}
