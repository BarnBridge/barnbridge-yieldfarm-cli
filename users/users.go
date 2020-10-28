package users

import (
	"fmt"
	"math/big"
	"sort"

	"github.com/landoop/tableprinter"
	"github.com/spf13/viper"

	"github.com/barnbridge/barnbrdige-yieldfarm-cli/types"
	"github.com/barnbridge/barnbrdige-yieldfarm-cli/utils"
)

type Config struct {
	Tokens types.Tokens
	Pool   string
}

type Users struct {
	config Config

	printer *tableprinter.Printer
}

func New(config Config) *Users {
	return &Users{
		config:  config,
		printer: utils.GetPrinter(),
	}
}

type Metrics struct {
	Name  string `header:"metric"`
	Value string `header:"value"`
}

func (u *Users) Run() {
	actions := fetchActions()

	switch u.config.Pool {
	case "stable":
		stableActions := u.getStableActions(actions)
		u.runOnActions(stableActions)
	case "lp":
		unilpActions := u.getUniLPActions(actions)
		u.runOnActions(unilpActions)
	default:
		stableActions := u.getStableActions(actions)
		u.runOnActions(stableActions)

		unilpActions := u.getUniLPActions(actions)
		u.runOnActions(unilpActions)
	}
}

type Stats struct {
	NumActions, NumDeposits, NumWithdrawals   int64
	TotalDepositedAmount, TotalWithdrewAmount *big.Int
	UniqueUsers                               int64
}

func (u *Users) runOnActions(actions []Action) {
	tokens := u.config.Tokens
	var uniqueUsers = make(map[string]*User)
	stats := Stats{
		NumActions:           int64(len(actions)),
		TotalDepositedAmount: new(big.Int),
		TotalWithdrewAmount:  new(big.Int),
	}

	for _, a := range actions {
		if _, exists := uniqueUsers[a.User]; !exists {
			uniqueUsers[a.User] = &User{
				Addr:    a.User,
				Balance: new(big.Int),
			}
		}

		amount, ok := new(big.Int).SetString(a.Amount, 10)
		if !ok {
			log.Fatal("could not convert amount to big int")
		}

		if tokens.TokenIs(a.Token, "usdc") {
			amount = new(big.Int).Mul(amount, tenPow12)
		}

		switch a.Type {
		case "DEPOSIT":
			stats.NumDeposits++
			stats.TotalDepositedAmount = new(big.Int).Add(stats.TotalDepositedAmount, amount)
			uniqueUsers[a.User].NumDeposit++
			uniqueUsers[a.User].Balance = new(big.Int).Add(uniqueUsers[a.User].Balance, amount)
		case "WITHDRAW":
			stats.NumWithdrawals++
			stats.TotalWithdrewAmount = new(big.Int).Add(stats.TotalWithdrewAmount, amount)
			uniqueUsers[a.User].NumWithdraw++
			uniqueUsers[a.User].Balance = new(big.Int).Sub(uniqueUsers[a.User].Balance, amount)
		}
	}

	stats.UniqueUsers = int64(len(uniqueUsers))

	var users []User
	for _, v := range uniqueUsers {
		users = append(users, *v)
	}

	n := viper.GetInt("top-n")
	if n < 1 {
		n = len(users)
	}
	if n > len(users) {
		n = len(users)
	}

	tvl := utils.ScaleInt(new(big.Int).Sub(stats.TotalDepositedAmount, stats.TotalWithdrewAmount), 18)
	for i := range users {
		users[i].PoolShare = new(big.Float).Quo(utils.ScaleInt(users[i].Balance, 18), tvl)
		users[i].PotentialReward = new(big.Float).Mul(users[i].PoolShare, big.NewFloat(32000))
	}

	printTopByActions(users, n, u.printer)
	printTopByBalance(users, n, u.printer)

	u.printStats(stats)
}

func (u *Users) printStats(stats Stats) {
	var m1 []Metrics
	m1 = append(m1, Metrics{
		Name:  "# of actions",
		Value: fmt.Sprintf("%d", stats.NumActions),
	})

	m1 = append(m1, Metrics{
		Name:  "# of deposits",
		Value: fmt.Sprintf("%d", stats.NumDeposits),
	})

	m1 = append(m1, Metrics{
		Name:  "# of withdrawals",
		Value: fmt.Sprintf("%d", stats.NumWithdrawals),
	})

	fmt.Printf("\nActions metrics\n")
	u.printer.Print(m1)

	var m2 []Metrics

	m2 = append(m2, Metrics{
		Name:  "Total deposits",
		Value: utils.FormatBigFloat(utils.ScaleInt(stats.TotalDepositedAmount, 18)),
	})

	m2 = append(m2, Metrics{
		Name:  "Total withdrawals",
		Value: utils.FormatBigFloat(utils.ScaleInt(stats.TotalWithdrewAmount, 18)),
	})

	tvl := new(big.Int).Sub(stats.TotalDepositedAmount, stats.TotalWithdrewAmount)

	m2 = append(m2, Metrics{
		Name:  "TVL",
		Value: utils.FormatBigFloat(utils.ScaleInt(tvl, 18)),
	})

	fmt.Printf("\nValue metrics\n")
	u.printer.Print(m2)

	var m3 []Metrics
	m3 = append(m3, Metrics{
		Name:  "Unique users",
		Value: fmt.Sprintf("%d", stats.UniqueUsers),
	})

	var avgDeposit = big.NewFloat(0)
	if stats.UniqueUsers > 0 {
		avgDeposit = utils.ScaleFloat(
			new(big.Float).Quo(
				new(big.Float).SetInt(stats.TotalDepositedAmount),
				big.NewFloat(float64(stats.UniqueUsers)),
			),
			18,
		)
	}

	m3 = append(m3, Metrics{
		Name:  "Avg deposit/user",
		Value: utils.FormatBigFloat(avgDeposit),
	})

	fmt.Printf("\nUsers metrics\n")
	u.printer.Print(m3)
}

func printTopByBalance(users []User, n int, printer *tableprinter.Printer) {
	fmt.Printf("\nTop N users by balance:\n")

	sort.Slice(users, func(i, j int) bool {
		return users[i].Balance.Cmp(users[j].Balance) > 0
	})

	var data2 []usersByBalance
	for i := 0; i < n; i++ {
		data2 = append(data2, usersByBalance{
			Position:        fmt.Sprintf("%d", i+1),
			Addr:            users[i].Addr,
			Balance:         fmt.Sprintf("%.2f", utils.ScaleInt(users[i].Balance, 18)),
			PoolShare:       fmt.Sprintf("%.10f", users[i].PoolShare),
			PotentialReward: fmt.Sprintf("%.10f", users[i].PotentialReward),
		})
	}
	printer.Print(data2)
}

func printTopByActions(users []User, n int, printer *tableprinter.Printer) {
	sort.Slice(users, func(i, j int) bool {
		return users[i].NumDeposit+users[i].NumWithdraw > users[j].NumDeposit+users[j].NumWithdraw
	})

	fmt.Printf("\nTop N users by # of actions:\n")

	var data []usersByActions
	for i := 0; i < n; i++ {
		data = append(data, usersByActions{
			Position:    fmt.Sprintf("%d", i+1),
			Addr:        users[i].Addr,
			NumDeposit:  users[i].NumDeposit,
			NumWithdraw: users[i].NumWithdraw,
			Balance:     fmt.Sprintf("%.2f", utils.ScaleInt(users[i].Balance, 18)),
		})
	}
	printer.Print(data)
}
