package cmd

import (
	"fmt"
	"os"
	"strings"

	"golang.org/x/term"

	nostr_lib "github.com/andrew-weber/lockstr/lib"
	"github.com/nbd-wtf/go-nostr"
	"github.com/nbd-wtf/go-nostr/nip04"

	"github.com/spf13/cobra"
	"github.com/spf13/viper"
)

// addCmd represents the add command
var addCmd = &cobra.Command{
	Use:   "add",
	Short: "Add a new password",
	Long:  `Add a new password to the password manager`,
	Run: func(cmd *cobra.Command, args []string) {
		if len(args) == 0 {
			fmt.Println("No key provided. Try 'lockstr add <key>'")
			return
		}

		fmt.Print("Enter Password: ")
		bpw, err := term.ReadPassword(int(os.Stdin.Fd()))
		if err != nil {
			panic(err)
		}

		fmt.Println("\nSaving Password...")

		sk := viper.GetString("KEY")
		pub, _ := nostr.GetPublicKey(sk)

		shared, _ := nip04.ComputeSharedSecret(pub, sk)
		pwd, _ := nip04.Encrypt(strings.TrimSpace(string(bpw)), shared)

		thing := &nostr_lib.Entry{
			Key:      strings.TrimSpace(strings.TrimSpace(string(args[0]))),
			Password: string(pwd),
		}

		nostr_lib.AddEntry(thing)
	},
}

func init() {
	rootCmd.AddCommand(addCmd)
}
