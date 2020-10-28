package stats

import (
	"fmt"
	"math/big"
	"strings"

	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/ethclient"
	"github.com/landoop/tableprinter"
	"github.com/sirupsen/logrus"
	"github.com/spf13/viper"

	"github.com/barnbridge/barnbrdige-yieldfarm-cli/contracts"
	"github.com/barnbridge/barnbrdige-yieldfarm-cli/types"
	"github.com/barnbridge/barnbrdige-yieldfarm-cli/utils"
)

var log = logrus.WithField("module", "stats")

type Config struct {
	RPCURL string
	User   string
	Tokens types.Tokens
}

type Stats struct {
	config  Config
	staking *contracts.Staking
	user    common.Address

	printer *tableprinter.Printer
}

func New(config Config) *Stats {
	var acc = common.HexToAddress(config.User)

	// Create an IPC based RPC connection to a remote node
	conn, err := ethclient.Dial(viper.GetString("rpc-url"))
	if err != nil {
		log.Fatalf("Failed to connect to the Ethereum client: %v", err)
	}

	s, err := contracts.NewStaking(common.HexToAddress(viper.GetString("staking.addr")), conn)
	if err != nil {
		log.Fatal(err)
	}

	return &Stats{
		config:  config,
		staking: s,
		user:    acc,
		printer: utils.GetPrinter(),
	}
}

func (s *Stats) Run() {
	currentEpoch, err := s.staking.GetCurrentEpoch(nil)
	if err != nil {
		log.Fatal(err)
	}

	fmt.Printf("Current staking epoch is: %s\n\n", currentEpoch.String())
	s.do(currentEpoch)

	fmt.Println()
	fmt.Println(strings.Repeat("─", 100))

	fmt.Printf("\nNext epoch:\n\n")
	nextEpoch := new(big.Int).Add(currentEpoch, big.NewInt(1))
	s.do(nextEpoch)
}

func (s *Stats) do(epoch *big.Int) {
	tokens := s.config.Tokens
	user := s.user
	var tableData []tokenBalances

	var balances = make(map[string]*big.Int)
	for _, v := range tokens.Keys() {
		token := tokens[v]
		totalBalance, err := s.staking.BalanceOf(nil, user, common.HexToAddress(token.Addr))
		if err != nil {
			log.Fatal(err)
		}

		epochBalance, err := s.staking.GetEpochUserBalance(nil, user, common.HexToAddress(v), epoch)
		if err != nil {
			log.Fatal(err)
		}

		balances[token.ID] = epochBalance

		tableData = append(tableData, tokenBalances{
			Token:            token.Name,
			TotalBalance:     utils.FormatBigFloat(utils.ScaleInt(totalBalance, token.Decimals)),
			EffectiveBalance: utils.FormatBigFloat(utils.ScaleInt(epochBalance, token.Decimals)),
		})
	}

	fmt.Println("Balances:")
	s.printer.Print(tableData)

	fmt.Printf("\nPool sizes:\n")

	var ps = make(map[string]*big.Int)
	for _, v := range tokens.Keys() {
		token := tokens[v]
		poolSize, err := s.staking.GetEpochPoolSize(nil, common.HexToAddress(token.Addr), epoch)
		if err != nil {
			log.Fatal(err)
		}

		ps[token.ID] = poolSize
	}

	var t2Data []poolSizes

	t2Data = append(t2Data, poolSizes{
		Pool: "USDC/DAI/sUSD",
		Size: utils.FormatBigFloat(getStablePoolSize(ps)),
		MyShare: utils.FormatBigFloat(
			getShareStable(ps, balances),
			10,
		),
		PotentialReward: utils.FormatBigFloat(getStableReward(ps, balances)),
	})

	t2Data = append(t2Data, poolSizes{
		Pool: "USDC_BOND_UNI_LP",
		Size: utils.FormatBigFloat(utils.ScaleInt(ps["unilp"], 18)),
		MyShare: utils.FormatBigFloat(
			getShareUniLP(ps, balances),
			10,
		),
		PotentialReward: utils.FormatBigFloat(getUniLPReward(ps, balances)),
	})

	s.printer.Print(t2Data)
}

func getStablePoolSize(ps map[string]*big.Int) *big.Float {
	size := new(big.Float)

	size = new(big.Float).Add(size, utils.ScaleInt(ps["usdc"], 6))
	size = new(big.Float).Add(size, utils.ScaleInt(ps["dai"], 18))
	size = new(big.Float).Add(size, utils.ScaleInt(ps["susd"], 18))

	return size
}

func getStablePoolBalance(balances map[string]*big.Int) *big.Float {
	balance := new(big.Float)

	balance = new(big.Float).Add(balance, utils.ScaleInt(balances["usdc"], 6))
	balance = new(big.Float).Add(balance, utils.ScaleInt(balances["dai"], 18))
	balance = new(big.Float).Add(balance, utils.ScaleInt(balances["susd"], 18))

	return balance
}

func getShareStable(ps map[string]*big.Int, balances map[string]*big.Int) *big.Float {
	size := getStablePoolSize(ps)
	if size.Cmp(big.NewFloat(0)) == 0 {
		return big.NewFloat(0)
	}

	balance := getStablePoolBalance(balances)

	return new(big.Float).Quo(balance, size)
}

func getShareUniLP(ps map[string]*big.Int, balances map[string]*big.Int) *big.Float {
	if ps["unilp"].Cmp(big.NewInt(0)) == 0 {
		return big.NewFloat(0)
	}

	return new(big.Float).Quo(utils.BigIntToBigFloat(balances["unilp"]), utils.BigIntToBigFloat(ps["unilp"]))
}

func getStableReward(ps map[string]*big.Int, balances map[string]*big.Int) *big.Float {
	var rewardsPerWeek = big.NewFloat(32000)

	share := getShareStable(ps, balances)

	return new(big.Float).Mul(share, rewardsPerWeek)
}

func getUniLPReward(ps map[string]*big.Int, balances map[string]*big.Int) *big.Float {
	var rewardsPerWeek = big.NewFloat(20000)

	share := getShareUniLP(ps, balances)

	return new(big.Float).Mul(share, rewardsPerWeek)
}
