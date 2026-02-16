package main

import (
	"encoding/base64"
	"fmt"
)

func main() {
	data := []byte("Hello~ Base 64 Encoding!")
	encoding := base64.StdEncoding.EncodeToString(data)
	fmt.Println("Base64 endoded :", encoding)
	decoded, err := base64.StdEncoding.DecodeString(encoding)
	if err!= nil {
		fmt.Println("Error in decoding", err)
		return
	}
	fmt.Println("Decoded :", string(decoded))

	urlSafeEncoded := base64.URLEncoding.EncodeToString(data)
	fmt.Println("URL Safe base64 encode :", urlSafeEncoded)

	urlDecoded, err := base64.URLEncoding.DecodeString(urlSafeEncoded)
	if err != nil {
		fmt.Println("Error Occured :", err)
	}
	fmt.Println("URL Safe decoded :", string(urlDecoded))
}