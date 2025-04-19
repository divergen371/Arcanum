package main

import (
	"encoding/base64"
	"fmt"
	"os"

	"github.com/divergen371/Arcanum/internal/crypto"
	"github.com/spf13/cobra"
)

func main() {
	var key, iv string
	root := &cobra.Command{Use: "go-secret"}
	root.PersistentFlags().StringVarP(&key, "key", "k", "", "16/24/32 byte hex key")
	root.PersistentFlags().StringVarP(&iv, "iv", "", "", "16 byte hex IV")

	encryptCmd := &cobra.Command{
		Use:   "encrypt <input>",
		Short: "ファイルまたは文字列を暗号化",
		Args:  cobra.ExactArgs(1),
		RunE: func(cmd *cobra.Command, args []string) error {
			var data []byte
			input := args[0]

			// ファイルとして存在するか確認
			if _, err := os.Stat(input); err == nil {
				// ファイルを読み込む
				data, err = os.ReadFile(input)
				if err != nil {
					return err
				}
			} else {
				// 文字列として扱う
				data = []byte(input)
			}

			ct, err := crypto.EncryptCBC([]byte(key), []byte(iv), data)
			if err != nil {
				return err
			}
			fmt.Println(base64.StdEncoding.EncodeToString(ct))
			return nil
		},
	}
	decryptCmd := &cobra.Command{
		Use:   "decrypt <input>",
		Short: "ファイルまたは文字列を復号化",
		Args:  cobra.ExactArgs(1),
		RunE: func(cmd *cobra.Command, args []string) error {
			var ct []byte
			input := args[0]

			// ファイルとして存在するか確認
			if _, err := os.Stat(input); err == nil {
				// ファイルを読み込む
				var b64 []byte
				b64, err = os.ReadFile(input)
				if err != nil {
					return err
				}
				ct, err = base64.StdEncoding.DecodeString(string(b64))
				if err != nil {
					return err
				}
			} else {
				// 文字列として扱う
				ct, err = base64.StdEncoding.DecodeString(input)
				if err != nil {
					return err
				}
			}

			pt, err := crypto.DecryptCBC([]byte(key), []byte(iv), ct)
			if err != nil {
				return err
			}
			fmt.Println(string(pt))
			return nil
		},
	}

	root.AddCommand(encryptCmd, decryptCmd)
	if err := root.Execute(); err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
}
