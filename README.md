### ThunderKitty-Ransomware
## Telegram:
- https://t.me/ebytelabs
The encryption process involves the following steps:

1. **Generate a User ID**: A unique user ID is generated to identify the victim.
2. **Generate an Encryption Key**: A 256-bit AES encryption key is generated randomly.
3. **Encrypt Files**: All files in the specified target directory are recursively encrypted using AES-GCM (Galois/Counter Mode). Each file is encrypted with a unique nonce.
4. **Send Notification**: A notification containing the user ID, encrypted directory, and encryption key is sent to a specified Discord webhook.
5. **Create Ransom Note**: A ransom note is created on the victim's desktop, providing instructions on how to pay the ransom and recover the encrypted files.

### Decryption

The decryption process involves the following steps:

1. **Obtain Decryption Key**: The victim must obtain the decryption key by following the instructions in the ransom note.
2. **Decrypt Files**: The victim uses the provided decryption script to decrypt the encrypted files using the obtained key.

## Usage

### Ransomware.go

Replace the placeholders in the script with appropriate values:

- `dischook`: Your Discord webhook URL.
- `targetDir`: The directory to encrypt.
- `xmraddr`: Your Monero address for receiving the ransom.
- `cashamt`: The ransom amount.
- `email`: Your contact email.

![image](https://github.com/EvilBytecode/ThunderKitty-Ransomware/assets/151552809/5ea603f7-0fa4-4e60-a467-5d5a9f432ed1)
![image](https://github.com/EvilBytecode/ThunderKitty-Ransomware/assets/151552809/5a64ae2c-2d07-4998-9876-a070050195bd)


## License
This project is licensed under the MIT License. See the LICENSE file for details.
