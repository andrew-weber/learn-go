package main

import (
	"context"
	"fmt"

	"github.com/nbd-wtf/go-nostr"
	"github.com/nbd-wtf/go-nostr/nip19"
)

func main() {
	relay, err := nostr.RelayConnect(context.Background(), "wss://relay.damus.io")
	if err != nil {
		panic(err)
	}

	npub := "npub1nqsaem8ejvznch6luqphvmq2tg84v7cjjhsavc8wtqm9807sxudscugesa"

	var filters nostr.Filters
	if _, v, err := nip19.Decode(npub); err == nil {
		pub := v.(string)
		filters = []nostr.Filter{{
			Kinds:   []int{1},
			Authors: []string{pub},
			Limit:   5,
		}}
	} else {
		panic(err)
	}

	ctx, _ := context.WithCancel(context.Background())
	sub := relay.Subscribe(ctx, filters)

	go func() {
		<-sub.EndOfStoredEvents
		// handle end of stored events (EOSE, see NIP-15)
	}()

	for ev := range sub.Events {
		// handle returned event.
		// channel will stay open until the ctx is cancelled (in this case, by calling cancel())

		fmt.Println(ev.Content)
	}

}
