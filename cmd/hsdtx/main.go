package main

import (
	"bufio"
	"fmt"
	"os"
	"encoding/hex"

	"goshake/primitives"
	"goshake/util"
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

	fmt.Printf("\n")

	reader := new(util.Reader)
	reader.Init(data[:])

	tx := new(primitives.TX)
	tx.Read(reader)
	tx.Print()
}
