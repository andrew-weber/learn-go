package cmd

import (
	"context"
	"fmt"
	"strings"
	"time"

	"github.com/nbd-wtf/go-nostr"
	"github.com/spf13/cobra"
	"github.com/spf13/viper"
)

// addCmd represents the add command
var addCmd = &cobra.Command{
	Use:   "add",
	Short: "Add a new password",
	Long:  `Add a new password to the password manager`,
	Run: func(cmd *cobra.Command, args []string) {
		sk := nostr.GeneratePrivateKey()
		pub, _ := nostr.GetPublicKey(sk)

		ev := nostr.Event{
			PubKey:    pub,
			CreatedAt: time.Now(),
			Kind:      1,
			Tags:      nil,
			Content:   "Hello World!",
		}

		// calling Sign sets the event ID field and the event Sig field
		ev.Sign(sk)

		// publish the event to two relays
		for _, url := range strings.Split(viper.GetString("RELAY"), ",") {
			relay, e := nostr.RelayConnect(context.Background(), strings.TrimSpace(url))
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
