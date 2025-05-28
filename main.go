package main

import (
	"crypto/rand"
	"encoding/base64"
	"fmt"
	"log"
	"os"

	"golang.org/x/crypto/argon2"
	"golang.org/x/crypto/bcrypt"
	"golang.org/x/crypto/ssh"
)

type User struct {
	Username     string
	PasswordHash string
	Salt         []byte
}

func generateSalt() ([]byte, error) {
	salt := make([]byte, 16)
	if _, err := rand.Read(salt); err != nil {
		return nil, err
	}
	return salt, nil
}

func hashPasswordWithArgon2(password string, salt []byte) string {
	// Using Argon2id with recommended parameters
	hash := argon2.IDKey([]byte(password), salt, 1, 64*1024, 4, 32)
	return base64.StdEncoding.EncodeToString(hash)
}

func hashPasswordWithBcrypt(password string) (string, error) {
	hash, err := bcrypt.GenerateFromPassword([]byte(password), bcrypt.DefaultCost)
	if err != nil {
		return "", err
	}
	return string(hash), nil
}

func verifyPassword(password, hash string) bool {
	err := bcrypt.CompareHashAndPassword([]byte(hash), []byte(password))
	return err == nil
}

func generateSSHKeyPair() (string, string, error) {
	privateKey, err := ssh.ParsePrivateKey([]byte(os.Getenv("SSH_PRIVATE_KEY")))
	if err != nil {
		return "", "", err
	}
	publicKey := privateKey.PublicKey()
	return string(ssh.MarshalAuthorizedKey(publicKey)), "", nil
}

func main() {
	// Create a new user
	user := &User{
		Username: "testuser",
	}

	// Generate salt for Argon2
	salt, err := generateSalt()
	if err != nil {
		log.Fatal("Failed to generate salt:", err)
	}
	user.Salt = salt

	// Hash password using both methods
	password := "SecurePassword123!"

	// Argon2 hashing
	argon2Hash := hashPasswordWithArgon2(password, salt)
	fmt.Printf("Argon2 Hash: %s\n", argon2Hash)

	// Bcrypt hashing
	bcryptHash, err := hashPasswordWithBcrypt(password)
	if err != nil {
		log.Fatal("Failed to hash password with bcrypt:", err)
	}
	fmt.Printf("Bcrypt Hash: %s\n", bcryptHash)

	// Verify password
	if verifyPassword(password, bcryptHash) {
		fmt.Println("Password verification successful!")
	}

	// Try with wrong password
	if !verifyPassword("WrongPassword", bcryptHash) {
		fmt.Println("Wrong password verification failed as expected")
	}

	// Generate SSH key pair (if SSH_PRIVATE_KEY environment variable is set)
	if os.Getenv("SSH_PRIVATE_KEY") != "" {
		publicKey, _, err := generateSSHKeyPair()
		if err != nil {
			log.Printf("Failed to generate SSH key pair: %v", err)
		} else {
			fmt.Printf("SSH Public Key: %s", publicKey)
		}
	}
}
