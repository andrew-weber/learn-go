package cmd

import (
	"bufio"
	"context"
	"encoding/json"
	"fmt"
	"os"
	"strings"
	"time"

	"golang.org/x/term"

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

		reader := bufio.NewReader(os.Stdin)

		fmt.Print("Enter Key: ")
		k, err := reader.ReadString('\n')
		if err != nil {
			panic(err)
		}

		fmt.Print("Enter Password: ")
		bpw, err := term.ReadPassword(int(os.Stdin.Fd()))
		if err != nil {
			panic(err)
		}

		type kvp struct {
			Key string
			Val string
		}

		thing := kvp{
			Key: strings.TrimSpace(string(k)),
			Val: strings.TrimSpace(string(bpw)),
		}

		result, _ := json.Marshal(thing)

		fmt.Println("\n\b", thing.Key, thing.Val, string(result))

		sk := viper.GetString("KEY")
		pub, _ := nostr.GetPublicKey(sk)

		shared, _ := nip04.ComputeSharedSecret(pub, sk)
		msg, _ := nip04.Encrypt(string(result), shared)

		ev := nostr.Event{
			PubKey:    pub,
			CreatedAt: time.Now(),
			Kind:      4,
			Tags:      nostr.Tags{nostr.Tag{"p", pub}},
			Content:   msg,
		}

		// calling Sign sets the event ID field and the event Sig field
		ev.Sign(sk)

		// publish the event to two relays
		for _, url := range strings.Split(viper.GetString("RELAY"), ",") {
			url := strings.TrimSpace(url)

			relay, e := nostr.RelayConnect(context.Background(), url)
			if e != nil {
				fmt.Println(e)
				continue
			}
			fmt.Println("published to ", url, relay.Publish(context.Background(), ev))
		}

	},
}

func init() {
	rootCmd.AddCommand(addCmd)

	// Here you will define your flags and configuration settings.

	// Cobra supports Persistent Flags which will work for this command
	// and all subcommands, e.g.:
	// addCmd.PersistentFlags().String("foo", "", "A help for foo")

	// Cobra supports local flags which will only run when this command
	// is called directly, e.g.:
	// addCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")
}
