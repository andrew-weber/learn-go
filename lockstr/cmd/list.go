package cmd

import (
	"context"
	"fmt"
	"strings"

	"github.com/nbd-wtf/go-nostr"
	"github.com/nbd-wtf/go-nostr/nip19"
	"github.com/spf13/cobra"
	"github.com/spf13/viper"
)

// listCmd represents the list command
var listCmd = &cobra.Command{
	Use:   "list",
	Short: "List password keys",
	Long:  ``,
	Run: func(cmd *cobra.Command, args []string) {

		for _, url := range strings.Split(viper.GetString("RELAY"), ",") {
			relay, err := nostr.RelayConnect(context.Background(), strings.TrimSpace(url))
			if err != nil {
				continue
			}

			var filter nostr.Filter
			if _, v, err := nip19.Decode(viper.GetString("NPUB")); err == nil {
				pub := v.(string)
				filter = nostr.Filter{
					Kinds:   []int{1},
					Authors: []string{pub},
				}
			} else {
				fmt.Println(err)
			}

			ctx, _ := context.WithCancel(context.Background())
			events := relay.QuerySync(ctx, filter)

			for _, ev := range events {
				fmt.Println(ev.ID, strings.TrimSpace(ev.Content))
			}
		}

	},
}

func init() {
	rootCmd.AddCommand(listCmd)

	// Here you will define your flags and configuration settings.

	// Cobra supports Persistent Flags which will work for this command
	// and all subcommands, e.g.:
	// listCmd.PersistentFlags().String("foo", "", "A help for foo")

	// Cobra supports local flags which will only run when this command
	// is called directly, e.g.:
	// listCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")
}
