package main

import (
	"flag"
	"fmt"
	"log"
	"strconv"
)

func main() {
	flag.Parse()
	inputList := flag.Args()

	bip39UpperBound := len(gBIP39English) - 1 //upperbounds is -1 as gBIP39English is a zero-based array
	for _, v := range inputList {
		bipPos, pErr := strconv.ParseInt(v, 10, 32)
		if pErr != nil {
			log.Printf("failed to parse %s: %+v", v, pErr)
			continue
		}

		bipPos--
		if bipPos < 0 || int(bipPos) > bip39UpperBound {
			log.Printf("value %s is out of range (must be between 0001 and %04d)", v, bip39UpperBound)
			continue
		}

		fmt.Printf("%04d - %s\n", bipPos+1, gBIP39English[bipPos])
	}
}
