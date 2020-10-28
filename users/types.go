package users

import (
	"math/big"

	"github.com/ALTree/bigfloat"
)

var tenPow12, _ = bigfloat.Pow(big.NewFloat(10), big.NewFloat(12)).Int(nil)

type Action struct {
	Type           string
	Amount         string
	BlockTimestamp int64
	Token          string
	TxHash         string
	User           string
}

type User struct {
	Addr            string
	NumDeposit      int64
	NumWithdraw     int64
	Balance         *big.Int
	PoolShare       *big.Float
	PotentialReward *big.Float
}

type usersByActions struct {
	Position    string `header:"#"`
	Addr        string `header:"address"`
	NumDeposit  int64  `header:"# deposit"`
	NumWithdraw int64  `header:"# withdraw"`
	Balance     string `header:"balance"`
}

type usersByBalance struct {
	Position        string `header:"#"`
	Addr            string `header:"address"`
	Balance         string `header:"balance"`
	PoolShare       string `header:"pool share"`
	PotentialReward string `header:"potential reward"`
}
