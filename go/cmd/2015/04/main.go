package main

import (
	"crypto/md5"
	"fmt"
	"strings"
)

func main() {
	fmt.Printf("%d\n", calcHash("ckczppom"))
	fmt.Printf("%d\n", calcHashSixZeroesTrailing("ckczppom"))
}

func calcHash(key string) int {
	var answer int
	prefix := "00000"

	for {
		hash := md5.Sum([]byte(fmt.Sprintf("%s%d", key, answer)))
		if strings.HasPrefix(fmt.Sprintf("%x", hash), prefix) {
			return answer
		}
		answer++
	}
}

func calcHashSixZeroesTrailing(key string) int {
	var answer int
	prefix := "000000"

	for {
		hash := md5.Sum([]byte(fmt.Sprintf("%s%d", key, answer)))
		if strings.HasPrefix(fmt.Sprintf("%x", hash), prefix) {
			return answer
		}
		answer++
	}
}
