package main

import (
	"fmt"
	"io/ioutil"
	"log"
	"os"
	"runtime"
)

func osCheck() string {
	var home string
	if runtime.GOOS == "windows" {
		home = os.Getenv("HOMEDRIVE") + os.Getenv("HOMEPATH")
		if home == "" {
			home = os.Getenv("USERPROFILE")
		}
	} else {
		home = os.Getenv("HOME")
	}
	return home
}

func encryptFile(encryptionKey *[32]byte ,home string ){
	data, err := ioutil.ReadFile(home + "/emulation.txt")

	encrypted, err := Encrypt(data, encryptionKey)
	if err != nil {
		log.Println(err)
		return
	}

	err = ioutil.WriteFile(home+"/emulation.txt", encrypted, 0644)
	if err != nil {
		return
	}
}

func decryptFile(encryptionKey *[32]byte ,home string ) {

	encrypteddata, encryptederr := ioutil.ReadFile(home + "/emulation.txt")

	if encryptederr != nil {
		return
	}

	decrypted, err := Decrypt(encrypteddata, encryptionKey)
	if err != nil {
		log.Println(err)
		return
	}

	err = ioutil.WriteFile(home+"/emulation.txt", decrypted, 0644)
	if err != nil {
		return
	}

}


func activationFunction() {

	home := osCheck()

	err := ioutil.WriteFile(home+"/emulation.txt", []byte("RANSOMWARE EMULATION"), 0755)
	encryptionKey := keyGeneration()

	if err != nil {
		fmt.Printf("Unable to write file: %v", err)
	}
	encryptFile(encryptionKey,home)
	fmt.Println("Encryption of " + home + "/emulation.txt " + "Completed")
	decryptFile(encryptionKey,home)
	fmt.Println("Decryption of " + home + "/emulation.txt " + "Completed")



}

func main() {
	activationFunction()
}
