package epochs

import (
	"fmt"
	"math/big"

	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/ethclient"
	"github.com/sirupsen/logrus"
	"github.com/spf13/viper"

	"github.com/barnbridge/barnbrdige-yieldfarm-cli/contracts"
	"github.com/barnbridge/barnbrdige-yieldfarm-cli/types"
)

var log = logrus.WithField("module", "epochs")

func Run(tokens types.Tokens) {
	// Create an IPC based RPC connection to a remote node
	conn, err := ethclient.Dial(viper.GetString("rpc-url"))
	if err != nil {
		log.Fatalf("Failed to connect to the Ethereum client: %v", err)
	}

	s, err := contracts.NewStaking(common.HexToAddress(viper.GetString("staking.addr")), conn)
	if err != nil {
		log.Fatal(err)
	}

	currentEpoch, err := s.GetCurrentEpoch(nil)
	if err != nil {
		log.Fatal(err)
	}

	fmt.Printf("Current staking epoch is: %s\n\n", currentEpoch.String())

	for _, t := range tokens.Keys() {
		token := tokens[t]
		fmt.Printf("%s: %d\n", token.Addr, findMostRecentEpoch(s, token.Addr, currentEpoch))
	}
}

func findMostRecentEpoch(staking *contracts.Staking, addr string, currentEpoch *big.Int) int64 {
	for i := currentEpoch.Int64() + 2; i >= 0; i-- {
		init, err := staking.EpochIsInitialized(nil, common.HexToAddress(addr), big.NewInt(i))
		if err != nil {
			log.Fatal(err)
		}

		if init {
			return i
		}
	}

	return -1
}
