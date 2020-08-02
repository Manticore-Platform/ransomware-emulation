package main

import (
	"crypto/rand"
	"crypto/rsa"
	"crypto/sha256"
	"fmt"
	"io"
	"math/big"
)

func fromBase10(base10 string) *big.Int {
	i, ok := new(big.Int).SetString(base10, 10)
	if !ok {
		panic("bad number: " + base10)
	}
	return i
}

func EncryptionKeyGeneration() *[32]byte {
	key := [32]byte{}
	_, err := io.ReadFull(rand.Reader, key[:])
	if err != nil {
		panic(err)
	}
	return &key
}

var Key rsa.PrivateKey


func keyGeneration()  *[32]byte{
	keys, _ := rsa.GenerateKey(rand.Reader, 2048)

	Key = rsa.PrivateKey{
		PublicKey: rsa.PublicKey{
			N: fromBase10(keys.N.String()), // yes, yes change all of those
			E: 65537,
		},
		D: fromBase10(keys.D.String()),
		Primes: []*big.Int{
			fromBase10(keys.Primes[0].String()),
			fromBase10(keys.Primes[1].String()),
		},
	}
	Key.Precompute()
	randomKey := EncryptionKeyGeneration()

	fmt.Println("AES KEY GENERATION : ",randomKey)

	encryptedKey, _ := rsa.EncryptOAEP(sha256.New(), rand.Reader, &Key.PublicKey, randomKey[:], nil)

	fmt.Println("ENCRYPTED AES KEY VIA GENERATED PUBLIC KEY : ",encryptedKey)

	aes_key, _ := rsa.DecryptOAEP(sha256.New(), rand.Reader, &Key, encryptedKey, nil)

	fmt.Println("DECRYPTED AES KEY VIA GENERATED PRIVATE KEY", aes_key)

	return randomKey

}
