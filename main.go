package main

import (
	"encoding/base64"
	"fmt"
)

func main() {
	var whatIsIt string
	secret := "aWFuZ25vVzpOQU06RU5JTDp0YTpzdTpuaW9K"

	sd, _ := base64.StdEncoding.DecodeString(secret)
	whatIsIt = string(reverse(sd))

	fmt.Println(whatIsIt)
}

func reverse(xs []byte) []byte {
	var result []byte

	for i := len(xs) - 1; i >= 0; i-- {
		result = append(result, xs[i])
	}

	return result
}
