package commands

import (
	"fmt"
	"strings"

	formatter "github.com/kwix/logrus-module-formatter"
	"github.com/sirupsen/logrus"
	"github.com/spf13/cobra"
	"github.com/spf13/viper"

	"github.com/barnbridge/barnbrdige-yieldfarm-cli/types"
)

func initLogging() {
	logging := viper.GetString("logging")

	if verbose {
		logging = "*=debug"
	}

	if vverbose {
		logging = "*=trace"
	}

	if logging == "" {
		logging = "*=info"
	}

	f, err := formatter.New(formatter.NewModulesMap(logging))
	if err != nil {
		panic(err)
	}

	logrus.SetFormatter(f)

	log.Debug("Debug mode")
}

func addContractsFlags(cmd *cobra.Command) {
	cmd.Flags().String("staking.addr", "", "Address of staking contract")
	cmd.Flags().String("yieldfarm.addr", "", "Address of yield farm contract")
	cmd.Flags().String("yieldfarmlp.addr", "", "Address of yield farm LP contract")
}

func bindViperToContractsFlags(cmd *cobra.Command) {
	viper.BindPFlag("staking.addr", cmd.Flag("staking.addr"))
	viper.BindPFlag("yieldfarm.addr", cmd.Flag("yieldfarm.addr"))
	viper.BindPFlag("yieldfarmlp.addr", cmd.Flag("yieldfarmlp.addr"))
}

func addRPCURLFlag(cmd *cobra.Command) {
	cmd.Flags().String("rpc-url", "", "Ethereum RPC url")
}

func bindViperToRPCURLFlag(cmd *cobra.Command) {
	viper.BindPFlag("rpc-url", cmd.Flag("rpc-url"))
}

func addTokensFlags(cmd *cobra.Command) {
	for _, t := range types.TokenIDs {
		cmd.Flags().String(
			fmt.Sprintf("tokens.%s.addr", t),
			"",
			fmt.Sprintf("Address for %s token", t),
		)

		cmd.Flags().String(
			fmt.Sprintf("tokens.%s.name", t),
			"",
			fmt.Sprintf("Name for %s token", t),
		)

		cmd.Flags().Int64(
			fmt.Sprintf("tokens.%s.decimals", t),
			18,
			fmt.Sprintf("Decimals for %s token", t),
		)
	}
}

func bindViperToTokensFlags(cmd *cobra.Command) {
	for _, t := range types.TokenIDs {
		flagName := fmt.Sprintf("tokens.%s.addr", t)
		viper.BindPFlag(flagName, cmd.Flag(flagName))
	}
}

func getTokens() types.Tokens {
	res := make(types.Tokens)

	for _, t := range types.TokenIDs {
		name := viper.GetString(fmt.Sprintf("tokens.%s.name", t))
		addr := viper.GetString(fmt.Sprintf("tokens.%s.addr", t))
		decimals := viper.GetInt64(fmt.Sprintf("tokens.%s.decimals", t))

		res[strings.ToLower(addr)] = types.Token{
			ID:       t,
			Addr:     addr,
			Name:     name,
			Decimals: decimals,
		}
	}

	return res
}
