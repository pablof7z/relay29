package main

import (
	"context"
	"fmt"

	"github.com/fiatjaf/khatru"
	"github.com/nbd-wtf/go-nostr"
)

func requireAuth(ctx context.Context, filter nostr.Filter) (reject bool, msg string) {
	pubkey := khatru.GetAuthed(ctx)

	if pubkey == "" {
		return true, "auth-required: something"
	}

	return false, ""
}

func requireKindAndSingleGroupID(ctx context.Context, filter nostr.Filter) (reject bool, msg string) {
	pubkey := khatru.GetAuthed(ctx)

	fmt.Println("pubkey:", pubkey)

	// if there is no pubkey, send back auth-required
	// if pubkey == "" {
	// 	return true, "auth-required: something"
	// }

	fmt.Println("pubkey:", pubkey)

	// isMeta := false
	isNormal := false
	for _, kind := range filter.Kinds {
		if kind < 10000 {
			isNormal = true
		} else if kind >= 30000 {
			// isMeta = true
		}
	}
	// if isNormal && isMeta {
	// 	return true, "cannot request both meta and normal events at the same time"
	// }
	// if !isNormal && !isMeta {
	// 	return true, "unexpected kinds requested"
	// }

	if isNormal {
		if tags, _ := filter.Tags["h"]; len(tags) == 0 {
			return true, "must have an 'h' tag"
		}
	}

	return false, ""
}
