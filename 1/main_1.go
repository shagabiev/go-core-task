package main

import (
	"crypto/sha256"
	"encoding/hex"
	"fmt"
	"reflect"
	"strings"
)

func main() {
	var numDecimal int = 42           // Десятичная система
	var numOctal int = 052            // Восьмеричная система
	var numHexadecimal int = 0x2A     // Шестнадцатиричная система
	var pi float64 = 3.14             // Тип float64
	var name string = "Golang"        // Тип string
	var isActive bool = true          // Тип bool
	var complexNum complex64 = 1 + 2i // Тип complex64

	fmt.Printf("numDecimal: %s\n", getType(numDecimal))
	fmt.Printf("numOctal: %s\n", getType(numOctal))
	fmt.Printf("numHexadecimal: %s\n", getType(numHexadecimal))
	fmt.Printf("pi: %s\n", getType(pi))
	fmt.Printf("name: %s\n", getType(name))
	fmt.Printf("isActive: %s\n", getType(isActive))
	fmt.Printf("complexNum: %s\n", getType(complexNum))

	result := concatToString(
		numDecimal,
		numOctal,
		numHexadecimal,
		pi,
		name,
		isActive,
		complexNum,
	)

	runes := stringToRunes(result)
	hash := hashSHA256WithSalt(string(runes), "go-2024")

	fmt.Printf("SHA256 with salt: %s\n", hash)

	fmt.Println(hashSHA256WithSalt("abc", "go-2024"))
}

func getType(value any) string {
	return reflect.TypeOf(value).String()
}

func concatToString(values ...any) string {
	var b strings.Builder

	for _, v := range values {
		b.WriteString(fmt.Sprint(v))
	}

	return b.String()
}

func stringToRunes(input string) []rune {
	return []rune(input)
}

func hashSHA256WithSalt(input, salt string) string {
	middle := len(input) / 2

	hasher := sha256.New()
	hasher.Write([]byte(input[:middle]))
	hasher.Write([]byte(salt))
	hasher.Write([]byte(input[middle:]))

	return hex.EncodeToString(hasher.Sum(nil))
}
