package main

import (
	"crypto/md5"
	"encoding/hex"
	"fmt"
)

func main() {
	key := "yzbqklnj"
	for i := 0; ; i++ {
		hash := getMD5Hash(fmt.Sprintf("%s%d", key, i))
		if hash[:6] == "000000" {
			fmt.Printf("The first number that produces a hash with five leading zeroes is: %d\n", i)
			break
		}
	}
}

func getMD5Hash(text string) string {
	hasher := md5.New()
	hasher.Write([]byte(text))
	hashBytes := hasher.Sum(nil)
	return hex.EncodeToString(hashBytes)
}
