package main

import (
	"crypto/aes"
	"crypto/cipher"
	"crypto/rand"
	"encoding/hex"
	"fmt"
	"io"
	"log"
	"net/http"
)

func main() {
	// Initialize the database connection
	err := InitDB("postgres://username:MysteriousNJCF$28@localhost:5432/AEShield%20DB?sslmode=disable")
	if err != nil {
		log.Fatal("Failed to connect to database: ", err)
	}
	defer postgresDB.Close()

	http.HandleFunc("/encrypt", handleEncrypt)
	http.HandleFunc("/decrypt", handleDecrypt)

	fmt.Println("Server running on http://localhost:8080")
	http.ListenAndServe(":8080", nil)
}

func handleEncrypt(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodPost {
		http.Error(w, "Method not allowed", http.StatusMethodNotAllowed)
		return
	}

	// Get the key and file from the form data
	keyHex := r.FormValue("key")
	file, header, err := r.FormFile("file")
	if err != nil {
		http.Error(w, "File required", http.StatusBadRequest)
		return
	}
	defer file.Close()

	key, err := hex.DecodeString(keyHex)
	if err != nil || len(key) != 32 {
		http.Error(w, "Key must be 64 hex characters (32 bytes)", http.StatusBadRequest)
		return
	}

	// Read file content
	data, err := io.ReadAll(file)
	if err != nil {
		http.Error(w, "Failed to read file", http.StatusInternalServerError)
		return
	}

	encrypted, err := encryptBytes(data, key)
	if err != nil {
		http.Error(w, "Encryption failed: "+err.Error(), http.StatusInternalServerError)
		return
	}

	// Store the encrypted file and log the request
	if err := StoreFile(header.Filename, "encrypt", encrypted); err != nil {
		log.Println("Failed to store file: ", err)
	}

	// Set download headers
	w.Header().Set("Content-Disposition", "attachment; filename=\""+header.Filename+".enc\"")
	w.Header().Set("Content-Type", "application/octet-stream")
	w.Write(encrypted)
}

func handleDecrypt(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodPost {
		http.Error(w, "Method not allowed", http.StatusMethodNotAllowed)
		return
	}

	keyHex := r.FormValue("key")
	file, header, err := r.FormFile("file")
	if err != nil {
		http.Error(w, "File required", http.StatusBadRequest)
		return
	}
	defer file.Close()

	key, err := hex.DecodeString(keyHex)
	if err != nil || len(key) != 32 {
		http.Error(w, "Key must be 64 hex characters (32 bytes)", http.StatusBadRequest)
		return
	}

	data, err := io.ReadAll(file)
	if err != nil {
		http.Error(w, "Failed to read file", http.StatusInternalServerError)
		return
	}

	decrypted, err := decryptBytes(data, key)
	if err != nil {
		http.Error(w, "Decryption failed: "+err.Error(), http.StatusInternalServerError)
		return
	}

	// Store the decrypted file and log the request
	if err := StoreFile(header.Filename, "decrypt", decrypted); err != nil {
		log.Println("Failed to store file: ", err)
	}

	// Set download headers
	w.Header().Set("Content-Disposition", "attachment; filename=\""+header.Filename+".dec\"")
	w.Header().Set("Content-Type", "application/octet-stream")
	w.Write(decrypted)
}

// Core logic from your CLI version, modified to work with byte slices

func encryptBytes(plaintext []byte, key []byte) ([]byte, error) {
	block, err := aes.NewCipher(key)
	if err != nil {
		return nil, err
	}

	aesGCM, err := cipher.NewGCM(block)
	if err != nil {
		return nil, err
	}

	nonce := make([]byte, aesGCM.NonceSize())
	if _, err = rand.Read(nonce); err != nil {
		return nil, err
	}

	ciphertext := aesGCM.Seal(nonce, nonce, plaintext, nil)
	return ciphertext, nil
}

func decryptBytes(ciphertext []byte, key []byte) ([]byte, error) {
	block, err := aes.NewCipher(key)
	if err != nil {
		return nil, err
	}

	aesGCM, err := cipher.NewGCM(block)
	if err != nil {
		return nil, err
	}

	nonceSize := aesGCM.NonceSize()
	if len(ciphertext) < nonceSize {
		return nil, fmt.Errorf("ciphertext too short")
	}

	nonce := ciphertext[:nonceSize]
	ciphertext = ciphertext[nonceSize:]

	return aesGCM.Open(nil, nonce, ciphertext, nil)
}
