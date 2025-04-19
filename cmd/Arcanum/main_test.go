package main

import (
	"bytes"
	"os"
	"strings"
	"testing"
)

func TestEncryptDecrypt(t *testing.T) {
	// テスト用のキーとIV
	key := "6368616e676520746869732070617373" // 16バイトのキー
	iv := "1234567890abcdef"                  // 16バイトのIV

	// テスト用の文字列
	plaintext := "This is a test message."

	// テスト用のファイルを作成
	inputFile := "test_input.txt"
	encryptedFile := "test_encrypted.txt"
	decryptedFile := "test_decrypted.txt"

	err := os.WriteFile(inputFile, []byte(plaintext), 0o600)
	if err != nil {
		t.Fatalf("Failed to create test file: %v", err)
	}

	defer os.Remove(inputFile)
	defer os.Remove(encryptedFile)
	defer os.Remove(decryptedFile)

	// ファイルを暗号化して別のファイルに保存
	os.Args = []string{"go-secret", "encrypt", inputFile, "--key", key, "--iv", iv}

	// 標準出力をキャプチャ
	r, w, _ := os.Pipe()
	stdout := os.Stdout
	os.Stdout = w

	main()

	w.Close()
	var buf bytes.Buffer
	buf.ReadFrom(r)
	os.Stdout = stdout

	// 暗号化された出力を取得し、改行を削除
	encryptedOutput := strings.TrimSpace(buf.String())

	// 暗号化されたテキストを復号化
	r, w, _ = os.Pipe()
	os.Stdout = w

	os.Args = []string{"go-secret", "decrypt", encryptedOutput, "--key", key, "--iv", iv}
	main()

	w.Close()
	buf.Reset()
	buf.ReadFrom(r)
	os.Stdout = stdout

	// 復号化された出力を取得し、改行を削除
	decryptedOutput := strings.TrimSpace(buf.String())

	// 復号化されたテキストを確認
	if decryptedOutput != plaintext {
		t.Errorf("Decrypted text does not match original plaintext.\nExpected: %s\nGot: %s", plaintext, decryptedOutput)
	}
}
