package main

import (
	"bufio"
	"crypto/aes"
	"crypto/cipher"
	"encoding/hex"
	"fmt"
	"io/ioutil"
	"os"
	"path/filepath"
	"strings"
)

func main() {
	var targetDir, keyString string

	fmt.Print("Enter the directory path containing encrypted files: ")
	scanner := bufio.NewScanner(os.Stdin)
	if scanner.Scan() {
		targetDir = scanner.Text()
	}

	fmt.Print("Enter the decryption key: ")
	if scanner.Scan() {
		keyString = scanner.Text()
	}

	globalKey, err := hex.DecodeString(keyString)
	if err != nil {
		fmt.Println("Error decoding key:", err)
		return
	}

	decryptDir(targetDir, globalKey)
}

func decryptDir(directory string, globalKey []byte) {
	files, err := ioutil.ReadDir(directory)
	if err != nil {
		fmt.Println("Error reading directory:", err)
		return
	}

	for _, file := range files {
		filePath := filepath.Join(directory, file.Name())
		if file.IsDir() {
			decryptDir(filePath, globalKey)
			continue
		}
		if strings.HasSuffix(file.Name(), ".thunderkitty.encrypted") {
			decrypt(filePath, globalKey)
		}
	}
}

func decrypt(filePath string, globalKey []byte) {
	encryptedData, err := ioutil.ReadFile(filePath)
	if err != nil {
		fmt.Println("Error reading encrypted file:", err)
		return
	}

	block, err := aes.NewCipher(globalKey)
	if err != nil {
		fmt.Println("Error creating cipher:", err)
		return
	}

	aesgcm, err := cipher.NewGCM(block)
	if err != nil {
		fmt.Println("Error creating AES GCM:", err)
		return
	}

	nonceSize := aesgcm.NonceSize()
	if len(encryptedData) < nonceSize {
		fmt.Println("Invalid encrypted data")
		return
	}

	nonce := encryptedData[:nonceSize]
	ciphertext := encryptedData[nonceSize:]

	plaintext, err := aesgcm.Open(nil, nonce, ciphertext, nil)
	if err != nil {
		fmt.Println("Error decrypting:", err)
		return
	}

	originalFilePath := strings.TrimSuffix(filePath, ".thunderkitty.encrypted")
	err = ioutil.WriteFile(originalFilePath, plaintext, 0644)
	if err != nil {
		fmt.Println("Error writing decrypted file:", err)
		return
	}

	fmt.Println("Decrypted file:", originalFilePath)

	err = os.Remove(filePath)
	if err != nil {
		fmt.Println("Error removing encrypted file:", err)
		return
	}
}
