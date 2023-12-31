package util

import "fmt"

const base62char = "0123456789abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ"

func EncodeToBase62(number int) string {
	base62 := ""
	for number > 0 {
		remainder := number % 62
		base62 = string(base62char[remainder]) + base62
		number /= 62
		fmt.Printf("Generated Base62 char: %v\n", base62)
	}

	return base62
}
