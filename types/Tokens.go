package types

import (
	"sort"
	"strings"
)

type Tokens map[string]Token

func (t Tokens) TokenIs(addr string, id string) bool {
	token, exists := t[strings.ToLower(addr)]
	if !exists {
		return false
	}

	return strings.ToLower(id) == strings.ToLower(token.ID)
}

func (t Tokens) Keys() []string {
	var keys []string

	for k := range t {
		keys = append(keys, k)
	}

	sort.Strings(keys)

	return keys
}
