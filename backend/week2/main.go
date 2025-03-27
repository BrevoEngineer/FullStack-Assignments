package main

import (
	"fmt"
)

const prime64v1 = 11400714785074694791
const prime64v2 = 14029467366897019727
const prime64v3 = 1609587929392839161
const prime64v4 = 9650029242287828579
const prime64v5 = 2870177450012600261

func rotateLeft(val uint64, shift uint64) uint64 {
	return (val << shift) | (val >> (64 - shift))
}

func xxhash64(input string) uint64 {
	var h64 uint64 = 0xAAAAAAAAAAAAAAAA
	lenInput := uint64(len(input))
	const limit = 32

	var i uint64
	for i = 0; i+limit <= lenInput; i += limit {
		block := input[i : i+limit]
		h64 += uint64(len(block))
		h64 = rotateLeft(h64, 27)
		h64 ^= mixBlock(block)
		h64 = h64*prime64v1 + prime64v2
	}

	remainder := input[i:]
	if len(remainder) > 0 {
		h64 ^= uint64(len(remainder))
		h64 = rotateLeft(h64, 23)
		h64 ^= mixBlock(remainder)
		h64 = h64 * prime64v3
	}

	h64 ^= h64 >> 33
	h64 *= prime64v2
	h64 ^= h64 >> 29
	h64 *= prime64v3
	h64 ^= h64 >> 32

	return h64
}

func mixBlock(block string) uint64 {
	var hash uint64 = 0
	for i := 0; i < len(block); i++ {
		hash ^= uint64(block[i])
		hash = rotateLeft(hash, 7)
		hash *= prime64v1
	}
	return hash
}

const mod uint64 = 62 * 62 * 62 * 62 * 62 * 62 * 62 * 62 * 62 * 62

func base62Conversion(num uint64) string {
	const base62Chars = "0123456789ABCDEFGHIJKLMNOPQRSTUVWXYZabcdefghijklmnopqrstuvwxyz"
	if num == 0 {
		return "0"
	}
	encoded := ""
	base := uint64(62)

	for num > 0 {
		remainder := num % base
		encoded = string(base62Chars[remainder]) + encoded
		num /= base
	}
	return encoded
}

func generateHash(input string) string {
	hashVal := xxhash64(input)
	reduced := hashVal % mod
	hashString := base62Conversion(reduced)

	for len(hashString) < 10 {
		hashString = "0" + hashString
	}

	return hashString
}

func main() {
	fmt.Println(generateHash("Brevo"))
}
