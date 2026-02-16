package intermediate

import (
	"crypto/rand"
	"crypto/sha256"
	"encoding/base64"
	"fmt"
	"io"
)

func main() {
	password := "Example@1234"
	// hash := sha256.Sum256([]byte(password))
	// fmt.Println("Password :", password)
	// fmt.Println("Hash :", hash)
	// fmt.Printf("SHA 256 Hex value : %x\n", hash)

	// hash512 := sha512.Sum512([]byte(password))
	// fmt.Printf("SHA 512 Hex value : %x\n", hash512)

	salt, err := generateSalt()

	if err != nil {
		fmt.Println("Error :", err)
		return
	}

	hash := hashPassword(password, salt)
	fmt.Println("Password :", password)
	fmt.Println("Hash :", string(hash))

	saltStr := base64.StdEncoding.EncodeToString(salt)
	fmt.Println("Salt String :", string(saltStr))

	decodedSalt, err := base64.StdEncoding.DecodeString(saltStr)
	if err != nil {
		fmt.Println("Error decoding salt :", err)
		return
	}
	
	loginPassword := enterPassword()
	loginHash := hashPassword(loginPassword, decodedSalt)
	fmt.Println("Login Hash :", loginHash)
	if loginHash == hash {
		fmt.Println("Correct Password, Logged In")
	} else {
		fmt.Println("Incorrect Password")
	}
}

func generateSalt() ([]byte, error){
	salt := make([]byte, 16)
	_, err := io.ReadFull(rand.Reader, salt)
	if err != nil {
		return nil, err
	}
	return salt, nil
}

func hashPassword(password string, salt []byte) string {
	saltedPassword := append(salt, []byte(password)...)
	hash := sha256.Sum256(saltedPassword)
	return base64.StdEncoding.EncodeToString(hash[:])
}

func enterPassword() string {
	var password string
	fmt.Print("Please enter Password : ")
	fmt.Scanln(&password)
	return password
}