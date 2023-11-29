package cmd

import (
	"fmt"
	"math/rand"
	"os"

	"github.com/spf13/cobra"
)

var (
	generateCmd = &cobra.Command{
		Use:   "generate",
		Short: "Generate random passwords",
		Long: `Generate random passwords with customizable options.
	For example:
	passgen generate -l 8 -q 3 -d -s`,
		Run: generatePassword,
	}

	charset      = "ABCDEFGHIJKLMNOPQRSTUVWXYZabcdefghijklmnopqrstuvwxyz"
	digits       = "0123456789"
	specialChars = "!@#$%^&*()_+{}[]|;:,.<>?-="
)

func Execute() {
	err := generateCmd.Execute()
	if err != nil {
		os.Exit(1)
	}
}

func init() {
	flags := generateCmd.Flags()
	flags.IntP("lenght", "l", 12, "Lenght of generated password")
	flags.IntP("quantity", "q", 1, "How many different passwords to generate")
	flags.BoolP("digits", "d", false, "Include digits in generated password")
	flags.BoolP("special-chars", "s", false, "Include special characters in generated password")
}

func generatePassword(cmd *cobra.Command, args []string) {

	flags := cmd.Flags()
	lenght, _ := flags.GetInt("lenght")
	quantity, _ := flags.GetInt("quantity")
	isDigits, _ := flags.GetBool("digits")
	isSpecialChars, _ := flags.GetBool("special-chars")

	if isDigits {
		charset += digits
	}

	if isSpecialChars {
		charset += specialChars
	}

	password := make([]byte, lenght)

	for i := 1; i <= quantity; i++ {
		for i := range password {
			password[i] = charset[rand.Intn(len(charset))]
		}

		fmt.Println(string(password))
	}

}
