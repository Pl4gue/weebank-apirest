package main

import (
	"math/rand"
	"time"
)

const charset = "abcdefghijklmnopqrstuvwxyz" +
	"0123456789"

var seededRand *rand.Rand = rand.New(
	rand.NewSource(time.Now().UnixNano()))

func StringWithCharset(length int) string {
	b := make([]byte, length)
	for i := range b {
		b[i] = charset[seededRand.Intn(len(charset))]
	}
	return string(b)
}

func reverseString(str string) string {
	byte_str := []rune(str)
	for i, j := 0, len(byte_str)-1; i < j; i, j = i+1, j-1 {
		byte_str[i], byte_str[j] = byte_str[j], byte_str[i]
	}
	return string(byte_str)
}

type User struct {
	Pub string `json:"pub"`
	Pvt string `json:"pvt"`
}

type Transaction struct {
	From  User   `json:"from"`
	To    string `json:"to"`
	Value int    `json:"value"`
}

func genarateKeys() User {
	pubKey := StringWithCharset(6)
	pvtKey := reverseString(pubKey)

	return User{
		Pub: pubKey,
		Pvt: pvtKey,
	}
}
