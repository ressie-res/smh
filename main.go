package main

import (
	"crypto/sha256"
	"encoding/hex"
	"fmt"
)

func main() {

	var testpassword string
	var hashespassword string

	hashthisfella := sha256.New()

	fmt.Scan(&testpassword)
	hashthisfella.Write([]byte(testpassword))
	hashespassword = hex.EncodeToString(hashthisfella.Sum(nil))

	fmt.Println(hashespassword)

}
