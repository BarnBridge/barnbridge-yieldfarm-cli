package types

type Token struct {
	ID       string
	Addr     string
	Name     string
	Decimals int64
}

var TokenIDs = []string{"dai", "susd", "usdc", "unilp"}
