package commands

import (
	"github.com/spf13/cobra"
	"github.com/spf13/viper"

	"github.com/barnbridge/barnbrdige-yieldfarm-cli/stats"
)

var statsCmd = &cobra.Command{
	Use: "stats",
	PreRun: func(cmd *cobra.Command, args []string) {
		bindViperToRPCURLFlag(cmd)
		bindViperToContractsFlags(cmd)
		bindViperToTokensFlags(cmd)
	},
	Run: func(cmd *cobra.Command, args []string) {
		s := stats.New(stats.Config{
			RPCURL: viper.GetString("rpc-url"),
			User:   viper.GetString("userAddr"),
			Tokens: getTokens(),
		})

		s.Run()
	},
}

func init() {
	RootCmd.AddCommand(statsCmd)

	addRPCURLFlag(statsCmd)
	addContractsFlags(statsCmd)
	addTokensFlags(statsCmd)

	statsCmd.Flags().String("userAddr", "", "Address of the user for which to call the functions")
	viper.BindPFlag("userAddr", statsCmd.Flag("userAddr"))
}
