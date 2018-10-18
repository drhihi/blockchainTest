package main

import (
	"crypto/aes"
	"crypto/cipher"
	"crypto/rand"
	"encoding/hex"
	"errors"
	"flag"
	"fmt"
	"io"
)

func main() {
	var key, data, method string
	flag.StringVar(&key,"key", "", "a privat key")
	flag.StringVar(&data,"data", "", "a data")
	flag.StringVar(&method,"m", "e", "e - encrypt (use default)\nd - decrypt")
	flag.Parse()

	if method == "e" {
		cipherdata, err := Encrypt([]byte(data), []byte(key))
		if err != nil {
			panic(err)
		}
		fmt.Printf("%s => %s\n", data, hex.EncodeToString(cipherdata))
	} else if method == "d" {
		cipherdata, _ := hex.DecodeString(data)
		plaindata, err := Decrypt(cipherdata, []byte(key))
		if err != nil {
			panic(err)
		}
		fmt.Printf("%s => %s\n", data, plaindata)
	} else {
		panic("method error")
	}

}

func Encrypt(data, key []byte) ([]byte, error){
	c, err := aes.NewCipher(key)

	if err != nil {
		return nil, err
	}

	gcm, err := cipher.NewGCM(c)
	if err != nil {
		return nil, err
	}

	nonce := make([]byte, gcm.NonceSize())
	if _, err = io.ReadFull(rand.Reader, nonce); err != nil {
		return nil, err
	}

	return gcm.Seal(nonce, nonce, data, nil), nil
}

func Decrypt(data, key []byte) ([]byte, error) {
	c, err := aes.NewCipher(key)
	if err != nil {
		return nil, err
	}

	gcm, err := cipher.NewGCM(c)
	if err != nil {
		return nil, err
	}

	nonceSize := gcm.NonceSize()
	if len(data) < nonceSize {
		return nil, errors.New("fail")
	}

	nonce, ciphertext := data[:nonceSize], data[nonceSize:]
	return gcm.Open(nil, nonce, ciphertext, nil)
}