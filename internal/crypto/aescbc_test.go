package crypto

import (
	"bytes"
	"encoding/hex"
	"testing"
)

func TestEncryptDecryptCBC(t *testing.T) {
	key, _ := hex.DecodeString("6368616e676520746869732070617373")
	iv, _ := hex.DecodeString("1234567890abcdef1234567890abcdef")
	plaintext := []byte("This is a test message.")

	// Encrypt the plaintext
	ciphertext, err := EncryptCBC(key, iv, plaintext)
	if err != nil {
		t.Fatalf("Encryption failed: %v", err)
	}

	// Decrypt the ciphertext
	decrypted, err := DecryptCBC(key, iv, ciphertext)
	if err != nil {
		t.Fatalf("Decryption failed: %v", err)
	}

	// Check if the decrypted text matches the original plaintext
	if !bytes.Equal(decrypted, plaintext) {
		t.Errorf("Decrypted text does not match original plaintext.\nGot: %s\nWant: %s", decrypted, plaintext)
	}
}
